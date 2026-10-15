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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApiBasicConfig(""),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiAddress("ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiAddress("ADDRESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_CaCertificate(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_ca_certificate"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiCaCertificate("CA_CERTIFICATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "CA_CERTIFICATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiCaCertificate("CA_CERTIFICATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "CA_CERTIFICATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_description"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiDescription("DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiDescription("DESCRIPTION_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiEnabled("ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiEnabled("ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_InsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_insecure_mode"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiInsecureMode("INSECURE_MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "INSECURE_MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiInsecureMode("INSECURE_MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "INSECURE_MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_name"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_output_data_format"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiOutputDataFormat("OUTPUT_DATA_FORMAT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "OUTPUT_DATA_FORMAT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiOutputDataFormat("OUTPUT_DATA_FORMAT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "OUTPUT_DATA_FORMAT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Port(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_port"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiPort("PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiPort("PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_tags"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiTags("TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiTags("TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApiResource_TransportProtocol(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_transport_protocol"
	var v cdc.DestinationSyslog

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApiTransportProtocol("TRANSPORT_PROTOCOL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "TRANSPORT_PROTOCOL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApiTransportProtocol("TRANSPORT_PROTOCOL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "TRANSPORT_PROTOCOL_UPDATE_REPLACE_ME"),
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

func testAccDestinationApiBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test" {
}
`)
}

func testAccDestinationApiAddress(address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_address" {
    address = %q
}
`, address)
}

func testAccDestinationApiCaCertificate(caCertificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_ca_certificate" {
    ca_certificate = %q
}
`, caCertificate)
}

func testAccDestinationApiDescription(description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_description" {
    description = %q
}
`, description)
}

func testAccDestinationApiEnabled(enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_enabled" {
    enabled = %q
}
`, enabled)
}

func testAccDestinationApiInsecureMode(insecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_insecure_mode" {
    insecure_mode = %q
}
`, insecureMode)
}

func testAccDestinationApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_name" {
    name = %q
}
`, name)
}

func testAccDestinationApiOutputDataFormat(outputDataFormat string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_output_data_format" {
    output_data_format = %q
}
`, outputDataFormat)
}

func testAccDestinationApiPort(port string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_port" {
    port = %q
}
`, port)
}

func testAccDestinationApiTags(tags string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_tags" {
    tags = %q
}
`, tags)
}

func testAccDestinationApiTransportProtocol(transportProtocol string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_transport_protocol" {
    transport_protocol = %q
}
`, transportProtocol)
}
