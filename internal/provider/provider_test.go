package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the HashiCups client is properly configured.
	// It is also possible to use the HASHICUPS_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = `
		variable "luzmo_api_key" {
			type      = string
			sensitive = true
		}
		variable "luzmo_api_token" {
			type      = string
			sensitive = true
		}
		provider "luzmo" {
			api_key     = var.luzmo_api_key
			api_token   = var.luzmo_api_token
			api_version = "0.1.0"
		}
	`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"luzmo": providerserver.NewProtocol6WithError(New("test")()),
	}
)
