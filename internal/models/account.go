package models

type Account struct {
	Id                       string
	Name                     string
	Description              string
	ProviderName             string
	Scope                    *string
	Host                     *string
	Active                   *bool
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
	Active                   *bool
	Port                     *string
	Cache                    int64
	DatasetsMetaSyncEnabled  bool
	DatasetsMetaSyncInterval *int32
}

func NewAccount(params NewAccountParams) *Account {
	account := Account(params)

	return &account
}
