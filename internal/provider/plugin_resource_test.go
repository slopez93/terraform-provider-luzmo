package provider

import (
	"fmt"
	"terraform-provider-luzmo/internal/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testPluginResourceConfig(pluginName string, slug string) string {
	return fmt.Sprintf(
		`%s
		resource "luzmo_plugin" "plugin_example" {
			name             = "%s"
			description      = "Plugin created by terraform provider test"
			slug             = "%s"
			base_url         = "https://lansweeper.com/luzmo-plugin-test"
			url              = "https://lansweeper.com/luzmo-plugin-test/documentation"
			protocol_version = "3.0.0"
			pushdown         = true

			supports_like          = true
			supports_distinctcount = false
			supports_order_limit   = false
			supports_join          = false
			supports_sql           = false
		}
		`,
		providerConfig,
		pluginName,
		slug,
	)
}

func TestAccPluginResource(t *testing.T) {
	randSlug := utils.RandomString(10)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testPluginResourceConfig("Plugin managed by TF", randSlug),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_plugin.plugin_example", "name", "Plugin managed by TF"),
					resource.TestCheckResourceAttr("luzmo_plugin.plugin_example", "description", "Plugin created by terraform provider test"),
					resource.TestCheckResourceAttrSet("luzmo_plugin.plugin_example", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testPluginResourceConfig("Plugin managed by TF updated", randSlug),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_plugin.plugin_example", "name", "Plugin managed by TF updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
