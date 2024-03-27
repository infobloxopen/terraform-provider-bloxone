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

func TestAccNamedListsResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_named_list.test"
	var v fw.AtcfwNamedList
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedListsBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListsExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedListsResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_named_list.test"
	var v fw.AtcfwNamedList
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNamedListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListsBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedListsExists(context.Background(), resourceName, &v),
					testAccCheckNamedListsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckNamedListsExists(ctx context.Context, resourceName string, v *fw.AtcfwNamedList) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		id, err := strconv.Atoi(rs.Primary.ID)
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			NamedListsReadNamedList(ctx, int32(id)).
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

func testAccCheckNamedListsDestroy(ctx context.Context, v *fw.AtcfwNamedList) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			NamedListsReadNamedList(ctx, *v.Id).
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

func testAccCheckNamedListsDisappears(ctx context.Context, v *fw.AtcfwNamedList) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			NamedListsAPI.
			NamedListsDeleteSingleNamedLists(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNamedListsBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test" {
	name = %q
	items_described = [
	{
		item = "tf-domain.com."
		description = "Exaample Domain"
	}
	]
}
`, name)
}
