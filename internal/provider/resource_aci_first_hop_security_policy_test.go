// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccResourceFhsBDPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFhsBDPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "source_guard", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "source_guard", "enabled-both"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFhsBDPolMinDependencyWithFvTenantAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFhsBDPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test", "source_guard", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.allow_test_2", "source_guard", "enabled-both"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFhsBDPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "source_guard", "enabled-both"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFhsBDPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "ip_inspection", "disabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "router_advertisement", "disabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "source_guard", "disabled"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFhsBDPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name", "test_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFhsBDPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "source_guard", "enabled-both"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_first_hop_security_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFhsBDPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "ip_inspection", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "router_advertisement", "enabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "source_guard", "enabled-both"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.description", "description_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.managed_config_check", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.managed_config_flag", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.max_hop_limit", "10"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.max_router_preference", "disabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.min_hop_limit", "1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.name", "name_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.other_config_check", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.other_config_flag", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_first_hop_security_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFhsBDPolChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.description", "description_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.managed_config_check", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.managed_config_flag", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.max_hop_limit", "10"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.max_router_preference", "disabled"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.min_hop_limit", "1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.name", "name_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.other_config_check", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.other_config_flag", "no"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "route_advertisement_guard_policy.owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFhsBDPolChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.#", "1"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("aci_first_hop_security_policy.test",
						tfjsonpath.New("route_advertisement_guard_policy"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"annotation":            knownvalue.Null(),
								"annotations":           knownvalue.Null(),
								"tags":                  knownvalue.Null(),
								"description":           knownvalue.Null(),
								"managed_config_check":  knownvalue.Null(),
								"managed_config_flag":   knownvalue.Null(),
								"max_hop_limit":         knownvalue.Null(),
								"max_router_preference": knownvalue.Null(),
								"min_hop_limit":         knownvalue.Null(),
								"name":                  knownvalue.Null(),
								"name_alias":            knownvalue.Null(),
								"other_config_check":    knownvalue.Null(),
								"other_config_flag":     knownvalue.Null(),
								"owner_key":             knownvalue.Null(),
								"owner_tag":             knownvalue.Null(),
							},
						),
					),
				},
			},
			// Update with all children removed
			{
				Config:             testConfigFhsBDPolChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_first_hop_security_policy.test", "tags.#", "0"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("aci_first_hop_security_policy.test",
						tfjsonpath.New("route_advertisement_guard_policy"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"annotation":            knownvalue.Null(),
								"annotations":           knownvalue.Null(),
								"tags":                  knownvalue.Null(),
								"description":           knownvalue.Null(),
								"managed_config_check":  knownvalue.Null(),
								"managed_config_flag":   knownvalue.Null(),
								"max_hop_limit":         knownvalue.Null(),
								"max_router_preference": knownvalue.Null(),
								"min_hop_limit":         knownvalue.Null(),
								"name":                  knownvalue.Null(),
								"name_alias":            knownvalue.Null(),
								"other_config_check":    knownvalue.Null(),
								"other_config_flag":     knownvalue.Null(),
								"owner_key":             knownvalue.Null(),
								"owner_tag":             knownvalue.Null(),
							},
						),
					),
				},
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigFhsBDPolMinDependencyWithFvTenantAllowExisting = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "allow_test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
resource "aci_first_hop_security_policy" "allow_test_2" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_first_hop_security_policy.allow_test]
}
`

const testConfigFhsBDPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigFhsBDPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "annotation"
  description = "description_1"
  ip_inspection = "disabled"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  router_advertisement = "disabled"
  source_guard = "disabled"
}
`

const testConfigFhsBDPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "orchestrator:terraform"
  description = ""
  ip_inspection = "enabled-both"
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  router_advertisement = "enabled"
  source_guard = "enabled-both"
}
`
const testConfigFhsBDPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
  route_advertisement_guard_policy = {
    annotation = "annotation_1"
    annotations = [
	  {
        key = "key_0"
        value = "value_1"
	  },
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    tags = [
	  {
        key = "key_0"
        value = "value_1"
	  },
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    description = "description_1"
    managed_config_check = "no"
    managed_config_flag = "no"
    max_hop_limit = "10"
    max_router_preference = "disabled"
    min_hop_limit = "1"
    name = "name_1"
    name_alias = "name_alias_1"
    other_config_check = "no"
    other_config_flag = "no"
    owner_key = "owner_key_1"
    owner_tag = "owner_tag_1"
  }
  tags = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
}
`

const testConfigFhsBDPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigFhsBDPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  route_advertisement_guard_policy = {}
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigFhsBDPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_first_hop_security_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = []
  route_advertisement_guard_policy = {}
  tags = []
}
`
