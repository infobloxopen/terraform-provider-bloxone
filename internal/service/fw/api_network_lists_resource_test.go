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

func TestAccNetworkListsResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_network_list.test"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListsBasicConfig(name, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", "156.2.3.0/24"),
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

func TestAccNetworkListsResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_network_list.test"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkListsBasicConfig(name, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					testAccCheckNetworkListsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkListsResource_Name(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_name"
	var v1, v2 fw.NetworkList
	name1 := acctest.RandomNameWithPrefix("nl")
	name2 := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListsName(name1, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListsName(name2, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkListsResource_Description(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_description"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListsDescription(name, "156.2.3.0/24", "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListsDescription(name, "156.2.3.0/24", "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkListsResource_Items(t *testing.T) {
	resourceName := "bloxone_td_network_list.test_items"
	var v fw.NetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkListsItems(name, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", "156.2.3.0/24"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkListsItems(name, "192.20.30.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items.0", "192.20.30.0/24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworkListsExists(ctx context.Context, resourceName string, v *fw.NetworkList) resource.TestCheckFunc {
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

func testAccCheckNetworkListsDestroy(ctx context.Context, v *fw.NetworkList) resource.TestCheckFunc {
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

func testAccCheckNetworkListsDisappears(ctx context.Context, v *fw.NetworkList) resource.TestCheckFunc {
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

func testAccNetworkListsBasicConfig(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test" {
	name = %q
	items = [%q]
}
`, name, item)
}

func testAccNetworkListsName(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_name" {
	name = %q
	items = [%q]
}

`, name, item)
}

func testAccNetworkListsDescription(name, item, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_description" {
	name = %q
	items = [%q]
	description = %q
}

`, name, item, description)
}

func testAccNetworkListsItems(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test_items" {
	name = %q
	items = [%q]
}

`, name, item)
}
