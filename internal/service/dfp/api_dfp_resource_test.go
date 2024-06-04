package dfp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"

	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - net_addr_policy_ids

func TestAccDfpResource_basic(t *testing.T) {
	var resourceName = "bloxone_dfp_service.test"
	var v dfp.Dfp
	hostName := acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpBasicConfig(hostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					resource.TestCheckResourceAttrSet(resourceName, "elb_ip_list.#"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_name"),
					resource.TestCheckResourceAttrSet(resourceName, "site_id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDfpResource_InternalDomainLists(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_internal_domain_lists"
	var v dfp.Dfp
	hostName := acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpInternalDomainLists(hostName, "test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.0", "data.bloxone_td_internal_domain_lists.default", "results.0.id"),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.1", "bloxone_td_internal_domain_list.test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDfpInternalDomainLists(hostName, "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.0", "data.bloxone_td_internal_domain_lists.default", "results.0.id"),
					resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists.1", "bloxone_td_internal_domain_list.test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDfpResource_ResolversAll(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_resolvers_all"
	var v dfp.Dfp
	hostName := acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpResolverAll(hostName, "1.1.1.1", "true", "false", "DO53"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.is_fallback", "true"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.is_local", "false"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.0", "DO53"),
				),
			},
			// Update and Read
			{
				Config: testAccDfpResolverAll(hostName, "10.10.10.1", "false", "true", "DOT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.address", "10.10.10.1"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.is_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.is_local", "true"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.0", "DOT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDfpResource_ResolversAll_Protocols_Multiple(t *testing.T) {
	resourceName := "bloxone_dfp_service.test_resolvers_all_protocols_multiple"
	var v dfp.Dfp
	hostName := acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDfpResolverAllProtocolMultiple(hostName, []string{"DO53", "DOT"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.0", "DO53"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.1", "DOT"),
				),
			},
			// Update and Read
			{
				Config: testAccDfpResolverAllProtocolMultiple(hostName, []string{"DOT", "DO53"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.0", "DOT"),
					resource.TestCheckResourceAttr(resourceName, "resolvers_all.0.protocols.1", "DO53"),
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

func testAccBaseWithInfraService(hostName, serviceType string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test_host" {
	display_name = %[1]q
}

resource "bloxone_infra_service" "example" {
	name = "%[1]s-%[2]s"
	pool_id = bloxone_infra_host.test_host.pool_id
	service_type = "%[2]s"
	desired_state = "start"
	wait_for_state = false
	depends_on = [bloxone_infra_host.test_host]
}
`, hostName, serviceType)
}

func testAccDfpBasicConfig(hostName string) string {
	config := `
resource "bloxone_dfp_service" "test" {
	service_id = bloxone_infra_service.example.id
	internal_domain_lists=[one(data.bloxone_td_internal_domain_lists.default.results).id]
}
`
	return strings.Join([]string{testAccBaseDfp(hostName), config}, "")
}

func testAccDfpInternalDomainLists(hostName, internalDomainList string) string {
	list1 := acctest.RandomNameWithPrefix("td-internal_domain_list")
	list2 := acctest.RandomNameWithPrefix("td-internal_domain_list")
	config := fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test1" {
	name = %q
	internal_domains = ["example.somedomain.com"]
}

resource "bloxone_td_internal_domain_list" "test2" {
	name = %q
	internal_domains = ["example.newdomain.com"]
}
resource "bloxone_dfp_service" "test_internal_domain_lists" {
	service_id = bloxone_infra_service.example.id
	internal_domain_lists = [one(data.bloxone_td_internal_domain_lists.default.results).id, bloxone_td_internal_domain_list.%s.id ]
}
`, list1, list2, internalDomainList)
	return strings.Join([]string{testAccBaseDfp(hostName), config}, "")
}

func testAccDfpResolverAll(hostName, resolverAddress, isFallback, isLocal, protocols string) string {
	config := fmt.Sprintf(`
resource "bloxone_dfp_service" "test_resolvers_all" {
	service_id = bloxone_infra_service.example.id
	internal_domain_lists = [one(data.bloxone_td_internal_domain_lists.default.results).id]
	resolvers_all = [
		{
			address = %q
			is_fallback = %q
			is_local = %q
			protocols = [%q]
		}
	]
}
`, resolverAddress, isFallback, isLocal, protocols)
	return strings.Join([]string{testAccBaseDfp(hostName), config}, "")
}

func testAccDfpResolverAllProtocolMultiple(hostName string, protocols []string) string {
	protocolsStr := ""
	for i, protocol := range protocols {
		if i > 0 {
			protocolsStr += ","
		}
		protocolsStr += fmt.Sprintf(`%q`, protocol)
	}

	config := fmt.Sprintf(`
resource "bloxone_dfp_service" "test_resolvers_all_protocols_multiple" {
	service_id = bloxone_infra_service.example.id
	internal_domain_lists = [one(data.bloxone_td_internal_domain_lists.default.results).id]
	resolvers_all = [
		{
			address = "2.2.2.2"
			is_fallback = "false"
			is_local = "true"
			protocols = [%s]
		}
	]
}
`, protocolsStr)
	return strings.Join([]string{testAccBaseDfp(hostName), config}, "")
}

func testAccBaseDfp(hostName string) string {
	// This is a workaround, ideally it would be nice to have the default internal domain list in the provider
	config := `
data "bloxone_td_internal_domain_lists" "default" {
	filters = {
		is_default = true
	}
}`
	return strings.Join([]string{testAccBaseWithInfraService(hostName, "dfp"), config}, "")
}
