package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type AccountResourceModel struct {
	ID                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	ProviderName             types.String `tfsdk:"provider_name"`
	Scope                    types.String `tfsdk:"scope"`
	Host                     types.String `tfsdk:"host"`
	Active                   types.Bool   `tfsdk:"active"`
	Port                     types.Int32  `tfsdk:"port"`
	Cache                    types.Int64  `tfsdk:"cache"`
	DatasetsMetaSyncEnabled  types.Bool   `tfsdk:"datasets_meta_sync_enabled"`
	DatasetsMetaSyncInterval types.Int32  `tfsdk:"datasets_meta_sync_interval"`
}

type LuzmoAccountDTO struct {
	Properties               map[string]string `json:"properties"`
	Date                     string            `json:"date"`
	Active                   bool              `json:"active"`
	Cache                    int64             `json:"cache"`
	Synced                   bool              `json:"synced"`
	DatasetsMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetsMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
	ProviderName             string            `json:"provider"`
	Name                     string            `json:"name"`
	Description              map[string]string `json:"description"`
	Id                       string            `json:"id"`
	UserId                   string            `json:"user_id"`
	UpdatedAt                string            `json:"updated_at"`
	CreatedAt                string            `json:"created_at"`
	Identifier               *string           `json:"identifier"`
	Expiry                   *string           `json:"expiry"`
	Scope                    *string           `json:"scope"`
	Host                     *string           `json:"host"`
	Port                     *int32            `json:"port"`
	Version                  *string           `json:"version"`
	PluginId                 string            `json:"plugin_id"`
	ShareableId              string            `json:"shareable_id"`
}

type LuzmoCreateAccountDTO struct {
	Name                     string            `json:"name"`
	Description              map[string]string `json:"description"`
	ProviderName             string            `json:"provider"`
	Scope                    *string           `json:"scope"`
	Host                     *string           `json:"host"`
	Active                   bool              `json:"active"`
	Port                     *int32            `json:"port"`
	Cache                    int64             `json:"cache"`
	DatasetsMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetsMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
}

type LuzmoUpdateAccountDTO struct {
	Name                     string            `json:"name"`
	Description              map[string]string `json:"description"`
	ProviderName             string            `json:"provider"`
	Scope                    *string           `json:"scope"`
	Host                     *string           `json:"host"`
	Active                   bool              `json:"active"`
	Port                     *int32            `json:"port"`
	Cache                    int64             `json:"cache"`
	DatasetsMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetsMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
}
