
resource "aci_igmp_snooping_policy" "example_tenant" {
  parent_dn = aci_tenant.example.id
  name      = "test_name"
}
