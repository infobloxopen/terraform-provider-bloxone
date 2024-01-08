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

// TODO:add tests
// The following require additional resource/data source objects to be supported.
// - auto_generate_records
// - addresses
// - host_names

// Test case
func TestAccIpamHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func TestAccIpamHostResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpamHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpamHostBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					testAccCheckIpamHostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpamHostResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test_comment"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostComment(name, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpamHostComment(name, "test comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpamHostResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test_tags"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccIpamHostTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpamHostResource_Addresses(t *testing.T) {
	var (
		resourceName     = "bloxone_ipam_host.test_address"
		resourceNameNaip = "bloxone_ipam_host.test_address_naip"
		spaceResource    = "bloxone_ipam_ip_space.test"
		spaceResource1   = "bloxone_ipam_ip_space.test1"
		spaceResource2   = "bloxone_ipam_ip_space.test2"
		name             = acctest.RandomNameWithPrefix("ipam_host")
		nameNaip         = acctest.RandomNameWithPrefix("ipam_host_naip")
		v                ipam.IpamsvcIpamHost
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read NaIP example
			{
				Config: testAccIpamHostAddressesNAIP(nameNaip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceNameNaip, &v),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttrPair(spaceResource, "id", resourceNameNaip, "addresses.0.space"),
				),
			},
			{
				Config: testAccIpamHostAddressesNAIPMultipleAddress(nameNaip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceNameNaip, &v),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.#", "3"),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttrPair(spaceResource, "id", resourceNameNaip, "addresses.0.space"),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.1.address", "192.168.1.1"),
					resource.TestCheckResourceAttrPair(spaceResource1, "id", resourceNameNaip, "addresses.1.space"),
					resource.TestCheckResourceAttr(resourceNameNaip, "addresses.2.address", "10.0.0.1"),
					resource.TestCheckResourceAttrPair(spaceResource2, "id", resourceNameNaip, "addresses.2.space"),
				),
			},
			// Create and Read
			{
				Config: testAccIpamHostAddresses(name, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttrPair(spaceResource, "id", resourceName, "addresses.0.space"),
				),
			},
			// Update and Read
			{
				Config: testAccIpamHostAddresses(name, "10.0.0.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttrPair(spaceResource, "id", resourceName, "addresses.0.space"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpamHostExists(ctx context.Context, resourceName string, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostRead(ctx, rs.Primary.ID).
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

func testAccCheckIpamHostDestroy(ctx context.Context, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostRead(ctx, *v.Id).
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

func testAccCheckIpamHostDisappears(ctx context.Context, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpamHostBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test" {
    name = "%s"
}
`, name)
}

func testAccIpamHostComment(name, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test_comment" {
    name = "%s"
    comment = "%s"
}
`, name, comment)
}

func testAccIpamHostTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s=%q
`, k, v)
	}
	tagsStr += "\t}"
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}

func testAccIpamHostAddresses(name, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_host" "test_address" {
    name = %q
    addresses = [
		{
			address = %q
			space = bloxone_ipam_ip_space.test.id
		}
	]
	depends_on = [bloxone_ipam_subnet.test]
}
`, name, address)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccIpamHostAddressesNAIP(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_host" "test_address_naip" {
    name = %q
	addresses = [
		{
			next_available_id = bloxone_ipam_subnet.test.id
		}
	]
}
`, name)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccIpamHostAddressesNAIPMultipleAddress(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_host" "test_address_naip" {
    name = %q
	addresses = [
		{
			next_available_id = bloxone_ipam_subnet.test.id
		},
		{
			next_available_id = bloxone_ipam_subnet.test1.id
		},
		{
			next_available_id = bloxone_ipam_subnet.test2.id
		}
	]
}
`, name)
	return strings.Join([]string{testAccMultipleIPSpaceAndSubnet(), config}, "")
}

func testAccMultipleIPSpaceAndSubnet() string {
	return `
	resource "bloxone_ipam_ip_space" "test" {
		name = "test"
	}
	resource "bloxone_ipam_subnet" "test" {
		address = "10.0.0.0"
		cidr = 24
		space = bloxone_ipam_ip_space.test.id
	}
	resource "bloxone_ipam_ip_space" "test1" {
		name = "test1"
	}
	resource "bloxone_ipam_subnet" "test1" {
		address = "192.168.1.0"
		cidr = 24
		space = bloxone_ipam_ip_space.test1.id
	}
	resource "bloxone_ipam_ip_space" "test2" {
		name = "test2"
	}
	resource "bloxone_ipam_subnet" "test2" {
		address = "10.0.0.0"
		cidr = 24
		space = bloxone_ipam_ip_space.test2.id
	}
`
}
