---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "marklogic_user Data Source - terraform-provider-marklogic"
subcategory: ""
description: |-
  A user in Marklogic.
---

# marklogic_user (Data Source)

A user in Marklogic.

## Example Usage

```terraform
data "marklogic_user" "test-user" {
  name = "test-user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The user's name.

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **description** (String) A description for the user.
- **roles** (Set of String) The roles assigned to the user.


