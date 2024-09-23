package models

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/utils"
)

type NewDashboardParams struct {
	Id          string
	Name        string
	Description string
	Subtitle    string
	Contents    string
	Css         *string
}

type Dashboard struct {
	Id          string
	Name        string
	Description string
	Subtitle    string
	Contents    map[string]interface{}
	Css         *string
}

func NewDashboard(params NewDashboardParams) (*Dashboard, error) {
	var dashboardContent map[string]interface{}
	err := json.Unmarshal([]byte(params.Contents), &dashboardContent)
	if err != nil {
		return nil, err
	}

	dashboard := Dashboard{
		Id:          params.Id,
		Name:        params.Name,
		Description: params.Description,
		Subtitle:    params.Subtitle,
		Contents:    dashboardContent,
	}

	if params.Css != nil {
		dashboard.Css = params.Css
	}

	return &dashboard, nil
}

func (d *Dashboard) NormalizeContents() string {
	normalized, err := utils.NormalizeMap(d.Contents)
	if err != nil {
		return ""
	}

	return normalized
}
