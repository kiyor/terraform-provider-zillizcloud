---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zillizcloud_databases Data Source - zillizcloud"
subcategory: ""
description: |-
  List databases of a given cluster by connect_address
---

# zillizcloud_databases (Data Source)

List databases of a given cluster by connect_address



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connect_address` (String) Cluster's connection address

### Read-Only

- `items` (Attributes List) List of databases (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `db_name` (String) Database name
