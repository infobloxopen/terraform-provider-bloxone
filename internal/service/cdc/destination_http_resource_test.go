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

// TODO: the acceptance tests cannot create credentials, so this must be the ID of existing credential on the Portal.

const testAccDestinationHttpCredentialID = "c3509ad0-9471-4584-92c5-e23e53e92b4b"

func TestAccDestinationHttpResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "credential_id", testAccDestinationHttpCredentialID),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "authentication", ""),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "ASIM"),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "HTTP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_http.test"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationHttpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					testAccCheckDestinationHttpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationHttpResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_address"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpAddress(name, fmt.Sprintf("https://%s.example.com", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("https://%s.example.com", name)),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpAddress(name, fmt.Sprintf("https://%s-updated.example.com", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("https://%s-updated.example.com", name)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_description"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_enabled"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_output_data_format"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpOutputDataFormat(name, "ASIM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "ASIM"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpOutputDataFormat(name, "CIM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CIM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_name"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_tags"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationHttpResource_TransportProtocol(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_http.test_transport_protocol"
	var v cdc.DestinationHttp
	name := acctest.RandomNameWithPrefix("destination-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationHttpTransportProtocol(name, "HTTP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "HTTP"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationHttpTransportProtocol(name, "HTTPS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "HTTPS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationHttpExists(ctx context.Context, resourceName string, v *cdc.DestinationHttp) resource.TestCheckFunc {
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
			ReadDestinationHttp(ctx, id).
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

func testAccCheckDestinationHttpDestroy(ctx context.Context, v *cdc.DestinationHttp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationHttp(ctx, *v.Id).
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

func testAccCheckDestinationHttpDisappears(ctx context.Context, v *cdc.DestinationHttp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationHttp(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDestinationHttpBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
}
`, name)
}

func testAccDestinationHttpAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_address" {
  name          = %[1]q
  address       = %[2]q
  credential_id = "`+testAccDestinationHttpCredentialID+`"
}
`, name, address)
}

func testAccDestinationHttpDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_description" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
  description   = %[2]q
}
`, name, description)
}

func testAccDestinationHttpEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_enabled" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
  enabled       = %[2]s
}
`, name, enabled)
}

func testAccDestinationHttpOutputDataFormat(name string, outputDataFormat string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_output_data_format" {
  name               = %[1]q
  address            = "https://%[1]s.example.com"
  credential_id      = "`+testAccDestinationHttpCredentialID+`"
  output_data_format = %[2]q
}
`, name, outputDataFormat)
}

func testAccDestinationHttpTransportProtocol(name string, transportProtocol string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_transport_protocol" {
  name               = %[1]q
  address            = "https://%[1]s.example.com"
  credential_id      = "`+testAccDestinationHttpCredentialID+`"
  transport_protocol = %[2]q
}
`, name, transportProtocol)
}

func testAccDestinationHttpName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_name" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
}
`, name)
}

func testAccDestinationHttpTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test_tags" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
  tags          = %[2]s
}
`, name, tagsStr)
}
