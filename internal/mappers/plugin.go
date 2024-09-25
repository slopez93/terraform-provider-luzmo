package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (m *Mapper) MapToPlugin(pluginDto dtos.LuzmoPluginDTO) (*models.Plugin, error) {
	plugin := models.Plugin{
		Id:                    pluginDto.Id,
		Name:                  pluginDto.Name[config.DefaultLang],
		Description:           pluginDto.Description[config.DefaultLang],
		Slug:                  pluginDto.Slug,
		ProtocolVersion:       models.ProtocolVersion(pluginDto.ProtocolVersion),
		BaseUrl:               pluginDto.BaseUrl,
		Url:                   pluginDto.Url,
		Pushdown:              *pluginDto.Pushdown,
		SupportsLike:          pluginDto.SupportsLike,
		SupportsDistinctcount: pluginDto.SupportsDistinctcount,
		SupportsOrderLimit:    pluginDto.SupportsOrderLimit,
		SupportsJoin:          pluginDto.SupportsJoin,
		SupportsSql:           pluginDto.SupportsSql,
		SupportsNestedFilters: pluginDto.SupportsNestedFilters,
	}

	return &plugin, nil
}
