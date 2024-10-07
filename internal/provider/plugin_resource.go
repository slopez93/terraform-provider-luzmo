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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &PluginResource{}
	_ resource.ResourceWithConfigure   = &PluginResource{}
	_ resource.ResourceWithImportState = &PluginResource{}
)

type PluginResource struct {
	lzService *services.LuzmoService
}

func NewPluginResource() resource.Resource {
	return &PluginResource{}
}

func (r *PluginResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin"
}

func (r *PluginResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Plugin.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "String identifier of the order.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the plugin.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the plugin.",
				Required:    true,
			},
			"slug": schema.StringAttribute{
				Description: "The slug of the plugin.",
				Required:    true,
			},
			"base_url": schema.StringAttribute{
				Description: "The Base URL of the plugin.",
				Required:    true,
			},
			"url": schema.StringAttribute{
				Description: "The URL of the plugin.",
				Optional:    true,
			},
			"protocol_version": schema.StringAttribute{
				Description: "The Protocol Version of the plugin.",
				Required:    true,
			},
			"pushdown": schema.BoolAttribute{
				Description: "The Pushdown of the plugin.",
				Optional:    true,
			},
			"supports_like": schema.BoolAttribute{
				Description: "The pushdown of the plugin.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"supports_distinctcount": schema.BoolAttribute{
				Description: "The supports_distinctcount of the plugin.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"supports_order_limit": schema.BoolAttribute{
				Description: "The supports_order_limit of the plugin.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"supports_join": schema.BoolAttribute{
				Description: "The supports_join of the plugin.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"supports_sql": schema.BoolAttribute{
				Description: "The supports_sql of the plugin.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"supports_nested_filters": schema.BoolAttribute{
				Description: "The supports_nested_filters of the plugin",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

// Create implements resource.Resource.
func (r *PluginResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dtos.PluginResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plugin := models.NewPlugin(models.NewPluginParams{
		Id:                    plan.ID.ValueString(),
		Name:                  plan.Name.ValueString(),
		Description:           plan.Description.ValueString(),
		Slug:                  plan.Slug.ValueString(),
		BaseUrl:               plan.BaseUrl.ValueString(),
		Url:                   plan.Url.ValueStringPointer(),
		Pushdown:              plan.Pushdown.ValueBoolPointer(),
		ProtocolVersion:       models.ProtocolVersion(plan.ProtocolVersion.ValueString()),
		SupportsLike:          plan.SupportsLike.ValueBool(),
		SupportsDistinctcount: plan.SupportsDistinctcount.ValueBool(),
		SupportsOrderLimit:    plan.SupportsOrderLimit.ValueBool(),
		SupportsJoin:          plan.SupportsJoin.ValueBool(),
		SupportsSql:           plan.SupportsSql.ValueBool(),
		SupportsNestedFilters: plan.SupportsNestedFilters.ValueBool(),
	})

	pluginResponse, err := r.lzService.CreatePlugin(*plugin)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Luzmo Plugin",
			"Could not create Plugin ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	plan = *r.lzService.Mapper.MapToPluginResource(*pluginResponse)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read implements resource.Resource.
func (r *PluginResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dtos.PluginResourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plugin, err := r.lzService.FindPluginById(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Finding Luzmo Plugin",
			"Could not read Plugin ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state = *r.lzService.Mapper.MapToPluginResource(*plugin)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update implements resource.Resource.
func (r *PluginResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan dtos.PluginResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plugin := models.NewPlugin(models.NewPluginParams{
		Id:                    plan.ID.ValueString(),
		Name:                  plan.Name.ValueString(),
		Description:           plan.Description.ValueString(),
		Slug:                  plan.Slug.ValueString(),
		BaseUrl:               plan.BaseUrl.ValueString(),
		Url:                   plan.Url.ValueStringPointer(),
		Pushdown:              plan.Pushdown.ValueBoolPointer(),
		ProtocolVersion:       models.ProtocolVersion(plan.ProtocolVersion.ValueString()),
		SupportsLike:          plan.SupportsLike.ValueBool(),
		SupportsDistinctcount: plan.SupportsDistinctcount.ValueBool(),
		SupportsOrderLimit:    plan.SupportsOrderLimit.ValueBool(),
		SupportsJoin:          plan.SupportsJoin.ValueBool(),
		SupportsSql:           plan.SupportsSql.ValueBool(),
		SupportsNestedFilters: plan.SupportsNestedFilters.ValueBool(),
	})
	updatedPlugin, err := r.lzService.UpdatePlugin(*plugin)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Luzmo Plugin",
			"Could not update Plugin ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	plan = *r.lzService.Mapper.MapToPluginResource(*updatedPlugin)

	// Set refreshed state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete implements resource.Resource.
func (r *PluginResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state dtos.PluginResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.lzService.DeletePlugin(state.ID.ValueString())
	if err != nil {
		tflog.Info(ctx, err.Error())
		resp.Diagnostics.AddError(
			"Error Deleting Luzmo Plugin",
			"Could not delete plugin, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *PluginResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Configure implements resource.ResourceWithConfigure.
func (r *PluginResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
