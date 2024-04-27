package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - policy_ids

func TestAccAccessCodesResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_access_code.test"
	var v fw.AccessCode
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
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "terraform_test"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "custom_list"),
					resource.TestCheckResourceAttr(resourceName, "activation", time.Now().UTC().Format(time.RFC3339)),
					resource.TestCheckResourceAttr(resourceName, "expiration", time.Now().UTC().Add(time.Hour).Format(time.RFC3339)),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					resource.TestCheckResourceAttrSet(resourceName, "access_key"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.action", ""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "rules.0.redirect_name", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_access_code.test"
	var v fw.AccessCode
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

func TestAccAccessCodesResource_Name(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_name"
	var v1, v2 fw.AccessCode
	name1 := acctest.RandomNameWithPrefix("ac")
	name2 := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Activation(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_activation"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	actTime1 := time.Now().UTC().Format(time.RFC3339)
	actTime2 := time.Now().UTC().Add(time.Minute * 10).Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesActivation(name, actTime1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "activation", actTime1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesActivation(name, actTime2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "activation", actTime2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Description(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_description"
	var v1, v2 fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesDescription(name, "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesDescription(name, "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Expiration(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_expiration"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	expTime1 := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	expTime2 := time.Now().UTC().Add(time.Hour * 2).Format(time.RFC3339)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesExpiration(name, expTime1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration", expTime1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesExpiration(name, expTime2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration", expTime2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Rules(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_rules"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesRules(name, "terraform_test", "custom_list"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "terraform_test"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "custom_list"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesRules(name, "suspicious", "named_feed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "suspicious"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAccessCodesExists(ctx context.Context, resourceName string, v *fw.AccessCode) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			ReadAccessCode(ctx, rs.Primary.ID).
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

func testAccCheckAccessCodesDestroy(ctx context.Context, v *fw.AccessCode) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			ReadAccessCode(ctx, *v.AccessKey).
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

func testAccCheckAccessCodesDisappears(ctx context.Context, v *fw.AccessCode) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			AccessCodesAPI.
			DeleteSingleAccessCodes(ctx, *v.AccessKey).
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
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
}

func testAccAccessCodesName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test_name" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
}

func testAccAccessCodesActivation(name, activation string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test_activation" {
	name = %[1]q
	activation = %q
	expiration = %[3]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

`, name, activation, time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
}

func testAccAccessCodesExpiration(name, expiration string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test_expiration" {
	name = %[1]q
	activation = %[2]q
	expiration = %q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), expiration)
}

func testAccAccessCodesDescription(name, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test_description" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	description = %[4]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), description)
}

func testAccAccessCodesRules(name, data, rulesType string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test_rules" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			data = %[4]q,
			type = %[5]q
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), data, rulesType)
}
