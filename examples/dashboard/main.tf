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

resource "luzmo_dashboard" "dashboard_example" {
  name        = "Dashboard managed by TF"
  description = "Dashboard created by terraform provider"
  subtitle    = "Terraform Dashboards"
  contents    = jsonencode(local.dashboard_content)
  css         = ".pivottable-value { font-size: 18px; }"
}

locals {
  dashboard_content = jsondecode(file("${path.module}/dashboard_contents.json"))
}