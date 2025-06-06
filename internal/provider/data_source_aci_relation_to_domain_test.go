// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsDomAttWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsDomAttDataSourceDependencyWithFvAEPg + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "target_dn", "uni/vmmp-VMware/dom-domain_1"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "class_preference", "encap"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "delimiter", ""),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "enable_netflow", "disabled"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "encapsulation", "unknown"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "encapsulation_mode", "auto"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "epg_cos", "Cos0"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "epg_cos_pref", "disabled"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "lag_policy_name", ""),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "netflow_direction", "both"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "primary_encapsulation", "unknown"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "primary_encapsulation_inner", "unknown"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "resolution_immediacy", "lazy"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "secondary_encapsulation_inner", "unknown"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "switching_mode", "native"),
					resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "untagged", "no"),
					composeAggregateTestCheckFuncWithVersion(t, "4.0(1h)-", "inside",
						resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "binding_type", "none"),
						resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "number_of_ports", "0"),
						resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "port_allocation", "none")),
					composeAggregateTestCheckFuncWithVersion(t, "4.2(3j)-", "inside",
						resource.TestCheckResourceAttr("data.aci_relation_to_domain.test", "custom_epg_name", "")),
				),
			},
			{
				Config:      testConfigFvRsDomAttNotExistingFvAEPg + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_domain data source"),
			},
		},
	})
}

const testConfigFvRsDomAttDataSourceDependencyWithFvAEPg = testConfigFvRsDomAttMinDependencyWithFvAEPg + `
data "aci_relation_to_domain" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "uni/vmmp-VMware/dom-domain_1"
  depends_on = [aci_relation_to_domain.test]
}
`

const testConfigFvRsDomAttNotExistingFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
data "aci_relation_to_domain" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  target_dn = "uni/vmmp-VMware/dom-domain_1_not_existing"
}
`
