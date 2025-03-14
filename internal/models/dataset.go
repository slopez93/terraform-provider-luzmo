package models

import "errors"

type Dataset struct {
	Id               string
	Name             string
	Description      string
	Subtitle         *string
	Subtype          string
	SourceDataset    string
	SourceSheet      string
	Transformation   *string
	Cache            int64
	UpdateMetadata   bool
	MetaSyncInterval int32
	MetaSyncInherit  bool
	MetaSyncEnabled  bool
	DatasetId        *string
	ProviderName     *string
}

type NewDatasetParams struct {
	Id               string
	Name             string
	Description      string
	SubTitle         *string
	SubType          string
	SourceDataset    string
	SourceSheet      string
	Transformation   *string
	Cache            int64
	UpdateMetadata   bool
	MetaSyncInterval int32
	MetaSyncInherit  bool
	MetaSyncEnabled  *bool
	DatasetId        *string
	ProviderName     *string
}

func NewDataset(params NewDatasetParams) *Dataset {
	dataset := Dataset{
		Id:               params.Id,
		Name:             params.Name,
		Description:      params.Description,
		Subtype:          params.SubType,
		SourceDataset:    params.SourceDataset,
		SourceSheet:      params.SourceSheet,
		Cache:            params.Cache,
		UpdateMetadata:   params.UpdateMetadata,
		MetaSyncInherit:  params.MetaSyncInherit,
		MetaSyncInterval: params.MetaSyncInterval,
		Subtitle:         params.SubTitle,
		Transformation:   params.Transformation,
		MetaSyncEnabled:  *params.MetaSyncEnabled,
		DatasetId:        params.DatasetId,
		ProviderName:     params.ProviderName,
	}

	return &dataset
}

func (d *Dataset) MustBeCreatedByDataProvider() (bool, error) {
	hasDataset := d.DatasetId != nil

	if hasDataset && d.ProviderName == nil {
		return false, errors.New("ProviderName must be set if Dataset was created by a data provider")
	}

	return hasDataset, nil
}

func (d *Dataset) SetDatasetId(datasetId string) {
	d.DatasetId = &datasetId
}
