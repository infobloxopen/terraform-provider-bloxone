package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNetworkListResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_network_list.test"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("network-list")
	ip := acctest.RandomIP() + "/32"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListBasicConfig(name, ip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", ip),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkListResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_network_list.test"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("network-list")
	ip := acctest.RandomIP() + "/32"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkListDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkListBasicConfig(name, ip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					testAccCheckNetworkListDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkListResource_Name(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_name"
	var v1, v2 fw.NetworkList
	name1 := acctest.RandomNameWithPrefix("network-list")
	name2 := acctest.RandomNameWithPrefix("network-list")
	ip := acctest.RandomIP() + "/32"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListName(name1, ip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListName(name2, ip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkListResource_Description(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_description"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("network-list")
	ip := acctest.RandomIP() + "/32"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListDescription(name, ip, "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListDescription(name, ip, "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkListResource_Items(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_items"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("network-list")
	ip1 := acctest.RandomIP() + "/32"
	ip2 := acctest.RandomIP() + "/32"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListItems(name, ip1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", ip1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListItems(name, ip2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", ip2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworkListExists(ctx context.Context, resourceName string, v *fw.NetworkList) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			NetworkListsAPI.
			ReadNetworkList(ctx, int32(id)).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResults()
		return nil
	}
}

func testAccCheckNetworkListDestroy(ctx context.Context, v *fw.NetworkList) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			NetworkListsAPI.
			ReadNetworkList(ctx, *v.Id).
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

func testAccCheckNetworkListDisappears(ctx context.Context, v *fw.NetworkList) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			NetworkListsAPI.
			DeleteSingleNetworkLists(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkListBasicConfig(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test" {
	name = %q
	items = [%q]
}
`, name, item)
}

func testAccNetworkListName(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_name" {
	name = %q
	items = [%q]
}

`, name, item)
}

func testAccNetworkListDescription(name, item, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_description" {
	name = %q
	items = [%q]
	description = %q
}

`, name, item, description)
}

func testAccNetworkListItems(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_items" {
	name = %q
	items = [%q]
}

`, name, item)
}
