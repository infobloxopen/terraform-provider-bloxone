package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAccessCodeResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_access_code.test"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeBasicConfig(name, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func TestAccAccessCodeResource_disappears(t *testing.T) {

	t.Skip("Test Skipped due to inconsistent error codes returned by the API [TDDFW-397]")

	resourceName := "bloxone_td_access_code.test"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAccessCodeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAccessCodeBasicConfig(name, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					testAccCheckAccessCodeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAccessCodeResource_Name(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_name"
	var v1, v2 fw.AccessCode
	name1 := acctest.RandomNameWithPrefix("ac")
	name2 := acctest.RandomNameWithPrefix("ac")
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeName(name1, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodeName(name2, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodeResource_Activation(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_activation"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	actTime1 := time.Now().UTC().Format(time.RFC3339)
	actTime2 := time.Now().UTC().Add(time.Minute * 10).Format(time.RFC3339)
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeActivation(name, actTime1, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "activation", actTime1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodeActivation(name, actTime2, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "activation", actTime2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodeResource_Description(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_description"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeDescription(name, "Test Description", namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodeDescription(name, "Updated Test Description", namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodeResource_Expiration(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_expiration"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	expTime1 := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	expTime2 := time.Now().UTC().Add(time.Hour * 2).Format(time.RFC3339)
	namedListName := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeExpiration(name, expTime1, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration", expTime1),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodeExpiration(name, expTime2, namedListName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expiration", expTime2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAccessCodeResource_Rules(t *testing.T) {
	resourceName := "bloxone_td_access_code.test_rules"
	var v fw.AccessCode
	name := acctest.RandomNameWithPrefix("ac")
	namedListName1 := acctest.RandomNameWithPrefix("named-list")
	namedListName2 := acctest.RandomNameWithPrefix("named-list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAccessCodeRules(name, "test", namedListName1, namedListName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.test", "name"),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.type", "bloxone_td_named_list.test", "type"),
				),
			},
			// Update and Read
			{
				Config: testAccAccessCodeRules(name, "test2", namedListName1, namedListName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.test2", "name"),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.type", "bloxone_td_named_list.test2", "type"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAccessCodeExists(ctx context.Context, resourceName string, v *fw.AccessCode) resource.TestCheckFunc {
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

func testAccCheckAccessCodeDestroy(ctx context.Context, v *fw.AccessCode) resource.TestCheckFunc {
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

func testAccCheckAccessCodeDisappears(ctx context.Context, v *fw.AccessCode) resource.TestCheckFunc {
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

func testAccBaseWithNamedList(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test" {
	name = %[1]q
	items_described = [
	{
		item = "example.com"
		description = "Example Domain"
	}
	]
	type = "custom_list"
}
`, name)
}

func testAccAccessCodeBasicConfig(name, namedListName string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "test" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccAccessCodeName(name, namedListName string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "test_name" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccAccessCodeActivation(name, activation, namedListName string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "test_activation" {
	name = %[1]q
	activation = %q
	expiration = %[3]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

`, name, activation, time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccAccessCodeExpiration(name, expiration, namedListName string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "test_expiration" {
	name = %[1]q
	activation = %[2]q
	expiration = %q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), expiration)
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccAccessCodeDescription(name, description, namedListName string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "test_description" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	description = %[4]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), description)
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccAccessCodeRules(name, rules, namedListName1, namedListName2 string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_named_list" "test2" {
	name = %[1]q
	items_described = [
	{
		item = "example2.com"
		description = "Example Domain"
	}
	]
	type = "custom_list"
}

resource "bloxone_td_access_code" "test_rules" {
	name = %[2]q
	activation = %[3]q
	expiration = %[4]q
	rules = [
		{
			data = bloxone_td_named_list.%[5]s.name,
			type = bloxone_td_named_list.%[5]s.type
		}
	]
}

`, namedListName2, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), rules)
	return strings.Join([]string{testAccBaseWithNamedList(namedListName1), config}, "")
}
