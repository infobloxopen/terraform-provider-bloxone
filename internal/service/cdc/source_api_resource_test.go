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

func TestAccSourceApiResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
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

func TestAccSourceApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_source_nios.test"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSourceApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					testAccCheckSourceApiDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSourceApiResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_address"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiAddress(name, name+"-1.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiAddress(name, name+"-2.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-2.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_CaCertificate(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_ca_certificate"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")
	cert := acctest.ReadTestdata(t, "ca_cert.pem")
	certUpdated := acctest.ReadTestdata(t, "ca_cert_updated.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiCaCertificate(name, cert),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", cert),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiCaCertificate(name, certUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", certUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Certificate(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_certificate"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")
	certKey := acctest.ReadTestdata(t, "certkey.pem")
	certKeyUpdated := acctest.ReadTestdata(t, "certkey_updated.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiCertificate(name, certKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate", certKey),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiCertificate(name, certKeyUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate", certKeyUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_description"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiDescription(name, "DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiDescription(name, "DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_enabled"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_InsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_insecure_mode"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")
	cert := acctest.ReadTestdata(t, "ca_cert.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiInsecureMode(name, cert, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiInsecureMode(name, cert, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_name"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_RpzLogInsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_rpz_log_insecure_mode"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiRpzLogInsecureMode(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_log_insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiRpzLogInsecureMode(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_log_insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_tags"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSourceApiExists(ctx context.Context, resourceName string, v *cdc.SourceNios) resource.TestCheckFunc {
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
			SourceApiAPI.
			ReadSourceNios(ctx, id).
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

func testAccCheckSourceApiDestroy(ctx context.Context, v *cdc.SourceNios) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			SourceApiAPI.
			ReadSourceNios(ctx, *v.Id).
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

func testAccCheckSourceApiDisappears(ctx context.Context, v *cdc.SourceNios) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			SourceApiAPI.
			DeleteSourceNios(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSourceApiBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, name)
}

func testAccSourceApiAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_address" {
    name          = %q
    address       = %q
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, address)
}

func testAccSourceApiCaCertificate(name string, caCertificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_ca_certificate" {
    name           = %q
    address        = "%s.example.com"
    nios_username  = "tf_auto"
    nios_password  = "tf_auto_password1!"
    ca_certificate = %q
}
`, name, name, caCertificate)
}

func testAccSourceApiCertificate(name string, certificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_certificate" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    certificate   = %q
}
`, name, name, certificate)
}

func testAccSourceApiDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_description" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    description   = %q
}
`, name, name, description)
}

func testAccSourceApiEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_enabled" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    enabled       = %s
}
`, name, name, enabled)
}

func testAccSourceApiInsecureMode(name string, cert string, insecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_insecure_mode" {
    name           = %q
    address        = "%s.example.com"
    nios_username  = "tf_auto"
    nios_password  = "tf_auto_password1!"
    ca_certificate = %q
    insecure_mode  = %s
}
`, name, name, cert, insecureMode)
}

func testAccSourceApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_name" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, name)
}

func testAccSourceApiRpzLogInsecureMode(name string, rpzLogInsecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_rpz_log_insecure_mode" {
    name                   = %q
    address                = "%s.example.com"
    nios_username          = "tf_auto"
    nios_password          = "tf_auto_password1!"
    rpz_log_insecure_mode  = %s
}
`, name, name, rpzLogInsecureMode)
}

func testAccSourceApiTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_tags" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    tags          = %s
}
`, name, name, tagsStr)
}
