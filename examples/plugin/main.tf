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

resource "luzmo_plugin" "plugin_example" {
  name             = "Plugin managed by TF"
  description      = "Plugin created by terraform provider test"
  slug             = "wow"
  base_url         = "https://lansweeper.com/luzmo-plugin-test"
  url              = "https://lansweeper.com/luzmo-plugin-test/documentation"
  protocol_version = "3.0.0"
  pushdown         = true

  supports_like           = true
  supports_distinctcount  = true
  supports_order_limit    = true
  supports_join           = true
  supports_sql            = true
  supports_nested_filters = true
}