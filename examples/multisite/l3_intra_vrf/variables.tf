# SHARED VARS

variable "tenant_name" {
  type    = string
  default = "multisite_l3_intra_vrf"
}

variable "vrf_name" {
  type    = string
  default = "shared_vrf"
}

variable "contract_name" {
  type    = string
  default = "shared_contract"
}

variable "contract_subject_name" {
  type    = string
  default = "shared_subject"
}

variable "filter_name" {
  type    = string
  default = "shared_filter"
}

variable "filter_entry_name" {
  type    = string
  default = "shared_filter_entry"
}

# SITE 1 VARS

variable "site_1_id" {
  type    = string
  default = "101"
}

variable "site_1_bd_name" {
  type    = string
  default = "bd_site_1"
}

variable "site_1_bd_subnet_ip" {
  type    = string
  default = "192.168.1.254/24"
}

variable "site_1_ap_name" {
  type    = string
  default = "ap_site_1"
}

variable "site_1_epg_name" {
  type    = string
  default = "epg_site_1"
}

variable "site_1_domain_target_dn" {
  type    = string
  default = "uni/vmmp-VMware/dom-VMware-VMM"
}

# SITE 2 VARS

variable "site_2_id" {
  type    = string
  default = "102"
}

variable "site_2_bd_name" {
  type    = string
  default = "bd_site_2"
}

variable "site_2_bd_subnet_ip" {
  type    = string
  default = "192.168.2.254/24"
}

variable "site_2_ap_name" {
  type    = string
  default = "ap_site_2"
}

variable "site_2_epg_name" {
  type    = string
  default = "epg_site_2"
}

variable "site_2_domain_target_dn" {
  type    = string
  default = "uni/phys-phys"
}

variable "site_2_static_path_encapsulation" {
  type    = string
  default = "vlan-105"
}

variable "site_2_static_path_target_dn" {
  type    = string
  default = "topology/pod-1/paths-101/pathep-[eth1/17]"
}