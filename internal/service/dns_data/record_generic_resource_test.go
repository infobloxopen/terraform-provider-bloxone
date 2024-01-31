package dns_data_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordGenericResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordGenericRdataPresentation(zoneFqdn, "TYPE256", "10 1 \"https://example.com\""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.0.type", "PRESENTATION"),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.0.value", "10 1 \"https://example.com\""),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE256"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordGenericRdataPresentation(zoneFqdn string, flags string, regexp string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_record" "test_rdata" {
	name_in_zone = "generic"
    type = %q
    rdata        = {
      subfields = [
        {
          type  = "PRESENTATION"
          value = %q
        }
      ]
    }
    zone = bloxone_dns_auth_zone.test.id
}
`, flags, regexp)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
