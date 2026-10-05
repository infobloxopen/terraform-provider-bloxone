package ipamfederation_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

// testAccDeleteAllFLDs returns a TestCheckFunc that deletes all ForwardLookingDelegations
// via the API. Used as a cleanup step before Terraform's final destroy phase, because the
// next-available FLD data source allocates FLDs via POST (side effect) and the parent
// FederatedBlock cannot be deleted while FLDs reference it.
func testAccDeleteAllFLDs(ctx context.Context) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
			List(ctx).Execute()
		if err != nil {
			return fmt.Errorf("list FLDs for cleanup: %w", err)
		}
		for _, fld := range apiRes.GetResults() {
			if fld.Id == nil {
				continue
			}
			if _, delErr := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
				Delete(ctx, *fld.Id).Execute(); delErr != nil {
				return fmt.Errorf("delete FLD %s: %w", *fld.Id, delErr)
			}
		}
		return nil
	}
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlock(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
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
			{
				// Remove data source so Terraform doesn't re-allocate FLDs on plan.
				// The Check deletes existing FLDs via API so the FederatedBlock can be
				// destroyed in the subsequent Terraform destroy phase.
				Config: testAccNextAvailableFLDBaseConfig(realmName, "10.10.0.0", 16),
				Check:  testAccDeleteAllFLDs(ctx),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockWithCount(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
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
			{
				Config: testAccNextAvailableFLDBaseConfig(realmName, "10.11.0.0", 16),
				Check:  testAccDeleteAllFLDs(ctx),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp4(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagKey := "test_fld_global"
	tagVal := acctest.RandomNameWithPrefix("ip4")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(realmName, "10.12.0.0", "ip4", 16, 26, tagKey, tagVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttr(dataSourceName, "protocol", "ip4"),
				),
			},
			{
				Config: testAccNextAvailableFLDBaseConfigWithTags(realmName, "10.12.0.0", 16, tagKey, tagVal),
				Check:  testAccDeleteAllFLDs(ctx),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp6(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagKey := "test_fld_global"
	tagVal := acctest.RandomNameWithPrefix("ip6")
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(realmName, "2001:db8::", "ip6", 32, 48, tagKey, tagVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "48"),
					resource.TestCheckResourceAttr(dataSourceName, "protocol", "ip6"),
				),
			},
			{
				Config: testAccNextAvailableFLDBaseConfigWithTags(realmName, "2001:db8::", 32, tagKey, tagVal),
				Check:  testAccDeleteAllFLDs(ctx),
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

// testAccNextAvailableFLDBaseConfigWithTags creates a FederatedRealm and FederatedBlock with
// tags, required for global FLD allocation which uses tags to identify eligible blocks.
func testAccNextAvailableFLDBaseConfigWithTags(realmName, address string, cidr int, tagKey, tagVal string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %q
    cidr            = %d
    tags = {
        %s = %q
    }
}
`, address, cidr, tagKey, tagVal)
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

// testAccNextAvailableFLDGlobalConfig creates a tagged block and a global FLD data source that
// uses matching tags to identify eligible blocks for allocation. protocol must be "ip4" or "ip6".
// depends_on forces Terraform to read the data source during apply (after the block exists),
// not during plan (before the block is created).
func testAccNextAvailableFLDGlobalConfig(realmName, address, protocol string, blockCidr, fldCidr int, tagKey, tagVal string) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = %q
    tags = {
        %s = %q
    }
    depends_on = [bloxone_federation_federated_block.test]
}
`, fldCidr, protocol, tagKey, tagVal)
	return strings.Join([]string{testAccNextAvailableFLDBaseConfigWithTags(realmName, address, blockCidr, tagKey, tagVal), config}, "")
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
