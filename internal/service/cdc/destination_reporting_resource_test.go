package cdc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationReportingResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_reporting.test"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationReportingDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationReportingBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					testAccCheckDestinationReportingDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationReportingResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_address"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingAddress(name, fmt.Sprintf("%s.example.com", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s.example.com", name)),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingAddress(name, fmt.Sprintf("%s-updated.example.com", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s-updated.example.com", name)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_description"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_enabled"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Password(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_password"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingPassword(name, "example_password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					// password is masked by the backend, so the configured value is kept in state
					resource.TestCheckResourceAttr(resourceName, "password", "example_password"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingPassword(name, "example_password_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "example_password_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_ReportingApplianceAddress(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_reporting_appliance_address"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingReportingApplianceAddress(name, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reporting_appliance_address", "10.0.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingReportingApplianceAddress(name, "10.0.0.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reporting_appliance_address", "10.0.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Username(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_username"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingUsername(name, "admin"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "admin"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingUsername(name, "admin_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "admin_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_name"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationReportingResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_reporting.test_tags"
	var v cdc.DestinationReporting
	name := acctest.RandomNameWithPrefix("destination-reporting")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationReportingTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationReportingTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationReportingExists(ctx context.Context, resourceName string, v *cdc.DestinationReporting) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationReporting(ctx, id).
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

func testAccCheckDestinationReportingDestroy(ctx context.Context, v *cdc.DestinationReporting) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationReporting(ctx, *v.Id).
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

func testAccCheckDestinationReportingDisappears(ctx context.Context, v *cdc.DestinationReporting) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationReporting(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDestinationReportingBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}
`, name)
}

func testAccDestinationReportingAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_address" {
  name                        = %[1]q
  address                     = %[2]q
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}
`, name, address)
}

func testAccDestinationReportingUsername(name string, username string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_username" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = %[2]q
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}
`, name, username)
}

func testAccDestinationReportingPassword(name string, password string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_password" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = %[2]q
  reporting_appliance_address = "10.0.0.1"
}
`, name, password)
}

func testAccDestinationReportingReportingApplianceAddress(name string, reportingApplianceAddress string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_reporting_appliance_address" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = %[2]q
}
`, name, reportingApplianceAddress)
}

func testAccDestinationReportingDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_description" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
  description                 = %[2]q
}
`, name, description)
}

func testAccDestinationReportingEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_enabled" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
  enabled                     = %[2]s
}
`, name, enabled)
}

func testAccDestinationReportingName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_name" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}
`, name)
}

func testAccDestinationReportingTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test_tags" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
  tags                        = %[2]s
}
`, name, tagsStr)
}
