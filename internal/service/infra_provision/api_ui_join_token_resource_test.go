package infra_provision_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccUIJoinTokenResource_basic(t *testing.T) {
	var resourceName = "bloxone_infra_join_token.test"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUIJoinTokenBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "token_id"),
					resource.TestCheckResourceAttrSet(resourceName, "join_token"),
					resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUIJoinTokenResource_disappears(t *testing.T) {
	resourceName := "bloxone_infra_join_token.test"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUIJoinTokenDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUIJoinTokenBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					testAccCheckUIJoinTokenDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccUIJoinTokenResource_Description(t *testing.T) {
	var resourceName = "bloxone_infra_join_token.test_description"
	name1 := acctest.RandomNameWithPrefix("jt")
	name2 := acctest.RandomNameWithPrefix("jt")
	var v1 infraprovision.JoinToken
	var v2 infraprovision.JoinToken

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUIJoinTokenDescription(name1, "description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
				),
			},
			// Update and Read
			// Since require replace is set, a new resource will be created and the previous one is destroyed
			{
				Config: testAccUIJoinTokenDescription(name2, "updated_description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenDestroy(context.Background(), &v1), // the previous should be destroyed
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "description", "updated_description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUIJoinTokenResource_Tags(t *testing.T) {
	var resourceName = "bloxone_infra_join_token.test_tags"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUIJoinTokenTags(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccUIJoinTokenTags(name, "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUIJoinTokenResource_ExpiresAt(t *testing.T) {
	var resourceName = "bloxone_infra_join_token.test_expires_at"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")
	expiresAt := time.Now().UTC().Add(24 * time.Hour)
	expiresAtUpdated := expiresAt.Add(24 * time.Hour)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUIJoinTokenExpiresAt(name, expiresAt.Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expires_at", expiresAt.Format(time.RFC3339)),
				),
			},
			// Update and Read
			{
				Config: testAccUIJoinTokenExpiresAt(name, expiresAtUpdated.Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expires_at", expiresAtUpdated.Format(time.RFC3339)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckUIJoinTokenExists(ctx context.Context, resourceName string, v *infraprovision.JoinToken) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.HostActivationAPI.
			UIJoinTokenAPI.
			Read(ctx, rs.Primary.ID).
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

func testAccCheckUIJoinTokenDestroy(ctx context.Context, v *infraprovision.JoinToken) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.HostActivationAPI.
			UIJoinTokenAPI.
			Read(ctx, *v.Id).
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

func testAccCheckUIJoinTokenDisappears(ctx context.Context, v *infraprovision.JoinToken) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.HostActivationAPI.
			UIJoinTokenAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccUIJoinTokenBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test" {
    name = %q
}
`, name)
}

func testAccUIJoinTokenDescription(name, description string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test_description" {
    name = %q
    description = %q
}
`, name, description)
}

func testAccUIJoinTokenTags(name, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test_tags" {
    name = %q
    tags = {
        tag1 = %q
	}
}
`, name, tagValue)
}

func testAccUIJoinTokenExpiresAt(name, expiresAt string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test_expires_at" {
    name = %q
    expires_at = %q
}
`, name, expiresAt)
}
