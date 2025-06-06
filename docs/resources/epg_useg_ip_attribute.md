---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_epg_useg_ip_attribute"
sidebar_current: "docs-aci-resource-aci_epg_useg_ip_attribute"
description: |-
  Manages ACI EPG uSeg IP Attribute
---

# aci_epg_useg_ip_attribute #

Manages ACI EPG uSeg IP Attribute



## API Information ##

* Class: [fvIpAttr](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvIpAttr/overview)

* Supported in ACI versions: 1.1(1j) and later.

* Distinguished Name Format: `uni/tn-{name}/ap-{name}/epg-{name}/crtrn/ipattr-{name}`

## GUI Information ##

* Location: `Tenants -> Application Profiles -> uSeg EPGs -> uSeg Attributes`

## Example Usage ##

The configuration snippet below creates a EPG uSeg IP Attribute with only required attributes.

```hcl

resource "aci_epg_useg_ip_attribute" "example_epg_useg_block_statement" {
  parent_dn = aci_epg_useg_block_statement.example.id
  ip        = "131.107.1.200"
  name      = "131"
}

```
The configuration snippet below shows all possible attributes of the EPG uSeg IP Attribute.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_epg_useg_ip_attribute" "full_example_epg_useg_block_statement" {
  parent_dn      = aci_epg_useg_block_statement.example.id
  annotation     = "annotation"
  description    = "description_1"
  ip             = "131.107.1.200"
  name           = "131"
  name_alias     = "name_alias_1"
  owner_key      = "owner_key_1"
  owner_tag      = "owner_tag_1"
  use_epg_subnet = "yes"
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

All examples for the EPG uSeg IP Attribute resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_epg_useg_ip_attribute) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_epg_useg_block_statement](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/epg_useg_block_statement) ([fvCrtrn](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCrtrn/overview))
* `ip` (ip) - (string) The device IP address of the EPG uSeg IP Attribute object.
* `name` (name) - (string) The name of the EPG uSeg IP Attribute object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the EPG uSeg IP Attribute object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the EPG uSeg IP Attribute object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the EPG uSeg IP Attribute object.
* `name_alias` (nameAlias) - (string) The name alias of the EPG uSeg IP Attribute object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `use_epg_subnet` (usefvSubnet) - (string) Parameter used to determine whether a previously configured subnet address should be used as the IP filter. This attribute is supported in ACI versions: 2.0(1m) and later.
  - Default: `no`
  - Valid Values: `no`, `yes`.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing EPG uSeg IP Attribute can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_epg_useg_ip_attribute.example_epg_useg_block_statement uni/tn-{name}/ap-{name}/epg-{name}/crtrn/ipattr-{name}
```

Starting in Terraform version 1.5, an existing EPG uSeg IP Attribute can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/epg-{name}/crtrn/ipattr-{name}"
  to = aci_epg_useg_ip_attribute.example_epg_useg_block_statement
}
```
