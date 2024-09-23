package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"terraform-provider-luzmo/internal/mappers"
)

const HostURL string = "https://api.luzmo.com"
const DefaultApiVersion string = "0.1.0"
const DashboardApiPath string = "securable"
const PluginApiPath string = "plugin"

type LuzmoService struct {
	ApiKey     string
	ApiToken   string
	ApiVersion string
	HttpClient *http.Client

	Mapper mappers.Mapper
}

func NewLuzmoService(apiKey string, apiToken string, apiVersion string) (*LuzmoService, error) {
	if apiVersion == "" {
		apiVersion = DefaultApiVersion
	}

	return &LuzmoService{
		ApiKey:     apiKey,
		ApiToken:   apiToken,
		ApiVersion: DefaultApiVersion,
		HttpClient: &http.Client{},

		Mapper: mappers.Mapper{},
	}, nil
}

func (ls *LuzmoService) doRequest(path string, payload interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", HostURL, ls.ApiVersion, path), strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}
