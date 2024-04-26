package keys_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccKerberosDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_keys_kerberoses.test"
	var v keys.KerberosKey

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKerberosDataSourceConfigFilters("DNS/ns.b1ddi.neo1.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKerberosExists(context.Background(), dataSourceName, &v),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "5"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.domain"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.principal", "DNS/ns.b1ddi.neo1.com"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.uploaded_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.version"),
				),
			},
		},
	})
}

func TestAccKerberosDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_keys_kerberoses.test_tag"
	var v keys.KerberosKey

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKerberosDataSourceConfigTagFilters("tf_acc_test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKerberosExists(context.Background(), dataSourceName, &v),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "5"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.algorithm"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.domain"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.principal", "DNS/ns.b1ddi.neo1.com"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.tags.used_for", "tf_acc_test"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.uploaded_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
				),
			},
		},
	})
}

func testAccCheckKerberosExists(ctx context.Context, dataSourceName string, v *keys.KerberosKey) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("not found: %s", dataSourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.KeysAPI.
			KerberosAPI.
			Read(ctx, rs.Primary.Attributes["results.0.id"]).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", dataSourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

// below all TestAcc functions

func testAccKerberosDataSourceConfigFilters(principal string) string {
	return fmt.Sprintf(`
data "bloxone_keys_kerberoses" "test" {
  filters = {
	principal = %q
  }
}
`, principal)
}

func testAccKerberosDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
data "bloxone_keys_kerberoses" "test_tag" {
  tag_filters = {
    used_for = %q
  }
}
`, tagValue)
}
