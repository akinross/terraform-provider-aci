// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsCtxToEigrpCtxAfPolWithFvCtx(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsCtxToEigrpCtxAfPolDataSourceDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_from_vrf_to_eigrp_address_family_context.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("data.aci_relation_from_vrf_to_eigrp_address_family_context.test", "eigrp_address_family_context_name", "test_tn_eigrp_ctx_af_pol_name"),
					resource.TestCheckResourceAttr("data.aci_relation_from_vrf_to_eigrp_address_family_context.test", "annotation", "orchestrator:terraform"),
				),
			},
			{
				Config:      testConfigFvRsCtxToEigrpCtxAfPolNotExistingFvCtx,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_from_vrf_to_eigrp_address_family_context data source"),
			},
		},
	})
}

const testConfigFvRsCtxToEigrpCtxAfPolDataSourceDependencyWithFvCtx = testConfigFvRsCtxToEigrpCtxAfPolMinDependencyWithFvCtx + `
data "aci_relation_from_vrf_to_eigrp_address_family_context" "test" {
  parent_dn = aci_vrf.test.id
  address_family = "ipv4-ucast"
  eigrp_address_family_context_name = "test_tn_eigrp_ctx_af_pol_name"
  depends_on = [aci_relation_from_vrf_to_eigrp_address_family_context.test]
}
`

const testConfigFvRsCtxToEigrpCtxAfPolNotExistingFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
data "aci_relation_from_vrf_to_eigrp_address_family_context" "test_non_existing" {
  parent_dn = aci_vrf.test.id
  address_family = "ipv6-ucast"
  eigrp_address_family_context_name = "non_existing_tn_eigrp_ctx_af_pol_name"
}
`
