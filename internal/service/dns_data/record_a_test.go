package dns_data_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordAResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordABasicConfig(zoneFqdn, "10.0.0.15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "10.0.0.15"),
					resource.TestCheckResourceAttrPair(resourceName, "zone", "bloxone_dns_auth_zone.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_a_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordABasicConfig(zoneFqdn, "10.0.0.15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					testAccCheckRecordDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordAResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_comment"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAComment(zoneFqdn, "10.0.0.1", "some comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAComment(zoneFqdn, "10.0.0.1", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_disabled"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordADisabled(zoneFqdn, "10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordADisabled(zoneFqdn, "10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_NameInZone(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_name_in_zone"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordANameInZone(zoneFqdn, "10.0.0.1", "xyz"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "xyz"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordANameInZone(zoneFqdn, "10.0.0.1", "abc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "abc"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordARdata(zoneFqdn, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "10.0.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordARdata(zoneFqdn, "10.0.0.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "10.0.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_tags"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordATags(zoneFqdn, "10.0.0.1", map[string]string{
					"tag1": "value1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordATags(zoneFqdn, "10.0.0.1", map[string]string{
					"tag1": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Ttl(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_ttl"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordATtl(zoneFqdn, "10.0.0.1", "60"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordATtl(zoneFqdn, "10.0.0.1", "90"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "90"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_view"
	var v1, v2 dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAView("10.0.0.1", "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAView("10.0.0.1", "bloxone_dns_view.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDestroy(context.Background(), &v1),
					testAccCheckRecordExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Zone(t *testing.T) {
	var resourceName = "bloxone_dns_a_record.test_zone"
	var v1, v2 dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAZone("10.0.0.1", "bloxone_dns_auth_zone.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "zone", "bloxone_dns_auth_zone.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAZone("10.0.0.1", "bloxone_dns_auth_zone.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDestroy(context.Background(), &v1),
					testAccCheckRecordExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "zone", "bloxone_dns_auth_zone.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordADataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_a_records.test"
	resourceName := "bloxone_dns_a_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("a")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordADataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordADataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_a_records.test"
	resourceName := "bloxone_dns_a_record.test"
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	var v dns_data.DataRecord
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordADataSourceConfigTagFilters(zoneFqdn, "10.0.0.15", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccCheckRecordResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "absolute_name_spec", dataSourceName, "results.0.absolute_name_spec"),
		resource.TestCheckResourceAttrPair(resourceName, "absolute_zone_name", dataSourceName, "results.0.absolute_zone_name"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "delegation", dataSourceName, "results.0.delegation"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_absolute_name_spec", dataSourceName, "results.0.dns_absolute_name_spec"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_absolute_zone_name", dataSourceName, "results.0.dns_absolute_zone_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name_in_zone", dataSourceName, "results.0.dns_name_in_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_rdata", dataSourceName, "results.0.dns_rdata"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "ipam_host", dataSourceName, "results.0.ipam_host"),
		resource.TestCheckResourceAttrPair(resourceName, "name_in_zone", dataSourceName, "results.0.name_in_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "results.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "provider_metadata", dataSourceName, "results.0.provider_metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "rdata", dataSourceName, "results.0.rdata"),
		resource.TestCheckResourceAttrPair(resourceName, "source", dataSourceName, "results.0.source"),
		resource.TestCheckResourceAttrPair(resourceName, "subtype", dataSourceName, "results.0.subtype"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "results.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "results.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "results.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "view_name", dataSourceName, "results.0.view_name"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "results.0.zone"),
	}
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

func testAccRecordABasicConfig(zoneFqdn string, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test" {
    rdata = {
		"address" = %q
	}
	zone = bloxone_dns_auth_zone.test.id
}
`, address)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordAComment(zoneFqdn string, address string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_comment" {
    rdata = {
		"address" = %q
	}
	zone = bloxone_dns_auth_zone.test.id
	comment = %q
}
`, address, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordADisabled(zoneFqdn string, address string, disabled bool) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_disabled" {
    rdata = {
		"address" = %q
	}
	zone = bloxone_dns_auth_zone.test.id
	disabled = %t
}
`, address, disabled)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordANameInZone(zoneFqdn string, address string, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_name_in_zone" {
    rdata = {
		"address" = %q
	}
    zone = bloxone_dns_auth_zone.test.id
    name_in_zone = %q
}
`, address, nameInZone)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordARdata(zoneFqdn string, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_rdata" {
    rdata = {
		"address" = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, address)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordATags(zoneFqdn string, address string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr = tagsStr + fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_tags" {
    rdata = {
		"address" = %q
	}
    zone = bloxone_dns_auth_zone.test.id
    tags = %s
}
`, address, tagsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordATtl(zoneFqdn string, address string, ttl string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test_ttl" {
    rdata = {
        "address" = %q
    }
    zone = bloxone_dns_auth_zone.test.id
    ttl = %q
}
`, address, ttl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordAView(address string, view string) string {
	view1name := acctest.RandomNameWithPrefix("view")
	view2name := acctest.RandomNameWithPrefix("view")
	return fmt.Sprintf(`
resource "bloxone_dns_view" "one" {
	name = %[1]q
}

resource "bloxone_dns_view" "two" {
	name = %[2]q
}

resource "bloxone_dns_auth_zone" "test" {
	fqdn = "test.com."
	view = %[4]s.id
	primary_type = "cloud"
}

resource "bloxone_dns_a_record" "test_view" {
	rdata = {
		"address" = %[3]q
	}
	absolute_name_spec = "a.test.com."
	view = %[4]s.id
	depends_on = [bloxone_dns_auth_zone.test]
}
`, view1name, view2name, address, view)
}

func testAccRecordAZone(address string, zone string) string {
	zone1fqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	zone2fqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "one" {
	fqdn = %q
	primary_type = "cloud"
}

resource "bloxone_dns_auth_zone" "two" {
	fqdn = %q
	primary_type = "cloud"
}

resource "bloxone_dns_a_record" "test_zone" {
    rdata = {
		"address" = %q
	}
    zone = %s.id
}
`, zone1fqdn, zone2fqdn, address, zone)
}

func testAccBaseWithZone(zoneFqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
    fqdn = %q
    primary_type = "cloud"
}
`, zoneFqdn)
}

func testAccRecordADataSourceConfigFilters(zoneFqdn, niz string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    address = "10.0.0.15"
  }
}

data "bloxone_dns_a_records" "test" {
  filters = {
    name_in_zone = %[1]q
    zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_a_record.test]
}`, niz)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}

func testAccRecordADataSourceConfigTagFilters(zoneFqdn, address string, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test" {
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    address = %[1]q
  }
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_a_records" "test" {
  tag_filters = {
	tag1 = bloxone_dns_a_record.test.tags.tag1
  }
}
`, address, tagValue)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
