# SHARED VARS

variable "tenant_name" {
  type    = string
  default = "multisite_l2_without_flooding"
}

variable "vrf_name" {
  type    = string
  default = "shared_vrf"
}

variable "bd_name" {
  type    = string
  default = "shared_bd"
}

variable "bd_subnet_ip" {
  type    = string
  default = "192.168.2.254/24"
}

variable "ap_name" {
  type    = string
  default = "shared_ap"
}

variable "epg_name" {
  type    = string
  default = "shared_epg"
}


# SITE 1 VARS

variable "site_1_id" {
  type    = string
  default = "101"
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