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

resource "luzmo_account" "account_example" {
  name          = "Account managed by TF"
  description   = "Account created by terraform provider"
  provider_name = "wow"
}
