package ipam_test

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "strings"
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
    "github.com/hashicorp/terraform-plugin-testing/terraform"
    "github.com/infobloxopen/bloxone-go-client/ipam"
    "github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

// TODO: add tests for the following
// Test to be enabled once DhcpOptions is implemented
// Inheritance sources - After _inherit support is added

func TestAccAddressBlockResource_basic(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockBasicConfig("192.168.0.0", "16"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    // TODO: check and validate these
                    resource.TestCheckResourceAttr(resourceName, "address", "192.168.0.0"),
                    resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.test", "id"),
                    // Test Read Only fields
                    resource.TestCheckResourceAttrSet(resourceName, "asm_config.%"),
                    resource.TestCheckResourceAttrSet(resourceName, "created_at"),
                    resource.TestCheckResourceAttrSet(resourceName, "id"),
                    resource.TestCheckResourceAttrSet(resourceName, "protocol"),
                    resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
                    resource.TestCheckResourceAttrSet(resourceName, "usage.#"),
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

func TestAccAddressBlockResource_disappears(t *testing.T) {
    resourceName := "bloxone_ipam_address_block.test"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        CheckDestroy:             testAccCheckAddressBlockDestroy(context.Background(), &v),
        Steps: []resource.TestStep{
            {
                Config: testAccAddressBlockBasicConfig("192.168.0.0", "16"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    testAccCheckAddressBlockDisappears(context.Background(), &v),
                ),
                ExpectNonEmptyPlan: true,
            },
        },
    })
}

