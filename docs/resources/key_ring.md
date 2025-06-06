---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "AAA"
layout: "aci"
page_title: "ACI: aci_key_ring"
sidebar_current: "docs-aci-resource-aci_key_ring"
description: |-
  Manages ACI Key Ring
---

# aci_key_ring #

Manages ACI Key Ring



## API Information ##

* Class: [pkiKeyRing](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/pkiKeyRing/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Formats:
  - `uni/tn-{name}/certstore/keyring-{name}`
  - `uni/userext/pkiext/keyring-{name}`

## GUI Information ##

* Locations:
  - `Admin -> AAA -> Security -> Key Rings`
  - `Cloud Network Controller -> Administrative -> Security -> Key Rings`

## Example Usage ##

The configuration snippet below creates a Key Ring with only required attributes.

```hcl

resource "aci_key_ring" "example" {
  name = "test_name"
}

// This example is only applicable to Cisco Cloud Network Controller
resource "aci_key_ring" "example_tenant" {
  parent_dn = aci_tenant.example.id
  name      = "test_name"
}

```
The configuration snippet below shows all possible attributes of the Key Ring.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_key_ring" "full_example" {
  admin_state           = "completed"
  annotation            = "annotation"
  certificate           = <<EOT
-----BEGIN CERTIFICATE-----
MIICODCCAaGgAwIBAgIJAIt8XMntue0VMA0GCSqGSIb3DQEBCwUAMDQxDjAMBgNV
BAMMBUFkbWluMRUwEwYDVQQKDAxZb3VyIENvbXBhbnkxCzAJBgNVBAYTAlVTMCAX
DTE4MDEwOTAwNTk0NFoYDzIxMTcxMjE2MDA1OTQ0WjA0MQ4wDAYDVQQDDAVBZG1p
bjEVMBMGA1UECgwMWW91ciBDb21wYW55MQswCQYDVQQGEwJVUzCBnzANBgkqhkiG
9w0BAQEFAAOBjQAwgYkCgYEAohG/7axtt7CbSaMP7r+2mhTKbNgh0Ww36C7Ta14i
v+VmLyKkQHnXinKGhp6uy3Nug+15a+eIu7CrgpBVMQeCiWfsnwRocKcQJWIYDrWl
XHxGQn31yYKR6mylE7Dcj3rMFybnyhezr5D8GcP85YRPmwG9H2hO/0Y1FUnWu9Iw
AQkCAwEAAaNQME4wHQYDVR0OBBYEFD0jLXfpkrU/ChzRvfruRs/fy1VXMB8GA1Ud
IwQYMBaAFD0jLXfpkrU/ChzRvfruRs/fy1VXMAwGA1UdEwQFMAMBAf8wDQYJKoZI
hvcNAQELBQADgYEAOmvre+5tgZ0+F3DgsfxNQqLTrGiBgGCIymPkP/cBXXkNuJyl
3ac7tArHQc7WEA4U2R2rZbEq8FC3UJJm4nUVtCPvEh3G9OhN2xwYev79yt6pIn/l
KU0Td2OpVyo0eLqjoX5u2G90IBWzhyjFbo+CcKMrSVKj1YOdG0E3OuiJf00=
-----END CERTIFICATE-----
EOT
  description           = "description_1"
  elliptic_curve        = "none"
  key                   = <<EOT
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKIRv+2sbbewm0mj
D+6/tpoUymzYIdFsN+gu02teIr/lZi8ipEB514pyhoaerstzboPteWvniLuwq4KQ
VTEHgoln7J8EaHCnECViGA61pVx8RkJ99cmCkepspROw3I96zBcm58oXs6+Q/BnD
/OWET5sBvR9oTv9GNRVJ1rvSMAEJAgMBAAECgYByu3QO0qF9h7X3JEu0Ld4cKBnB
giQ2uJC/et7KxIJ/LOvw9GopBthyt27KwG1ntBkJpkTuAaQHkyNns7vLkNB0S0IR
+owVFEcKYq9VCHTaiQU8TDp24gN+yPTrpRuH8YhDVq5SfVdVuTMgHVQdj4ya4VlF
Gj+a7+ipxtGiLsVGrQJBAM7p0Fm0xmzi+tBOASUAcVrPLcteFIaTBFwfq16dm/ON
00Khla8Et5kMBttTbqbukl8mxFjBEEBlhQqb6EdQQ0sCQQDIhHx1a9diG7y/4DQA
4KvR3FCYwP8PBORlSamegzCo+P1OzxiEo0amX7yQMA5UyiP/kUsZrme2JBZgna8S
p4R7AkEAr7rMhSOPUnMD6V4WgsJ5g1Jp5kqkzBaYoVUUSms5RASz4+cwJVCwTX91
Y1jcpVIBZmaaY3a0wrx13ajEAa0dOQJBAIpjnb4wqpsEh7VpmJqOdSdGxb1XXfFQ
sA0T1OQYqQnFppWwqrxIL+d9pZdiA1ITnNqyvUFBNETqDSOrUHwwb2cCQGArE+vu
ffPUWQ0j+fiK+covFG8NL7H+26NSGB5+Xsn9uwOGLj7K/YT6CbBtr9hJiuWjM1Al
0V4ltlTuu2mTMaw=
-----END PRIVATE KEY-----
EOT
  key_type              = "RSA"
  modulus               = "mod1024"
  name                  = "test_name"
  name_alias            = "name_alias_1"
  owner_key             = "owner_key_1"
  owner_tag             = "owner_tag_1"
  regenerate            = "no"
  certificate_authority = "test_name"
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
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
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
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
  ]
}

