package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAccessCodesResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_access_code.test"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_access_code.test"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAccessCodesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAccessCodesBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					testAccCheckAccessCodesDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckAccessCodesExists(ctx context.Context, resourceName string, v *fw.AtcfwAccessCode) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			AccessCodesReadAccessCode(ctx, rs.Primary.ID).
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

func testAccCheckAccessCodesDestroy(ctx context.Context, v *fw.AtcfwAccessCode) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			AccessCodesReadAccessCode(ctx, *v.AccessKey).
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

func testAccCheckAccessCodesDisappears(ctx context.Context, v *fw.AtcfwAccessCode) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			AccessCodesDeleteSingleAccessCodes(ctx, *v.AccessKey).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAccessCodesBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test" {
	name = %q
}
`, name)
}
