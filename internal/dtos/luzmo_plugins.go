package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type PluginResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	Slug                  types.String `tfsdk:"slug"`
	BaseUrl               types.String `tfsdk:"base_url"`
	Url                   types.String `tfsdk:"url"`
	Pushdown              types.Bool   `tfsdk:"pushdown"`
	ProtocolVersion       types.String `tfsdk:"protocol_version"`
	SupportsLike          types.Bool   `tfsdk:"supports_like"`
	SupportsDistinctcount types.Bool   `tfsdk:"supports_distinctcount"`
	SupportsOrderLimit    types.Bool   `tfsdk:"supports_order_limit"`
	SupportsJoin          types.Bool   `tfsdk:"supports_join"`
	SupportsSql           types.Bool   `tfsdk:"supports_sql"`
	SupportsNestedFilters types.Bool   `tfsdk:"supports_nested_filters"`
}

type LuzmoPluginDTO struct {
	Id              string            `json:"id"`
	Name            map[string]string `json:"name"`
	Description     map[string]string `json:"description"`
	Slug            string            `json:"slug"`
	ProtocolVersion string            `json:"protocol_version"`
	BaseUrl         string            `json:"base_url"`
	Url             *string           `json:"url"`
	Color           string            `json:"color"`
	License         string            `json:"license"`
	Authorize       string            `json:"authorize"`
	Tiles           bool              `json:"tiles"`
	Search          bool              `json:"search"`
	Public          bool              `json:"public"`
	Reviewed        bool              `json:"reviewed"`
	Pushdown        *bool             `json:"pushdown"`

	SupportsLike          bool `json:"supports_like"`
	SupportsDistinctcount bool `json:"supports_distinctcount"`
	SupportsOrderLimit    bool `json:"supports_order_limit"`
	SupportsJoin          bool `json:"supports_join"`
	SupportsSql           bool `json:"supports_sql"`
	SupportsNestedFilters bool `json:"supports_nested_filters"`
}

type LuzmoCreatePluginDTO struct {
	Name                  map[string]string `json:"name"`
	Description           map[string]string `json:"description"`
	Slug                  string            `json:"slug"`
	BaseUrl               string            `json:"base_url"`
	Url                   *string           `json:"url"`
	ProtocolVersion       string            `json:"protocol_version"`
	Pushdown              *bool             `json:"pushdown"`
	SupportsLike          bool              `json:"supports_like"`
	SupportsDistinctcount bool              `json:"supports_distinctcount"`
	SupportsOrderLimit    bool              `json:"supports_order_limit"`
	SupportsJoin          bool              `json:"supports_join"`
	SupportsSql           bool              `json:"supports_sql"`
	SupportsNestedFilters bool              `json:"supports_nested_filters"`
}

type LuzmoUpdatePluginDTO struct {
	Name                  map[string]string `json:"name"`
	Description           map[string]string `json:"description"`
	BaseUrl               string            `json:"base_url"`
	Url                   *string           `json:"url"`
	ProtocolVersion       string            `json:"protocol_version"`
	Pushdown              *bool             `json:"pushdown"`
	SupportsLike          bool              `json:"supports_like"`
	SupportsDistinctcount bool              `json:"supports_distinctcount"`
	SupportsOrderLimit    bool              `json:"supports_order_limit"`
	SupportsJoin          bool              `json:"supports_join"`
	SupportsSql           bool              `json:"supports_sql"`
	SupportsNestedFilters bool              `json:"supports_nested_filters"`
}
