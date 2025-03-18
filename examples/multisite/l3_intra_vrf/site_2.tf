# SHARED CONFIGURATION

resource "aci_tenant" "shared_tenant_site_2" {
  provider = aci.site_2
  name     = var.tenant_name
}

resource "aci_vrf" "shared_vrf_site_2" {
  provider  = aci.site_2
  tenant_dn = aci_tenant.shared_tenant_site_2.id
  name      = var.vrf_name
}

resource "aci_associated_site" "associated_site_of_vrf_site_2" {
  provider  = aci.site_2
  parent_dn = aci_vrf.shared_vrf_site_2.id
  site_id   = var.site_2_id
  remote_sites = [
    {
      remote_pc_tag     = aci_vrf.shared_vrf_site_1.scope
      remote_vrf_pc_tag = aci_vrf.shared_vrf_site_1.pc_tag
      site_id           = var.site_1_id
    }
  ]
}

resource "aci_contract" "shared_contract_site_2" {
  provider  = aci.site_2
  tenant_dn = aci_tenant.shared_tenant_site_2.id
  name      = var.contract_name
}

resource "aci_contract_subject" "shared_contract_subject_site_2" {
  provider    = aci.site_2
  contract_dn = aci_contract.shared_contract_site_2.id
  name        = var.contract_subject_name
}

resource "aci_contract_subject_filter" "shared_contract_subject_filter_site_2" {
  provider            = aci.site_2
  contract_subject_dn = aci_contract_subject.shared_contract_subject_site_2.id
  filter_dn           = aci_filter.shared_filter_site_2.id
}

resource "aci_filter" "shared_filter_site_2" {
  provider  = aci.site_2
  tenant_dn = aci_tenant.shared_tenant_site_1.id
  name      = var.filter_name
}

resource "aci_filter_entry" "shared_filter_entry_site_2" {
  provider  = aci.site_2
  filter_dn = aci_filter.shared_filter_site_2.id
  name      = var.filter_entry_name
}

# SITE SPECIFIC CONFIGURATION

resource "aci_bridge_domain" "bd_site_2" {
  provider  = aci.site_2
  parent_dn = aci_tenant.shared_tenant_site_2.id
  name      = var.site_2_bd_name
  relation_to_vrf = {
    vrf_name = aci_vrf.shared_vrf_site_2.name
  }
  arp_flooding                 = "yes"
  enable_intersite_bum_traffic = "yes"
  intersite_l2_stretch         = "yes"
  optimize_wan_bandwidth       = "yes"
}

resource "aci_subnet" "subnet_bd_site_2" {
  provider  = aci.site_2
  parent_dn = aci_bridge_domain.bd_site_2.id
  ip        = var.site_2_bd_subnet_ip
}

resource "aci_associated_site" "associated_site_of_bd_site_2" {
  provider  = aci.site_2
  parent_dn = aci_bridge_domain.bd_site_2.id
  site_id   = var.site_2_id
  remote_sites = [
    {
      remote_pc_tag = aci_bridge_domain.shadow_bd_site_2.segment
      site_id       = var.site_1_id
    }
  ]
}

resource "aci_application_profile" "ap_site_2" {
  provider  = aci.site_2
  tenant_dn = aci_tenant.shared_tenant_site_2.id
  name      = var.site_2_ap_name
}

resource "aci_application_epg" "epg_site_2" {
  provider  = aci.site_2
  parent_dn = aci_application_profile.ap_site_2.id
  name      = var.site_2_epg_name
  relation_to_bridge_domain = {
    bridge_domain_name = aci_bridge_domain.bd_site_2.name
  }
  relation_to_consumed_contracts = [
    {
      contract_name = aci_contract.shared_contract_site_2.name
    }
  ]
  relation_to_domains = [
    {
      deployment_immediacy = "immediate"
      resolution_immediacy = "immediate"
      target_dn            = var.site_2_domain_target_dn
    }
  ]
  relation_to_static_paths = [
    {
      encapsulation        = var.site_2_static_path_encapsulation
      deployment_immediacy = "immediate"
      mode                 = "untagged"
      target_dn            = var.site_2_static_path_target_dn
    }
  ]
}

resource "aci_associated_site" "associated_site_of_epg_site_2" {
  provider  = aci.site_2
  parent_dn = aci_application_epg.epg_site_2.id
  site_id   = var.site_2_id
  remote_sites = [
    {
      remote_pc_tag = aci_application_epg.shadow_epg_site_2.pc_tag
      site_id       = var.site_1_id
    }
  ]
}

# SHADOW CONFIGURATION OF SITE 1

resource "aci_bridge_domain" "shadow_bd_site_1" {
  provider   = aci.site_2
  annotation = "orchestrator:msc-shadow:yes"
  parent_dn  = aci_tenant.shared_tenant_site_2.id
  name       = aci_bridge_domain.bd_site_1.name
  relation_to_vrf = {
    vrf_name = aci_vrf.shared_vrf_site_2.name
  }
  arp_flooding                 = "yes"
  enable_intersite_bum_traffic = "yes"
  intersite_l2_stretch         = "yes"
  optimize_wan_bandwidth       = "yes"
}

resource "aci_subnet" "shadow_subnet_bd_site_1" {
  provider   = aci.site_2
  annotation = "orchestrator:msc-shadow:yes"
  parent_dn  = aci_bridge_domain.shadow_bd_site_1.id
  ip         = aci_subnet.subnet_bd_site_1.ip
}

resource "aci_associated_site" "associated_site_of_shadow_bd_site_1" {
  provider  = aci.site_2
  parent_dn = aci_bridge_domain.shadow_bd_site_1.id
  site_id   = var.site_2_id
  remote_sites = [
    {
      remote_pc_tag = aci_bridge_domain.bd_site_1.segment
      site_id       = var.site_1_id
    }
  ]
}

resource "aci_application_profile" "shadow_ap_site_1" {
  provider   = aci.site_2
  annotation = "orchestrator:msc-shadow:yes"
  tenant_dn  = aci_tenant.shared_tenant_site_2.id
  name       = aci_application_profile.ap_site_1.name

}

resource "aci_application_epg" "shadow_epg_site_1" {
  provider   = aci.site_2
  annotation = "orchestrator:msc-shadow:yes"
  parent_dn  = aci_application_profile.shadow_ap_site_1.id
  relation_to_bridge_domain = {
    annotation         = "orchestrator:msc-shadow:yes"
    bridge_domain_name = aci_bridge_domain.shadow_bd_site_1.name
  }
  relation_to_provided_contracts = [
    {
      annotation    = "orchestrator:msc-shadow:yes"
      contract_name = aci_contract.shared_contract_site_2.name
    }
  ]
  name = aci_application_epg.epg_site_1.name
}

resource "aci_associated_site" "shadow_associated_site_site_1" {
  provider  = aci.site_2
  parent_dn = aci_application_epg.shadow_epg_site_1.id
  site_id   = var.site_2_id
  remote_sites = [
    {
      remote_pc_tag = aci_application_epg.epg_site_1.pc_tag
      site_id       = var.site_1_id
    }
  ]
}