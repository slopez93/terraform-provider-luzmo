package provider

import (
	"fmt"
	"terraform-provider-luzmo/internal/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccountResourceConfig(accountName string, provider string) string {
	return fmt.Sprintf(
		`%s
		resource "luzmo_account" "account_example" {
			name             = "%s"
			description      = "Account created by terraform provider test"
			provider             = "%s"
		}
		`,
		providerConfig,
		accountName,
		provider,
	)
}

func TestAccAccountResource(t *testing.T) {
	randProvider := utils.RandomString(10)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccountResourceConfig("Account managed by TF", randProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_account.account_example", "name", "Account managed by TF"),
					resource.TestCheckResourceAttr("luzmo_account.account_example", "description", "Account created by terraform provider test"),
					resource.TestCheckResourceAttrSet("luzmo_account.account_example", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccountResourceConfig("Account managed by TF updated", randProvider),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_account.account_example", "name", "Account managed by TF updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
