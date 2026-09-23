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

// createTestFederatedRealm creates a FederatedRealm via API and registers cleanup.
func createTestFederatedRealm(t *testing.T, ctx context.Context, name string) string {
	t.Helper()
	acctest.PreCheck(t)
	body := ipamfederation.FederatedRealm{Name: name}
	res, _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedRealmAPI.
		Create(ctx).Body(body).Execute()
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

// createTestFederatedBlock creates a FederatedBlock via API and registers cleanup.
func createTestFederatedBlock(t *testing.T, ctx context.Context, realmId, name string) string {
	t.Helper()
	body := ipamfederation.NewFederatedBlock(realmId)
	body.SetName(name)
	body.SetAddress("10.10.0.0")
	body.SetCidr(16)
	res, _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedBlockAPI.
		Create(ctx).Body(*body).Execute()
	if err != nil {
		t.Fatalf("create federated block: %v", err)
	}
	result := res.GetResult()
	id := result.GetId()
	t.Cleanup(func() {
		// Delete all FLDs in this block first, then delete the block.
		deleteFLDsByBlockPrefix(t, ctx, "10.10.0.0/16")
		if _, err := acctest.BloxOneClient.IPAMFederationAPI.FederatedBlockAPI.
			Delete(ctx, id).Execute(); err != nil {
			t.Logf("cleanup federated block %s: %v", id, err)
		}
	})
	return id
}

// deleteFLDsByBlockPrefix deletes all ForwardLookingDelegations whose address starts within the given block.
func deleteFLDsByBlockPrefix(t *testing.T, ctx context.Context, _ string) {
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

	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"))
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 26),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "results.#"),
					resource.TestCheckResourceAttr(dataSourceName, "cidr", "26"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockWithCount(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"))
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockWithCount(blockId, 26, 2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttr(dataSourceName, "fld_count", "2"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp4(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"))
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(26, "ip4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "results.#"),
					resource.TestCheckResourceAttr(dataSourceName, "protocol", "ip4"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_globalIp6(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-ipv6"))
	dataSourceName := "data.bloxone_federation_next_available_forward_looking_delegations.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(48, "ip6"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "results.#"),
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
				Config:      testAccNextAvailableFLDGlobalConfig(26, "ip5"),
				ExpectError: regexp.MustCompile(`Attribute protocol value must be one of`),
			},
		},
	})
}

func testAccNextAvailableFLDGlobalConfig(cidr int, protocol string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr     = %d
    protocol = %q
}
`, cidr, protocol)
}

func testAccNextAvailableFLDByBlockConfig(blockId string, cidr int) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    cidr               = %d
}
`, blockId, cidr)
}

func testAccNextAvailableFLDByBlockWithCount(blockId string, cidr, fldCount int) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    cidr               = %d
    fld_count          = %d
}
`, blockId, cidr, fldCount)
}
