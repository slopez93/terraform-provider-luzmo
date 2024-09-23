package services

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (ls *LuzmoService) CreateDashboard(dashboard models.Dashboard) (*models.Dashboard, error) {
	payload := dtos.LuzmoCreateRequest[dtos.LuzmoCreateDashboardDTO]{
		Action:  "create",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Properties: dtos.LuzmoCreateDashboardDTO{
			Name: map[string]string{
				config.DefaultLang: dashboard.Name,
			},
			Subtitle: map[string]string{
				config.DefaultLang: dashboard.Subtitle,
			},
			Description: map[string]string{
				config.DefaultLang: dashboard.Description,
			},
			Contents: dashboard.Contents,
			Css:      dashboard.Css,
			Type:     "dashboard",
		},
	}

	body, err := ls.doRequest(DashboardApiPath, payload)
	if err != nil {
		return nil, err
	}

	dashboardDTO := dtos.LuzmoDashboardDTO{}
	err = json.Unmarshal(body, &dashboardDTO)
	if err != nil {
		return nil, err
	}

	d, _ := ls.Mapper.MapToDashboard(dashboardDTO)

	return d, nil
}

func (ls *LuzmoService) UpdateDashboard(dashboard models.Dashboard) (*models.Dashboard, error) {
	payload := dtos.LuzmoUpdateRequest[dtos.LuzmoUpdateDashboardDTO]{
		Action:  "update",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      dashboard.Id,
		Properties: dtos.LuzmoUpdateDashboardDTO{
			Name: map[string]string{
				config.DefaultLang: dashboard.Name,
			},
			Subtitle: map[string]string{
				config.DefaultLang: dashboard.Subtitle,
			},
			Description: map[string]string{
				config.DefaultLang: dashboard.Description,
			},
			Contents: dashboard.Contents,
			Css:      dashboard.Css,
		},
	}

	body, err := ls.doRequest(DashboardApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoDashboardDTO{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	dashboardUpdated, _ := ls.Mapper.MapToDashboard(responseDTO)

	return dashboardUpdated, nil
}

func (ls *LuzmoService) FindDashboardById(id string) (*models.Dashboard, error) {
	payload := dtos.LuzmoFindRequest{
		Action:  "get",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Find:    dtos.Find{Where: dtos.Where{ID: id}},
	}

	body, err := ls.doRequest(DashboardApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoFindResponse[dtos.LuzmoDashboardDTO]{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	dashboard, _ := ls.Mapper.MapToDashboard(responseDTO.Rows[0])

	return dashboard, nil
}

func (ls *LuzmoService) DeleteDashboard(id string) error {
	payload := dtos.LuzmoDeleteRequest{
		Action:  "delete",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      id,
	}

	_, err := ls.doRequest(DashboardApiPath, payload)
	if err != nil {
		return err
	}

	return nil
}
