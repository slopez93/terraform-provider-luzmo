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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &AccountResource{}
	_ resource.ResourceWithConfigure   = &AccountResource{}
	_ resource.ResourceWithImportState = &AccountResource{}
)

type AccountResource struct {
	lzService *services.LuzmoService
}

func NewAccountResource() resource.Resource {
	return &AccountResource{}
}

func (r *AccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (r *AccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a account.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "String identifier of the order.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the account.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the account.",
				Required:    true,
			},
			"provider_name": schema.StringAttribute{
				Description: "The slug of your own plugin or one of a database.",
				Required:    true,
			},
			"scope": schema.StringAttribute{
				Description: "Provider-specific description of services used.",
				Optional:    true,
			},
			"host": schema.StringAttribute{
				Description: "Endpoint of this account. For relational database connections, this corresponds to the hostname of the database.",
				Optional:    true,
			},
			"active": schema.BoolAttribute{
				Description: "Indicates whether queries may be sent to this database or plugin connection.",
				Optional:    true,
			},
			"port": schema.StringAttribute{
				Description: "Port of this connection. For relational database connections, this corresponds to the port of the database.",
				Optional:    true,
			},
			"cache": schema.Int64Attribute{
				Description: "Number of seconds queries to this data connector are cached in Luzmo's caching layer. Use 0 to disable caching.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"datasets_meta_sync_enabled": schema.BoolAttribute{
				Description: "Indicates whether automatic metadata sync is enabled for all connection datasets with meta_sync_inherit=true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"datasets_meta_sync_interval": schema.Int32Attribute{
				Description: " Metadata sync interval in hours for connection datasets with meta_sync_inherit=true.",
				Optional:    true,
				Computed:    true,
				Default:     int32default.StaticInt32(1),
			},
		},
	}
}

// Create implements resource.Resource.
func (r *AccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dtos.AccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	account := models.NewAccount(models.NewAccountParams{
		Id:                       plan.ID.ValueString(),
		Name:                     plan.Name.ValueString(),
		Description:              plan.Description.ValueString(),
		ProviderName:             plan.ProviderName.ValueString(),
		Scope:                    plan.Scope.ValueStringPointer(),
		Host:                     plan.Host.ValueStringPointer(),
		Active:                   plan.Active.ValueBoolPointer(),
		Port:                     plan.Port.ValueStringPointer(),
		Cache:                    plan.Cache.ValueInt64(),
		DatasetsMetaSyncEnabled:  plan.DatasetsMetaSyncEnabled.ValueBool(),
		DatasetsMetaSyncInterval: plan.DatasetsMetaSyncInterval.ValueInt32Pointer(),
	})

	accountResponse, err := r.lzService.CreateAccount(*account)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Luzmo Account",
			"Could not create Account ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	planResult := *r.lzService.Mapper.MapToAccountResource(*accountResponse)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, planResult)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read implements resource.Resource.
func (r *AccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dtos.AccountResourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	account, err := r.lzService.FindAccountById(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Account",
			"Could not read Account ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	stateResult := *r.lzService.Mapper.MapToAccountResource(*account)

	// Set refreshed state
	diags = resp.State.Set(ctx, &stateResult)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update implements resource.Resource.
func (r *AccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan dtos.AccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	luzmoAccount, err := r.lzService.FindAccountById(plan.ID.ValueString())
	if err != nil || luzmoAccount == nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Account",
			"Could not read Account ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	account := models.NewAccount(models.NewAccountParams{
		Id:                       plan.ID.ValueString(),
		Name:                     plan.Name.ValueString(),
		Description:              plan.Description.ValueString(),
		ProviderName:             plan.ProviderName.ValueString(),
		Scope:                    plan.Scope.ValueStringPointer(),
		Host:                     plan.Host.ValueStringPointer(),
		Active:                   plan.Active.ValueBoolPointer(),
		Port:                     plan.Port.ValueStringPointer(),
		Cache:                    plan.Cache.ValueInt64(),
		DatasetsMetaSyncEnabled:  plan.DatasetsMetaSyncEnabled.ValueBool(),
		DatasetsMetaSyncInterval: plan.DatasetsMetaSyncInterval.ValueInt32Pointer(),
	})
	updatedAccount, err := r.lzService.UpdateAccount(*account)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Luzmo Account",
			"Could not update Account ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	planResult := *r.lzService.Mapper.MapToAccountResource(*updatedAccount)

	// Set refreshed state
	diags = resp.State.Set(ctx, planResult)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete implements resource.Resource.
func (r *AccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state dtos.AccountResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.lzService.DeleteAccount(state.ID.ValueString())
	if err != nil {
		tflog.Info(ctx, err.Error())
		resp.Diagnostics.AddError(
			"Error Deleting Luzmo Account",
			"Could not delete Account, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *AccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Configure implements resource.ResourceWithConfigure.
func (r *AccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
