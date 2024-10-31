package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type AccountResourceModel struct {
	ID                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	Provider                types.String `tfsdk:"provider"`
	Scope                   types.String `tfsdk:"scope"`
	Host                    types.String `tfsdk:"host"`
	Active                  types.Bool   `tfsdk:"active"`
	Invalid                 types.Bool   `tfsdk:"invalid"`
	Port                    types.Int32  `tfsdk:"port"`
	Cache                   types.Int64  `tfsdk:"cache"`
	DatasetMetaSyncEnabled  types.Bool   `tfsdk:"dataset_meta_sync_enabled"`
	DatasetMetaSyncInterval types.Int32  `tfsdk:"dataset_meta_sync_interval"`
}

type LuzmoAccountDTO struct {
	Properties              map[string]string `json:"properties"`
	Date                    string            `json:"date"`
	Active                  bool              `json:"active"`
	Invalid                 bool              `json:"invalid"`
	Cache                   int64             `json:"cache"`
	Synced                  bool              `json:"synced"`
	DatasetMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
	Provider                string            `json:"provider"`
	Name                    string            `json:"name"`
	Description             map[string]string `json:"description"`
	Id                      string            `json:"id"`
	UserId                  string            `json:"user_id"`
	UpdatedAt               string            `json:"updated_at"`
	CreatedAt               string            `json:"created_at"`
	Identifier              *string           `json:"identifier"`
	Expiry                  *string           `json:"expiry"`
	Scope                   *string           `json:"scope"`
	Host                    *string           `json:"host"`
	Port                    *int32            `json:"port"`
	Version                 *string           `json:"version"`
	PluginId                string            `json:"plugin_id"`
	ShareableId             string            `json:"shareable_id"`
}

type LuzmoCreateAccountDTO struct {
	Name                    string            `json:"name"`
	Description             map[string]string `json:"description"`
	Provider                string            `json:"provider"`
	Scope                   *string           `json:"scope"`
	Host                    *string           `json:"host"`
	Active                  bool              `json:"active"`
	Invalid                 bool              `json:"invalid"`
	Port                    *int32            `json:"port"`
	Cache                   int64             `json:"cache"`
	DatasetMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
}

type LuzmoUpdateAccountDTO struct {
	Name                    string            `json:"name"`
	Description             map[string]string `json:"description"`
	Provider                string            `json:"provider"`
	Scope                   *string           `json:"scope"`
	Host                    *string           `json:"host"`
	Active                  bool              `json:"active"`
	Invalid                 bool              `json:"invalid"`
	Port                    *int32            `json:"port"`
	Cache                   int64             `json:"cache"`
	DatasetMetaSyncEnabled  bool              `json:"dataset_meta_sync_enabled"`
	DatasetMetaSyncInterval *int32            `json:"dataset_meta_sync_interval"`
}
