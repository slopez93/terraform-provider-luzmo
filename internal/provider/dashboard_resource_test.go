package provider

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testDashboardResourceConfig(dashboardName string) string {
	contentAbsPath, err := filepath.Abs("../../examples/dashboard/dashboard_contents.json")
	if err != nil {
		log.Fatalf("File dashboard content not exist: %v", err)
	}

	return fmt.Sprintf(
		`%s
		locals {
			dashboard_content = jsondecode(file("%s"))
		}				
		resource "luzmo_dashboard" "test" {
			name        = "%s"
			description = "Dashboard created by terraform provider"
			subtitle    = "Terraform Dashboards"
			contents    = jsonencode(local.dashboard_content)
			css         = ".pivottable-value { font-size: 18px; }"
		}
		`,
		providerConfig,
		contentAbsPath,
		dashboardName,
	)
}

func TestAccDashboardResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testDashboardResourceConfig("Dashboard managed by TF"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_dashboard.test", "name", "Dashboard managed by TF"),
					resource.TestCheckResourceAttr("luzmo_dashboard.test", "description", "Dashboard created by terraform provider"),
					resource.TestCheckResourceAttrSet("luzmo_dashboard.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testDashboardResourceConfig("Dashboard managed by TF updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_dashboard.test", "name", "Dashboard managed by TF updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
