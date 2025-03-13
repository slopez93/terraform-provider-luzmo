package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type DatasetResourceModel struct {
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	Subtitle         types.String `tfsdk:"subtitle"`
	SubType          types.String `tfsdk:"subtype"`
	SourceDataset    types.String `tfsdk:"source_dataset"`
	SourceSheet      types.String `tfsdk:"source_sheet"`
	Transformation   types.String `tfsdk:"transformation"`
	Cache            types.Int64  `tfsdk:"cache"`
	UpdateMetadata   types.Bool   `tfsdk:"update_metadata"`
	MetaSyncInterval types.Int32  `tfsdk:"meta_sync_interval"`
	MetaSyncInherit  types.Bool   `tfsdk:"meta_sync_inherit"`
	MetaSyncEnabled  types.Bool   `tfsdk:"meta_sync_enabled"`
	DatasetId        types.String `tfsdk:"dataset_id"`
	ProviderName     types.String `tfsdk:"provider_name"`
}

type LuzmoDatasetDTO struct {
	Id                         string                 `json:"id"`
	Subtype                    string                 `json:"subtype"`
	Name                       map[string]string      `json:"name"`
	Description                map[string]string      `json:"description"`
	Subtitle                   *map[string]string     `json:"subtitle"`
	Contents                   map[string]interface{} `json:"contents"`
	Css                        *string                `json:"css"`
	SourceDataset              string                 `json:"source_dataset"`
	SourceSheet                string                 `json:"source_sheet"`
	Transformation             *string                `json:"transformation"`
	Cache                      int64                  `json:"cache"`
	Storage                    string                 `json:"storage"`
	IsVariant                  bool                   `json:"is_variant"`
	MigratedRows               *int64                 `json:"migrated_rows"`
	UsesClickhouseExperimental bool                   `json:"uses_clickhouse_experimental"`
	MetaSyncInterval           int32                  `json:"meta_sync_interval"`
	MetaSyncInherit            bool                   `json:"meta_sync_inherit"`
	MetaSyncEnabled            bool                   `json:"meta_sync_enabled"`
	CreatedAt                  string                 `json:"created_at"`
	UpdatedAt                  string                 `json:"updated_at"`
	AccelerationId             *string                `json:"acceleration_id"`
	AccountId                  string                 `json:"account_id"`
	TemplateId                 *string                `json:"template_id"`
	OriginalId                 *string                `json:"original_id"`
	ModifierId                 *string                `json:"modifier_id"`
	OwnerId                    string                 `json:"owner_id"`
	SourceTemplateId           *string                `json:"source_template_id"`
	UpdateMetadata             bool                   `json:"update_metadata"`
}

type LuzmoCreateDatasetDTO struct {
	Name             map[string]string  `json:"name"`
	Type             string             `json:"type"`
	Description      map[string]string  `json:"description"`
	SubTitle         map[string]*string `json:"sub_title"`
	SubType          string             `json:"sub_type"`
	SourceDataset    string             `json:"source_dataset"`
	SourceSheet      string             `json:"source_sheet"`
	Transformation   *string            `json:"transformation"`
	Cache            int64              `json:"cache"`
	UpdateMetadata   bool               `json:"update_metadata"`
	MetaSyncInterval int32              `json:"meta_sync_interval"`
	MetaSyncInherit  bool               `json:"meta_sync_inherit"`
	MetaSyncEnabled  bool               `json:"meta_sync_enabled"`
}

type LuzmoUpdateDatasetDTO struct {
	Name             map[string]string `json:"name"`
	Description      map[string]string `json:"description"`
	SubTitle         map[string]string `json:"sub_title"`
	SubType          string            `json:"sub_type"`
	Cache            int64             `json:"cache"`
	UpdateMetadata   bool              `json:"update_metadata"`
	MetaSyncInterval int32             `json:"meta_sync_interval"`
	MetaSyncInherit  bool              `json:"meta_sync_inherit"`
	MetaSyncEnabled  bool              `json:"meta_sync_enabled"`
}

type LuzmoDataProviderCreateDatasetsDTO struct {
	Action    string    `json:"action"`
	AccountId string    `json:"account_id"`
	Provider  string    `json:"provider"`
	Datasets  *[]string `json:"datasets"`
}

type LuzmoDataProviderCreateDatasetsResponseDTO struct {
	Count int `json:"count"`
	Data  []LuzmoDatasetDTO
}