// This example is only applicable to Cisco Cloud Network Controller
resource "aci_key_ring" "full_example_tenant" {
  parent_dn             = aci_tenant.example.id
  admin_state           = "completed"
  annotation            = "annotation"
  certificate           = <<EOT
-----BEGIN CERTIFICATE-----
MIICODCCAaGgAwIBAgIJAIt8XMntue0VMA0GCSqGSIb3DQEBCwUAMDQxDjAMBgNV
BAMMBUFkbWluMRUwEwYDVQQKDAxZb3VyIENvbXBhbnkxCzAJBgNVBAYTAlVTMCAX
DTE4MDEwOTAwNTk0NFoYDzIxMTcxMjE2MDA1OTQ0WjA0MQ4wDAYDVQQDDAVBZG1p
bjEVMBMGA1UECgwMWW91ciBDb21wYW55MQswCQYDVQQGEwJVUzCBnzANBgkqhkiG
9w0BAQEFAAOBjQAwgYkCgYEAohG/7axtt7CbSaMP7r+2mhTKbNgh0Ww36C7Ta14i
v+VmLyKkQHnXinKGhp6uy3Nug+15a+eIu7CrgpBVMQeCiWfsnwRocKcQJWIYDrWl
XHxGQn31yYKR6mylE7Dcj3rMFybnyhezr5D8GcP85YRPmwG9H2hO/0Y1FUnWu9Iw
AQkCAwEAAaNQME4wHQYDVR0OBBYEFD0jLXfpkrU/ChzRvfruRs/fy1VXMB8GA1Ud
IwQYMBaAFD0jLXfpkrU/ChzRvfruRs/fy1VXMAwGA1UdEwQFMAMBAf8wDQYJKoZI
hvcNAQELBQADgYEAOmvre+5tgZ0+F3DgsfxNQqLTrGiBgGCIymPkP/cBXXkNuJyl
3ac7tArHQc7WEA4U2R2rZbEq8FC3UJJm4nUVtCPvEh3G9OhN2xwYev79yt6pIn/l
KU0Td2OpVyo0eLqjoX5u2G90IBWzhyjFbo+CcKMrSVKj1YOdG0E3OuiJf00=
-----END CERTIFICATE-----
EOT
  description           = "description_1"
  elliptic_curve        = "none"
  key                   = <<EOT
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKIRv+2sbbewm0mj
D+6/tpoUymzYIdFsN+gu02teIr/lZi8ipEB514pyhoaerstzboPteWvniLuwq4KQ
VTEHgoln7J8EaHCnECViGA61pVx8RkJ99cmCkepspROw3I96zBcm58oXs6+Q/BnD
/OWET5sBvR9oTv9GNRVJ1rvSMAEJAgMBAAECgYByu3QO0qF9h7X3JEu0Ld4cKBnB
giQ2uJC/et7KxIJ/LOvw9GopBthyt27KwG1ntBkJpkTuAaQHkyNns7vLkNB0S0IR
+owVFEcKYq9VCHTaiQU8TDp24gN+yPTrpRuH8YhDVq5SfVdVuTMgHVQdj4ya4VlF
Gj+a7+ipxtGiLsVGrQJBAM7p0Fm0xmzi+tBOASUAcVrPLcteFIaTBFwfq16dm/ON
00Khla8Et5kMBttTbqbukl8mxFjBEEBlhQqb6EdQQ0sCQQDIhHx1a9diG7y/4DQA
4KvR3FCYwP8PBORlSamegzCo+P1OzxiEo0amX7yQMA5UyiP/kUsZrme2JBZgna8S
p4R7AkEAr7rMhSOPUnMD6V4WgsJ5g1Jp5kqkzBaYoVUUSms5RASz4+cwJVCwTX91
Y1jcpVIBZmaaY3a0wrx13ajEAa0dOQJBAIpjnb4wqpsEh7VpmJqOdSdGxb1XXfFQ
sA0T1OQYqQnFppWwqrxIL+d9pZdiA1ITnNqyvUFBNETqDSOrUHwwb2cCQGArE+vu
ffPUWQ0j+fiK+covFG8NL7H+26NSGB5+Xsn9uwOGLj7K/YT6CbBtr9hJiuWjM1Al
0V4ltlTuu2mTMaw=
-----END PRIVATE KEY-----
EOT
  key_type              = "RSA"
  modulus               = "mod1024"
  name                  = "test_name"
  name_alias            = "name_alias_1"
  owner_key             = "owner_key_1"
  owner_tag             = "owner_tag_1"
  regenerate            = "no"
  certificate_authority = "test_name"
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
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
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
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
  ]
}

