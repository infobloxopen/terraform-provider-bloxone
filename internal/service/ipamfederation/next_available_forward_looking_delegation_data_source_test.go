package ipamfederation_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

const nextAvailableFLDDataSourceName = "data.bloxone_federation_next_available_forward_looking_delegations.test"

// This data source creates Forward Looking Delegations every time it is read, so Terraform cannot delete them.
// The tests create the parent objects through the API and register cleanup that removes the FLDs (selected by a
// unique tag) before the block and realm.

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlock(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.201.0.0", 16, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 26, 1, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip4"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.tags.tf_test", runTag),
					resource.TestCheckResourceAttrSet(nextAvailableFLDDataSourceName, "results.0.id"),
					resource.TestCheckResourceAttrSet(nextAvailableFLDDataSourceName, "results.0.address"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockWithCount(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.202.0.0", 16, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 26, 2, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "fld_count", "2"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "26"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.1.cidr", "26"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIPv4(t *testing.T) {

	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.203.0.0", 16, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(26, "ip4", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "protocol", "ip4"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip4"),
					resource.TestCheckResourceAttrSet(nextAvailableFLDDataSourceName, "results.0.id"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIPv6(t *testing.T) {

	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool-ipv6"), "ip6")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-ipv6"), "2001:db8:203::", 48, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(56, "ip6", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "protocol", "ip6"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip6"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "56"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_invalidProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDGlobalConfig(26, "ip5", "unused"),
				ExpectError: regexp.MustCompile(`value must be one of`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_invalidCidr(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDGlobalConfig(129, "ip6", "unused"),
				ExpectError: regexp.MustCompile(`value must be between`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_invalidCount(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockConfig("block-id", 26, 0, "unused"),
				ExpectError: regexp.MustCompile(`value must be at least`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_blockAndPoolConflict(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = "block-id"
    federated_pool_id  = "pool-id"
    cidr               = 26
}
`,
				ExpectError: regexp.MustCompile(`cannot be specified when`),
			},
		},
	})
}

// createTestFederatedRealm creates a FederatedRealm through the API and registers cleanup. A bare struct is used
// because NewFederatedRealm sets the read-only provider field.
func createTestFederatedRealm(t *testing.T, ctx context.Context, name string) string {
	t.Helper()
	res, _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedRealmAPI.
		Create(ctx).Body(ipamfederation.FederatedRealm{Name: name}).Execute()
	if err != nil {
		t.Fatalf("create federated realm: %v", err)
	}
	result := res.GetResult()
	id := result.GetId()
	t.Cleanup(func() {
		if _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedRealmAPI.
			Delete(ctx, id).Execute(); err != nil {
			t.Logf("cleanup federated realm %s: %v", id, err)
		}
	})
	return id
}

// createTestFederatedPool creates a FederatedPool through the API and registers cleanup. Global next-available
// allocation only considers blocks that belong to a pool. The pool is registered for cleanup before the block,
// so it is deleted after the block.
func createTestFederatedPool(t *testing.T, ctx context.Context, realmId, name, protocol string) string {
	t.Helper()
	pool := ipamfederation.FederatedPool{Name: &name, FederatedRealm: &realmId, Protocol: &protocol}
	res, _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedPoolAPI.
		Create(ctx).Body(pool).Execute()
	if err != nil {
		t.Fatalf("create federated pool: %v", err)
	}
	result := res.GetResult()
	id := result.GetId()
	t.Cleanup(func() {
		if _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedPoolAPI.
			Delete(ctx, id).Execute(); err != nil {
			t.Logf("cleanup federated pool %s: %v", id, err)
		}
	})
	return id
}

// createTestFederatedBlock creates a FederatedBlock through the API. The block is tagged with runTag because the
// global scope uses the request tags to select the blocks to allocate from. Cleanup deletes the FLDs tagged with runTag
// before deleting the block.
func createTestFederatedBlock(t *testing.T, ctx context.Context, realmId, name, address string, cidr int64, poolId, runTag string) string {
	t.Helper()
	body := ipamfederation.NewFederatedBlock(realmId)
	if poolId != "" {
		body.SetFederatedPoolId(poolId)
	}
	body.SetName(name)
	body.SetAddress(address)
	body.SetCidr(cidr)
	body.SetTags(map[string]interface{}{"tf_test": runTag})
	res, _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedBlockAPI.
		Create(ctx).Body(*body).Execute()
	if err != nil {
		t.Fatalf("create federated block: %v", err)
	}
	result := res.GetResult()
	id := result.GetId()
	t.Cleanup(func() {
		deleteTestFLDsByTag(t, ctx, runTag)
		if _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedBlockAPI.
			Delete(ctx, id).Execute(); err != nil {
			t.Logf("cleanup federated block %s: %v", id, err)
		}
	})
	return id
}

// deleteTestFLDsByTag deletes only the ForwardLookingDelegations tagged with this test run.
func deleteTestFLDsByTag(t *testing.T, ctx context.Context, runTag string) {
	t.Helper()
	apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
		List(ctx).Tfilter(fmt.Sprintf("tf_test==%q", runTag)).Execute()
	if err != nil {
		t.Logf("list FLDs for cleanup: %v", err)
		return
	}
	for _, fld := range apiRes.GetResults() {
		if fld.Id == nil {
			continue
		}
		if _, err := acctest.BloxOneClient.IPAMFederationAPI.ForwardLookingDelegationAPI.
			Delete(ctx, *fld.Id).Execute(); err != nil {
			t.Logf("delete FLD %s: %v", *fld.Id, err)
		}
	}
}

func testAccNextAvailableFLDGlobalConfig(cidr int, protocol, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = %q
    tags = {
        tf_test = %q
    }
}
`, cidr, protocol, runTag)
}

func testAccNextAvailableFLDByBlockConfig(blockId string, cidr, fldCount int, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    cidr               = %d
    fld_count          = %d
    tags = {
        tf_test = %q
    }
}
`, blockId, cidr, fldCount, runTag)
}
