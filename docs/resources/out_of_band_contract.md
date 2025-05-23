---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Contract"
layout: "aci"
page_title: "ACI: aci_out_of_band_contract"
sidebar_current: "docs-aci-resource-aci_out_of_band_contract"
description: |-
  Manages ACI Out Of Band Contract
---

# aci_out_of_band_contract #

Manages ACI Out Of Band Contract



## API Information ##

* Class: [vzOOBBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzOOBBrCP/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-mgmt/oobbrc-{name}`

## GUI Information ##

* Location: `Tenants (mgmt) -> Contracts -> Out-Of-Band Contracts`

## Example Usage ##

The configuration snippet below creates a Out Of Band Contract with only required attributes.

```hcl

resource "aci_out_of_band_contract" "example" {
  name = "test_name"
}

```
The configuration snippet below shows all possible attributes of the Out Of Band Contract.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_out_of_band_contract" "full_example" {
  annotation  = "annotation"
  description = "description_1"
  intent      = "estimate_add"
  name        = "test_name"
  name_alias  = "name_alias_1"
  owner_key   = "owner_key_1"
  owner_tag   = "owner_tag_1"
  priority    = "level1"
  scope       = "application-profile"
  target_dscp = "AF11"
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}

```

All examples for the Out Of Band Contract resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_out_of_band_contract) folder.

## Schema ##

### Required ###

* `name` (name) - (string) The name of the Out Of Band Contract object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Out Of Band Contract object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the Out Of Band Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the Out Of Band Contract object.
* `intent` (intent) - (string) The Install Rules or Estimate Number of Rules. This attribute is supported in ACI versions: 4.2(1i) and later.
  - Default: `install`
  - Valid Values: `estimate_add`, `estimate_delete`, `install`.
* `name_alias` (nameAlias) - (string) The name alias of the Out Of Band Contract object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
  - Default: `unspecified`
  - Valid Values:
    * `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.
    * Or a value in the range of `0` to `9`.
* `scope` (scope) - (string) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.
  - Default: `context`
  - Valid Values: `application-profile`, `context`, `global`, `tenant`.
* `target_dscp` (targetDscp) - (string) The target DSCP value of the Out Of Band Contract object. This attribute is supported in ACI versions: 1.2(2g) and later.
  - Default: `unspecified`
  - Valid Values:
    * `AF11`, `AF12`, `AF13`, `AF21`, `AF22`, `AF23`, `AF31`, `AF32`, `AF33`, `AF41`, `AF42`, `AF43`, `CS0`, `CS1`, `CS2`, `CS3`, `CS4`, `CS5`, `CS6`, `CS7`, `EF`, `VA`, `unspecified`.
    * Or a value in the range of `0` to `64`.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Out Of Band Contract can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_out_of_band_contract.example uni/tn-mgmt/oobbrc-{name}
```

Starting in Terraform version 1.5, an existing Out Of Band Contract can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-mgmt/oobbrc-{name}"
  to = aci_out_of_band_contract.example
}
```
