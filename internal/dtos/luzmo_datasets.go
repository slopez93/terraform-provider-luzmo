package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type DatasetResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Description        types.String `tfsdk:"description"`
	SubTitle           types.String `tfsdk:"subTitle"`
	Type               types.String `tfsdk:"type"`
	SubType            types.String `tfsdk:"subType"`
	SourceDataset      types.String `tfsdk:"source_dataset"`
	SourceSheet        types.String `tfsdk:"source_sheet"`
	Transformation     types.String `tfsdk:"transformation"`
	Cache              types.Int64  `tfsdk:"cache"`
	UpdateMetadata     types.Bool   `tfsdk:"update_metadata"`
	MetaSyncInterval   types.Int32  `tfsdk:"meta_sync_interval"`
	MetaSyncInherit    types.Bool   `tfsdk:"meta_sync_inherit"`
	MetaSyncEnabled    types.Bool   `tfsdk:"meta_sync_enabled"`
	LastMetadataSyncAt types.String `tfsdk:"last_metadata_sync_at"`
}

type LuzmoDatasetDTO struct {
	Id                         string                 `json:"id"`
	Type                       string                 `json:"type"`
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
	MetaSyncInterval           *int32                 `json:"meta_sync_interval"`
	MetaSyncInherit            bool                   `json:"meta_sync_inherit"`
	MetaSyncEnabled            bool                   `json:"meta_sync_enabled"`
	LastMetadataSyncAt         *string                `json:"last_metadata_sync_at"`
	LastMetadataSyncAttemptAt  *string                `json:"last_metadata_sync_attempt_at"`
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
	Name               map[string]string `json:"name"`
	Description        map[string]string `json:"description"`
	SubTitle           map[string]string `json:"sub_title"`
	Type               *string           `json:"type"`
	SubType            *string           `json:"sub_type"`
	SourceDataset      string            `json:"source_dataset"`
	SourceSheet        string            `json:"source_sheet"`
	Transformation     *string           `json:"transformation"`
	Cache              int64             `json:"cache"`
	UpdateMetadata     bool              `json:"update_metadata"`
	MetaSyncInterval   *int32            `json:"meta_sync_interval"`
	MetaSyncInherit    bool              `json:"meta_sync_inherit"`
	MetaSyncEnabled    bool              `json:"meta_sync_enabled"`
	LastMetadataSyncAt *string           `json:"last_metadata_sync_at"`
}

type LuzmoUpdateDatasetDTO struct {
	Name               map[string]string `json:"name"`
	Description        map[string]string `json:"description"`
	SubTitle           map[string]string `json:"sub_title"`
	Type               *string           `json:"type"`
	SubType            *string           `json:"sub_type"`
	SourceDataset      string            `json:"source_dataset"`
	SourceSheet        string            `json:"source_sheet"`
	Transformation     *string           `json:"transformation"`
	Cache              int64             `json:"cache"`
	UpdateMetadata     bool              `json:"update_metadata"`
	MetaSyncInterval   *int32            `json:"meta_sync_interval"`
	MetaSyncInherit    bool              `json:"meta_sync_inherit"`
	MetaSyncEnabled    bool              `json:"meta_sync_enabled"`
	LastMetadataSyncAt *string           `json:"last_metadata_sync_at"`
}
