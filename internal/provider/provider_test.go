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
		provider "luzmo" {
			api_key     = "bc088ce2-5440-41ba-bf85-c3fc2e0ff6a3"
			api_token   = "DVkbaZbf0eLHi0GvMEIdMAuMfvzm8Q1UawSO4WkBUUkYAvdgXDERyW6ShFDN2Xu1V4skucIe5vLocnIaTYfDnIeXCiJFRiowpRbOPPSdjy2Cx4QMAOZjrgmhx8N4f51KU9Y8LETPJNpMyAk7EECfY6"
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
