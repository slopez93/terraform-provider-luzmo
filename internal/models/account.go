package models

type Account struct {
	Id                      string
	Name                    string
	Description             string
	Provider                string
	Scope                   *string
	Host                    *string
	Active                  bool
	Invalid                 bool
	Port                    *int32
	Cache                   int64
	DatasetMetaSyncEnabled  bool
	DatasetMetaSyncInterval *int32
}

type NewAccountParams struct {
	Id                      string
	Name                    string
	Description             string
	Provider                string
	Scope                   *string
	Host                    *string
	Active                  bool
	Invalid                 bool
	Port                    *int32
	Cache                   int64
	DatasetMetaSyncEnabled  bool
	DatasetMetaSyncInterval *int32
}

func NewAccount(params NewAccountParams) *Account {
	account := Account{
		Id:                      params.Id,
		Name:                    params.Name,
		Description:             params.Description,
		Provider:                params.Provider,
		Active:                  params.Active,
		Invalid:                 params.Invalid,
		Cache:                   params.Cache,
		DatasetMetaSyncEnabled:  params.DatasetMetaSyncEnabled,
		DatasetMetaSyncInterval: params.DatasetMetaSyncInterval,
	}

	return &account
}
