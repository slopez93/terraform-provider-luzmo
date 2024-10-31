package provider

import (
	"fmt"
	"terraform-provider-luzmo/internal/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testDatasetResourceConfig(datasetName string, sourceDataset string, sourceSheet string) string {
	return fmt.Sprintf(
		`%s
		resource "luzmo_dataset" "dataset_example" {
			name            = "%s"
			description     = "Dataset created by terraform provider test"
			source_dataset 	= "%s"
			source_sheet 	= "%s"
		}
		`,
		providerConfig,
		datasetName,
		sourceDataset,
		sourceSheet,
	)
}

func TestAccDatasetResource(t *testing.T) {
	randSourceDataset := utils.RandomString(10)
	randSourceSheet := utils.RandomString(10)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testDatasetResourceConfig("Dataset managed by TF", randSourceDataset, randSourceSheet),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_dataset.dataset_example", "name", "Dataset managed by TF"),
					resource.TestCheckResourceAttr("luzmo_dataset.dataset_example", "description", "Dataset created by terraform provider test"),
					resource.TestCheckResourceAttr("luzmo_dataset.dataset_example", "source_dataset", randSourceDataset),
					resource.TestCheckResourceAttr("luzmo_dataset.dataset_example", "source_sheet", "source_sheet"),
					resource.TestCheckResourceAttrSet("luzmo_dataset.dataset_example", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testDatasetResourceConfig("Dataset managed by TF updated", randSourceDataset, randSourceSheet),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("luzmo_dataset.dataset_example", "name", "Dataset managed by TF updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
