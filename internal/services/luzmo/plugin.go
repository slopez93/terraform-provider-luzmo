package services

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (ls *LuzmoService) CreatePlugin(plugin models.Plugin) (*models.Plugin, error) {
	payload := dtos.LuzmoCreateRequest[dtos.LuzmoCreatePluginDTO]{
		Action:  "create",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Properties: dtos.LuzmoCreatePluginDTO{
			Name: map[string]string{
				config.DefaultLang: plugin.Name,
			},
			Description: map[string]string{
				config.DefaultLang: plugin.Description,
			},
			Slug:                  plugin.Slug,
			BaseUrl:               plugin.BaseUrl,
			Url:                   plugin.Url,
			ProtocolVersion:       string(plugin.ProtocolVersion),
			Pushdown:              &plugin.Pushdown,
			SupportsLike:          plugin.SupportsLike,
			SupportsDistinctcount: plugin.SupportsDistinctcount,
			SupportsOrderLimit:    plugin.SupportsOrderLimit,
			SupportsJoin:          plugin.SupportsJoin,
			SupportsSql:           plugin.SupportsSql,
			SupportsNestedFilters: plugin.SupportsNestedFilters,
		},
	}

	body, err := ls.doRequest(PluginApiPath, payload)
	if err != nil {
		return nil, err
	}

	pluginDto := dtos.LuzmoPluginDTO{}
	err = json.Unmarshal(body, &pluginDto)
	if err != nil {
		return nil, err
	}

	d, _ := ls.Mapper.MapToPlugin(pluginDto)

	return d, nil
}

func (ls *LuzmoService) UpdatePlugin(plugin models.Plugin) (*models.Plugin, error) {
	payload := dtos.LuzmoUpdateRequest[dtos.LuzmoUpdatePluginDTO]{
		Action:  "update",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      plugin.Id,
		Properties: dtos.LuzmoUpdatePluginDTO{
			Name: map[string]string{
				config.DefaultLang: plugin.Name,
			},
			Description: map[string]string{
				config.DefaultLang: plugin.Description,
			},
			BaseUrl:               plugin.BaseUrl,
			Url:                   plugin.Url,
			ProtocolVersion:       string(plugin.ProtocolVersion),
			Pushdown:              &plugin.Pushdown,
			SupportsLike:          plugin.SupportsLike,
			SupportsDistinctcount: plugin.SupportsDistinctcount,
			SupportsOrderLimit:    plugin.SupportsOrderLimit,
			SupportsJoin:          plugin.SupportsJoin,
			SupportsSql:           plugin.SupportsSql,
			SupportsNestedFilters: plugin.SupportsNestedFilters,
		},
	}

	body, err := ls.doRequest(PluginApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoPluginDTO{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	pluginUpdated, _ := ls.Mapper.MapToPlugin(responseDTO)

	return pluginUpdated, nil
}

func (ls *LuzmoService) FindPluginById(id string) (*models.Plugin, error) {
	payload := dtos.LuzmoFindRequest{
		Action:  "get",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Find:    dtos.Find{Where: dtos.Where{ID: id}},
	}

	body, err := ls.doRequest(PluginApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoFindResponse[dtos.LuzmoPluginDTO]{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	plugin, _ := ls.Mapper.MapToPlugin(responseDTO.Rows[0])

	return plugin, nil
}

func (ls *LuzmoService) DeletePlugin(id string) error {
	payload := dtos.LuzmoDeleteRequest{
		Action:  "delete",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      id,
	}

	_, err := ls.doRequest(PluginApiPath, payload)
	if err != nil {
		return err
	}

	return nil
}
