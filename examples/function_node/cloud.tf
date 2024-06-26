# terraform plan for cloud APICs

data "aci_tenant" "tf_tenant" {
  name = "tf_ansible_test"
}

data "aci_vrf" "tf_vrf" {
  tenant_dn = data.aci_tenant.tf_tenant.id
  name      = "tf_vrf"
}

data "aci_cloud_context_profile" "ccp1" {
  tenant_dn = data.aci_tenant.tf_tenant.id
  name      = "tf_ccp"
}

data "aci_cloud_cidr_pool" "cidr1" {
  cloud_context_profile_dn = data.aci_cloud_context_profile.ccp1.id
  addr                     = "10.20.0.0/25"
}

data "aci_cloud_subnet" "cs1" {
  cloud_cidr_pool_dn = data.aci_cloud_cidr_pool.cidr1.id
  ip                 = "10.20.0.0/25"
}

# Create Logical Firewall Representation (3rd party example)

resource "aci_cloud_l4_l7_third_party_device" "third_party_fw" {
  tenant_dn                     = data.aci_tenant.tf_tenant.id
  name                          = "tf_third_party_fw"
  relation_cloud_rs_ldev_to_ctx = data.aci_vrf.tf_vrf.id
  interface_selectors {
    allow_all = "yes"
    name      = "trust"
    end_point_selectors {
      match_expression = "custom:internal=='trust'"
      name             = "trust"
    }
  }
  interface_selectors {
    allow_all = "yes"
    name      = "untrust"
    end_point_selectors {
      match_expression = "custom:external=='untrust'"
      name             = "untrust"
    }
  }
}

# Create Native Network Load Balancer for Firewall

resource "aci_cloud_l4_l7_native_load_balancer" "cloud_nlb" {
  tenant_dn                              = data.aci_tenant.tf_tenant.id
  name                                   = "tf_cloud_nlb"
  relation_cloud_rs_ldev_to_cloud_subnet = [data.aci_cloud_subnet.cs1.id]
  allow_all                              = "yes"
  is_static_ip                           = "yes"
  static_ip_addresses                    = ["10.20.0.0"]
  scheme                                 = "internal"
  cloud_l4l7_load_balancer_type          = "network"
}

#  1. Create first L4-L7 Service Graph Template 'sg1' with type cloud
#    Add two nodes with basic parameters
#     - The first node 'N0' with type ADC_ONE_ARM and relation to the cloud native load balancer
#     - The second node 'N1' with type FW_ROUTED and relation to the 3rd party firewall
#    Create the connection between the templates ('T1 - consumer' and 'T2 - provider') and nodes ('N0' and 'N1').

resource "aci_l4_l7_service_graph_template" "sg1" {
  tenant_dn                         = data.aci_tenant.tf_tenant.id
  name                              = "tf_sg_1"
  l4_l7_service_graph_template_type = "cloud"
}

# Create a function node with type ADC_ONE_ARM and relation to the cloud native load balancer
resource "aci_function_node" "function_node_nlb" {
  l4_l7_service_graph_template_dn     = aci_l4_l7_service_graph_template.sg1.id
  name                                = "N0"
  func_template_type                  = "ADC_ONE_ARM"
  relation_vns_rs_node_to_cloud_l_dev = aci_cloud_l4_l7_native_load_balancer.cloud_nlb.id
  managed                             = "yes"
  func_type                           = "GoTo"
  is_copy                             = "no"
  sequence_number                     = "0"
}

# Create a function node with type FW_ROUTED and relation to the 3rd party firewall
resource "aci_function_node" "function_node_fw" {
  l4_l7_service_graph_template_dn      = aci_function_node.function_node_nlb.l4_l7_service_graph_template_dn
  name                                 = "N1"
  func_template_type                   = "FW_ROUTED"
  relation_vns_rs_node_to_cloud_l_dev  = aci_cloud_l4_l7_third_party_device.third_party_fw.id
  l4_l7_device_interface_consumer_name = "trust"
  l4_l7_device_interface_provider_name = "untrust"
  managed                              = "no"
}

# Create L4-L7 Service Graph connection with template T1 and the first node N0.
resource "aci_connection" "sg1_t1-n0" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg1.id
  name                            = "CON0"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_l4_l7_service_graph_template.sg1.term_cons_dn,
    aci_function_node.function_node_nlb.conn_consumer_dn
  ]
}

