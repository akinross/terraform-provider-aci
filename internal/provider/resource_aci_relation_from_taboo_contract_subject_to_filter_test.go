// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceVzRsDenyRuleWithVzTSubj(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigVzRsDenyRuleMinDependencyWithVzTSubjAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "directives.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "directives.#", "0"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigVzRsDenyRuleMinDependencyWithVzTSubjAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigVzRsDenyRuleMinDependencyWithVzTSubjAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test", "directives.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.allow_test_2", "directives.#", "0"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigVzRsDenyRuleMinDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.#", "0"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigVzRsDenyRuleAllDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.0", "log"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.1", "no_stats"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigVzRsDenyRuleMinDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "filter_name", "test_tn_vz_filter_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigVzRsDenyRuleResetDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.#", "0"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_from_taboo_contract_subject_to_filter.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigVzRsDenyRuleChildrenDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "filter_name", "test_tn_vz_filter_name"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "directives.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_from_taboo_contract_subject_to_filter.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigVzRsDenyRuleChildrenRemoveFromConfigDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigVzRsDenyRuleChildrenRemoveOneDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigVzRsDenyRuleChildrenRemoveAllDependencyWithVzTSubj,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_from_taboo_contract_subject_to_filter.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testDependencyConfigVzRsDenyRule = `
resource "aci_filter" "test_filter_0" {
  tenant_dn = aci_tenant.test.id
  name = "filter_name_1"
}
`

const testConfigVzRsDenyRuleMinDependencyWithVzTSubjAllowExisting = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "allow_test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
}
resource "aci_relation_from_taboo_contract_subject_to_filter" "allow_test_2" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
  depends_on = [aci_relation_from_taboo_contract_subject_to_filter.allow_test]
}
`

const testConfigVzRsDenyRuleMinDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
}
`

const testConfigVzRsDenyRuleAllDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
  annotation = "annotation"
  directives = ["log", "no_stats"]
}
`

const testConfigVzRsDenyRuleResetDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
  annotation = "orchestrator:terraform"
  directives = []
}
`
const testConfigVzRsDenyRuleChildrenDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
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
}
`

const testConfigVzRsDenyRuleChildrenRemoveFromConfigDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
}
`

const testConfigVzRsDenyRuleChildrenRemoveOneDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigVzRsDenyRuleChildrenRemoveAllDependencyWithVzTSubj = testDependencyConfigVzRsDenyRule + testConfigVzTSubjMinDependencyWithVzTaboo + `
resource "aci_relation_from_taboo_contract_subject_to_filter" "test" {
  parent_dn = aci_taboo_contract_subject.test.id
  filter_name = "test_tn_vz_filter_name"
  annotations = []
  tags = []
}
`
