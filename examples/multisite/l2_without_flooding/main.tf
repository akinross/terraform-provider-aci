terraform {
  required_providers {
    aci = {
      source = "ciscodevnet/aci"
    }
  }
}

# PROVIDER CONFIGURATION FOR SITE 1: ansible_test
provider "aci" {
  alias    = "site_1"
  username = ""
  password = ""
  url      = ""
}

# PROVIDER CONFIGURATION FOR SITE 2: ansible_test_2
provider "aci" {
  alias    = "site_2"
  username = ""
  password = ""
  url      = ""
}