# Create L4-L7 Service Graph connection with current node N0 and next node N1.
resource "aci_connection" "sg1_n0-n1" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg1.id
  name                            = "CON1"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_function_node.function_node_nlb.conn_provider_dn,
    aci_function_node.function_node_fw.conn_consumer_dn
  ]
}

# Create L4-L7 Service Graph connection with the last node N1 and template T2.
resource "aci_connection" "sg1_n1-t1" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg1.id
  name                            = "CON2"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_function_node.function_node_fw.conn_provider_dn,
    aci_l4_l7_service_graph_template.sg1.term_prov_dn
  ]
}

#  2. Create second L4-L7 Service Graph Template 'sg2' with type cloud
#    Add two nodes with additional parameters
#     - The first node 'N0' with type ADC_ONE_ARM and relation to the cloud native load balancer
#     - The second node 'N1' with type FW_ROUTED and relation to the 3rd party firewall
#    Create the connection between the templates ('T1 - consumer' and 'T2 - provider') and nodes ('N0' and 'N1').

resource "aci_l4_l7_service_graph_template" "sg2" {
  tenant_dn                         = data.aci_tenant.tf_tenant.id
  name                              = "tf_sg_2"
  l4_l7_service_graph_template_type = "cloud"
}

# Create a function node with type ADC_ONE_ARM and relation to the cloud native load balancer
# Additional parameters are added to the function node to demonstrate the different options available.
resource "aci_function_node" "function_node_nlb2" {
  l4_l7_service_graph_template_dn                = aci_l4_l7_service_graph_template.sg2.id
  name                                           = "N0"
  func_template_type                             = "ADC_ONE_ARM"
  routing_mode                                   = "Redirect"
  relation_vns_rs_node_to_cloud_l_dev            = aci_cloud_l4_l7_native_load_balancer.cloud_nlb.id
  managed                                        = "yes"
  l4_l7_device_interface_consumer_connector_type = "none"
  l4_l7_device_interface_provider_connector_type = "redir"
}

# Create a function node with type FW_ROUTED and relation to the 3rd party firewall
# Additional parameters are added to the function node to demonstrate the different options available.
resource "aci_function_node" "function_node_fw2" {
  l4_l7_service_graph_template_dn                         = aci_function_node.function_node_nlb2.l4_l7_service_graph_template_dn
  name                                                    = "N1"
  func_template_type                                      = "FW_ROUTED"
  relation_vns_rs_node_to_cloud_l_dev                     = aci_cloud_l4_l7_third_party_device.third_party_fw.id
  l4_l7_device_interface_consumer_name                    = "trust"
  l4_l7_device_interface_provider_name                    = "untrust"
  l4_l7_device_interface_consumer_connector_type          = "redir"
  l4_l7_device_interface_provider_connector_type          = "snat"
  l4_l7_device_interface_consumer_attachment_notification = "no"
  l4_l7_device_interface_provider_attachment_notification = "yes"
  managed                                                 = "no"
}

# Create L4-L7 Service Graph sg2 connection with template T1 and the first node N0.
resource "aci_connection" "sg2_t1-n0" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg2.id
  name                            = "CON0"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_l4_l7_service_graph_template.sg2.term_cons_dn,
    aci_function_node.function_node_nlb2.conn_consumer_dn
  ]
}

# Create L4-L7 Service Graph sg2 connection with current node N0 and next node N1.
resource "aci_connection" "sg2_n0-n1" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg2.id
  name                            = "CON1"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_function_node.function_node_nlb2.conn_provider_dn,
    aci_function_node.function_node_fw2.conn_consumer_dn
  ]
}

# Create L4-L7 Service Graph sg2 connection with the last node N1 and template T2.
resource "aci_connection" "sg2_n1-t1" {
  l4_l7_service_graph_template_dn = aci_l4_l7_service_graph_template.sg2.id
  name                            = "CON2"
  adj_type                        = "L3"
  conn_dir                        = "provider"
  conn_type                       = "external"
  direct_connect                  = "no"
  unicast_route                   = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_function_node.function_node_fw2.conn_provider_dn,
    aci_l4_l7_service_graph_template.sg2.term_prov_dn
  ]
}