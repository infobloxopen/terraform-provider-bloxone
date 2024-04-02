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

/*
TODO:
 - Add unit test for dhcp_options
*/

func TestAccFixedAddressResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressBasicConfig(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "match_type", "mac"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa:aa:aa:aa:aa:aa"),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_fixed_address.test"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedAddressBasicConfig(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					testAccCheckFixedAddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFixedAddressResource_Address(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_address"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressAddress(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressAddress(spaceName, "10.0.0.11", "mac", "bb:bb:bb:bb:bb:bb"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.11"),
				),
			},
			// Next available IP test
			{
				Config: testAccFixedAddressAddressNAIP(spaceName, "mac", "cc:cc:cc:cc:cc:cc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_comment"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressComment(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "this range is created by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this range is created by terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressComment(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "update: this range is created by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "update: this range is created by terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_DisableDhcp(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_disable_dhcp"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressDisableDhcp(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressDisableDhcp(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_DhcpOptions(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_dhcp_options"
	var v1 ipam.IpamsvcFixedAddress
	optionSpaceName := acctest.RandomNameWithPrefix("os")
	fixedAddressName := acctest.RandomNameWithPrefix("fa")
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressDhcpOptionsOption(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", fixedAddressName, optionSpaceName, "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressDhcpOptionsGroup(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", fixedAddressName, optionSpaceName, "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_HeaderOptionFilename(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_header_option_filename"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressHeaderOptionFilename(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "header_option_filename"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "header_option_filename"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressHeaderOptionFilename(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "header_option_filename_update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "header_option_filename_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_header_option_server_address"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressHeaderOptionServerAddress(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "10.0.0.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "10.0.0.12"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressHeaderOptionServerAddress(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "10.0.0.13"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "10.0.0.13"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_header_option_server_name"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressHeaderOptionServerName(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "header_option_server_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "header_option_server_name"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressHeaderOptionServerName(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "header_option_server_name_update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "header_option_server_name_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Hostname(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_hostname"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressHostname(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "hostname1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname1"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressHostname(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "hostname2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_inheritance_sources"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressInheritanceSources(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressInheritanceSources(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "override"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_IpSpace(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_ip_space"
	var v1, v2 ipam.IpamsvcFixedAddress
	spaceName1 := acctest.RandomNameWithPrefix("ip-space")
	spaceName2 := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressIpSpace(spaceName1, spaceName2, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "bloxone_ipam_ip_space.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressIpSpace(spaceName1, spaceName2, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "bloxone_ipam_ip_space.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressDestroy(context.Background(), &v1),
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_MatchType_And_MatchValue(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_match_type"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "client_hex", "aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "client_hex"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "client_hex", "bb"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "client_hex"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "bb"),
				),
			},
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "client_text", "clienttext"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "client_text"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "clienttext"),
				),
			},
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "mac"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa:aa:aa:aa:aa:aa"),
				),
			},
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "relay_hex", "aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "relay_hex"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa"),
				),
			},
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "relay_text", "relaytext"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "relay_text"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "relaytext"),
				),
			},
			{
				Config: testAccFixedAddressMatchTypeAndMatchValue(spaceName, "10.0.0.10", "client_hex", "aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "client_hex"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_name"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressName(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "example_fixed_address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "example_fixed_address"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressName(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "example_fixed_address_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "example_fixed_address_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_tags"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressTags(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressTags(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFixedAddressExists(ctx context.Context, resourceName string, v *ipam.IpamsvcFixedAddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			FixedAddressAPI.
			FixedAddressRead(ctx, rs.Primary.ID).
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

func testAccCheckFixedAddressDestroy(ctx context.Context, v *ipam.IpamsvcFixedAddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			FixedAddressAPI.
			FixedAddressRead(ctx, *v.Id).
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

func testAccCheckFixedAddressDisappears(ctx context.Context, v *ipam.IpamsvcFixedAddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			FixedAddressAPI.
			FixedAddressDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFixedAddressBasicConfig(spaceName, address, matchType, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressAddress(spaceName, address string, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_address" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressAddressNAIP(spaceName, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_address" {
    ip_space = bloxone_ipam_ip_space.test.id
    next_available_id = bloxone_ipam_subnet.test.id
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressComment(spaceName, address string, matchType string, matchValue string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_comment" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    comment = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, comment)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressDisableDhcp(spaceName, address string, matchType string, matchValue string, disableDhcp string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_disable_dhcp" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    disable_dhcp = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, disableDhcp)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressDhcpOptionsOption(spaceName, address string, matchType string, matchValue string, name, optionSpace, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_dhcp_options" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
 	name = %q
    depends_on = [bloxone_ipam_subnet.test]
    dhcp_options = [
      {
       type = %q
       option_code = bloxone_dhcp_option_code.test.id
       option_value = %q
      }
    ]
}
`, address, matchType, matchValue, name, optionItemType, optValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), testAccBaseWithOptionSpaceAndCode("og-"+optionSpace, optionSpace, "ip4"), config}, "")
}

func testAccFixedAddressDhcpOptionsGroup(spaceName, address string, matchType string, matchValue string, name, optionSpace, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_dhcp_options" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    name = %q
    depends_on = [bloxone_ipam_subnet.test]
    dhcp_options = [
      {
       type = %q
       group = bloxone_dhcp_option_group.test.id
      }
    ]
}
`, address, matchType, matchValue, name, optionItemType)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), testAccBaseWithOptionSpaceAndCode("og-"+optionSpace, optionSpace, "ip4"), config}, "")
}

func testAccFixedAddressHeaderOptionFilename(spaceName, address string, matchType string, matchValue string, headerOptionFilename string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_header_option_filename" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    header_option_filename = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, headerOptionFilename)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressHeaderOptionServerAddress(spaceName, address string, matchType string, matchValue string, headerOptionServerAddress string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_header_option_server_address" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    header_option_server_address = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, headerOptionServerAddress)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressHeaderOptionServerName(spaceName, address string, matchType string, matchValue string, headerOptionServerName string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_header_option_server_name" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    header_option_server_name = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, headerOptionServerName)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressHostname(spaceName, address string, matchType string, matchValue string, hostname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_hostname" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    hostname = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, hostname)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressIpSpace(spaceName1, spaceName2, address string, matchType string, matchValue string, ipSpace string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = "10.0.0.0"
    cidr = 24
    space = %s.id
}

resource "bloxone_dhcp_fixed_address" "test_ip_space" {
    address = %q
    match_type = %q
    match_value = %q
    ip_space = %s.id
    depends_on = [bloxone_ipam_subnet.test]
}
`, ipSpace, address, matchType, matchValue, ipSpace)
	return strings.Join([]string{testAccBaseWithTwoIPSpace(spaceName1, spaceName2), config}, "")
}

func testAccFixedAddressMatchTypeAndMatchValue(spaceName, address string, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_match_type" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressName(spaceName, address string, matchType string, matchValue string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_name" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    name = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, name)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressTags(spaceName, address string, matchType string, matchValue string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_tags" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    tags = %s
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, matchType, matchValue, tagsStr)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressInheritanceSources(spaceName, address string, matchType string, matchValue, action string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_inheritance_sources" {
	ip_space = bloxone_ipam_ip_space.test.id
	address = %[1]q
	match_type = %[2]q
	match_value = %[3]q
	inheritance_sources = {
		header_option_filename = {
			action = %[4]q
		}
		header_option_server_address = {
			action = %[4]q
		}
		header_option_server_name = {
			action = %[4]q
		}
	}
	header_option_filename = "header_option_filename"
	header_option_server_address = "10.0.0.12"
	header_option_server_name = "header_option_server_name"
	depends_on = [bloxone_ipam_subnet.test]
		

}
`, address, matchType, matchValue, action)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}
