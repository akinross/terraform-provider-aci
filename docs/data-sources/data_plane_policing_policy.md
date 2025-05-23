---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Tenant Policies"
layout: "aci"
page_title: "ACI: aci_data_plane_policing_policy"
sidebar_current: "docs-aci-data-source-aci_data_plane_policing_policy"
description: |-
  Data source for ACI Data Plane Policing Policy
---

# aci_data_plane_policing_policy #

Data source for ACI Data Plane Policing Policy

## API Information ##

* Class: [qosDppPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/qosDppPol/overview)

* Supported in ACI versions: 1.2(2g) and later.

* Distinguished Name Formats:
  - `uni/infra/qosdpppol-{name}`
  - `uni/tn-{name}/qosdpppol-{name}`

## GUI Information ##

* Location: `Tenants -> Policies -> Protocol -> Data Plane Policing`

## Example Usage ##

```hcl

data "aci_data_plane_policing_policy" "example_tenant" {
  parent_dn = aci_tenant.example.id
  name      = "test_name"
}

```

## Schema ##

### Required ###

* `name` (name) - (string) The name of the Data Plane Policing Policy object.

### Optional ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_tenant](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tenant) ([fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [infraInfra](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraInfra/overview)

  - Default: `uni/infra`

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Data Plane Policing Policy object.
* `admin_state` (adminSt) - (string) The administrative state of the Data Plane Policing Policy object.
* `annotation` (annotation) - (string) The annotation of the Data Plane Policing Policy object. This attribute is supported in ACI versions: 3.2(1l) and later.
* `excessive_burst` (be) - (string) The excessive burst size of the Data Plane Policing Policy object. Only applicable for 2R3C policer.
* `excessive_burst_unit` (beUnit) - (string) The excessive burst size unit of the Data Plane Policing Policy object. Only applicable for 2R3C policer.
* `burst` (burst) - (string) The burst size of the Data Plane Policing Policy object.
* `burst_unit` (burstUnit) - (string) The burst size unit of the Data Plane Policing Policy object.
* `conform_action` (conformAction) - (string) The conform action of the Data Plane Policing Policy object.
* `conform_mark_cos` (conformMarkCos) - (string) The conform mark class of service (CoS) of the Data Plane Policing Policy object.
* `conform_mark_dscp` (conformMarkDscp) - (string) The conform mark differentiated services code point (DSCP) of the Data Plane Policing Policy object.
* `description` (descr) - (string) The description of the Data Plane Policing Policy object.
* `exceed_action` (exceedAction) - (string) The exceed action of the Data Plane Policing Policy object.
* `exceed_mark_cos` (exceedMarkCos) - (string) The exceed mark class of service (CoS) of the Data Plane Policing Policy object.
* `exceed_mark_dscp` (exceedMarkDscp) - (string) The exceed mark differentiated services code point (DSCP) of the Data Plane Policing Policy object.
* `mode` (mode) - (string) Policer mode - bytes or packet policer.
* `name_alias` (nameAlias) - (string) The name alias of the Data Plane Policing Policy object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `peak_rate` (pir) - (string) The peak information rate (PIR) of the Data Plane Policing Policy object. Only applicable for 2R3C policer.
* `peak_rate_unit` (pirUnit) - (string) The peak information rate (PIR) unit of the Data Plane Policing Policy object. Only applicable for 2R3C policer.
* `rate` (rate) - (string) The rate of the Data Plane Policing Policy object.
* `rate_unit` (rateUnit) - (string) The rate unit of the Data Plane Policing Policy object.
* `sharing_mode` (sharingMode) - (string) The sharing mode of the Data Plane Policing Policy object. This attribute is supported in ACI versions: 3.1(1i) and later.
* `type` (type) - (string) The type of the Data Plane Policing Policy object.
* `violate_action` (violateAction) - (string) The violate action of the Data Plane Policing Policy object.
* `violate_mark_cos` (violateMarkCos) - (string) The violate mark class of service (CoS) of the Data Plane Policing Policy object.
* `violate_mark_dscp` (violateMarkDscp) - (string) The violate mark differentiated services code point (DSCP) of the Data Plane Policing Policy object.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
