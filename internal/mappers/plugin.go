package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (m *Mapper) MapToPluginResource(plugin models.Plugin) *dtos.PluginResourceModel {
	var pluginModel dtos.PluginResourceModel

	pluginModel.ID = types.StringValue(plugin.Id)
	pluginModel.Name = types.StringValue(plugin.Name)
	pluginModel.Description = types.StringValue(plugin.Description)
	pluginModel.Slug = types.StringValue(plugin.Slug)
	pluginModel.BaseUrl = types.StringValue(plugin.BaseUrl)
	pluginModel.Url = types.StringPointerValue(plugin.Url)
	pluginModel.Pushdown = types.BoolValue(plugin.Pushdown)
	pluginModel.ProtocolVersion = types.StringValue(string(plugin.ProtocolVersion))
	pluginModel.SupportsLike = types.BoolValue(plugin.SupportsLike)
	pluginModel.SupportsDistinctcount = types.BoolValue(plugin.SupportsDistinctcount)
	pluginModel.SupportsOrderLimit = types.BoolValue(plugin.SupportsOrderLimit)
	pluginModel.SupportsJoin = types.BoolValue(plugin.SupportsJoin)
	pluginModel.SupportsSql = types.BoolValue(plugin.SupportsSql)
	pluginModel.SupportsNestedFilters = types.BoolValue(plugin.SupportsNestedFilters)

	return &pluginModel
}
