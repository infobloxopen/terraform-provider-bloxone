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

func TestAccAccessCodesResource_Name(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_name"
	var v1, v2 fw.AtcfwAccessCode
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesDestroy(context.Background(), &v1),
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Activation(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_activation"
	var v1, v2 fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesActivation(name, time.Now().UTC().Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "activation", time.Now().UTC().Format(time.RFC3339)),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesActivation(name, time.Now().UTC().Add(time.Minute*10).Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesDestroy(context.Background(), &v1),
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "activation", time.Now().UTC().Add(time.Minute*10).Format(time.RFC3339)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Description(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_description"
	var v1, v2 fw.AtcfwAccessCode
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
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesDescription(name, "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Expiration(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_expiration"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesExpiration(name, time.Now().UTC().Add(time.Hour).Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "expiration", time.Now().UTC().Add(time.Hour).Format(time.RFC3339)),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodesExpiration(name, time.Now().UTC().Add(time.Hour*2).Format(time.RFC3339)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
					resource.TestCheckResourceAttr(resourceName, "expiration", time.Now().UTC().Add(time.Hour*2).Format(time.RFC3339)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodesResource_Rules(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_rules"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodesRules(name, "antimalware", "named_feed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.data", "antimalware"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.type", "named_feed"),
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
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
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
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
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
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
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
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
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
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
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
			action = "" ,
			data = %[4]q,
			description = "",
			redirect_name = "",
			type = %[5]q
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), data, rulesType)
}
