package services

import (
	"encoding/json"
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"
)

func (ls *LuzmoService) CreateAccount(account models.Account) (*models.Account, error) {
	payload := dtos.LuzmoCreateRequest[dtos.LuzmoCreateAccountDTO]{
		Action:  "create",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Properties: dtos.LuzmoCreateAccountDTO{
			Name: account.Name,
			Description: map[string]string{
				config.DefaultLang: account.Description,
			},
			Provider:                account.Provider,
			Scope:                   account.Scope,
			Host:                    account.Host,
			Active:                  account.Active,
			Invalid:                 account.Invalid,
			Port:                    account.Port,
			Cache:                   account.Cache,
			DatasetMetaSyncEnabled:  account.DatasetMetaSyncEnabled,
			DatasetMetaSyncInterval: account.DatasetMetaSyncInterval,
		},
	}

	body, err := ls.doRequest(AccountApiPath, payload)
	if err != nil {
		return nil, err
	}

	accountDto := dtos.LuzmoAccountDTO{}
	err = json.Unmarshal(body, &accountDto)
	if err != nil {
		return nil, err
	}

	d, _ := ls.Mapper.MapToAccount(accountDto)

	return d, nil
}

func (ls *LuzmoService) FindAccountById(id string) (*models.Account, error) {
	payload := dtos.LuzmoFindRequest{
		Action:  "get",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Find:    dtos.Find{Where: dtos.Where{ID: id}},
	}

	body, err := ls.doRequest(AccountApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoFindResponse[dtos.LuzmoAccountDTO]{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	account, _ := ls.Mapper.MapToAccount(responseDTO.Rows[0])

	return account, nil
}

func (ls *LuzmoService) UpdateAccount(account models.Account) (*models.Account, error) {
	payload := dtos.LuzmoUpdateRequest[dtos.LuzmoUpdateAccountDTO]{
		Action:  "update",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      account.Id,
		Properties: dtos.LuzmoUpdateAccountDTO{
			Name:                    account.Name,
			Description:             map[string]string{config.DefaultLang: account.Description},
			Provider:                account.Provider,
			Scope:                   account.Scope,
			Host:                    account.Host,
			Active:                  account.Active,
			Invalid:                 account.Invalid,
			Port:                    account.Port,
			Cache:                   account.Cache,
			DatasetMetaSyncEnabled:  account.DatasetMetaSyncEnabled,
			DatasetMetaSyncInterval: account.DatasetMetaSyncInterval,
		},
	}

	body, err := ls.doRequest(AccountApiPath, payload)
	if err != nil {
		return nil, err
	}

	responseDTO := dtos.LuzmoAccountDTO{}
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return nil, err
	}

	accountUpdated, _ := ls.Mapper.MapToAccount(responseDTO)

	return accountUpdated, nil
}

func (ls *LuzmoService) DeleteAccount(id string) error {
	payload := dtos.LuzmoDeleteRequest{
		Action:  "delete",
		Version: ls.ApiVersion,
		Key:     ls.ApiKey,
		Token:   ls.ApiToken,
		Id:      id,
	}

	_, err := ls.doRequest(AccountApiPath, payload)
	if err != nil {
		return err
	}

	return nil
}
