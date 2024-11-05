package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *Mapper) MapToDataset(datasetDto dtos.LuzmoDatasetDTO) (*models.Dataset, error) {
	var subtitle string

	if datasetDto.Subtitle != nil {
		subtitle = (*datasetDto.Subtitle)[config.DefaultLang]
	}

	dataset := models.Dataset{
		Id:                 datasetDto.Id,
		Name:               datasetDto.Name[config.DefaultLang],
		Description:        datasetDto.Description[config.DefaultLang],
		Subtitle:           &subtitle,
		Subtype:            datasetDto.Subtype,
		SourceDataset:      datasetDto.SourceDataset,
		SourceSheet:        datasetDto.SourceSheet,
		Transformation:     datasetDto.Transformation,
		Cache:              datasetDto.Cache,
		UpdateMetadata:     datasetDto.UpdateMetadata,
		MetaSyncInterval:   datasetDto.MetaSyncInterval,
		MetaSyncInherit:    datasetDto.MetaSyncInherit,
		MetaSyncEnabled:    datasetDto.MetaSyncEnabled,
		LastMetadataSyncAt: datasetDto.LastMetadataSyncAt,
	}

	return &dataset, nil
}

func (m *Mapper) MapToDatasetResource(dataset models.Dataset) *dtos.DatasetResourceModel {
	var datasetModel dtos.DatasetResourceModel

	datasetModel.ID = types.StringValue(dataset.Id)
	datasetModel.Name = types.StringValue(dataset.Name)
	datasetModel.Description = types.StringValue(dataset.Description)
	datasetModel.Subtitle = types.StringPointerValue(dataset.Subtitle)
	datasetModel.SubType = types.StringValue(dataset.Subtype)
	datasetModel.SourceDataset = types.StringValue(dataset.SourceDataset)
	datasetModel.SourceSheet = types.StringValue(dataset.SourceSheet)
	datasetModel.Transformation = types.StringPointerValue(dataset.Transformation)
	datasetModel.Cache = types.Int64Value(dataset.Cache)
	datasetModel.UpdateMetadata = types.BoolValue(dataset.UpdateMetadata)
	datasetModel.MetaSyncInterval = types.Int32PointerValue(&dataset.MetaSyncInterval)
	datasetModel.MetaSyncInherit = types.BoolValue(dataset.MetaSyncInherit)
	datasetModel.MetaSyncEnabled = types.BoolPointerValue(&dataset.MetaSyncEnabled)
	datasetModel.LastMetadataSyncAt = types.StringPointerValue(dataset.LastMetadataSyncAt)

	return &datasetModel
}
