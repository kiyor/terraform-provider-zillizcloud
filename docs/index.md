---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zillizcloud Provider"
subcategory: ""
description: |-
  
---

# zillizcloud Provider



## Example Usage

```terraform
terraform {
  required_providers {
    zillizcloud = {
      source = "zilliztech/zillizcloud"
    }
  }
}

provider "zillizcloud" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive) Zilliz Cloud API Key
- `host_address` (String) Zilliz Cloud Host Address
- `region_id` (String) Zilliz Cloud Region Id