```

All examples for the Key Ring resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_key_ring) folder.

## Schema ##

### Required ###

* `name` (name) - (string) The name of the Key Ring object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Key Ring object.

### Optional ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_tenant](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/fvTenant) ([fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview))
  - Default: `uni/userext/pkiext`
* `admin_state` (adminState) - (string) The current administrative state of the certificate request process.
  - Default: `started`
  - Valid Values: `completed`, `created`, `reqCreated`, `started`, `tpSet`.
* `annotation` (annotation) - (string) The annotation of the Key Ring object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `certificate` (cert) - (string) A certificate contains a device's public key along with signed information verifying the identity of the device.
* `description` (descr) - (string) The description of the Key Ring object.
* `elliptic_curve` (eccCurve) - (string) The elliptic curve used by the provided key. This attribute is supported in ACI versions: 6.0(2h) and later.
  - Valid Values: `none`, `prime256v1`, `secp384r1`, `secp521r1`.
* `key` (key) - (string) The private key of the certificate. This sensitive value is excluded from the resource's lifecycle configuration and is not tracked by Terraform.
* `key_type` (keyType) - (string) The type used by the provided key. This attribute is supported in ACI versions: 6.0(2h) and later.
  - Default: `RSA`
  - Valid Values: `ECC`, `RSA`, `indeterminate`.
* `modulus` (modulus) - (string) The length of the encryption keys. A longer key length increases the difficulty of breaking the key.
  - Default: `mod2048`
  - Valid Values: `mod1024`, `mod1536`, `mod2048`, `mod3072`, `mod4096`, `mod512`, `none`.
* `name_alias` (nameAlias) - (string) The name alias of the Key Ring object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `regenerate` (regen) - (string) Forces regeneration of the keypair. Each PKI device holds a pair of asymmetric Rivest-Shamir-Adleman (RSA) or Elliptic Curve Cryptography (ECC) encryption keys, one kept private and one made public, stored in an internal key ring.
  - Default: `no`
  - Valid Values: `no`, `yes`.
* `certificate_authority` (tp) - (string) The certificate of the Certificate Authority (CA) that issued the certificate provided in the 'certificate' attribute. The CA can be a root CA, an intermediate CA, or a trust anchor in a chain of trust leading to a root CA.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Key Ring can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_key_ring.example uni/tn-{name}/certstore/keyring-{name}
```

Starting in Terraform version 1.5, an existing Key Ring can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/certstore/keyring-{name}"
  to = aci_key_ring.example
}
```
