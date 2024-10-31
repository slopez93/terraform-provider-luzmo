package mappers

import (
	"terraform-provider-luzmo/internal/config"
	"terraform-provider-luzmo/internal/dtos"
	"terraform-provider-luzmo/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *Mapper) MapToAccount(accountDto dtos.LuzmoAccountDTO) (*models.Account, error) {
	account := models.Account{
		Id:                      accountDto.Id,
		Name:                    accountDto.Name,
		Description:             accountDto.Description[config.DefaultLang],
		Provider:                accountDto.Provider,
		Scope:                   accountDto.Scope,
		Host:                    accountDto.Host,
		Active:                  accountDto.Active,
		Invalid:                 accountDto.Invalid,
		Port:                    accountDto.Port,
		Cache:                   accountDto.Cache,
		DatasetMetaSyncEnabled:  accountDto.DatasetMetaSyncEnabled,
		DatasetMetaSyncInterval: accountDto.DatasetMetaSyncInterval,
	}

	return &account, nil
}

func (m *Mapper) MapToAccountResource(account models.Account) *dtos.AccountResourceModel {
	var accountResource dtos.AccountResourceModel

	accountResource.ID = types.StringValue(account.Id)
	accountResource.Name = types.StringValue(account.Name)
	accountResource.Description = types.StringValue(account.Description)
	accountResource.Provider = types.StringValue(account.Provider)
	accountResource.Active = types.BoolValue(account.Active)
	accountResource.Invalid = types.BoolValue(account.Invalid)
	accountResource.Cache = types.Int64Value(account.Cache)
	accountResource.DatasetMetaSyncEnabled = types.BoolValue(account.DatasetMetaSyncEnabled)

	if account.Scope != nil {
		accountResource.Scope = types.StringValue(*account.Scope)
	}

	if account.Host != nil {
		accountResource.Host = types.StringValue(*account.Host)
	}

	if account.Port != nil {
		accountResource.Port = types.Int32Value(*account.Port)
	}

	if account.DatasetMetaSyncInterval != nil {
		accountResource.DatasetMetaSyncInterval = types.Int32Value(*account.DatasetMetaSyncInterval)
	}

	return &accountResource
}
