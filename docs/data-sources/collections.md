---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zillizcloud_collections Data Source - zillizcloud"
subcategory: ""
description: |-
  List collections of a given database by connect_address and db_name
---

# zillizcloud_collections (Data Source)

List collections of a given database by connect_address and db_name



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connect_address` (String) Cluster's connection address
- `db_name` (String) Database name

### Read-Only

- `items` (Attributes List) List of collections (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `collection_name` (String) Collection name
