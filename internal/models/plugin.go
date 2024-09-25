package models

type ProtocolVersion string

const (
	Version_1     ProtocolVersion = "1.0.0"
	Version_1_1_0 ProtocolVersion = "1.1.0"
	Version_2     ProtocolVersion = "2.0.0"
	Version_3     ProtocolVersion = "3.0.0"
)

type Plugin struct {
	Id                    string
	Name                  string
	Description           string
	BaseUrl               string
	Url                   *string
	Slug                  string
	Pushdown              bool
	ProtocolVersion       ProtocolVersion
	SupportsLike          bool
	SupportsDistinctcount bool
	SupportsOrderLimit    bool
	SupportsJoin          bool
	SupportsSql           bool
	SupportsNestedFilters bool
}

type NewPluginParams struct {
	Id                    string
	Name                  string
	Description           string
	BaseUrl               string
	Url                   *string
	Slug                  string
	Pushdown              *bool
	ProtocolVersion       ProtocolVersion
	SupportsLike          bool
	SupportsDistinctcount bool
	SupportsOrderLimit    bool
	SupportsJoin          bool
	SupportsSql           bool
	SupportsNestedFilters bool
}

func NewPlugin(params NewPluginParams) *Plugin {
	var pushdown *bool = params.Pushdown

	plugin := Plugin{
		Id:                    params.Id,
		Name:                  params.Name,
		Description:           params.Description,
		BaseUrl:               params.BaseUrl,
		Url:                   params.Url,
		Slug:                  params.Slug,
		ProtocolVersion:       params.ProtocolVersion,
		SupportsLike:          params.SupportsLike,
		SupportsDistinctcount: params.SupportsDistinctcount,
		SupportsOrderLimit:    params.SupportsOrderLimit,
		SupportsJoin:          params.SupportsJoin,
		SupportsSql:           params.SupportsSql,
		SupportsNestedFilters: params.SupportsNestedFilters,
	}

	if params.Pushdown == nil {
		*pushdown = false
	}

	plugin.Pushdown = *pushdown

	return &plugin
}

func (p *Plugin) IsV3() bool {
	return p.ProtocolVersion == Version_3
}
