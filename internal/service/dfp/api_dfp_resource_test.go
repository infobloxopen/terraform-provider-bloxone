package dfp_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDfpResource_basic(t *testing.T) {
	var resourceName = "bloxone_dfp_service.test"
	var v dfp.Dfp
	var name = acctest.RandomNameWithPrefix("dfp_service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//func TestAccDfpResource_disappears(t *testing.T) {
//	resourceName := "bloxone_dfp_service.test"
//	var v dfp.Dfp
//	var name = acctest.RandomNameWithPrefix("dfp_service")
//
//	resource.Test(t, resource.TestCase{
//		PreCheck:                 func() { acctest.PreCheck(t) },
//		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
//		CheckDestroy:             testAccCheckDfpDestroy(context.Background(), &v),
//		Steps: []resource.TestStep{
//			{
//				Config: testAccDfpBasicConfig(name),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckDfpExists(context.Background(), resourceName, &v),
//					//testAccCheckDfpDisappears(context.Background(), &v),
//				),
//				ExpectNonEmptyPlan: true,
//			},
//		},
//	})
//}

func TestAccSecurityPoliciesResource_Name(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_name"
	var v dfp.Dfp
	var name1 = acctest.RandomNameWithPrefix("dfp_service")
	var name2 = acctest.RandomNameWithPrefix("dfp_service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccDfpName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Host(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_name"
	var v dfp.Dfp
	var name = acctest.RandomNameWithPrefix("dfp_service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpHost(name, "dfp_host1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "host.0.legacy_host_id", "bloxone_infra_hosts.dfp_host1", "legacy_id"),
				),
			},
			// Update and Read
			{
				Config: testAccDfpHost(name, "dfp_host2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "host.0.legacy_host_id", "bloxone_infra_hosts.dfp_host2", "legacy_id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDfpResource_InternalDomainLists(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_internal_domain_lists"
	var v dfp.Dfp
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpInternalDomainLists(name, "test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.0", "bloxone_td_internal_domain_lists.test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDfpInternalDomainLists(name, "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.0", "bloxone_td_internal_domain_lists.test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDfpExists(ctx context.Context, resourceName string, v *dfp.Dfp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSForwardingProxyAPI.
			InfraServicesAPI.
			ReadDfpService(ctx, utils.ExtractResourceId(rs.Primary.Attributes["service_id"])).
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

func testAccCheckDfpDestroy(ctx context.Context, v *dfp.Dfp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSForwardingProxyAPI.
			InfraServicesAPI.
			ReadDfpService(ctx, *v.ServiceId).
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

//func testAccCheckDfpDisappears(ctx context.Context, v *dfp.Dfp) resource.TestCheckFunc {
//	// Delete the resource externally to verify disappears test
//	return func(state *terraform.State) error {
//		_, err := acctest.BloxOneClient.InfraManagementAPI.
//			.
//			De(ctx, *v.Id).
//			Execute()
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//}

func testAccDfpBasicConfig(name string) string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "test_host_1" {
  filters = {
    display_name = "TF_TEST_HOST_01"
  }
}

resource "bloxone_infra_service" "example" {
  name         = "example_dfp_service_new"
  pool_id      = data.bloxone_infra_hosts.test_host_1.results.0.pool_id
  service_type = "dfp"
  desired_state = "start"
  wait_for_state = false

}

resource "bloxone_dfp_service" "test" {
  service_id = bloxone_infra_service.example.id
}
`)
}

func testAccDfpName(name string) string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "dfp_host" {
	filters = {
		display_name = "TF_TEST_HOST_01"
	}
}

resource "bloxone_dfp_service" "test_name" {
	name = %q
	host = [
		{
			legacy_host_id = data.bloxone_infra_hosts.dfp_host.results.0.legacy_host_id
		}
	]
}
`, name)
}

func testAccDfpHost(name, host string) string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "dfp_host1" {
	filters = {
		display_name = "TF_TEST_HOST_01"
	}
}

data "bloxone_infra_hosts" "dfp_host2" {
	filters = {
		display_name = "TF_TEST_HOST_01"
	}
}

resource "bloxone_dfp_service" "test_hsot" {
	name = %q
	host = [
		{
			legacy_host_id = data.bloxone_infra_hosts.%s.results.0.legacy_host_id
		}
	]
}
`, name, host)
}

func testAccDfpInternalDomainLists(name, internalDomainList string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test1" {
	name = "internal_domain_list_1"
	internal_domains = "example.somedomain.com"
}

resource "bloxone_td_internal_domain_list" "test2" {
	name = "internal_domain_list_2"
	internal_domains = "example.newdomain.com"
}

data "bloxone_infra_hosts" "dfp_host" {
	filters = {
		display_name = "TF_TEST_HOST_01"
	}
}

resource "bloxone_dfp_service" "test_internal_domain_lists" {
	name = %q
	host = [
		{
			legacy_host_id = data.bloxone_infra_hosts.dfp_host.results.0.legacy_host_id
		}
	]
	internal_domain_lists = [resource.bloxone_td_internal_domain_list.%s.id]
}
`, name, internalDomainList)
}
