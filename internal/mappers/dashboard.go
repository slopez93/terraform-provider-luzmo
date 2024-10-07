package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (m *Mapper) MapToDashboardResource(dashboard models.Dashboard) *dtos.DashboardResourceModel {
	var dashboardResource dtos.DashboardResourceModel

	dashboardResource.ID = types.StringValue(dashboard.Id)
	dashboardResource.Name = types.StringValue(dashboard.Name)
	dashboardResource.Description = types.StringValue(dashboard.Description)
	dashboardResource.Subtitle = types.StringValue(dashboard.Subtitle)
	dashboardResource.Contents = types.StringValue(dashboard.NormalizeContents())
	if dashboard.Css != nil {
		dashboardResource.Css = types.StringValue(*dashboard.Css)
	}

	return &dashboardResource
}
