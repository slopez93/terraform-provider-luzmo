# Terraform Provider: Luzmo (Terraform Plugin SDK)

_This template repository is built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk). The template repository built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) can be found at [terraform-provider-scaffolding-framework](https://github.com/hashicorp/terraform-provider-scaffolding-framework). See [Which SDK Should I Use?](https://www.terraform.io/docs/plugin/which-sdk.html) in the Terraform documentation for additional information._

----

This is a terraform provider plugin for managing Luzmo plugins and dashboards in a simple way.

_Note_: There are more luzmo resource, for this initial version this plugin just handle plugin and dashboards.


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```bash
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider


Definining provider. The port should be the Clickhouse native protocol port (9000 by default, and 9440 for Clickhouse Cloud)

```hcl
provider "luzmo" {
  api_key           = ******
  api_token         = ******
  api_version       = "0.1.0" # Optional: By default is 0.1.0
}
```

In order to definte url, username and password in a safety way it is possible to define them using env vars:

```config
TF_luzmo_api_key    = ""
TF_luzmo_api_token  = ""
```

```hcl
resource "luzmo_plugin" "plugin_example" {
  name             = "Plugin managed by TF"
  description      = "Plugin created by terraform provider"
  slug             = "wow"
  base_url         = "https://example.com/luzmo-plugin-test"
  url              = "https://example.com/luzmo-plugin-test/documentation"
  protocol_version = "3.0.0"
  pushdown         = true

  supports_like          = true
  supports_distinctcount = false
  supports_order_limit   = false
  supports_join          = false
  supports_sql           = false
}
```

### Clustered server

Configuring provider

```hcl
provider "luzmo" {
  api_key     = var.luzmo_api_key
  api_token   = var.luzmo_api_token
  api_version = "0.1.0"
}
```

Creating a Plugin

```hcl
resource "luzmo_plugin" "plugin_example" {
  name             = "Plugin managed by TF"
  description      = "Plugin created by terraform provider"
  slug             = "wow"
  base_url         = "https://example.com/luzmo-plugin-test"
  url              = "https://example.com/luzmo-plugin-test/documentation"
  protocol_version = "3.0.0"
  pushdown         = true

  supports_like          = true
  supports_distinctcount = false
  supports_order_limit   = false
  supports_join          = false
  supports_sql           = false
}
```

Creating Dashboard

```hcl
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
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```