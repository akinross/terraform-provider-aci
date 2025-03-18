---
subcategory: "Guides"
page_title: "Cisco ACI Multi-Site"
sidebar_current: "docs-aci-guides-cisco_aci_multisite"
description: |-
  How to setup Multi-Site with in ACI
---

# Cisco ACI Multi-Site #

## Introduction ##

Cisco ACI Multi-Site architecture provides interconnectivity between fabrics (APIC cluster domains). Each fabric represents a different site that is part of the ACI Multi-Site domain. This architecture allows for multi-tenant Layer 2 and Layer 3 network connectivity across sites, with extension of the policy domain throughout the Multi-Site domain. 

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_architecture.png"/>
  <figcaption><b>Figure 1:</b> Multi-Site Architecture</figcaption>
</figure>

In the ACI Multi-Site architecture, endpoint reachability information is exchanged between sites with a Multiprotocol-BGP (MP-BGP) Ethernet VPN (EVPN) control plane. The MP-BGP EVPN sessions are established between the spine nodes of sites. Virtual Extensible LAN (VXLAN) tunnels are created when a session is established between the interconnected spine nodes across the Inter-Site Network (ISN). This provides Layer 2 and Layer 3 network connectivity between endpoints.

The VXLAN Network Identifier (VNID) can be identified in several ways:

* **L2-VNI:** Bridge Domain (BD) for Layer 2 communication. 
* **L3-VNI:** Virtual Routing and Forwarding (VRF) instance for Layer 3 communication of the endpoint sourcing the traffic (Intra-VRF). 

