package services

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (ls *LuzmoService) CreateDataset(dataset models.Dataset) (*models.Dataset, error) {
	payload := dtos.LuzmoCreateRequest[dtos.LuzmoCreateDatasetDTO]{
		Action:  "create",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Properties: dtos.LuzmoCreateDatasetDTO{
			Name: map[string]string{
				config.DefaultLang: dataset.Name,
			},
			Description: map[string]string{
				config.DefaultLang: dataset.Description,
			},
			SubTitle: map[string]*string{
				config.DefaultLang: dataset.Subtitle,
			},
			Type:               "dataset",
			SubType:            dataset.Subtype,
			SourceDataset:      dataset.SourceDataset,
			SourceSheet:        dataset.SourceSheet,
			Transformation:     dataset.Transformation,
			Cache:              dataset.Cache,
			UpdateMetadata:     dataset.UpdateMetadata,
			MetaSyncInterval:   dataset.MetaSyncInterval,
			MetaSyncInherit:    dataset.MetaSyncInherit,
			MetaSyncEnabled:    dataset.MetaSyncEnabled,
			LastMetadataSyncAt: dataset.LastMetadataSyncAt,
		},
	}

	body, err := ls.doRequest(DatasetApiPath, payload)
	if err != nil {
		return nil, err
	}

	datasetDto := dtos.LuzmoDatasetDTO{}
	err = json.Unmarshal(body, &datasetDto)
	if err != nil {
		return nil, err
	}

	d, _ := ls.Mapper.MapToDataset(datasetDto)

	return d, nil
}

func (ls *LuzmoService) FindDatasetById(id string) (*models.Dataset, error) {
	payload := dtos.LuzmoFindRequest{
		Action:  "get",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Find:    dtos.Find{Where: dtos.Where{ID: id}},
	}

	body, err := ls.doRequest(DatasetApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoFindResponse[dtos.LuzmoDatasetDTO]{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	dataset, _ := ls.Mapper.MapToDataset(responseDTO.Rows[0])

	return dataset, nil
}

func (ls *LuzmoService) UpdateDataset(dataset models.Dataset) (*models.Dataset, error) {
	payload := dtos.LuzmoUpdateRequest[dtos.LuzmoUpdateDatasetDTO]{
		Action:  "update",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      dataset.Id,
		Properties: dtos.LuzmoUpdateDatasetDTO{
			Name: map[string]string{
				config.DefaultLang: dataset.Name,
			},
			Description: map[string]string{
				config.DefaultLang: dataset.Description,
			},
			SubTitle: map[string]string{
				config.DefaultLang: *dataset.Subtitle,
			},
			SubType:          dataset.Subtype,
			Cache:            dataset.Cache,
			UpdateMetadata:   dataset.UpdateMetadata,
			MetaSyncInterval: dataset.MetaSyncInterval,
			MetaSyncInherit:  dataset.MetaSyncInherit,
			MetaSyncEnabled:  dataset.MetaSyncEnabled,
		},
	}

	body, err := ls.doRequest(DatasetApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoDatasetDTO{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	datasetUpdated, _ := ls.Mapper.MapToDataset(responseDTO)

	return datasetUpdated, nil
}

func (ls *LuzmoService) DeleteDataset(id string) error {
	payload := dtos.LuzmoDeleteRequest{
		Action:  "delete",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      id,
	}

	_, err := ls.doRequest(DatasetApiPath, payload)
	if err != nil {
		return err
	}

	return nil
}
