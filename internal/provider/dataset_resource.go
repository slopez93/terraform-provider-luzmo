package provider

import (
	"context"
	"fmt"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
	services "terraform-provider-luzmo/internal/services/luzmo"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &DatasetResource{}
	_ resource.ResourceWithConfigure   = &DatasetResource{}
	_ resource.ResourceWithImportState = &DatasetResource{}
)

type DatasetResource struct {
	lzService *services.LuzmoService
}

func NewDatasetResource() resource.Resource {
	return &DatasetResource{}
}

func (r *DatasetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dataset"
}

func (r *DatasetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a dataset.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "String identifier of the order.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the dataset.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the dataset.",
				Required:    true,
			},
			"subtitle": schema.StringAttribute{
				Description: "The subtitle of the dataset.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
			},
			"subtype": schema.StringAttribute{
				Description: "The subtype of the dataset.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("api"),
			},
			"source_dataset": schema.StringAttribute{
				Description: "The source dataset of the dataset.",
				Required:    true,
			},
			"source_sheet": schema.StringAttribute{
				Description: "The source sheet of the dataset.",
				Required:    true,
			},
			"transformation": schema.StringAttribute{
				Description: "The transformation of the dataset.",
				Optional:    true,
			},
			"cache": schema.Int64Attribute{
				Description: "Number of seconds queries to this data connector are cached in Luzmo's caching layer. Use 0 to disable caching.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"update_metadata": schema.BoolAttribute{
				Description: "Virtual property is used to trigger manual update for dataset metadata.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"meta_sync_interval": schema.Int32Attribute{
				Description: " Configure Metadata sync interval in hours for the dataset when meta_sync_inherit=false.",
				Optional:    true,
				Computed:    true,
				Default:     int32default.StaticInt32(1),
			},
			"meta_sync_inherit": schema.BoolAttribute{
				Description: "Indicates whether automatic metadata sync is enabled for this dataset.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"meta_sync_enabled": schema.BoolAttribute{
				Description: "Indicates whether automatic metadata sync is enabled for the dataset when meta_sync_inherit=false. ",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"last_metadata_sync_at": schema.StringAttribute{
				Description: " Last time metadata was synced was successful for the dataset.",
				Optional:    true,
			},
			"dataset_id": schema.StringAttribute{
				Description: "The dataset id.",
				Optional:    true,
			},
			"provider_name": schema.StringAttribute{
				Description: "The slug of your own plugin or one of a database. This is required for data provider datasets.",
				Optional:    true,
			},
		},
	}
}

// Create implements resource.Resource.
func (r *DatasetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dtos.DatasetResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dataset := models.NewDataset(models.NewDatasetParams{
		Id:                 plan.ID.ValueString(),
		Name:               plan.Name.ValueString(),
		Description:        plan.Description.ValueString(),
		SubTitle:           plan.Subtitle.ValueStringPointer(),
		SubType:            plan.SubType.ValueString(),
		SourceDataset:      plan.SourceDataset.ValueString(),
		SourceSheet:        plan.SourceSheet.ValueString(),
		Transformation:     plan.Transformation.ValueStringPointer(),
		Cache:              plan.Cache.ValueInt64(),
		UpdateMetadata:     plan.UpdateMetadata.ValueBool(),
		MetaSyncInterval:   *plan.MetaSyncInterval.ValueInt32Pointer(),
		MetaSyncInherit:    plan.MetaSyncInherit.ValueBool(),
		MetaSyncEnabled:    plan.MetaSyncEnabled.ValueBoolPointer(),
		LastMetadataSyncAt: plan.LastMetadataSyncAt.ValueStringPointer(),
		DatasetId:          plan.DatasetId.ValueStringPointer(),
		ProviderName:       plan.ProviderName.ValueStringPointer(),
	})

	mustBeCreatedByDataProvider, err := dataset.MustBeCreatedByDataProvider()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Luzmo Dataset",
			err.Error(),
		)
		return
	}

	if mustBeCreatedByDataProvider {
		dataProviderResponse, err := r.lzService.CreateDatasets(*dataset)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Creating Luzmo Dataset",
				"Could not create Dataset ID "+plan.ID.ValueString()+": "+err.Error(),
			)
			return
		}

		dataset.Id = dataProviderResponse[0].Id

		updatedDataset, err := r.lzService.UpdateDataset(*dataset)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Updating Luzmo Dataset",
				"Could not update Dataset ID "+dataset.Id+": "+err.Error(),
			)
			return
		}

		updatedDataset.DatasetId = dataset.DatasetId
		updatedDataset.ProviderName = dataset.ProviderName

		plan = *r.lzService.Mapper.MapToDatasetResource(*updatedDataset)
	} else {
		datasetResponse, err := r.lzService.CreateDataset(*dataset)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Creating Luzmo Dataset",
				"Could not create Dataset ID "+plan.ID.ValueString()+": "+err.Error(),
			)
			return
		}

		plan = *r.lzService.Mapper.MapToDatasetResource(*datasetResponse)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read implements resource.Resource.
