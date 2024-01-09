package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - dhcp_config.filters
// - dhcp_config.filters_v6
// - dhcp_config.ignore_items
// - dhcp_options
// - vendor_specific_option
// - inheritance_sources - Currently inheritance sources is always nil

func TestAccIpSpaceResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceBasicConfig("ip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "ip_space_name"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "client"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost"),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "-"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-zA-Z0-9_.]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpSpaceBasicConfig("ip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					testAccCheckIpSpaceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpSpaceResource_AsmConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_asm_config"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceAsmConfig("ip_space_name", 70, true, true, 12, 40, "count", 40, 30, 30, "2020-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "70"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "12"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "count"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.history", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2020-01-10T10:11:22Z")),
			},
			// Update and Read
			{
				Config: testAccIpSpaceAsmConfig("ip_space_name", 80, false, false, 10, 50, "percent", 50, 10, 10, "2021-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "80"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "false"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "50"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "percent"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.history", "50"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2021-01-10T10:11:22Z"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_comment"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceComment("ip_space_name", "some comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceComment("ip_space_name", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsClientUpdate(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_client_update"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsClientUpdate("ip_space_name", "server"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "server"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsClientUpdate("ip_space_name", "over_client_update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "over_client_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsConflictResolutionMode(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_conflict_resolution_mode"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsConflictResolutionMode("ip_space_name", false, "check_exists_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_exists_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsConflictResolutionMode("ip_space_name", true, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsDomain(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_domain"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsDomain("ip_space_name", "abc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "abc"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsDomain("ip_space_name", "xyz"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "xyz"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsGenerateName(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_generate_name"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsGenerateName("ip_space_name", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsGenerateName("ip_space_name", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsGeneratedPrefix(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_generated_prefix"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsGeneratedPrefix("ip_space_name", "host-prefix"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-prefix"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsGeneratedPrefix("ip_space_name", "host-another-prefix"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-another-prefix"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsSendUpdates(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_send_updates"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsSendUpdates("ip_space_name", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsSendUpdates("ip_space_name", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsTtlPercent(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_ttl_percent"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsTtlPercent("ip_space_name", "20"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsTtlPercent("ip_space_name", "40"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "40"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsUpdateOnRenew(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_update_on_renew"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsUpdateOnRenew("ip_space_name", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsUpdateOnRenew("ip_space_name", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DhcpConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_dhcp_config"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDhcpConfig("ip_space_name", 1000, 2000, true, true, true, 50, 60),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time", "1000"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time_v6", "2000"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "50"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDhcpConfig("ip_space_name", 1500, 2500, false, false, false, 55, 65),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time", "1500"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time_v6", "2500"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "55"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "65"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HeaderOptionFilename(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_filename"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionFilename("ip_space_name", "HEADER_OPTION_FILEip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILEip_space_name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionFilename("ip_space_name", "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_server_address"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionServerAddress("ip_space_name", "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionServerAddress("ip_space_name", "12.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_server_name"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionServerName("ip_space_name", "HEADER_OPTION_SERVER_ip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_ip_space_name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionServerName("ip_space_name", "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteChar(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_char"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteChar("ip_space_name", "+"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "+"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteChar("ip_space_name", "/"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "/"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteEnabled(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_enabled"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteEnabled("ip_space_name", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteEnabled("ip_space_name", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteRegex(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_regex"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteRegex("ip_space_name", "[^a-z]"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z]"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteRegex("ip_space_name", "[^0-9]"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^0-9]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_tags"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceTags("ip_space_name", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceTags("ip_space_name", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpSpaceExists(ctx context.Context, resourceName string, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceRead(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckIpSpaceDestroy(ctx context.Context, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceRead(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckIpSpaceDisappears(ctx context.Context, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpSpaceBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test" {
    name = %q
}
`, name)
}

func testAccIpSpaceAsmConfig(name string, asmThreshold int, enable, enableNotification bool, forecastPeriod, growthFactor int, growthType string, history, minTotal, minUnused int, reenableDate string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_asm_config" {
    name = %q
    asm_config = {
		asm_threshold = %d
		enable = %t
		enable_notification = %t
		forecast_period = %d
		growth_factor = %d
		growth_type = %q
		history = %d
		min_total = %d
		min_unused = %d
		reenable_date = %q
	}
}
`, name, asmThreshold, enable, enableNotification, forecastPeriod, growthFactor, growthType, history, minTotal, minUnused, reenableDate)
}

func testAccIpSpaceComment(name, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpSpaceDdnsClientUpdate(name, ddnsClientUpdate string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_client_update" {
    name = %q
    ddns_client_update = %q
}
`, name, ddnsClientUpdate)
}

func testAccIpSpaceDdnsConflictResolutionMode(name string, ddnsUseConflictResolution bool, ddnsConflictResolutionMode string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_conflict_resolution_mode" {
    name = %q
	ddns_use_conflict_resolution = %t
    ddns_conflict_resolution_mode = %q
}
`, name, ddnsUseConflictResolution, ddnsConflictResolutionMode)
}

func testAccIpSpaceDdnsDomain(name, ddnsDomain string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_domain" {
    name = %q
    ddns_domain = %q
}
`, name, ddnsDomain)
}

func testAccIpSpaceDdnsGenerateName(name string, ddnsGenerateName bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_generate_name" {
    name = %q
    ddns_generate_name = %t
}
`, name, ddnsGenerateName)
}

func testAccIpSpaceDdnsGeneratedPrefix(name, ddnsGeneratedPrefix string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_generated_prefix" {
    name = %q
    ddns_generated_prefix = %q
}
`, name, ddnsGeneratedPrefix)
}

func testAccIpSpaceDdnsSendUpdates(name string, ddnsSendUpdates bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_send_updates" {
    name = %q
    ddns_send_updates = %t
}
`, name, ddnsSendUpdates)
}

func testAccIpSpaceDdnsTtlPercent(name, ddnsTtlPercent string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_ttl_percent" {
    name = %q
    ddns_ttl_percent = %s
}
`, name, ddnsTtlPercent)
}

func testAccIpSpaceDdnsUpdateOnRenew(name string, ddnsUpdateOnRenew bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_update_on_renew" {
    name = %q
    ddns_update_on_renew = %t
}
`, name, ddnsUpdateOnRenew)
}

func testAccIpSpaceDhcpConfig(name string,
	abandonedReclaimTime, abandonedReclaimTimeV6 int,
	allowUnknown, allowUnknownV6, ignoreClientUid bool,
	leaseTime, leaseTimeV6 int) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_config" {
    name = %q
    dhcp_config = {
		abandoned_reclaim_time = %d
		abandoned_reclaim_time_v6 = %d
		allow_unknown = %t
		allow_unknown_v6 = %t
		ignore_client_uid = %t
		lease_time = %d
		lease_time_v6 = %d
	}
}
`, name, abandonedReclaimTime, abandonedReclaimTimeV6, allowUnknown, allowUnknownV6, ignoreClientUid, leaseTime, leaseTimeV6)
}

func testAccIpSpaceHeaderOptionFilename(name, headerOptionFilename string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_filename" {
    name = %q
    header_option_filename = %q
}
`, name, headerOptionFilename)
}

func testAccIpSpaceHeaderOptionServerAddress(name, headerOptionServerAddress string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_server_address" {
    name = %q
    header_option_server_address = %q
}
`, name, headerOptionServerAddress)
}

func testAccIpSpaceHeaderOptionServerName(name, headerOptionServerName string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_server_name" {
    name = %q
    header_option_server_name = %q
}
`, name, headerOptionServerName)
}

func testAccIpSpaceHostnameRewriteChar(name, hostnameRewriteChar string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_char" {
    name = %q
    hostname_rewrite_char = %q
}
`, name, hostnameRewriteChar)
}

func testAccIpSpaceHostnameRewriteEnabled(name string, hostnameRewriteEnabled bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_enabled" {
    name = %q
    hostname_rewrite_enabled = %t
}
`, name, hostnameRewriteEnabled)
}

func testAccIpSpaceHostnameRewriteRegex(name, hostnameRewriteRegex string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_regex" {
    name = %q
    hostname_rewrite_regex = %q
}
`, name, hostnameRewriteRegex)
}

func testAccIpSpaceTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