func TestAccAddressBlockResource_Address(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test"
    var v1, v2 ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockBasicConfig("192.168.0.0", "16"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v1),
                    resource.TestCheckResourceAttr(resourceName, "address", "192.168.0.0"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockBasicConfig("10.0.0.0", "16"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockDestroy(context.Background(), &v1),
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v2),
                    resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_AsmConfig(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_asm_config"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockAsmConfig("10.0.0.0", "16", 70, true, true, 12, 40, "count", 40, 30, 30, "2020-01-10T10:11:22Z"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "70"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "true"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "true"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "12"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "40"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "count"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.history", "40"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "30"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "30"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2020-01-10T10:11:22Z"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockAsmConfig("10.0.0.0", "16", 90, false, false, 14, 60, "count", 40, 60, 50, "2020-01-10T10:11:22Z"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "90"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "false"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "false"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "14"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "60"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "count"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.history", "40"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "60"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "50"),
                    resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2020-01-10T10:11:22Z"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_Cidr(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_cidr"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockCidr("192.168.0.0", "16"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockCidr("192.168.0.0", "24"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_Comment(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_comment"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockComment("192.168.0.0", "16", "This address block is created through Terraform"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "comment", "This address block is created through Terraform"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockComment("192.168.0.0", "16", "This address block was created through Terraform"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "comment", "This address block was created through Terraform"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsClientUpdate(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_client_update"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsClientUpdate("192.168.0.0", "16", "client"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "client"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsClientUpdate("192.168.0.0", "16", "over_no_update"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "over_no_update"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsDomain(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_domain"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsDomain("192.168.0.0", "16", "test.com"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_domain", "test.com"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsDomain("192.168.0.0", "16", "test123.com"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_domain", "test123.com"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsGenerateName(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_generate_name"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsGenerateName("192.168.0.0", "16", "false"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsGenerateName("192.168.0.0", "16", "true"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "true"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsGeneratedPrefix(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_generated_prefix"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsGeneratedPrefix("192.168.0.0", "16", "ut"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "ut"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsGeneratedPrefix("192.168.0.0", "16", "ut-ut"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "ut-ut"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsSendUpdates(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_send_updates"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsSendUpdates("192.168.0.0", "16", "true"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsSendUpdates("192.168.0.0", "16", "false"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsTtlPercent(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_ttl_percent"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsTtlPercent("192.168.0.0", "16", "25"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "25"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsTtlPercent("192.168.0.0", "16", "75"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "75"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsUpdateOnRenew(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_update_on_renew"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsUpdateOnRenew("192.168.0.0", "16", "false"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsUpdateOnRenew("192.168.0.0", "16", "true"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "true"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DdnsUseConflictResolution(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_ddns_use_conflict_resolution"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDdnsUseConflictResolution("192.168.0.0", "16", "true"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDdnsUseConflictResolution("192.168.0.0", "16", "false"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_DhcpConfig(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_dhcp_config"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockDhcpConfig("192.168.0.0", "16", true, true, true, 50, 60),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "true"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "true"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "50"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "60"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockDhcpConfig("192.168.0.0", "16", false, true, false, 150, 160),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "false"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "false"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "150"),
                    resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "160"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HeaderOptionFilename(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_header_option_filename"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHeaderOptionFilename("192.168.0.0", "16", "testfile"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_filename", "testfile"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHeaderOptionFilename("192.168.0.0", "16", "testfile1"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_filename", "testfile1"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HeaderOptionServerAddress(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_header_option_server_address"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHeaderOptionServerAddress("192.168.0.0", "16", "1.1.1.1"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "1.1.1.1"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHeaderOptionServerAddress("192.168.0.0", "16", "2.2.2.2"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "2.2.2.2"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HeaderOptionServerName(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_header_option_server_name"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHeaderOptionServerName("192.168.0.0", "16", "test"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "test"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHeaderOptionServerName("192.168.0.0", "16", "test-1"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "test-1"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HostnameRewriteChar(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_hostname_rewrite_char"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHostnameRewriteChar("192.168.0.0", "16", "a"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "a"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHostnameRewriteChar("192.168.0.0", "16", "c"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "c"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HostnameRewriteEnabled(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_hostname_rewrite_enabled"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHostnameRewriteEnabled("192.168.0.0", "16", "true"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "true"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHostnameRewriteEnabled("192.168.0.0", "16", "false"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_HostnameRewriteRegex(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_hostname_rewrite_regex"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockHostnameRewriteRegex("192.168.0.0", "16", "[^a-z]"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z]"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockHostnameRewriteRegex("192.168.0.0", "16", "[^g-hG-H0-9_.]"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^g-hG-H0-9_.]"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_InheritanceSources(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_inheritance_sources"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockInheritanceSources("192.168.0.0", "16", "inherit"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_enabled.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_hostname_block.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_block.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "inherit"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.hostname_rewrite_block.action", "inherit"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockInheritanceSources("192.168.0.0", "16", "override"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_hostname_block.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_block.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "override"),
                    resource.TestCheckResourceAttr(resourceName, "inheritance_sources.hostname_rewrite_block.action", "override"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_Name(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_name"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockName("192.168.0.0", "16", "test_name"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "name", "test_name"),
                ),
            },
            // Update and read
            {
                Config: testAccAddressBlockName("192.168.0.0", "16", "test_name_1"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v),
                    resource.TestCheckResourceAttr(resourceName, "name", "test_name_1"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_Space(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_space"
    var v1, v2 ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockSpace("192.168.0.0", "16", "bloxone_ipam_ip_space.one"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v1),
                    resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.one", "id"),
                ),
            },
            // Update and Read
            {
                Config: testAccAddressBlockSpace("192.168.0.0", "16", "bloxone_ipam_ip_space.two"),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockDestroy(context.Background(), &v1),
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v2),
                    resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.two", "id"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_Tags(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_tags"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockTags("192.168.0.0", "16", map[string]string{
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
                PreConfig: func() {
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v)
                },
                Config: testAccAddressBlockTags("192.168.0.0", "16", map[string]string{
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

func TestAccAddressBlockResource_WithDefaultTags(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_tags"
    var v ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockWithDefaultTags("192.168.0.0", "16", map[string]string{
                    "tag1": "value1",
                    "tag2": "value2",
                }),
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
                    resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
                    acctest.VerifyDefaultTag(resourceName),
                ),
            },
            // Update and Read
            {
                PreConfig: func() {
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v)
                },
                Config: testAccAddressBlockWithDefaultTags("192.168.0.0", "16", map[string]string{
                    "tag2": "value2changed",
                    "tag3": "value3",
                }),
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
                    resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
                    acctest.VerifyDefaultTag(resourceName),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func TestAccAddressBlockResource_NextAvailable_AddressBlock(t *testing.T) {
    var resourceName = "bloxone_ipam_address_block.test_next_available"
    var v1 ipam.IpamsvcAddressBlock
    var v2 ipam.IpamsvcAddressBlock

    resource.Test(t, resource.TestCase{
        PreCheck:                 func() { acctest.PreCheck(t) },
        ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccAddressBlockNextAvailableInAB("10.0.0.0", 24, 26),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v1),
                    resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_address_block.test", "id"),
                    resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
                    resource.TestCheckResourceAttr(resourceName, "cidr", "26"),
                    resource.TestCheckResourceAttrSet(resourceName, "id"),
                ),
            },
            // Update and Read
            // Update of next_available_id will destroy existing resource and create a new resource
            {
                Config: testAccAddressBlockNextAvailableInAB("12.0.0.0", 8, 16),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckAddressBlockDestroy(context.Background(), &v1),
                    testAccCheckAddressBlockExists(context.Background(), resourceName, &v2),
                    resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_address_block.test", "id"),
                    resource.TestCheckResourceAttr(resourceName, "address", "12.0.0.0"),
                    resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
                    resource.TestCheckResourceAttrSet(resourceName, "id"),
                ),
            },
            // Delete testing automatically occurs in TestCase
        },
    })
}

func testAccCheckAddressBlockExists(ctx context.Context, resourceName string, v *ipam.IpamsvcAddressBlock) resource.TestCheckFunc {
    // Verify the resource exists in the cloud
    return func(state *terraform.State) error {
        rs, ok := state.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("not found: %s", resourceName)
        }
        apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
            AddressBlockAPI.
            AddressBlockRead(ctx, rs.Primary.ID).
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

func testAccCheckAddressBlockDestroy(ctx context.Context, v *ipam.IpamsvcAddressBlock) resource.TestCheckFunc {
    // Verify the resource was destroyed
    return func(state *terraform.State) error {
        _, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
            AddressBlockAPI.
            AddressBlockRead(ctx, *v.Id).
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

func testAccCheckAddressBlockDisappears(ctx context.Context, v *ipam.IpamsvcAddressBlock) resource.TestCheckFunc {
    // Delete the resource externally to verify disappears test
    return func(state *terraform.State) error {
        _, err := acctest.BloxOneClient.IPAddressManagementAPI.
            AddressBlockAPI.
            AddressBlockDelete(ctx, *v.Id).
            Execute()
        if err != nil {
            return err
        }
        return nil
    }
}

func testAccAddressBlockBasicConfig(address, cidr string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
}
`, address, cidr)

    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockAsmConfig(address, cidr string, asmThreshold int, enable, enableNotification bool, forecastPeriod, growthFactor int, growthType string, history, minTotal, minUnused int, reenableDate string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_asm_config" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
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
`, address, cidr, asmThreshold, enable, enableNotification, forecastPeriod, growthFactor, growthType, history, minTotal, minUnused, reenableDate)

    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockCidr(address string, cidr string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_cidr" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
} 
`, address, cidr)

    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockComment(address string, cidr string, comment string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_comment" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    comment = %q
}
`, address, cidr, comment)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsClientUpdate(address string, cidr string, ddnsClientUpdate string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_client_update" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_client_update = %q
}
`, address, cidr, ddnsClientUpdate)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsDomain(address string, cidr string, ddnsDomain string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_domain" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_domain = %q
}
`, address, cidr, ddnsDomain)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsGenerateName(address string, cidr string, ddnsGenerateName string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_generate_name" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_generate_name = %q
}
`, address, cidr, ddnsGenerateName)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsGeneratedPrefix(address, cidr, ddnsGeneratedPrefix string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_generated_prefix" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_generated_prefix = %q
}
`, address, cidr, ddnsGeneratedPrefix)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsSendUpdates(address, cidr, ddnsSendUpdates string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_send_updates" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_send_updates = %q
}
`, address, cidr, ddnsSendUpdates)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsTtlPercent(address, cidr, ddnsTtlPercent string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_ttl_percent" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_ttl_percent = %q
}
`, address, cidr, ddnsTtlPercent)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsUpdateOnRenew(address, cidr, ddnsUpdateOnRenew string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_update_on_renew" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_update_on_renew = %q
}
`, address, cidr, ddnsUpdateOnRenew)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDdnsUseConflictResolution(address, cidr, ddnsUseConflictResolution string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_ddns_use_conflict_resolution" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    ddns_use_conflict_resolution = %q
}
`, address, cidr, ddnsUseConflictResolution)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHeaderOptionFilename(address, cidr, headerOptionFilename string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_header_option_filename" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    header_option_filename = %q
}
`, address, cidr, headerOptionFilename)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHeaderOptionServerAddress(address, cidr, headerOptionServerAddress string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_header_option_server_address" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    header_option_server_address = %q
}
`, address, cidr, headerOptionServerAddress)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHeaderOptionServerName(address, cidr, headerOptionServerName string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_header_option_server_name" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    header_option_server_name = %q
}
`, address, cidr, headerOptionServerName)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHostnameRewriteChar(address, cidr, hostnameRewriteChar string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_hostname_rewrite_char" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_char = %q
}
`, address, cidr, hostnameRewriteChar)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHostnameRewriteEnabled(address, cidr, hostnameRewriteEnabled string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_hostname_rewrite_enabled" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_enabled = %q
}
`, address, cidr, hostnameRewriteEnabled)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockHostnameRewriteRegex(address, cidr, hostnameRewriteRegex string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_hostname_rewrite_regex" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_regex = %q
}
`, address, cidr, hostnameRewriteRegex)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockDhcpConfig(address, cidr string, allowUnknown, allowUnknownV6, ignoreClientUid bool,
        leaseTime, leaseTimeV6 int) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_dhcp_config" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    dhcp_config = {
		allow_unknown = %t
		allow_unknown_v6 = %t
		ignore_client_uid = %t
		lease_time = %d
		lease_time_v6 = %d
	}
}
`, address, cidr, allowUnknown, allowUnknownV6, ignoreClientUid, leaseTime, leaseTimeV6)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockInheritanceSources(address, cidr, action string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_inheritance_sources" {
    address = %[1]q
    cidr = %[2]q
    space = bloxone_ipam_ip_space.test.id
	inheritance_sources = {
		asm_config = {
			action = %[3]q
			asm_enable_block = {
				action = %[3]q
			}
			asm_growth_block = {
				action = %[3]q
			}
			asm_threshold = {
				action = %[3]q
			}
			forecast_period = {
				action = %[3]q
			}
			history = {
				action = %[3]q
			}
			min_total = {
				action = %[3]q
			}
			min_unused = {
				action = %[3]q
			}
		}
		ddns_client_update = {
			action = %[3]q
		}
		ddns_conflict_resolution_mode = {
			action = %[3]q
		}
		ddns_enabled = {
			action = "inherit"
		}
		ddns_hostname_block = {
			action = %[3]q
		}
		ddns_ttl_percent = {
			action = %[3]q
		}
		ddns_update_block = {
			action = %[3]q
		}
		ddns_update_on_renew = {
			action = %[3]q
		}
		ddns_use_conflict_resolution = {
			action = %[3]q
		}
		header_option_filename = {
			action = %[3]q
		}
		header_option_server_address = {
			action = %[3]q
		}
		header_option_server_name = {
			action = %[3]q
		}
		hostname_rewrite_block = {
			action = %[3]q
		}
	}

}
`, address, cidr, action)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockName(address, cidr, name string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_name" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    name = %q
}
`, address, cidr, name)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockSpace(address, cidr, space string) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_space" {
    address = %q
    cidr = %q
    space = %s.id
}
`, address, cidr, space)
    return strings.Join([]string{testAccBaseWithTwoIPSpace(), config}, "")
}

func testAccAddressBlockTags(address, cidr string, tags map[string]string) string {
    tagsStr := "{\n"
    for k, v := range tags {
        tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
    }
    tagsStr += "\t}"

    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_tags" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    tags = %s
}
`, address, cidr, tagsStr)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockWithDefaultTags(address, cidr string, tags map[string]string) string {
    tagsStr := "{\n"
    for k, v := range tags {
        tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
    }
    tagsStr += "\t}"

    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test_tags" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    tags = %s
}
`, address, cidr, tagsStr)
    return strings.Join([]string{acctest.TestAccBase_ProviderWithDefaultTags(), testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressBlockNextAvailableInAB(address string, cidr, wantedCidr int) string {
    config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}

resource "bloxone_ipam_address_block" "test_next_available" {
    next_available_id = bloxone_ipam_address_block.test.id
    cidr = %d 
    space = bloxone_ipam_ip_space.test.id
    depends_on = [bloxone_ipam_address_block.test]
}
`, address, cidr, wantedCidr)
    return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}
