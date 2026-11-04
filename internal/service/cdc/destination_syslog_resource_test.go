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

func TestAccDestinationSyslogResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CEF"),
					resource.TestCheckResourceAttr(resourceName, "port", "514"),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "TCP"),
				),
			},
			// Import and Read
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSyslogDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSyslogBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					testAccCheckDestinationSyslogDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationSyslogResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_address"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogAddress(name, name+"-1.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogAddress(name, name+"-2.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-2.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_CaCertificate(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_ca_certificate"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")
	cert := acctest.ReadTestdata(t, "bloxone_cdc", "ca_cert.pem")
	certUpdated := acctest.ReadTestdata(t, "bloxone_cdc", "ca_cert_updated.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogCaCertificate(name, cert),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", cert),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogCaCertificate(name, certUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", certUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_description"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogDescription(name, "Test syslog destination description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test syslog destination description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogDescription(name, "Updated test syslog destination description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test syslog destination description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_enabled"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_InsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_insecure_mode"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")
	cert := acctest.ReadTestdata(t, "bloxone_cdc", "ca_cert.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogInsecureMode(name, cert, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogInsecureMode(name, cert, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_name"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_output_data_format"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogOutputDataFormat(name, "CEF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CEF"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogOutputDataFormat(name, "LEEF"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "LEEF"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_Port(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_port"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogPort(name, "514"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "514"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogPort(name, "6514"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "6514"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSyslogResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_tags"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
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

func TestAccDestinationSyslogResource_TransportProtocol(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_syslog.test_transport_protocol"
	var v cdc.DestinationSyslog
	name := acctest.RandomNameWithPrefix("destination")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSyslogTransportProtocol(name, "TCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "TCP"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSyslogTransportProtocol(name, "UDP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSyslogExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport_protocol", "UDP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationSyslogExists(ctx context.Context, resourceName string, v *cdc.DestinationSyslog) resource.TestCheckFunc {
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

func testAccCheckDestinationSyslogDestroy(ctx context.Context, v *cdc.DestinationSyslog) resource.TestCheckFunc {
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

func testAccCheckDestinationSyslogDisappears(ctx context.Context, v *cdc.DestinationSyslog) resource.TestCheckFunc {
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

func testAccDestinationSyslogBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test" {
    name    = %q
    address = "%s.example.com"
}
`, name, name)
}

func testAccDestinationSyslogAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_address" {
    name    = %q
    address = %q
}
`, name, address)
}

func testAccDestinationSyslogCaCertificate(name string, caCertificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_ca_certificate" {
    name           = %q
    address        = "%s.example.com"
    insecure_mode  = false
    ca_certificate = %q
}
`, name, name, caCertificate)
}

func testAccDestinationSyslogDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_description" {
    name        = %q
    address     = "%s.example.com"
    description = %q
}
`, name, name, description)
}

func testAccDestinationSyslogEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_enabled" {
    name    = %q
    address = "%s.example.com"
    enabled = %s
}
`, name, name, enabled)
}

func testAccDestinationSyslogInsecureMode(name string, cert string, insecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_insecure_mode" {
    name           = %q
    address        = "%s.example.com"
    ca_certificate = %q
    insecure_mode  = %s
}
`, name, name, cert, insecureMode)
}

func testAccDestinationSyslogName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_name" {
    name    = %q
    address = "%s.example.com"
}
`, name, name)
}

func testAccDestinationSyslogOutputDataFormat(name string, outputDataFormat string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_output_data_format" {
    name                = %q
    address             = "%s.example.com"
    output_data_format  = %q
}
`, name, name, outputDataFormat)
}

func testAccDestinationSyslogPort(name string, port string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_port" {
    name    = %q
    address = "%s.example.com"
    port    = %s
}
`, name, name, port)
}

func testAccDestinationSyslogTags(name string, tags map[string]string) string {
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

func testAccDestinationSyslogTransportProtocol(name string, transportProtocol string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_transport_protocol" {
    name                = %q
    address             = "%s.example.com"
    transport_protocol  = %q
}
`, name, name, transportProtocol)
}
