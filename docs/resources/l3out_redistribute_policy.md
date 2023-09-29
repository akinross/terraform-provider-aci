---
subcategory: "L3Out"
layout: "aci"
page_title: "ACI: l3out_redistribute_policy"
sidebar_current: "docs-aci-resource-l3out_redistribute_policy"
description: |-
  Manages ACI L3out Redistribute Policy
---

# aci_l3out_redistribute_policy #

Manages ACI L3out Redistribute Policy

## API Information ##

* `Class` - `l3extRsRedistributePol`

* `Distinguished Name Formats`
  - `uni/tn-{name}/out-{name}/rsredistributePol-[{tnRtctrlProfileName}]-{src}`

## GUI Information ##

* `Location` - `Tenants -> Networking -> L3Outs -> Redistribute Policies`

## Example Usage ##

```hcl

resource "aci_l3out_redistribute_policy" "example" {
  parent_dn                  = aci_l3_outside.example.id
  src                        = "direct"
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  annotations = [
    {
      key   = "annotations_1"
      value = "value_1"
    }
  ]
}

```

## Schema

### Required

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_l3_outside](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/l3_outside) (`l3extOut`)
* `src` - (string) The source IP address.
  - Valid Values: `attached-host`, `direct`, `static`.
* `route_control_profile_name` - (string) The name of the route profile associated with this object.

### Read-Only

* `id` - (string) The distinguished name (DN) of the L3out Redistribute Policy object.

### Optional
  
* `annotation` - (string) The annotation of the L3out Redistribute Policy object.
  - Default: `orchestrator:terraform`

* `annotations` - (list) A list of Annotations objects `tagAnnotation` which can be configured using the [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource.
  
  #### Required
  
  * `key` - (string) The key or password used to uniquely identify this configuration object.
  * `value` - (string) The value of the property.

## Importing ##

An existing L3out Redistribute Policy can be [imported](https://www.terraform.io/docs/import/index.html) into this resource via its distinguished name (DN), via the following command:

```
terraform import aci_l3out_redistribute_policy.example uni/tn-{name}/out-{name}/rsredistributePol-[{tnRtctrlProfileName}]-{src}
```

Starting in Terraform version 1.5, an existing L3out Redistribute Policy can be imported 
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/out-{name}/rsredistributePol-[{tnRtctrlProfileName}]-{src}"
  to = aci_l3out_redistribute_policy.example
}
```