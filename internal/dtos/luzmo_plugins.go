package dtos

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

	SupportsLike          *bool `json:"supports_like"`
	SupportsDistinctcount *bool `json:"supports_distinctcount"`
	SupportsOrderLimit    *bool `json:"supports_order_limit"`
	SupportsJoin          *bool `json:"supports_join"`
	SupportsSql           *bool `json:"supports_sql"`
	Pushdown              *bool `json:"pushdown"`
}

type LuzmoCreatePluginDTO struct {
	Name                  map[string]string `json:"name"`
	Description           map[string]string `json:"description"`
	Slug                  string            `json:"slug"`
	BaseUrl               string            `json:"base_url"`
	Url                   *string           `json:"url"`
	ProtocolVersion       string            `json:"protocol_version"`
	Pushdown              *bool             `json:"pushdown"`
	SupportsLike          *bool             `json:"supports_like"`
	SupportsDistinctcount *bool             `json:"supports_distinctcount"`
	SupportsOrderLimit    *bool             `json:"supports_order_limit"`
	SupportsJoin          *bool             `json:"supports_join"`
	SupportsSql           *bool             `json:"supports_sql"`
}

type LuzmoUpdatePluginDTO struct {
	Name                  map[string]string `json:"name"`
	Description           map[string]string `json:"description"`
	BaseUrl               string            `json:"base_url"`
	Url                   *string           `json:"url"`
	ProtocolVersion       string            `json:"protocol_version"`
	Pushdown              *bool             `json:"pushdown"`
	SupportsLike          *bool             `json:"supports_like"`
	SupportsDistinctcount *bool             `json:"supports_distinctcount"`
	SupportsOrderLimit    *bool             `json:"supports_order_limit"`
	SupportsJoin          *bool             `json:"supports_join"`
	SupportsSql           *bool             `json:"supports_sql"`
}
