package dfp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDfpResource_basic(t *testing.T) {
	var resourceName = "bloxone_dfp_service.test"
	var v dfp.AtcdfpDfp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDfpResource_disappears(t *testing.T) {
	resourceName := "bloxone_dfp_service.test"
	var v dfp.AtcdfpDfp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDfpBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					testAccCheckDfpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckDfpExists(ctx context.Context, resourceName string, v *dfp.AtcdfpDfp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSForwardingProxyAPI.
			DfpAPI.
			DfpReadDfp(ctx, rs.Primary.ID).
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

func testAccCheckDfpDestroy(ctx context.Context, v *dfp.AtcdfpDfp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSForwardingProxyAPI.
			DfpAPI.
			DfpRead(ctx, *v.Id).
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

func testAccCheckDfpDisappears(ctx context.Context, v *dfp.AtcdfpDfp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSForwardingProxyAPI.
			DfpAPI.
			DfpDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDfpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dfp_service" "test" {
}
`)
}
