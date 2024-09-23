package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (m *Mapper) MapToDashboard(dashboardDto dtos.LuzmoDashboardDTO) (*models.Dashboard, error) {
	dashboard := models.Dashboard{
		Id:          dashboardDto.Id,
		Name:        dashboardDto.Name[config.DefaultLang],
		Subtitle:    dashboardDto.Subtitle[config.DefaultLang],
		Description: dashboardDto.Description[config.DefaultLang],
		Contents:    dashboardDto.Contents,
		Css:         dashboardDto.Css,
	}

	return &dashboard, nil
}
