package dtos

import "github.com/hashicorp/terraform-plugin-framework/types"

type DashboardResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Subtitle    types.String `tfsdk:"subtitle"`
	Contents    types.String `tfsdk:"contents"`
	Css         types.String `tfsdk:"css"`
}

type LuzmoDashboardDTO struct {
	Id          string                 `json:"id"`
	Name        map[string]string      `json:"name"`
	Description map[string]string      `json:"description"`
	Subtitle    map[string]string      `json:"subtitle"`
	Contents    map[string]interface{} `json:"contents"`
	Css         *string                `json:"css"`
	Type        string                 `json:"type"`
}

type LuzmoCreateDashboardDTO struct {
	Name        map[string]string      `json:"name"`
	Description map[string]string      `json:"description"`
	Subtitle    map[string]string      `json:"subtitle"`
	Contents    map[string]interface{} `json:"contents"`
	Css         *string                `json:"css"`
	Type        string                 `json:"type"`
}

type LuzmoUpdateDashboardDTO struct {
	Name        map[string]string      `json:"name"`
	Description map[string]string      `json:"description"`
	Subtitle    map[string]string      `json:"subtitle"`
	Contents    map[string]interface{} `json:"contents"`
	Css         *string                `json:"css"`
}
