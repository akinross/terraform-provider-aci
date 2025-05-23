---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Multi-Site"
layout: "aci"
page_title: "ACI: aci_remote_site"
sidebar_current: "docs-aci-resource-aci_remote_site"
description: |-
  Manages ACI Remote Site
---

# aci_remote_site #

Manages ACI Remote Site

  -> This resource should only be used to accomplish multi-site orchestration.


## API Information ##

* Class: [fvRemoteId](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRemoteId/overview)

* Supported in ACI versions: 3.0(1k) and later.

* Distinguished Name Formats:
  - `uni/tn-{name}/BD-{name}/stAsc/site-{siteId}`
  - `uni/tn-{name}/ap-{name}/epg-{name}/stAsc/site-{siteId}`
  - `uni/tn-{name}/ctx-{name}/stAsc/site-{siteId}`
  - `uni/tn-{name}/mscGraphXlateCont/epgDefXlate-[{epgDefDn}]/stAsc/site-{siteId}`
  - `uni/tn-{name}/out-{name}/instP-{name}/stAsc/site-{siteId}`

## GUI Information ##

* Location: `Not shown in UI`

## Example Usage ##

The configuration snippet below creates a Remote Site with only required attributes.

```hcl

resource "aci_remote_site" "example_associated_site" {
  parent_dn = aci_associated_site.example.id
  site_id   = "0"
}

```
The configuration snippet below shows all possible attributes of the Remote Site.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_remote_site" "full_example_associated_site" {
  parent_dn         = aci_associated_site.example.id
  annotation        = "annotation"
  description       = "description_1"
  name              = "name_1"
  name_alias        = "name_alias_1"
  owner_key         = "owner_key_1"
  owner_tag         = "owner_tag_1"
  remote_vrf_pc_tag = "any"
  remote_pc_tag     = "16387"
  site_id           = "0"
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

All examples for the Remote Site resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_remote_site) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_associated_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/associated_site) ([fvSiteAssociated](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvSiteAssociated/overview))
* `site_id` (siteId) - (string) The remote site identifier that is associated with the object of the primary/local site as an integer.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Remote Site object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the Remote Site object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the Remote Site object.
* `name` (name) - (string) The name of the Remote Site object.
* `name_alias` (nameAlias) - (string) The name alias of the Remote Site object.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `remote_vrf_pc_tag` (remoteCtxPcTag) - (string) The policy class tag (pcTag) of the remote VRF. This attribute can only be present when the object for site association is a VRF.
* `remote_pc_tag` (remotePcTag) - (string) The policy class tag (pcTag) of the remote object.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Remote Site can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_remote_site.example_associated_site uni/tn-{name}/ap-{name}/epg-{name}/stAsc/site-{siteId}
```

Starting in Terraform version 1.5, an existing Remote Site can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/epg-{name}/stAsc/site-{siteId}"
  to = aci_remote_site.example_associated_site
}
```
