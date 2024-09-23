package dtos

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
