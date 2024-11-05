---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "luzmo_dataset Resource - terraform-provider-luzmo"
subcategory: ""
description: |-
  Manages a dataset.
---

# luzmo_dataset (Resource)

Manages a dataset.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the dataset.
- `name` (String) The name of the dataset.
- `source_dataset` (String) The source dataset of the dataset.
- `source_sheet` (String) The source sheet of the dataset.

### Optional

- `cache` (Number) Number of seconds queries to this data connector are cached in Luzmo's caching layer. Use 0 to disable caching.
- `last_metadata_sync_at` (String) Last time metadata was synced was successful for the dataset.
- `meta_sync_enabled` (Boolean) Indicates whether automatic metadata sync is enabled for the dataset when meta_sync_inherit=false.
- `meta_sync_inherit` (Boolean) Indicates whether automatic metadata sync is enabled for this dataset.
- `meta_sync_interval` (Number) Configure Metadata sync interval in hours for the dataset when meta_sync_inherit=false.
- `subtitle` (String) The subtitle of the dataset.
- `subtype` (String) The subtype of the dataset.
- `transformation` (String) The transformation of the dataset.
- `update_metadata` (Boolean) Virtual property is used to trigger manual update for dataset metadata.

### Read-Only

- `id` (String) String identifier of the order.