func (r *DatasetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dtos.DatasetResourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dataset, err := r.lzService.FindDatasetById(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Dataset",
			"Could not read Dataset ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	datasetToCheck := models.NewDataset(models.NewDatasetParams{
		Id:                 dataset.Id,
		Name:               dataset.Name,
		Description:        dataset.Description,
		SubTitle:           dataset.Subtitle,
		SubType:            dataset.Subtype,
		SourceDataset:      dataset.SourceDataset,
		SourceSheet:        dataset.SourceSheet,
		Transformation:     dataset.Transformation,
		Cache:              dataset.Cache,
		UpdateMetadata:     dataset.UpdateMetadata,
		MetaSyncInterval:   dataset.MetaSyncInterval,
		MetaSyncInherit:    dataset.MetaSyncInherit,
		MetaSyncEnabled:    &dataset.MetaSyncEnabled,
		LastMetadataSyncAt: dataset.LastMetadataSyncAt,
		DatasetId:          state.DatasetId.ValueStringPointer(),
		ProviderName:       state.ProviderName.ValueStringPointer(),
	})

	mustBeCreatedByDataProvider, err := datasetToCheck.MustBeCreatedByDataProvider()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Checking Dataset Creation Requirement",
			err.Error(),
		)
		return
	}

	if mustBeCreatedByDataProvider {
		dataset.ProviderName = state.ProviderName.ValueStringPointer()
		dataset.DatasetId = state.DatasetId.ValueStringPointer()
	}

	state = *r.lzService.Mapper.MapToDatasetResource(*dataset)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update implements resource.Resource.
func (r *DatasetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan dtos.DatasetResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	luzmoDataset, err := r.lzService.FindDatasetById(plan.ID.ValueString())
	if err != nil || luzmoDataset == nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Dataset",
			"Could not read Dataset ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	dataset := models.NewDataset(models.NewDatasetParams{
		Id:               plan.ID.ValueString(),
		Name:             plan.Name.ValueString(),
		Description:      plan.Description.ValueString(),
		SubTitle:         plan.Subtitle.ValueStringPointer(),
		SubType:          plan.SubType.ValueString(),
		Cache:            plan.Cache.ValueInt64(),
		UpdateMetadata:   plan.UpdateMetadata.ValueBool(),
		MetaSyncInterval: *plan.MetaSyncInterval.ValueInt32Pointer(),
		MetaSyncInherit:  plan.MetaSyncInherit.ValueBool(),
		MetaSyncEnabled:  plan.MetaSyncEnabled.ValueBoolPointer(),
		DatasetId:        plan.DatasetId.ValueStringPointer(),
		ProviderName:     plan.ProviderName.ValueStringPointer(),
	})

	updatedDataset, err := r.lzService.UpdateDataset(*dataset)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Luzmo Dataset",
			"Could not update Dataset ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	mustBeCreatedByDataProvider, err := dataset.MustBeCreatedByDataProvider()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Checking Dataset Creation Requirement",
			err.Error(),
		)
		return
	}

	if mustBeCreatedByDataProvider {
		updatedDataset.DatasetId = plan.DatasetId.ValueStringPointer()
		updatedDataset.ProviderName = plan.ProviderName.ValueStringPointer()
	}

	plan = *r.lzService.Mapper.MapToDatasetResource(*updatedDataset)

	// Set refreshed state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete implements resource.Resource.
func (r *DatasetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state dtos.DatasetResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.lzService.DeleteDataset(state.ID.ValueString())
	if err != nil {
		tflog.Info(ctx, err.Error())
		resp.Diagnostics.AddError(
			"Error Deleting Luzmo Dataset",
			"Could not delete Dataset, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *DatasetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Configure implements resource.ResourceWithConfigure.
func (r *DatasetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	lzService, ok := req.ProviderData.(*services.LuzmoService)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.lzService = lzService
}
