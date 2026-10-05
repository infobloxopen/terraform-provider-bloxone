package ipamfederation_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

// deleteFLDsByBlockPrefix deletes all ForwardLookingDelegations.
// The next-available FLD data source allocates FLDs via POST; they must be cleaned up
// explicitly because Terraform does not manage them as resources.
func deleteFLDsByBlockPrefix(t *testing.T, ctx context.Context) {
	t.Helper()
	apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
		List(ctx).Execute()
	if err != nil {
		t.Logf("list FLDs for cleanup: %v", err)
		return
	}
	for _, fld := range apiRes.GetResults() {
		if fld.Id == nil {
			continue
		}
		if _, delErr := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
			Delete(ctx, *fld.Id).Execute(); delErr != nil {
			t.Logf("delete FLD %s: %v", *fld.Id, delErr)
		}
	}
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlock(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	t.Cleanup(func() { deleteFLDsByBlockPrefix(t, ctx) })

	realmName := acctest.RandomNameWithPrefix("federated-realm")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(realmName, "10.10.0.0", 16, 26),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttr(dataSourceName, "cidr", "26"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockWithCount(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	t.Cleanup(func() { deleteFLDsByBlockPrefix(t, ctx) })

	realmName := acctest.RandomNameWithPrefix("federated-realm")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockWithCount(realmName, "10.11.0.0", 16, 26, 2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.1.cidr", "26"),
					resource.TestCheckResourceAttr(dataSourceName, "fld_count", "2"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp4(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	t.Cleanup(func() { deleteFLDsByBlockPrefix(t, ctx) })

	realmName := acctest.RandomNameWithPrefix("federated-realm")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalIp4Config(realmName, "10.12.0.0", 16, 26),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttr(dataSourceName, "protocol", "ip4"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp6(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	t.Cleanup(func() { deleteFLDsByBlockPrefix(t, ctx) })

	realmName := acctest.RandomNameWithPrefix("federated-realm")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalIp6Config(realmName, "2001:db8::", 32, 48),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "48"),
					resource.TestCheckResourceAttr(dataSourceName, "protocol", "ip6"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_invalidProtocol(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDDataSourceOnlyConfig(26, "ip5", "", ""),
				ExpectError: regexp.MustCompile(`Attribute protocol value must be one of`),
			},
		},
	})
}

// testAccNextAvailableFLDBaseConfig creates a FederatedRealm and FederatedBlock as Terraform resources.
func testAccNextAvailableFLDBaseConfig(realmName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %q
    cidr            = %d
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccNextAvailableFLDByBlockConfig(realmName, address string, blockCidr, fldCidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = bloxone_federation_federated_block.test.id
    cidr               = %d
}
`, fldCidr)
	return strings.Join([]string{testAccNextAvailableFLDBaseConfig(realmName, address, blockCidr), config}, "")
}

func testAccNextAvailableFLDByBlockWithCount(realmName, address string, blockCidr, fldCidr, fldCount int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = bloxone_federation_federated_block.test.id
    cidr               = %d
    fld_count          = %d
}
`, fldCidr, fldCount)
	return strings.Join([]string{testAccNextAvailableFLDBaseConfig(realmName, address, blockCidr), config}, "")
}

func testAccNextAvailableFLDGlobalIp4Config(realmName, address string, blockCidr, fldCidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = "ip4"
}
`, fldCidr)
	return strings.Join([]string{testAccNextAvailableFLDBaseConfig(realmName, address, blockCidr), config}, "")
}

func testAccNextAvailableFLDGlobalIp6Config(realmName, address string, blockCidr, fldCidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = "ip6"
}
`, fldCidr)
	return strings.Join([]string{testAccNextAvailableFLDBaseConfig(realmName, address, blockCidr), config}, "")
}

// testAccNextAvailableFLDDataSourceOnlyConfig is used by unit tests that validate schema without
// connecting to the API (no base infrastructure needed).
func testAccNextAvailableFLDDataSourceOnlyConfig(cidr int, protocol, tagKey, tagVal string) string {
	tagsBlock := ""
	if tagKey != "" && tagVal != "" {
		tagsBlock = fmt.Sprintf(`
    tags = {
        %s = %q
    }`, tagKey, tagVal)
	}
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = %q%s
}
`, cidr, protocol, tagsBlock)
}
