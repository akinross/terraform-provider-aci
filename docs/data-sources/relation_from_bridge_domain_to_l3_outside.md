---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Networking"
layout: "aci"
page_title: "ACI: aci_relation_from_bridge_domain_to_l3_outside"
sidebar_current: "docs-aci-data-source-aci_relation_from_bridge_domain_to_l3_outside"
description: |-
  Data source for ACI Relation From Bridge Domain To L3 Outside
---

# aci_relation_from_bridge_domain_to_l3_outside #

Data source for ACI Relation From Bridge Domain To L3 Outside

## API Information ##

* Class: [fvRsBDToOut](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsBDToOut/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Formats:
  - `uni/tn-{name}/BD-{name}/rsBDToOut-{tnL3extOutName}`
  - `uni/tn-{name}/rsBDToOut-{tnL3extOutName}`
  - `uni/tn-{name}/svcBD-{name}/rsBDToOut-{tnL3extOutName}`

## GUI Information ##

* Location: `Tenants -> Networking -> Bridge Domains -> Policy -> L3 Configurations -> Associated L3 Outs`

## Example Usage ##

```hcl

data "aci_relation_from_bridge_domain_to_l3_outside" "example_bridge_domain" {
  parent_dn       = aci_bridge_domain.example.id
  l3_outside_name = aci_l3_outside.example.name
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_bridge_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/bridge_domain) ([fvBD](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvBD/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [fvSvcBD](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvSvcBD/overview)

* `l3_outside_name` (tnL3extOutName) - (string) The name of the L3 Outside object. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/l3_outside) with `aci_l3_outside.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/l3_outside) with `data.aci_l3_outside.example.name`.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation From Bridge Domain To L3 Outside object.
* `annotation` (annotation) - (string) The annotation of the Relation From Bridge Domain To L3 Outside object. This attribute is supported in ACI versions: 3.2(1l) and later.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
