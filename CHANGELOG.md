# Changelog

## v1.3.0 (Jun 24, 2024)

FEATURES:
* **New feature:** `Support Default Tags` ([#124](https://github.com/infobloxopen/terraform-provider-bloxone/pull/124))
* **New feature:** `Assign Service Instance in DHCP Config Profile` ([#121](https://github.com/infobloxopen/terraform-provider-bloxone/pull/121))
* **New Module:** `Implementation for anycast module` ([#120](https://github.com/infobloxopen/terraform-provider-bloxone/pull/120))

## v1.2.0 (May 28, 2024)

FEATURES:
* **New Resource and Data Source:** `bloxone_td_application_filter`, `bloxone_td_application_filters` ([#113](https://github.com/infobloxopen/terraform-provider-bloxone/pull/113))
* **New Resource and Data Source:** `bloxone_td_category_filter`, `bloxone_td_category_filters` ([#113](https://github.com/infobloxopen/terraform-provider-bloxone/pull/113))
* **New Resource and Data Source:** `bloxone_td_custom_redirect`, `bloxone_td_custom_redirects` ([#113](https://github.com/infobloxopen/terraform-provider-bloxone/pull/113))
* **New Data Source:** `bloxone_td_content_categories` ([#113](https://github.com/infobloxopen/terraform-provider-bloxone/pull/113))
* **New Data Source:** `bloxone_td_threat_feeds` ([#113](https://github.com/infobloxopen/terraform-provider-bloxone/pull/113))

FIXES:
* Fixed inconsistent result error when applying `bloxone_td_internal_domain_list` and `bloxone_dfp_service` caused due to changes in the list order ([#114](https://github.com/infobloxopen/terraform-provider-bloxone/pull/114))  

## v1.1.0 (May 10, 2024)

NOTES:
* Upgraded to Go 1.22 ([#107](https://github.com/infobloxopen/terraform-provider-bloxone/pull/107))

FEATURES:
* **New Data Source:** `bloxone_td_pop_regions` ([#83](https://github.com/infobloxopen/terraform-provider-bloxone/pull/83))
* **New Resource and Data Source:** `bloxone_td_internal_domain_list`, `bloxone_td_internal_domain_lists` ([#83](https://github.com/infobloxopen/terraform-provider-bloxone/pull/83))
* **New Resource and Data Source:** `bloxone_td_access_code`, `bloxone_td_access_codes` ([#90](https://github.com/infobloxopen/terraform-provider-bloxone/pull/90))
* **New Resource and Data Source:** `bloxone_td_named_list`, `bloxone_td_named_lists` ([#90](https://github.com/infobloxopen/terraform-provider-bloxone/pull/90))
* **New Resource and Data Source:** `bloxone_td_network_list`, `bloxone_td_network_lists` ([#90](https://github.com/infobloxopen/terraform-provider-bloxone/pull/90))
* **New Resource and Data Source:** `bloxone_td_security_policy`, `bloxone_td_security_policies` ([#94](https://github.com/infobloxopen/terraform-provider-bloxone/pull/94))
* **New Resource and Data Source:** `bloxone_dfp_service`, `bloxone_dfp_services` ([#102](https://github.com/infobloxopen/terraform-provider-bloxone/pull/102))
* **New Resource and Data Source:** `bloxone_anycast_config`, `bloxone_anycast_configs` ([#92](https://github.com/infobloxopen/terraform-provider-bloxone/pull/92))
* **New Resource:** `bloxone_anycast_host` ([#95](https://github.com/infobloxopen/terraform-provider-bloxone/pull/95))

ENHANCEMENTS:
* Added example for DHCP HA group with Anycast ([#111](https://github.com/infobloxopen/terraform-provider-bloxone/pull/111))
* Added `dfp` and `anycast` as valid service types for modules ([#111](https://github.com/infobloxopen/terraform-provider-bloxone/pull/111))
* Added `echo_client_id` to `bloxone_ipam_ip_space.dhcp_config`, `bloxone_dhcp_server.dhcp_config` ([#93](https://github.com/infobloxopen/terraform-provider-bloxone/pull/93))
* Added example for ddns in `bloxone_dhcp_server` ([#88](https://github.com/infobloxopen/terraform-provider-bloxone/pull/88))

FIXES:
* Updating IP space fails with error - "The value of inheritance action field is not valid" ([#93](https://github.com/infobloxopen/terraform-provider-bloxone/pull/93))
* `abandonded_reclaim_time`, `abandoned_reclaim_time_v6` in `bloxone_ipam_subnet` and `bloxone_ipam_address_block` changed from "optional" to "read only". ([#96](https://github.com/infobloxopen/terraform-provider-bloxone/pull/96))
* Next available subnet/AB returns same block when count is used ([#100](https://github.com/infobloxopen/terraform-provider-bloxone/pull/100))

## v1.0.1 (March 20, 2024)

FEATURES:
* `bloxone_infra_host_azure` module for provisioning BloxOne host in Azure ([#67](https://github.com/infobloxopen/terraform-provider-bloxone/pull/67))
* `bloxone_infra_host_gcp` module for provisioning BloxOne host in GCP ([#82](https://github.com/infobloxopen/terraform-provider-bloxone/pull/82))

FIXES:
* Error when using options in `bloxone_dns_record` resources ([#86](https://github.com/infobloxopen/terraform-provider-bloxone/pull/86))
* Data source filter doesn't work for number fields ([#81](https://github.com/infobloxopen/terraform-provider-bloxone/pull/81))

## v1.0.0 (February 29, 2024)

First stable release of the BloxOne Terraform Provider

NOTES: 
* Added quickstart guides for DNS and DHCP ([#68](https://github.com/infobloxopen/terraform-provider-bloxone/pull/68), [#73](https://github.com/infobloxopen/terraform-provider-bloxone/pull/73))

FEATURES:
* **New Resource and Data Source:** `bloxone_dns_acl`, `bloxone_dns_acls` ([#64](https://github.com/infobloxopen/terraform-provider-bloxone/pull/64))

ENHANCEMENTS:
* Next available address block ([#63](https://github.com/infobloxopen/terraform-provider-bloxone/pull/63))
* Added retry to DNS and DHCP hosts data sources ([#71](https://github.com/infobloxopen/terraform-provider-bloxone/pull/71))
* Added acceptance test for DHCP options in various resources that supports it ([#65](https://github.com/infobloxopen/terraform-provider-bloxone/pull/65))

FIXES:
* AMI search in `bloxone_infra_host_aws` module ([#72](https://github.com/infobloxopen/terraform-provider-bloxone/pull/72))
* Make `fqdn` optional in `bloxone_dns_forward_zone.external_forwarders` ([#76](https://github.com/infobloxopen/terraform-provider-bloxone/pull/76))
* Unable to unset optional string fields ([#77](https://github.com/infobloxopen/terraform-provider-bloxone/pull/77))
* Acceptance Tests ([#79](https://github.com/infobloxopen/terraform-provider-bloxone/pull/79), [#80](https://github.com/infobloxopen/terraform-provider-bloxone/pull/80))

## v0.1.0 (January 31, 2024)

NOTES:
* Added migration guide to help users migrate from B1DDI provider ([#61](https://github.com/infobloxopen/terraform-provider-bloxone/pull/61))

FEATURES:
* **New Resource and Data Source:** `bloxone_infra_join_token`, `bloxone_infra_join_tokens` ([#17](https://github.com/infobloxopen/terraform-provider-bloxone/pull/17))
* **New Resource and Data Source:** `bloxone_infra_host`, `bloxone_infra_host` ([#18](https://github.com/infobloxopen/terraform-provider-bloxone/pull/18))
* **New Resource and Data Source:** `bloxone_infra_service`, `bloxone_infra_services` ([#19](https://github.com/infobloxopen/terraform-provider-bloxone/pull/19))
* **New Resource and Data Source:** `bloxone_ipam_ip_space`, `bloxone_ipam_ip_spaces` ([#16](https://github.com/infobloxopen/terraform-provider-bloxone/pull/16))
* **New Resource and Data Source:** `bloxone_ipam_address_block`, `bloxone_ipam_address_blocks` ([#24](https://github.com/infobloxopen/terraform-provider-bloxone/pull/24))
* **New Resource and Data Source:** `bloxone_ipam_subnet`, `bloxone_ipam_subnets` ([#21](https://github.com/infobloxopen/terraform-provider-bloxone/pull/21))
* **New Resource and Data Source:** `bloxone_ipam_range`, `bloxone_ipam_ranges` ([#25](https://github.com/infobloxopen/terraform-provider-bloxone/pull/25))
* **New Resource and Data Source:** `bloxone_ipam_host`, `bloxone_ipam_hosts` ([#20](https://github.com/infobloxopen/terraform-provider-bloxone/pull/20))
* **New Resource and Data Source:** `bloxone_ipam_address`, `bloxone_ipam_addresses` ([#23](https://github.com/infobloxopen/terraform-provider-bloxone/pull/23))
* **New Resource and Data Source:** `bloxone_dhcp_fixed_address`, `bloxone_dhcp_fixed_addresss` ([#28](https://github.com/infobloxopen/terraform-provider-bloxone/pull/28))
* **New Resource and Data Source:** `bloxone_dhcp_ha_group`, `bloxone_dhcp_ha_groups` ([#36](https://github.com/infobloxopen/terraform-provider-bloxone/pull/36))
* **New Resource and Data Source:** `bloxone_dhcp_server`, `bloxone_dhcp_servers` ([#37](https://github.com/infobloxopen/terraform-provider-bloxone/pull/37))
* **New Resource and Data Source:** `bloxone_dhcp_option_space`, `bloxone_dhcp_option_spaces` ([#46](https://github.com/infobloxopen/terraform-provider-bloxone/pull/46))
* **New Resource and Data Source:** `bloxone_dhcp_option_code`, `bloxone_dhcp_option_codes` ([#46](https://github.com/infobloxopen/terraform-provider-bloxone/pull/46))
* **New Resource and Data Source:** `bloxone_dhcp_option_group`, `bloxone_dhcp_option_groups` ([#46](https://github.com/infobloxopen/terraform-provider-bloxone/pull/46))
* **New Data Source:** `bloxone_dhcp_hosts` ([#28](https://github.com/infobloxopen/terraform-provider-bloxone/pull/28))
* **New Data Source:** `bloxone_dns_hosts` ([#26](https://github.com/infobloxopen/terraform-provider-bloxone/pull/26))
* **New Resource and Data Source:** `bloxone_dns_view`, `bloxone_dns_views` ([#22](https://github.com/infobloxopen/terraform-provider-bloxone/pull/22))
* **New Resource and Data Source:** `bloxone_dns_auth_zone`, `bloxone_dns_auth_zones` ([#27](https://github.com/infobloxopen/terraform-provider-bloxone/pull/27))
* **New Resource and Data Source:** `bloxone_dns_forward_zone`, `bloxone_dns_forward_zones` ([#34](https://github.com/infobloxopen/terraform-provider-bloxone/pull/34))
* **New Resource and Data Source:** `bloxone_dns_delegation`, `bloxone_dns_delegations` ([#32](https://github.com/infobloxopen/terraform-provider-bloxone/pull/32))
* **New Resource and Data Source:** `bloxone_dns_auth_nsg`, `bloxone_dns_auth_nsgs` ([#30](https://github.com/infobloxopen/terraform-provider-bloxone/pull/30))
* **New Resource and Data Source:** `bloxone_dns_forward_nsg`, `bloxone_dns_forward_nsgs` ([#31](https://github.com/infobloxopen/terraform-provider-bloxone/pull/31))
* **New Resource and Data Source:** `bloxone_dns_record`, `bloxone_dns_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_a_record`, `bloxone_dns_a_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_aaaa_record`, `bloxone_dns_aaaa_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_caa_record`, `bloxone_dns_caa_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_cname_record`, `bloxone_dns_cname_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_delegation_record`, `bloxone_dns_delegation_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_dname_record`, `bloxone_dns_dname_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_https_record`, `bloxone_dns_https_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_mx_record`, `bloxone_dns_mx_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_naptr_record`, `bloxone_dns_naptr_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_ns_record`, `bloxone_dns_ns_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_ptr_record`, `bloxone_dns_ptr_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_srv_record`, `bloxone_dns_srv_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_svcb_record`, `bloxone_dns_svcb_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_dns_txt_record`, `bloxone_dns_txt_records` ([#41](https://github.com/infobloxopen/terraform-provider-bloxone/pull/41))
* **New Resource and Data Source:** `bloxone_keys_tsig`, `bloxone_keys_tsigs` ([#33](https://github.com/infobloxopen/terraform-provider-bloxone/pull/33))
* **New Data Source:** `bloxone_keys_kerberos` ([#40](https://github.com/infobloxopen/terraform-provider-bloxone/pull/40))
* `bloxone_infra_host_aws` module for provisioning BloxOne host in AWS ([#53](https://github.com/infobloxopen/terraform-provider-bloxone/pull/54))
