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

func TestAccNamedListResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_named_list.test"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListBasicConfig(name, item),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.item", item),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.description", "Example Domain"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					resource.TestCheckResourceAttrSet(resourceName, "item_count"),
					resource.TestCheckResourceAttrSet(resourceName, "confidence_level"),
					resource.TestCheckResourceAttrSet(resourceName, "threat_level"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_named_list.test"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNamedListDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListBasicConfig(name, item),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					testAccCheckNamedListDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNamedListResource_Name(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_name"
	var v1, v2 fw.NamedList
	name1 := acctest.RandomNameWithPrefix("named-list")
	name2 := acctest.RandomNameWithPrefix("named-list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListName(name1, item),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListName(name2, item),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_ItemsDescribed(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_items_described"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named-list")
	item1 := acctest.RandomNameWithPrefix("named-list") + ".com"
	item2 := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListItemsDescribed(name, item1, "Example Item 1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.item", item1),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.description", "Example Item 1"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListItemsDescribed(name, item2, "Example Item 2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.item", item2),
					resource.TestCheckResourceAttr(resourceName, "items_described.0.description", "Example Item 2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_Description(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_description"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListDescription(name, item, "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListDescription(name, item, "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_Confidence(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_confidence"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListConfidence(name, item, "HIGH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "confidence_level", "HIGH"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListConfidence(name, item, "MEDIUM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "confidence_level", "MEDIUM"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListConfidence(name, item, "LOW"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "confidence_level", "LOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_Type(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_type"
	var v1 fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListType(name, item, "custom_list"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "type", "custom_list"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListResource_ThreatLevel(t *testing.T) {
	resourceName := "bloxone_td_named_list.test_threat_level"
	var v fw.NamedList
	name := acctest.RandomNameWithPrefix("named_list")
	item := acctest.RandomNameWithPrefix("named-list") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListThreatLevel(name, item, "HIGH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threat_level", "HIGH"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListThreatLevel(name, item, "MEDIUM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threat_level", "MEDIUM"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedListThreatLevel(name, item, "LOW"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threat_level", "LOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNamedListExists(ctx context.Context, resourceName string, v *fw.NamedList) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			ReadNamedList(ctx, int32(id)).
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

func testAccCheckNamedListDestroy(ctx context.Context, v *fw.NamedList) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			ReadNamedList(ctx, *v.Id).
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

func testAccCheckNamedListDisappears(ctx context.Context, v *fw.NamedList) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			DeleteSingleNamedLists(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNamedListBasicConfig(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test" {
	name = %[1]q
	items_described = [
	{
		item = %[2]q
		description = "Example Domain"
	}
	]
	type = "custom_list"
}
`, name, item)
}

func testAccNamedListName(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_name" {
	name = %q
	items_described = [
	{
		item = %q
		description = "Example Domain"
	}
	]
	type = "custom_list"
}
`, name, item)
}

func testAccNamedListDescription(name, item, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_description" {
	name = %q
	items_described = [
	{
		item = %q
		description = "Example Domain"
	}
	]
	description = %q
	type = "custom_list"
}

`, name, item, description)
}

func testAccNamedListConfidence(name, item, confidence string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_confidence" {
	name = %q
	items_described = [
	{
		item = %q
		description = "Example Domain"
	}
	]
	confidence_level = %q
	type = "custom_list"
}

`, name, item, confidence)
}

func testAccNamedListItemsDescribed(name, item, itemsDescription string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_items_described" {
	name = %q
	items_described = [
	{
		item = %q
		description = %q
	}
	]
	type = "custom_list"
}

`, name, item, itemsDescription)
}

func testAccNamedListThreatLevel(name, item, threatLevel string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_threat_level" {
	name = %q
	items_described = [
	{
		item = %q
		description = "Example Domain"
	}
	]
	threat_level = %q
	type = "custom_list"
}

`, name, item, threatLevel)
}

func testAccNamedListType(name, item, listType string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test_type" {
	name = %q
	items_described = [
	{
		item = %q
		description = "Example Domain"
	}
	]
	type = %q
}

`, name, item, listType)
}
