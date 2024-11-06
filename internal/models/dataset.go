package models

type Dataset struct {
	Id                 string
	Name               string
	Description        string
	Subtitle           *string
	Subtype            string
	SourceDataset      string
	SourceSheet        string
	Transformation     *string
	Cache              int64
	UpdateMetadata     bool
	MetaSyncInterval   int32
	MetaSyncInherit    bool
	MetaSyncEnabled    bool
	LastMetadataSyncAt *string
}

type NewDatasetParams struct {
	Id                 string
	Name               string
	Description        string
	SubTitle           *string
	SubType            string
	SourceDataset      string
	SourceSheet        string
	Transformation     *string
	Cache              int64
	UpdateMetadata     bool
	MetaSyncInterval   int32
	MetaSyncInherit    bool
	MetaSyncEnabled    *bool
	LastMetadataSyncAt *string
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
	}

	if params.SubTitle != nil {
		dataset.Subtitle = params.SubTitle
	}

	if params.Transformation != nil {
		dataset.Transformation = params.Transformation
	}

	if params.MetaSyncEnabled != nil {
		dataset.MetaSyncEnabled = *params.MetaSyncEnabled
	}

	if params.LastMetadataSyncAt != nil {
		dataset.LastMetadataSyncAt = params.LastMetadataSyncAt
	}

	return &dataset
}
