package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceDataRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
			{
				Config: `
					data "b1ddi_dns_records" "tf_acc_dns_records" {
						filters = {
							# Check string filter
							"name_in_zone" = "tf_acc_test_a_record"
							"type" = "A"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.type", "A"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_records.tf_acc_dns_records", "results.0.rdata.address", "192.168.1.15"),
				),
			},
		},
	})
}
