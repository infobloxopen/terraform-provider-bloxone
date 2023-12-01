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

func TestAccAddressResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressBasicConfig("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.1"),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "state", "used"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
					resource.TestCheckResourceAttr(resourceName, "usage.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "usage.0", "IPAM RESERVED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_address.test"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAddressBasicConfig("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					testAccCheckAddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAddressResource_Address(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_address"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressAddress("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressAddress("10.0.0.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_comment"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressComment("10.0.0.1", "some comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressComment("10.0.0.1", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Hwaddr(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_hwaddr"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressHwaddr("10.0.0.1", "00:11:22:33:44:55"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hwaddr", "00:11:22:33:44:55"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressHwaddr("10.0.0.1", "55:44:33:22:11:00"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hwaddr", "55:44:33:22:11:00"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Interface(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_interface"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressInterface("10.0.0.1", "eth0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interface", "eth0"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressInterface("10.0.0.1", "eth1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interface", "eth1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Names(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_names"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressNames("10.0.0.1", "name1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "names.0.name", "name1"),
					resource.TestCheckResourceAttr(resourceName, "names.0.type", "user"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressNames("10.0.0.1", "name2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "names.0.name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "names.0.type", "user"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Space(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_space"
	var v1 ipam.IpamsvcAddress
	var v2 ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressSpace("10.0.0.1", "bloxone_ipam_ip_space.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressSpace("10.0.0.1", "bloxone_ipam_ip_space.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressDestroy(context.Background(), &v1),
					testAccCheckAddressExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_tags"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressTags("10.0.0.1", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccAddressTags("10.0.0.1", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_NextAvailable_Subnet(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_next_available"
	var v1 ipam.IpamsvcAddress
	var v2 ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressNextAvailableInSubnet("10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.1"), // first address after broadcast address
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Update and Read
			// Update of next_available_id will destroy existing resource and create a new resource
			{
				Config: testAccAddressNextAvailableInSubnet("12.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressDestroy(context.Background(), &v1),
					testAccCheckAddressExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "12.0.0.1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_NextAvailable_AddressBlock(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_next_available"
	var v1 ipam.IpamsvcAddress
	var v2 ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressNextAvailableInAddressBlock("10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Update and Read
			// Update of next_available_id will destroy existing resource and create a new resource
			{
				Config: testAccAddressNextAvailableInAddressBlock("12.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressDestroy(context.Background(), &v1),
					testAccCheckAddressExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "12.0.0.1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAddressResource_NextAvailable_Range(t *testing.T) {
	var resourceName = "bloxone_ipam_address.test_next_available"
	var v1, v2 ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAddressNextAvailableInRange("10.0.0.10", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Update and Read
			// Update of next_available_id will destroy existing resource and create a new resource
			{
				Config: testAccAddressNextAvailableInRange("10.0.0.16", "10.0.0.26"),
				Taint:  []string{resourceName}, // Forces the recreation of the object
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAddressExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_ipam_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.16"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAddressExists(ctx context.Context, resourceName string, v *ipam.IpamsvcAddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			AddressAPI.
			AddressRead(ctx, rs.Primary.ID).
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

func testAccCheckAddressDestroy(ctx context.Context, v *ipam.IpamsvcAddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			AddressAPI.
			AddressRead(ctx, *v.Id).
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

func testAccCheckAddressDisappears(ctx context.Context, v *ipam.IpamsvcAddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			AddressAPI.
			AddressDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithIPSpaceAndSubnet() string {
	name := acctest.RandomNameWithPrefix("ip-space")
	return fmt.Sprintf(`
data "bloxone_dhcp_hosts" "dhcp_host_by_tag" {
  tag_filters = {
    used_for = "Terraform Acceptance Testing"
  }
}
resource "bloxone_ipam_ip_space" "test" {
    name = %q
}
resource "bloxone_ipam_subnet" "test" {
    address = "10.0.0.0"
    cidr = 24
    space = bloxone_ipam_ip_space.test.id
    dhcp_host = data.bloxone_dhcp_hosts.dhcp_host_by_tag.results.0.id
}
`, name)
}

func testAccAddressBasicConfig(address string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    depends_on = [bloxone_ipam_subnet.test]
}
`, address)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressAddress(address string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_address" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    depends_on = [bloxone_ipam_subnet.test]
}
`, address)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressComment(address string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_comment" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    comment = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, comment)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressHwaddr(address string, hwaddr string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_hwaddr" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    hwaddr = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, hwaddr)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressInterface(address string, interface_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_interface" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    interface = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, interface_)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressNames(address string, names ...string) string {
	quotedNames := make([]string, len(names))
	for i, n := range names {
		quotedNames[i] = fmt.Sprintf(`
        {
			name = %q
            type = "user"
        }
		`, n)
	}
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_names" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    names = [%s]
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, strings.Join(quotedNames, ","))
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressSpace(address string, space string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "one" {
    name = %[1]q
}
resource "bloxone_ipam_subnet" "one" {
    address = "10.0.0.0"
    cidr = 24
	space = bloxone_ipam_ip_space.one.id
}
resource "bloxone_ipam_ip_space" "two" {
    name = %[2]q
}
resource "bloxone_ipam_subnet" "two" {
    address = "10.0.0.0"
    cidr = 24
	space = bloxone_ipam_ip_space.two.id
}
resource "bloxone_ipam_address" "test_space" {
    address = %[3]q
    space = %[4]s.id
    depends_on = [bloxone_ipam_subnet.one, bloxone_ipam_subnet.two]
}
`, acctest.RandomNameWithPrefix("ip-space"), acctest.RandomNameWithPrefix("ip-space"), address, space)
	return config
}

func testAccAddressTags(address string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test_tags" {
    address = %q
    space = bloxone_ipam_ip_space.test.id
    tags = %s
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, tagsStr)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressNextAvailableInSubnet(address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}

resource "bloxone_ipam_address" "test_next_available" {
    next_available_id = bloxone_ipam_subnet.test.id
    space = bloxone_ipam_ip_space.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressNextAvailableInAddressBlock(address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}

resource "bloxone_ipam_subnet" "test" {
    address = %q
    cidr = 26
    space = bloxone_ipam_ip_space.test.id
}

resource "bloxone_ipam_address" "test_next_available" {
    next_available_id = bloxone_ipam_address_block.test.id
    space = bloxone_ipam_ip_space.test.id
    depends_on = [bloxone_ipam_subnet.test]
}
`, address, cidr, address)
	return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccAddressNextAvailableInRange(start, end string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test" {
    start = %q
    end = %q
    space = bloxone_ipam_ip_space.test.id
    depends_on = [bloxone_ipam_subnet.test]
}

resource "bloxone_ipam_address" "test_next_available" {
    next_available_id = bloxone_ipam_range.test.id
    space = bloxone_ipam_ip_space.test.id
}
`, start, end)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}
