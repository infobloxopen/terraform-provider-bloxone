package redirect_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/redirect"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccCustomRedirectsResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_custom_redirect.test"
	var v redirect.CustomRedirect
	name := acctest.RandomNameWithPrefix("custom-redirect")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCustomRedirectsBasicConfig(name, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCustomRedirectsResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_custom_redirect.test"
	var v redirect.CustomRedirect
	name := acctest.RandomNameWithPrefix("custom-redirect")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCustomRedirectsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRedirectsBasicConfig(name, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					testAccCheckCustomRedirectsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccCustomRedirectsResource_Data(t *testing.T) {
	var resourceName = "bloxone_td_custom_redirect.test_data"
	var v redirect.CustomRedirect
	name := acctest.RandomNameWithPrefix("custom-redirect")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCustomRedirectsData(name, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data", "156.2.3.10"),
				),
			},
			// Update and Read
			{
				Config: testAccCustomRedirectsData(name, "198.3.2.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data", "198.3.2.10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCustomRedirectsResource_Name(t *testing.T) {
	var resourceName = "bloxone_td_custom_redirect.test_name"
	var v redirect.CustomRedirect
	name1 := acctest.RandomNameWithPrefix("custom-redirect")
	name2 := acctest.RandomNameWithPrefix("custom-redirect")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCustomRedirectsName(name1, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccCustomRedirectsName(name2, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckCustomRedirectsExists(ctx context.Context, resourceName string, v *redirect.CustomRedirect) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.Atoi(rs.Primary.ID)
		apiRes, _, err := acctest.BloxOneClient.RedirectAPI.
			CustomRedirectsAPI.
			ReadCustomRedirect(ctx, int32(id)).
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

func testAccCheckCustomRedirectsDestroy(ctx context.Context, v *redirect.CustomRedirect) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.RedirectAPI.
			CustomRedirectsAPI.
			ReadCustomRedirect(ctx, *v.Id).
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

func testAccCheckCustomRedirectsDisappears(ctx context.Context, v *redirect.CustomRedirect) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.RedirectAPI.
			CustomRedirectsAPI.
			DeleteSingleCustomRedirect(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCustomRedirectsBasicConfig(name, data string) string {
	return fmt.Sprintf(`
resource "bloxone_td_custom_redirect" "test" {
	name = %q
	data = %q
}
`, name, data)
}

func testAccCustomRedirectsData(name, data string) string {
	return fmt.Sprintf(`
resource "bloxone_td_custom_redirect" "test_data" {
	name = %q
	data = %q
}
`, name, data)
}

func testAccCustomRedirectsName(name, data string) string {
	return fmt.Sprintf(`
resource "bloxone_td_custom_redirect" "test_name" {
	name = %q
	data = %q
}
`, name, data)
}
