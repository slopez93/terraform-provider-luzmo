package provider

import (
	"context"
	services "terraform-provider-luzmo/internal/services/luzmo"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &luzmoProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &luzmoProvider{
			version: version,
		}
	}
}

type luzmoProvider struct {
	version string
}

type luzmoProviderConfig struct {
	ApiKey     types.String `tfsdk:"api_key"`
	ApiToken   types.String `tfsdk:"api_token"`
	ApiUrl     types.String `tfsdk:"api_url"`
	ApiVersion types.String `tfsdk:"api_version"`
}

func (p *luzmoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "luzmo"
	resp.Version = p.version
}

func (p *luzmoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Description: "",
				Required:    true,
			},
			"api_token": schema.StringAttribute{
				Description: "",
				Required:    true,
			},
			"api_url": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"api_version": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
		},
	}
}

func (p *luzmoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Luzmo client")

	var config luzmoProviderConfig
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	lzService, err := services.NewLuzmoService(config.ApiKey.ValueString(), config.ApiToken.ValueString(), config.ApiVersion.ValueString(), config.ApiUrl.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Luzmo API Client",
			"An unexpected error occurred when creating the Luzmo API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Luzmo Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = lzService
	resp.ResourceData = lzService

	tflog.Info(ctx, "Configured Luzmo client", map[string]any{"success": true})
}

func (p *luzmoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *luzmoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDashboardResource,
		NewPluginResource,
	}
}

func (p *luzmoProvider) Functions(_ context.Context) []func() function.Function {
	return nil
}
