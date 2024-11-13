package services

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (ls *LuzmoService) CreateDatasets(dataset models.Dataset) ([]*models.Dataset, error) {
	payload := dtos.LuzmoCreateRequest[dtos.LuzmoDataProviderCreateDatasetsDTO]{
		Action:  "create",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Properties: dtos.LuzmoDataProviderCreateDatasetsDTO{
			Action:    "create",
			AccountId: dataset.SourceDataset,
			Provider:  *dataset.ProviderName,
			Datasets:  &[]string{*dataset.DatasetId},
		},
	}

	body, err := ls.doRequest(DataProviderApiPath, payload)
	if err != nil {
		return nil, err
	}

	datasetsDto := dtos.LuzmoDataProviderCreateDatasetsResponseDTO{}
	err = json.Unmarshal(body, &datasetsDto)
	if err != nil {
		return nil, err
	}

	d, _ := ls.Mapper.MapDataProviderResponseToDatasetResource(datasetsDto)

	return d, nil
}
