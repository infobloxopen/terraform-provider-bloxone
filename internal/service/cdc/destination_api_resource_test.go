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

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccDestinationApiResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					testAccCheckDestinationApiDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationApiResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_address"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiAddress(name, name+"-1.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiAddress(name, name+"-2.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-2.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_CaCertificate(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_ca_certificate"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiCaCertificate(name, testAccSourceApiCertPEM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", testAccSourceApiCertPEM),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiCaCertificate(name, testAccSourceApiCertPEMUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", testAccSourceApiCertPEMUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_description"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiDescription(name, "DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiDescription(name, "DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_enabled"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_InsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_insecure_mode"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiInsecureMode(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiInsecureMode(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_name"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_output_data_format"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiOutputDataFormat(name, "CEF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CEF"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiOutputDataFormat(name, "LEEF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "LEEF"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Port(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_port"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiPort(name, "514"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "514"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiPort(name, "6514"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "6514"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_tags"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_TransportProtocol(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_transport_protocol"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiTransportProtocol(name, "TCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "TCP"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiTransportProtocol(name, "UDP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "UDP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationApiExists(ctx context.Context, resourceName string, v *cdc.DestinationSyslog) resource.TestCheckFunc {
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
			ReadDestinationSyslog(ctx, id).
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

func testAccCheckDestinationApiDestroy(ctx context.Context, v *cdc.DestinationSyslog) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationSyslog(ctx, *v.Id).
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

func testAccCheckDestinationApiDisappears(ctx context.Context, v *cdc.DestinationSyslog) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationSyslog(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDestinationApiBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test" {
    name    = %q
    address = "%s.example.com"
}
`, name, name)
}

func testAccDestinationApiAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_address" {
    name    = %q
    address = %q
}
`, name, address)
}

func testAccDestinationApiCaCertificate(name string, caCertificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_ca_certificate" {
    name           = %q
    address        = "%s.example.com"
    insecure_mode  = false
    ca_certificate = %q
}
`, name, name, caCertificate)
}

func testAccDestinationApiDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_description" {
    name        = %q
    address     = "%s.example.com"
    description = %q
}
`, name, name, description)
}

func testAccDestinationApiEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_enabled" {
    name    = %q
    address = "%s.example.com"
    enabled = %s
}
`, name, name, enabled)
}

func testAccDestinationApiInsecureMode(name string, insecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_insecure_mode" {
    name           = %q
    address        = "%s.example.com"
    ca_certificate = %q
    insecure_mode  = %s
}
`, name, name, testAccSourceApiCertPEM, insecureMode)
}

func testAccDestinationApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_name" {
    name    = %q
    address = "%s.example.com"
}
`, name, name)
}

func testAccDestinationApiOutputDataFormat(name string, outputDataFormat string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_output_data_format" {
    name                = %q
    address             = "%s.example.com"
    output_data_format  = %q
}
`, name, name, outputDataFormat)
}

func testAccDestinationApiPort(name string, port string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_port" {
    name    = %q
    address = "%s.example.com"
    port    = %s
}
`, name, name, port)
}

func testAccDestinationApiTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_tags" {
    name    = %q
    address = "%s.example.com"
    tags    = %s
}
`, name, name, tagsStr)
}

func testAccDestinationApiTransportProtocol(name string, transportProtocol string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_transport_protocol" {
    name                = %q
    address             = "%s.example.com"
    transport_protocol  = %q
}
`, name, name, transportProtocol)
}
