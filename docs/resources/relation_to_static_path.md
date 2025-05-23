---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_relation_to_static_path"
sidebar_current: "docs-aci-resource-aci_relation_to_static_path"
description: |-
  Manages ACI Relation To Static Path
---

# aci_relation_to_static_path #

Manages ACI Relation To Static Path



## API Information ##

* Class: [fvRsPathAtt](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsPathAtt/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-{name}/ap-{name}/epg-{name}/rspathAtt-[{tDn}]`

## GUI Information ##

* Location: `Tenants -> Application Profiles -> Application EPGs -> Static Ports`

## Example Usage ##

The configuration snippet below creates a Relation To Static Path with only required attributes.

```hcl

resource "aci_relation_to_static_path" "example_application_epg" {
  parent_dn     = aci_application_epg.example.id
  encapsulation = "vlan-201"
  target_dn     = "topology/pod-1/paths-101/pathep-[eth1/1]"
}

```
The configuration snippet below shows all possible attributes of the Relation To Static Path.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_relation_to_static_path" "full_example_application_epg" {
  parent_dn             = aci_application_epg.example.id
  annotation            = "annotation"
  description           = "description_1"
  encapsulation         = "vlan-201"
  deployment_immediacy  = "immediate"
  mode                  = "native"
  primary_encapsulation = "vlan-203"
  target_dn             = "topology/pod-1/paths-101/pathep-[eth1/1]"
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

All examples for the Relation To Static Path resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_relation_to_static_path) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_application_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_epg) ([fvAEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAEPg/overview))
* `encapsulation` (encap) - (string) The VLAN encapsulation of the Relation To Static Path object.
* `target_dn` (tDn) - (string) null.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation To Static Path object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the Relation To Static Path object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the Relation To Static Path object. This attribute is supported in ACI versions: 1.0(4g) and later.
* `deployment_immediacy` (instrImedcy) - (string) The deployment immediacy of the Relation To Static Path object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).
  - Default: `lazy`
  - Valid Values: `immediate`, `lazy`.
* `mode` (mode) - (string) The static association mode of the Relation To Static Path object.
  - Default: `regular`
  - Valid Values: `native`, `regular`, `untagged`.
* `primary_encapsulation` (primaryEncap) - (string) The primary VLAN encapsulation of the Relation To Static Path object. This attribute is supported in ACI versions: 1.2(2g) and later.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Relation To Static Path can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_relation_to_static_path.example_application_epg uni/tn-{name}/ap-{name}/epg-{name}/rspathAtt-[{tDn}]
```

Starting in Terraform version 1.5, an existing Relation To Static Path can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/epg-{name}/rspathAtt-[{tDn}]"
  to = aci_relation_to_static_path.example_application_epg
}
```
