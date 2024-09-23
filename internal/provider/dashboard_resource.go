package provider

import (
	"context"
	"fmt"
	"terraform-provider-luzmo/internal/models"
	services "terraform-provider-luzmo/internal/services/luzmo"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &DashboardResource{}
	_ resource.ResourceWithConfigure   = &DashboardResource{}
	_ resource.ResourceWithImportState = &DashboardResource{}
)

type DashboardResource struct {
	lzService *services.LuzmoService
}

type DashboardResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Subtitle    types.String `tfsdk:"subtitle"`
	Contents    types.String `tfsdk:"contents"`
	Css         types.String `tfsdk:"css"`
}

func NewDashboardResource() resource.Resource {
	return &DashboardResource{}
}

func (r *DashboardResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dashboard"
}

func (r *DashboardResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a dashboard.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the dashboard.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the dashboard.",
				Required:    true,
			},
			"subtitle": schema.StringAttribute{
				Description: "The description of the dashboard.",
				Required:    true,
			},
			"contents": schema.StringAttribute{
				Description: "The contents of the dashboard.",
				Required:    true,
			},
			"css": schema.StringAttribute{
				Description: "The css of the dashboard.",
				Optional:    true,
			},
		},
	}
}

// Create implements resource.Resource.
func (r *DashboardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DashboardResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboard, err := models.NewDashboard(models.NewDashboardParams{
		Id:          plan.ID.ValueString(),
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		Subtitle:    plan.Subtitle.ValueString(),
		Contents:    plan.Contents.ValueString(),
		Css:         plan.Css.ValueStringPointer(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Dashboard",
			"Could not create Dashboard ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	responseDashboard, err := r.lzService.CreateDashboard(*dashboard)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Luzmo Dashboard",
			"Could not create Dashboard ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(responseDashboard.Id)
	plan.Contents = types.StringValue(responseDashboard.NormalizeContents())

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read implements resource.Resource.
func (r *DashboardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DashboardResourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboard, err := r.lzService.FindDashboardById(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Dashboard",
			"Could not create Dashboard ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state.ID = types.StringValue(dashboard.Id)
	state.Name = types.StringValue(dashboard.Name)
	state.Description = types.StringValue(dashboard.Description)
	state.Subtitle = types.StringValue(dashboard.Subtitle)
	state.Contents = types.StringValue(dashboard.NormalizeContents())
	if dashboard.Css != nil {
		state.Css = types.StringValue(*dashboard.Css)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update implements resource.Resource.
func (r *DashboardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DashboardResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboard, err := models.NewDashboard(models.NewDashboardParams{
		Id:          plan.ID.ValueString(),
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		Subtitle:    plan.Subtitle.ValueString(),
		Contents:    plan.Contents.ValueString(),
		Css:         plan.Css.ValueStringPointer(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Dashboard Model",
			"Could not create Dashboard ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	responseDashboard, err := r.lzService.UpdateDashboard(*dashboard)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Luzmo Dashboard",
			"Could not update Dashboard ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(responseDashboard.Id)
	plan.Contents = types.StringValue(responseDashboard.NormalizeContents())

	// Set refreshed state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete implements resource.Resource.
func (r *DashboardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state DashboardResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.lzService.DeleteDashboard(state.ID.ValueString())
	if err != nil {
		tflog.Info(ctx, err.Error())
		resp.Diagnostics.AddError(
			"Error Deleting Luzmo Dashboard",
			"Could not delete dashboard, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *DashboardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Configure implements resource.ResourceWithConfigure.
func (r *DashboardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