Cisco Nexus Dashboard Orchestrator (NDO) offers a centralized interface for configuring and monitoring policies across different sites. While using NDO is the simplest way to connect these sites, it is not the only option. This guide will not go into detail about the Cisco ACI Multi-Site architecture. For more information, please see the [Cisco ACI Multi-Site Architecture White Paper](https://www.cisco.com/c/en/us/solutions/collateral/data-center-virtualization/application-centric-infrastructure/white-paper-c11-739609.html).

The aim of this guide is to explain the configuration that can be managed by the Terraform Provider ACI to provide connectivity between endpoints in different sites over an existing ISN.

  -> Not all features of NDO are currently available through the Terraform Provider. If you need additional functionality, please [submit a new issue](https://github.com/CiscoDevNet/terraform-provider-aci/issues/new/choose). This will allow us to track, evaluate, and prioritize the missing features accordingly.

## ACI Translation Objects ##

Endpoint Groups (EPGs) allow for policy (contract) enforcement of the source object. These type of ACI objects expose a class ID ([pcTag](https://www.cisco.com/c/en/us/support/docs/cloud-systems-management/application-policy-infrastructure-controller-apic/217302-application-centric-infrastructure-all.html)) attribute which is a unique identifier for that object within a site. Multi-Site communication is achieved by applying a a translation function (name-space normalization) before the traffic is forwarded inside the receiving site. This translation functionality ensures the usage of the correct local significant values inside the receiving site for the object exposing the v (pcTag), BD segment (L2-VNI) and VRF scope or segment (L3-VNI).

  -> The `scope` and `segment` attributes are identical for a VRF ([fvCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCtx/overview)). For the remainder of this guide we will refer to the `scope` attribute of a VRF.

In the ACI model the [fvSiteAssociated](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvSiteAssociated/overview) and the [fvRemoteId](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRemoteId/overview) objects are required for translation. The associated site object (fvSiteAssociated) assigns the ID of the local site to the EPG, BD, or VRF. The remote site(s) object (fvRemoteId) contain the translation values and are attached to the associated site object.

For VRFs the L3-VNI is required to be mapped. This is done by setting the `scope` attribute of the remote VRF in the `remotePcTag` attribute, the `pcTag` of the remote VRF in the `remoteCtxPcTag`, and the remote site ID in the `siteId` attribute.

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_mapping_vrf.png"/>
  <figcaption><b>Figure 2:</b> Multi-Site VRF Site Mapping</figcaption>
</figure>

For BDs the L2-VNI is required to be mapped. This is done by setting the `segment` attribute of the remote BD in the `remotePcTag` attribute, and the remote site ID in the `siteId` attribute.

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_mapping_bd.png"/>
  <figcaption><b>Figure 3:</b> Multi-Site BD Site Mapping</figcaption>
</figure>

For Application EPGs the class ID is required to be mapped. This is done by setting the `pcTag` attribute of the remote Application EPG in the `remotePcTag` attribute, and the remote site ID in the `siteId` attribute.

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_mapping_epg.png"/>
  <figcaption><b>Figure 4:</b> Multi-Site EPG Site Mapping</figcaption>
</figure>

## Multi-Site Resources & Datasources ##

Management of [fvSiteAssociated](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvSiteAssociated/overview) ([aci_associated_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/associated_site)) and [fvRemoteId](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRemoteId/overview) ([aci_remote_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/remote_site)) through dedicated resources is supported in the Terraform Provider ACI version 2.16. 

  -> In earlier versions the [aci_rest](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/rest.html) or [aci_rest_managed](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/rest_managed.html) can be used.

Remote sites can also be managed from the [aci_associated_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/associated_site) resource with the `remote_sites` attribute. This reduces the amount of resources required to be managed, and thus reduces the amount of REST API calls to the APIC controller.

  !> You should not manage [aci_remote_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/remote_site) resources in combincation with `remote_sites` attribute in the [aci_associated_site](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/associated_site) resource. Doing so will result in unexpected behaviour.

## Pre-Requirements Examples ##

In this section, we provide practical examples to help understand how to configure and manage multi-site resources using the Terraform Provider for ACI. These examples are designed to demonstrate the key concepts and configurations required to manage your ACI environment across multiple sites.

Each example follows a structured approach to ensure clarity and ease of understanding:

1. **main.tf**: This file contains the provider configuration, authentication details, and sets the alias for `site_1` and `site_2`. You will need to set the `username`, `password`, and `url` attributes for each site in this file.
2. **site_1.tf** and **site_2.tf**: These files contain the resource configuration for the two sites. They are pre-populated with variables and should not be modified for the examples.
3. **vars.tf**: Contains variables that are required to be set for the example to work.

In order for the examples to be used a few requirements and assumptions had to be made:

* An environment must be setup that has [terraform installed](https://developer.hashicorp.com/terraform/install) with reachability to both sites. Terraform installation can be validated with the following command. 

    ```
    terraform version
    ```

    The output should be similar to the output below.

    ```
    Terraform v1.9.8
    on linux_arm64

    Your version of Terraform is out of date! The latest version
    is 1.10.0. You can update by downloading from https://www.terraform.io/downloads.html
    ```

* The ISN must be configured, and both sites must be able to communicate over the ISN.
* One endpoint is connected to a VMM domain in `site_1` and another to a physical domain in `site_2`. This setup is chosen to demonstrate different configuration options for endpoints. If you do not have access to a endpoint in the VMM domain and a endpoint in the physical domain, you can manually adjust the examples to reflect the appropriate binding configuration for your specific setup.
* For each example, it is required to retrieve the site ID for both sites. The easiest approach is to navigate to "Intrasite/Intersite Profile - Fabric Ext Connection Policy" on one of the sites. As shown in Figure 5, the local site and the remote sites are displayed there.

    <figure>
      <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/site_id_from_ui.png"/>
      <figcaption><b>Figure 5:</b> Site ID from UI</figcaption>
    </figure>

    Another option is shown in Figure 6, where [fvFabricExtConnP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvFabricExtConnP/overview) is inspected using the ACI object store at https://[APIC_IP_FOR_EACH_SITE]/visore.html#/&class=fvFabricExtConnP.

    <figure>
      <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/site_id_from_object_store.png"/>
      <figcaption><b>Figure 6:</b> Site ID from ACI object store</figcaption>
    </figure>

## Example 1: [Layer 2 Communication without Flooding example](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/multisite/l2_without_flooding) ##

In this example all the objects are stretched across the sites, and Broadcast, Unknown-unicast, or Multicast (BUM) flooding is not allowed in the BD configuration across the sites. This is achieved by setting `enable_intersite_bum_traffic` to `no` (default when not provided in the Terraform Provider ACI) in the [aci_bridge_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/bridge_domain) configuration. Please refer to topology 1 for a visual representation of the objects. The objects are similar in each site, because all the objects are stretched.

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_topology_l2_without_flooding.png"/>
  <figcaption><b>Topology 1:</b> Layer 2 Communication without Flooding</figcaption>
</figure>

Ensure all the pre-requirements are met and execute the below steps:

1. Configure a interface with IP address from the BD subnet on both endpoints. The output should be similar to the output below.

```
user@site_2:~$ ip a
...
3: enp1s0f1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 7c:69:f6:90:0e:33 brd ff:ff:ff:ff:ff:ff
    inet 192.168.2.3/24 brd 192.168.2.255 scope global noprefixroute enp1s0f1
       valid_lft forever preferred_lft forever
    inet6 fe80::7e69:f6ff:fe90:e33/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
```

-> On the VMM domain endpoint the IP address information will be visible when the distributed port group is set for the endpoint.

2. Configure a route via the IP address of the BD subnet on both endpoints. The output should be similar to the output below, where the BD subnet IP address is 192.168.2.254.

```
user@site_2:~$ ip route
...
192.168.0.0/16 via 192.168.2.254 dev enp1s0f1 proto static metric 101 
192.168.2.0/24 dev enp1s0f1 proto kernel scope link src 192.168.2.3 metric 101 
```

3. Create a directory and download the files from the example on the environment where Terraform is installed 

```
mkdir l2_without_flooding && \
cd l2_without_flooding && \
curl \
-o main.tf  https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/examples/multisite/l2_without_flooding/main.tf \
-o site_1.tf https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/examples/multisite/l2_without_flooding/site_1.tf \
-o site_2.tf https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/examples/multisite/l2_without_flooding/site_2.tf \
-o variables.tf https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/examples/multisite/l2_without_flooding/variables.tf
```

4. Adjust the default values in the downloaded `variables.tf` (format of input should remain the same):

    * **bd_subnet_ip**: IP address with mask that should be used in the BD subnet configuration in `site_1` and `site_2`.
    * **site_1_id**: The site id that should be used in `site_1`.
    * **site_1_domain_target_dn**: The VMM domain that should be used in `site_1`.
    * **site_2_id**: The site id that should be used in `site_2`.
    * **site_2_domain_target_dn**: The physical domain that should be used in `site_2`.
    * **site_2_static_path_encapsulation**: The encapsulation that should be used for the `static path binding` in `site_2`.
    * **site_2_static_path_target_dn**: The interface that should be used for the `static path binding` in `site_2`.

5. Inside the `l2_without_flooding` directory execute a Terraform plan.

```
terraform plan
```

The output should end similar to the output below.

```
...
  # aci_vrf.shared_vrf_site_2 will be created
  + resource "aci_vrf" "shared_vrf_site_2" {
      + annotation                              = "orchestrator:terraform"
      + bd_enforced_enable                      = (known after apply)
      + description                             = (known after apply)
      + id                                      = (known after apply)
      + ip_data_plane_learning                  = (known after apply)
      + knw_mcast_act                           = (known after apply)
      + name                                    = "shared_vrf"
      + name_alias                              = (known after apply)
      + pc_enf_dir                              = (known after apply)
      + pc_enf_pref                             = (known after apply)
      + pc_tag                                  = (known after apply)
      + relation_fv_rs_bgp_ctx_pol              = (known after apply)
      + relation_fv_rs_ctx_to_ep_ret            = (known after apply)
      + relation_fv_rs_ctx_to_ext_route_tag_pol = (known after apply)
      + relation_fv_rs_ospf_ctx_pol             = (known after apply)
      + relation_fv_rs_vrf_validation_pol       = (known after apply)
      + scope                                   = (known after apply)
      + tenant_dn                               = (known after apply)
    }

Plan: 18 to add, 0 to change, 0 to destroy.

─────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to
take exactly these actions if you run "terraform apply" now.
```

6. Inside the `l2_without_flooding` directory apply the Terraform configuration.

```
terraform apply -auto-approve
```

The output should end similar to the output below.

```
...
  # aci_vrf.shared_vrf_site_2 will be created
  + resource "aci_vrf" "shared_vrf_site_2" {
      + annotation                              = "orchestrator:terraform"
      + bd_enforced_enable                      = (known after apply)
      + description                             = (known after apply)
      + id                                      = (known after apply)
      + ip_data_plane_learning                  = (known after apply)
      + knw_mcast_act                           = (known after apply)
      + name                                    = "shared_vrf"
      + name_alias                              = (known after apply)
      + pc_enf_dir                              = (known after apply)
      + pc_enf_pref                             = (known after apply)
      + pc_tag                                  = (known after apply)
      + relation_fv_rs_bgp_ctx_pol              = (known after apply)
      + relation_fv_rs_ctx_to_ep_ret            = (known after apply)
      + relation_fv_rs_ctx_to_ext_route_tag_pol = (known after apply)
      + relation_fv_rs_ospf_ctx_pol             = (known after apply)
      + relation_fv_rs_vrf_validation_pol       = (known after apply)
      + scope                                   = (known after apply)
      + tenant_dn                               = (known after apply)
    }

Plan: 18 to add, 0 to change, 0 to destroy.
aci_tenant.shared_tenant_site_1: Creating...
aci_tenant.shared_tenant_site_2: Creating...
aci_tenant.shared_tenant_site_2: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding]
aci_tenant.shared_tenant_site_1: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding]
aci_application_profile.shared_ap_site_1: Creating...
aci_vrf.shared_vrf_site_1: Creating...
aci_vrf.shared_vrf_site_2: Creating...
aci_application_profile.shared_ap_site_2: Creating...
aci_application_profile.shared_ap_site_1: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap]
aci_application_profile.shared_ap_site_2: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap]
aci_vrf.shared_vrf_site_2: Creation complete after 3s [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf]
aci_bridge_domain.shared_bd_site_2: Creating...
aci_vrf.shared_vrf_site_1: Creation complete after 4s [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf]
aci_associated_site.associated_site_of_shared_vrf_site_2: Creating...
aci_associated_site.associated_site_of_shared_vrf_site_1: Creating...
aci_bridge_domain.shared_bd_site_1: Creating...
aci_associated_site.associated_site_of_shared_vrf_site_2: Creation complete after 0s [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf/stAsc]
aci_bridge_domain.shared_bd_site_1: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd]
aci_subnet.shared_subnet_bd_site_1: Creating...
aci_application_epg.shared_epg_site_1: Creating...
aci_associated_site.associated_site_of_shared_vrf_site_1: Creation complete after 2s [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf/stAsc]
aci_bridge_domain.shared_bd_site_2: Creation complete after 3s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd]
aci_subnet.shared_subnet_bd_site_2: Creating...
aci_associated_site.associated_site_of_shared_bd_site_1: Creating...
aci_associated_site.associated_site_of_shared_bd_site_2: Creating...
aci_application_epg.shared_epg_site_2: Creating...
aci_associated_site.associated_site_of_shared_bd_site_1: Creation complete after 2s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/stAsc]
aci_application_epg.shared_epg_site_1: Creation complete after 3s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg]
aci_associated_site.associated_site_of_shared_bd_site_2: Creation complete after 2s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/stAsc]
aci_application_epg.shared_epg_site_2: Creation complete after 2s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg]
aci_associated_site.associated_site_of_shared_epg_site_2: Creating...
aci_associated_site.associated_site_of_shared_epg_site_1: Creating...
aci_associated_site.associated_site_of_shared_epg_site_1: Creation complete after 1s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg/stAsc]
aci_associated_site.associated_site_of_shared_epg_site_2: Creation complete after 2s [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg/stAsc]
aci_subnet.shared_subnet_bd_site_2: Creation complete after 4s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/subnet-[192.168.2.254/24]]
aci_subnet.shared_subnet_bd_site_1: Creation complete after 5s [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/subnet-[192.168.2.254/24]]

Apply complete! Resources: 18 added, 0 changed, 0 destroyed.
```

7. From the endpoint in `site_2` test verify ping reachability to the BD subnet IP address is successful.

```
user@site_2:~$ ping -c 5 192.168.2.254
PING 192.168.2.254 (192.168.2.254) 56(84) bytes of data.
64 bytes from 192.168.2.254: icmp_seq=1 ttl=64 time=0.243 ms
64 bytes from 192.168.2.254: icmp_seq=2 ttl=64 time=0.279 ms
64 bytes from 192.168.2.254: icmp_seq=3 ttl=64 time=0.286 ms
64 bytes from 192.168.2.254: icmp_seq=4 ttl=64 time=0.201 ms
64 bytes from 192.168.2.254: icmp_seq=5 ttl=64 time=0.209 ms

--- 192.168.2.254 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4090ms
rtt min/avg/max/mdev = 0.201/0.243/0.286/0.034 ms
```

7. From the endpoint in `site_2` test ping reachability to the endpoint in `site_1` is unsuccessful.

```
user@site_2:~$ ping -c 5 192.168.2.2
PING 192.168.2.2 (192.168.2.2) 56(84) bytes of data.
From 192.168.2.3 icmp_seq=1 Destination Host Unreachable
From 192.168.2.3 icmp_seq=2 Destination Host Unreachable
From 192.168.2.3 icmp_seq=3 Destination Host Unreachable
From 192.168.2.3 icmp_seq=4 Destination Host Unreachable
From 192.168.2.3 icmp_seq=5 Destination Host Unreachable
```

 -> If this ping is successful assure that the distributed port group `multisite_l2_without_flooding|shared_ap|shared_epg` is not set for the endpoint.

8. In vSphere Client, set the distributed port group `multisite_l2_without_flooding|shared_ap|shared_epg` for the endpoint.

9. From the endpoint in `site_2` test ping reachability to the endpoint in `site_1` is successful.

```
user@site_2:~$ ping -c 5 192.168.2.2
PING 192.168.2.2 (192.168.2.2) 56(84) bytes of data.
64 bytes from 192.168.2.2: icmp_seq=1 ttl=64 time=1036 ms
64 bytes from 192.168.2.2: icmp_seq=2 ttl=64 time=0.581 ms
64 bytes from 192.168.2.2: icmp_seq=3 ttl=64 time=0.401 ms
64 bytes from 192.168.2.2: icmp_seq=4 ttl=64 time=0.361 ms
64 bytes from 192.168.2.2: icmp_seq=5 ttl=64 time=0.331 ms

--- 192.168.2.2 ping statistics ---
5 packets transmitted, 5 received, 0% packet loss, time 4107ms
rtt min/avg/max/mdev = 0.331/207.504/1035.849/414.172 ms, pipe 2
```

10. Inside the `l2_without_flooding` directory clean-up the environment by destroying the Terrform configuration.

```
terraform destroy -auto-approve
```

The output should end similar to the output below.

```
...
  # aci_vrf.shared_vrf_site_2 will be destroyed
  - resource "aci_vrf" "shared_vrf_site_2" {
      - annotation                              = "orchestrator:terraform" -> null
      - bd_enforced_enable                      = "no" -> null
      - id                                      = "uni/tn-multisite_l2_without_flooding/ctx-shared_vrf" -> null
      - ip_data_plane_learning                  = "enabled" -> null
      - knw_mcast_act                           = "permit" -> null
      - name                                    = "shared_vrf" -> null
      - pc_enf_dir                              = "ingress" -> null
      - pc_enf_pref                             = "enforced" -> null
      - pc_tag                                  = "16386" -> null
      - scope                                   = "2392066" -> null
      - tenant_dn                               = "uni/tn-multisite_l2_without_flooding" -> null
        # (8 unchanged attributes hidden)
    }

Plan: 0 to add, 0 to change, 18 to destroy.
aci_subnet.shared_subnet_bd_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/subnet-[192.168.2.254/24]]
aci_associated_site.associated_site_of_shared_bd_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/stAsc]
aci_associated_site.associated_site_of_shared_vrf_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf/stAsc]
aci_associated_site.associated_site_of_shared_vrf_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf/stAsc]
aci_associated_site.associated_site_of_shared_bd_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/stAsc]
aci_associated_site.associated_site_of_shared_epg_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg/stAsc]
aci_associated_site.associated_site_of_shared_epg_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg/stAsc]
aci_subnet.shared_subnet_bd_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd/subnet-[192.168.2.254/24]]
aci_associated_site.associated_site_of_shared_vrf_site_1: Destruction complete after 1s
aci_subnet.shared_subnet_bd_site_2: Destruction complete after 1s
aci_associated_site.associated_site_of_shared_bd_site_1: Destruction complete after 1s
aci_associated_site.associated_site_of_shared_bd_site_2: Destruction complete after 1s
aci_associated_site.associated_site_of_shared_epg_site_1: Destruction complete after 2s
aci_associated_site.associated_site_of_shared_vrf_site_2: Destruction complete after 2s
aci_subnet.shared_subnet_bd_site_1: Destruction complete after 2s
aci_associated_site.associated_site_of_shared_epg_site_2: Destruction complete after 2s
aci_application_epg.shared_epg_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg]
aci_application_epg.shared_epg_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap/epg-shared_epg]
aci_application_epg.shared_epg_site_1: Destruction complete after 0s
aci_application_epg.shared_epg_site_2: Destruction complete after 0s
aci_application_profile.shared_ap_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap]
aci_application_profile.shared_ap_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/ap-shared_ap]
aci_bridge_domain.shared_bd_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd]
aci_bridge_domain.shared_bd_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/BD-shared_bd]
aci_application_profile.shared_ap_site_1: Destruction complete after 0s
aci_bridge_domain.shared_bd_site_1: Destruction complete after 0s
aci_vrf.shared_vrf_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf]
aci_application_profile.shared_ap_site_2: Destruction complete after 0s
aci_bridge_domain.shared_bd_site_2: Destruction complete after 0s
aci_vrf.shared_vrf_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding/ctx-shared_vrf]
aci_vrf.shared_vrf_site_1: Destruction complete after 0s
aci_tenant.shared_tenant_site_1: Destroying... [id=uni/tn-multisite_l2_without_flooding]
aci_vrf.shared_vrf_site_2: Destruction complete after 1s
aci_tenant.shared_tenant_site_2: Destroying... [id=uni/tn-multisite_l2_without_flooding]
aci_tenant.shared_tenant_site_1: Destruction complete after 1s
aci_tenant.shared_tenant_site_2: Destruction complete after 0s

Destroy complete! Resources: 18 destroyed.
```


## Example 2: [Intra VRF example](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/multisite/l3_intra_vrf) ##

In this example the source EPG and destination EPG belong to different bridge domains mapped to the same VRF instance inside the same tenant. The tenant, contract and VRF instance are stretched across the two sites, and MP-BGP EVPN allows the exchange of host routing information, enabling intersite communication.

<figure>
  <img src="https://raw.githubusercontent.com/akinross/terraform-provider-aci/refs/heads/multisite_guide/docs/images/multisite_topology_intra_vrf.png"/>
  <figcaption><b>Topology 2:</b> Intra VRF Topology</figcaption>
</figure>

Configuring a contract between `epg_site_1` and `epg_site_2` in NDO would result in the creation of the corresponding shadow objects in the two sites. From each site NDO retrieves the specific identifiers (pcTag, L2VNI, and L3VNI) assigned to the local and shadow objects, and instructs the site to program translation rules in the local spines. The end result is that the configured policy can then correctly be applied on the leaf node before sending the traffic to the destination endpoint.

The Terraform Provider" ACI does not contain the logic of NDO, and thus the shadow objects and all the translation objects need to be specified in the Terraform configuration files. The shadow objects can be hidden by setting the annotation attribute of the object to `orchestrator:msc-shadow:yes`. This is not needed for the interconnectivity between sites but in the example it is included to align with the UI behaviour of NDO.

... in depth details of what is going on ...

## Contribute ##

.. details on how contributions can be provided for additional examples ...
