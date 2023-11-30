package keys_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccTsigResource_basic(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "secret"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol_name"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac_sha256"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTsigResource_disappears(t *testing.T) {
	resourceName := "bloxone_keys_tsig.test"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTsigDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTsigBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					testAccCheckTsigDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccTsigResource_Algorithm(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test_algorithm"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigAlgorithm(name, "hmac_sha512"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac_sha512"),
				),
			},
			// Update and Read
			{
				Config: testAccTsigAlgorithm(name, "hmac_sha1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "algorithm", "hmac_sha1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTsigResource_Comment(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test_comment"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigComment(name, "key created"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "key created"),
				),
			},
			// Update and Read
			{
				Config: testAccTsigComment(name, "key updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "key updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTsigResource_Name(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test_name"
	var v keys.KeysTSIGKey
	name1 := acctest.RandomNameWithPrefix("key") + "."
	name2 := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccTsigName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTsigResource_Secret(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test_secret"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigSecret(name, "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secret", "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="),
				),
			},
			// Update and Read
			{
				Config: testAccTsigSecret(name, "FzpyuZuQAHxLmwZVGlYcwaPB7Ow9MSWqSyyJlNR1XUc="),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secret", "FzpyuZuQAHxLmwZVGlYcwaPB7Ow9MSWqSyyJlNR1XUc="),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTsigResource_Tags(t *testing.T) {
	var resourceName = "bloxone_keys_tsig.test_tags"
	var v keys.KeysTSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTsigTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccTsigTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsigExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckTsigExists(ctx context.Context, resourceName string, v *keys.KeysTSIGKey) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.KeysAPI.
			TsigAPI.
			TsigRead(ctx, rs.Primary.ID).
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

func testAccCheckTsigDestroy(ctx context.Context, v *keys.KeysTSIGKey) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.KeysAPI.
			TsigAPI.
			TsigRead(ctx, *v.Id).
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

func testAccCheckTsigDisappears(ctx context.Context, v *keys.KeysTSIGKey) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.KeysAPI.
			TsigAPI.
			TsigDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccTsigBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test" {
    name = %q
}
`, name)
}

func testAccTsigAlgorithm(name string, algorithm string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test_algorithm" {
    name = %q
    algorithm = %q
}
`, name, algorithm)
}

func testAccTsigComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccTsigName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test_name" {
    name = %q
}
`, name)
}

func testAccTsigSecret(name string, secret string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test_secret" {
    name = %q
    secret = %q
}
`, name, secret)
}

func testAccTsigTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
