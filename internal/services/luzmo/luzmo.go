package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"terraform-provider-luzmo/internal/mappers"
)

const DefaultApiUrl string = "https://api.luzmo.com"
const DefaultApiVersion string = "0.1.0"
const DashboardApiPath string = "securable"
const PluginApiPath string = "plugin"
const AccountApiPath string = "account"
const DatasetApiPath string = "securable"
const DataProviderApiPath string = "dataprovider"

type LuzmoService struct {
	ApiKey     string
	ApiToken   string
	ApiUrl     string
	ApiVersion string
	HttpClient *http.Client

	Mapper mappers.Mapper
}

func NewLuzmoService(apiKey string, apiToken string, apiVersion string, apiUrl string) (*LuzmoService, error) {
	if apiVersion == "" {
		apiVersion = DefaultApiVersion
	}

	if apiUrl == "" {
		apiUrl = DefaultApiUrl
	}

	return &LuzmoService{
		ApiKey:     apiKey,
		ApiToken:   apiToken,
		ApiUrl:     apiUrl,
		ApiVersion: apiVersion,
		HttpClient: &http.Client{},

		Mapper: mappers.Mapper{},
	}, nil
}

func (ls *LuzmoService) doRequest(path string, payload interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", ls.ApiUrl, ls.ApiVersion, path), strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := ls.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}
