terraform {
  required_providers {
    luzmo = {
      source = "hashicorp.com/slopez93/luzmo"
    }
  }
}

provider "luzmo" {
  api_key     = var.luzmo_api_key
  api_token   = var.luzmo_api_token
  api_version = "0.1.0"
}

resource "luzmo_dataset" "dataset_example" {
  name           = "Dataset managed by TF"
  description    = "Dataset created by terraform provider"
  source_dataset = "8b7954e2-3664-4437-b8cb-4b05b6ddee91"
  source_sheet   = "example_dataset"
}
