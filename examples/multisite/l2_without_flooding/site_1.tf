resource "aci_tenant" "shared_tenant_site_1" {
  provider = aci.site_1
  name     = var.tenant_name
}

resource "aci_vrf" "shared_vrf_site_1" {
  provider  = aci.site_1
  tenant_dn = aci_tenant.shared_tenant_site_1.id
  name      = var.vrf_name
}

resource "aci_associated_site" "associated_site_of_shared_vrf_site_1" {
  provider  = aci.site_1
  parent_dn = aci_vrf.shared_vrf_site_1.id
  site_id   = var.site_1_id
  remote_sites = [
    {
      remote_pc_tag     = aci_vrf.shared_vrf_site_2.scope
      remote_vrf_pc_tag = aci_vrf.shared_vrf_site_2.pc_tag
      site_id           = var.site_2_id
    }
  ]
}

resource "aci_bridge_domain" "shared_bd_site_1" {
  provider  = aci.site_1
  parent_dn = aci_tenant.shared_tenant_site_1.id
  name      = var.bd_name
  relation_to_vrf = {
    vrf_name = aci_vrf.shared_vrf_site_1.name
  }
  intersite_l2_stretch = "yes"
}

resource "aci_subnet" "shared_subnet_bd_site_1" {
  provider  = aci.site_1
  parent_dn = aci_bridge_domain.shared_bd_site_1.id
  ip        = var.bd_subnet_ip
}

resource "aci_associated_site" "associated_site_of_shared_bd_site_1" {
  provider  = aci.site_1
  parent_dn = aci_bridge_domain.shared_bd_site_1.id
  site_id   = var.site_1_id
  remote_sites = [
    {
      remote_pc_tag = aci_bridge_domain.shared_bd_site_2.segment
      site_id       = var.site_2_id
    }
  ]
}

resource "aci_application_profile" "shared_ap_site_1" {
  provider  = aci.site_1
  tenant_dn = aci_tenant.shared_tenant_site_1.id
  name      = var.ap_name
}

resource "aci_application_epg" "shared_epg_site_1" {
  provider  = aci.site_1
  parent_dn = aci_application_profile.shared_ap_site_1.id
  name      = var.epg_name
  relation_to_bridge_domain = {
    bridge_domain_name = aci_bridge_domain.shared_bd_site_1.name
  }
  relation_to_domains = [
    {
      deployment_immediacy = "immediate"
      resolution_immediacy = "immediate"
      target_dn            = var.site_1_domain_target_dn
    }
  ]
}

resource "aci_associated_site" "associated_site_of_shared_epg_site_1" {
  provider  = aci.site_1
  parent_dn = aci_application_epg.shared_epg_site_1.id
  site_id   = var.site_1_id
  remote_sites = [
    {
      remote_pc_tag = aci_application_epg.shared_epg_site_2.pc_tag
      site_id       = var.site_2_id
    }
  ]
}
