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
// Add unit test for dhcp_options
// Add unit test for NextAvailableIP
// Add unit tests for inheritance
*/

func TestAccFixedAddressResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressBasicConfig("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "match_type", "mac"),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa:aa:aa:aa:aa:aa"),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					//resource.TestCheckResourceAttrSet(resourceName, "inheritance_assigned_hosts"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedAddressBasicConfig("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressAddress("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressAddress("10.0.0.11", "mac", "bb:bb:bb:bb:bb:bb"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.11"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_comment"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressComment("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "this range is created by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this range is created by terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressComment("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "update: this range is created by terraform"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressDisableDhcp("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressDisableDhcp("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Hostname(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_hostname"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressHostname("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "hostname1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname1"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressHostname("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "hostname2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_MatchType(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_match_type"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressMatchType("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "mac"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressMatchType("10.0.0.10", "client_text", "abcd"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_type", "client_text"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_MatchValue(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_match_value"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressMatchValue("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_value", "aa:aa:aa:aa:aa:aa"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressMatchValue("10.0.0.10", "mac", "bb:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_value", "bb:aa:aa:aa:aa:aa"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedAddressResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_fixed_address.test_name"
	var v ipam.IpamsvcFixedAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressName("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "example_fixed_address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "example_fixed_address"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressName("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", "example_fixed_address_updated"),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedAddressTags("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedAddressTags("10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
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

func testAccBaseWithIPSpaceAndSubnetAndRange() string {
	name := acctest.RandomNameWithPrefix("ip-space")
	return fmt.Sprintf(`

data "bloxone_dhcp_hosts" "dhcp_host_by_name" {
  filters = {
    name = "TF_TEST_HOST_01"
  }
}

resource "bloxone_ipam_ip_space" "test" {
  name    = %q
  comment = "Example IP space create for FA Acceptance Test"
  tags = {
    "Test" = "Acceptance"
  }
}

resource "bloxone_ipam_subnet" "test" {
  name    = "example"
  space   = bloxone_ipam_ip_space.test.id
  address = "10.0.0.0"
  cidr    = 24
  comment = "Example Subnet create for FA Acceptance Test"
  dhcp_host = data.bloxone_dhcp_hosts.dhcp_host_by_name.results.0.id
  tags = {
    "Test" = "Acceptance"
  }
}

resource "bloxone_ipam_range" "test" {
  start = "10.0.0.2"
  end   = "10.0.0.20"
  space = bloxone_ipam_ip_space.test.id

  # Other optional fields
  name    = "example"
  comment = "Example Range create for FA Acceptance Test"
  tags = {
    "Test" = "Acceptance"
  }
  depends_on = [bloxone_ipam_subnet.test]
}
`, name)
}

func testAccFixedAddressBasicConfig(address, matchType, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressAddress(address string, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_address" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressComment(address string, matchType string, matchValue string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_comment" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    comment = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue, comment)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressDisableDhcp(address string, matchType string, matchValue string, disableDhcp string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_disable_dhcp" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    disable_dhcp = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue, disableDhcp)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressHostname(address string, matchType string, matchValue string, hostname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_hostname" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    hostname = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue, hostname)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressMatchType(address string, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_match_type" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressMatchValue(address string, matchType string, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_match_value" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressName(address string, matchType string, matchValue string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test_name" {
    ip_space = bloxone_ipam_ip_space.test.id
    address = %q
    match_type = %q
    match_value = %q
    name = %q
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue, name)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}

func testAccFixedAddressTags(address string, matchType string, matchValue string, tags map[string]string) string {
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
    depends_on = [bloxone_ipam_range.test]
}
`, address, matchType, matchValue, tagsStr)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnetAndRange(), config}, "")
}
