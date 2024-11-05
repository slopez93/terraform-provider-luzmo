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
		variable "LUZMO_API_KEY" {
			type      = string
			sensitive = true
		}
		variable "LUZMO_API_TOKEN" {
			type      = string
			sensitive = true
		}
		provider "luzmo" {
			api_key     = var.LUZMO_API_KEY
			api_token   = var.LUZMO_API_TOKEN
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
