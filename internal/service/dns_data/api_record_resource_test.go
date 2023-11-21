package dns_data_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - inheritance_sources - Currently inheritance sources is always nil

// TODO: create test for absolute_name_spec

func TestAccRecordResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_record.test"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordBasicConfig("10.0.0.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "10.0.0.4"),
					resource.TestCheckResourceAttr(resourceName, "type", "A"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "absolute_name_spec"),
					resource.TestCheckResourceAttrSet(resourceName, "absolute_zone_name"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "dns_absolute_name_spec"),
					resource.TestCheckResourceAttrSet(resourceName, "dns_absolute_zone_name"),
					resource.TestCheckResourceAttrSet(resourceName, "dns_rdata"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "view_name"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_record.test"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordBasicConfig("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					testAccCheckRecordDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_comment"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordComment("10.0.0.1", "some comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordComment("10.0.0.1", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_disabled"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDisabled("10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDisabled("10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_NameInZone(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_name_in_zone"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNameInZone("10.0.0.1", "xyz"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "NAME_IN_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNameInZone("10.0.0.1", "abc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "NAME_IN_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Options(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_options"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordOptions("10.0.0.1", "OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordOptions("10.0.0.1", "OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_rdata"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRdata("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata", "RDATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRdata("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata", "RDATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_tags"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTags("10.0.0.1", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTags("10.0.0.1", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Ttl(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_ttl"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTtl("10.0.0.1", "TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTtl("10.0.0.1", "TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Type(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_type"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordType("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordType("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_view"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordView("10.0.0.1", "VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordView("10.0.0.1", "VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordResource_Zone(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_zone"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordZone("10.0.0.1", "ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone", "ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordZone("10.0.0.1", "ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone", "ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordExists(ctx context.Context, resourceName string, v *dns_data.DataRecord) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSDataAPI.
			RecordAPI.
			RecordRead(ctx, rs.Primary.ID).
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

func testAccCheckRecordDestroy(ctx context.Context, v *dns_data.DataRecord) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSDataAPI.
			RecordAPI.
			RecordRead(ctx, *v.Id).
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

func testAccCheckRecordDisappears(ctx context.Context, v *dns_data.DataRecord) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSDataAPI.
			RecordAPI.
			RecordDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordBasicConfig(address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test" {
    rdata = {
		"address" = %q
	}
    type = "A"
	zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
}
`, address)
}

func testAccRecordComment(address string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_comment" {
    rdata = {
		"address" = %q
	}
    type = "A"
	zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
	comment = %q
}
`, address, comment)
}

func testAccRecordDisabled(address string, disabled bool) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_disabled" {
    rdata = {
		"address" = %q
	}
    type = "A"
	zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
	disabled = %t
}
`, address, disabled)
}

func testAccRecordNameInZone(address string, nameInZone string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_name_in_zone" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    name_in_zone = %q
}
`, address, nameInZone)
}

func testAccRecordOptions(address string, options string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_options" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    options = %q
}
`, address, options)
}

func testAccRecordRdata(address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_rdata" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
 `, address)
}

func testAccRecordTags(address string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_tags" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    tags = %q
}
`, address, tags)
}

func testAccRecordTtl(address string, ttl string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_ttl" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    ttl = %q
}
`, address, ttl)
}

func testAccRecordType(address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_type" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
}
`, address)
}

func testAccRecordView(address string, view string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_view" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    view = %q
}
`, address, view)
}

func testAccRecordZone(address string, zone string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test_zone" {
    rdata = {
		"address" = %q
	}
    type = "A"
    zone = "dns/auth_zone/565d39e2-0591-4f2b-975d-29e1f73834d9"
    zone = %q
}
`, address, zone)
}
