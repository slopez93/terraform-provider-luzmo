package models

type Account struct {
	Id                       string
	Name                     string
	Description              string
	ProviderName             string
	Scope                    *string
	Host                     *string
	Active                   bool
	Port                     *string
	Cache                    int64
	DatasetsMetaSyncEnabled  bool
	DatasetsMetaSyncInterval *int32
}

type NewAccountParams struct {
	Id                       string
	Name                     string
	Description              string
	ProviderName             string
	Scope                    *string
	Host                     *string
	Active                   bool
	Port                     *string
	Cache                    int64
	DatasetsMetaSyncEnabled  bool
	DatasetsMetaSyncInterval *int32
}

func NewAccount(params NewAccountParams) *Account {
	account := Account{
		Id:                      params.Id,
		Name:                    params.Name,
		Description:             params.Description,
		ProviderName:            params.ProviderName,
		Active:                  params.Active,
		Cache:                   params.Cache,
		DatasetsMetaSyncEnabled: params.DatasetsMetaSyncEnabled,
	}

	if params.Scope != nil {
		account.Scope = params.Scope
	}

	if params.Host != nil {
		account.Host = params.Host
	}

	if params.Port != nil {
		account.Port = params.Port
	}

	if params.DatasetsMetaSyncInterval != nil {
		account.DatasetsMetaSyncInterval = params.DatasetsMetaSyncInterval
	}

	return &account
}
