package ipamfederation_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
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

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockIPv6(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-ipv6"), "2001:db8:204::", 48, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 56, 2, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip6"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "56"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.1.cidr", "56"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.tags.tf_test", runTag),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockWithNameAndComment(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	fldName := acctest.RandomNameWithPrefix("next-fld-name")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.204.0.0", 16, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockNameAndCommentConfig(blockId, 26, fldName, "next available FLD comment", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.name", fldName),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.comment", "next available FLD comment"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byPool(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.205.0.0", 16, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByPoolConfig(poolId, 26, 1, "ip4", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "federated_pool_id", poolId),
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

func TestAccNextAvailableForwardLookingDelegationDataSource_byPoolWithCount(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.206.0.0", 16, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByPoolConfig(poolId, 26, 3, "ip4", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "fld_count", "3"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "3"),
				),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_byPoolIPv6(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool-ipv6"), "ip6")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-ipv6"), "2001:db8:205::", 48, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByPoolConfig(poolId, 56, 1, "ip6", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip6"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "56"),
				),
			},
		},
	})
}

// A block that does not exist makes the API call fail, which must surface as a diagnostic instead of an empty result.
func TestAccNextAvailableForwardLookingDelegationDataSource_blockNotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockConfig("ipam/federated_block/00000000-0000-0000-0000-000000000000", 26, 1, "unused"),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*404`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_poolNotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByPoolConfig("ipam/federated_pool/00000000-0000-0000-0000-000000000000", 26, 1, "ip4", "unused"),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*404`),
			},
		},
	})
}

// The API rejects a protocol that differs from the protocol of the blocks in the pool.
func TestAccNextAvailableForwardLookingDelegationDataSource_poolProtocolMismatch(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.207.0.0", 16, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByPoolConfig(poolId, 56, 1, "ip6", runTag),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*400`),
			},
		},
	})
}

// Global allocation only considers blocks matching the tags, so tags that match no block are rejected.
func TestAccNextAvailableForwardLookingDelegationDataSource_globalNoMatchingBlock(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDGlobalConfig(26, "ip4", acctest.RandomNameWithPrefix("no-such-tag")),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*400`),
			},
		},
	})
}

// The requested cidr must be longer than the cidr of the parent block.
func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockCidrTooSmall(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.208.0.0", 24, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockConfig(blockId, 16, 1, runTag),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*400`),
			},
		},
	})
}

// A /24 block holds four /26 delegations, so asking for five fails with a conflict.
func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockInsufficientSpace(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.209.0.0", 24, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockConfig(blockId, 26, 5, runTag),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*409`),
			},
		},
	})
}

// Blocks of both protocols carry the same tags: the protocol filter selects the matching block and the other one is
// skipped.
func TestAccNextAvailableForwardLookingDelegationDataSource_globalMixedProtocols(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolV4 := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	poolV6 := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool-ipv6"), "ip6")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.218.0.0", 24, poolV4, runTag)
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-ipv6"), "2001:db8:218::", 48, poolV6, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalConfig(26, "ip4", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip4"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "26"),
				),
			},
			{
				Config: testAccNextAvailableFLDGlobalConfig(56, "ip6", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip6"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "56"),
				),
			},
		},
	})
}

// Several blocks carry the same tags: a request with a count larger than the first block can hold is spread across them.
func TestAccNextAvailableForwardLookingDelegationDataSource_globalSpansBlocks(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolA := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool-a"), "ip4")
	poolB := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool-b"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-a"), "10.216.0.0", 30, poolA, runTag)
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block-b"), "10.217.0.0", 24, poolB, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// The /30 block holds two /31 delegations, so the other two come from the /24 block.
				Config: testAccNextAvailableFLDGlobalCountConfig(31, 4, "ip4", runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "4"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "31"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.3.cidr", "31"),
				),
			},
		},
	})
}

// protocol is optional in global mode: with only IPv4 blocks matching the tags the protocol is derived from the block.
func TestAccNextAvailableForwardLookingDelegationDataSource_globalWithoutProtocol(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.219.0.0", 24, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDGlobalWithoutProtocolConfig(26, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(nextAvailableFLDDataSourceName, "protocol"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip4"),
				),
			},
		},
	})
}

// The pool request accepts protocol as an optional value: it is derived from the pool when omitted.
func TestAccNextAvailableForwardLookingDelegationDataSource_byPoolWithoutProtocol(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")
	_ = createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.220.0.0", 24, poolId, runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByPoolWithoutProtocolConfig(poolId, 26, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.protocol", "ip4"),
				),
			},
		},
	})
}

// protocol cannot be combined with federated_block_id, because the block supplies the protocol.
func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockProtocolConflict(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockProtocolConfig("federation/federated_block/does-not-matter", 26, "ip6"),
				ExpectError: regexp.MustCompile(`(?s)protocol.*federated_block_id`),
			},
		},
	})
}

// The longest IPv4 prefix is a valid request.
func TestAccNextAvailableForwardLookingDelegationDataSource_byBlockMaxCidr(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.222.0.0", 24, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 32, 1, runTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(nextAvailableFLDDataSourceName, "results.0.cidr", "32"),
				),
			},
		},
	})
}

// A pool without any Federated Block has nothing to allocate from.
func TestAccNextAvailableForwardLookingDelegationDataSource_poolWithoutBlocks(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	poolId := createTestFederatedPool(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-pool"), "ip4")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByPoolConfig(poolId, 26, 1, "ip4", "unused"),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*400`),
			},
		},
	})
}

// The name is limited to 256 characters.
func TestAccNextAvailableForwardLookingDelegationDataSource_nameTooLong(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.223.0.0", 24, "", runTag)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByBlockNameAndCommentConfig(blockId, 26, strings.Repeat("a", 257), "", runTag),
				ExpectError: regexp.MustCompile(`(?s)Client Error.*400`),
			},
		},
	})
}

// Every read allocates new delegations: a second step with the same configuration returns different addresses.
func TestAccNextAvailableForwardLookingDelegationDataSource_repeatedReadAllocatesNew(t *testing.T) {
	ctx := context.Background()
	acctest.PreCheck(t)

	runTag := acctest.RandomNameWithPrefix("next-fld")
	realmId := createTestFederatedRealm(t, ctx, acctest.RandomNameWithPrefix("federated-realm"))
	blockId := createTestFederatedBlock(t, ctx, realmId, acctest.RandomNameWithPrefix("federated-block"), "10.224.0.0", 24, "", runTag)

	var firstAddress string
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 28, 1, runTag),
				Check: resource.TestCheckResourceAttrWith(nextAvailableFLDDataSourceName, "results.0.address", func(value string) error {
					firstAddress = value
					return nil
				}),
			},
			{
				Config: testAccNextAvailableFLDByBlockConfig(blockId, 28, 1, runTag),
				Check: resource.TestCheckResourceAttrWith(nextAvailableFLDDataSourceName, "results.0.address", func(value string) error {
					if value == firstAddress {
						return fmt.Errorf("expected a newly allocated address, got the previous one: %s", value)
					}
					return nil
				}),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_emptyProtocol(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDGlobalConfig(26, "", "unused"),
				ExpectError: regexp.MustCompile(`value must be one of`),
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

func TestAccNextAvailableForwardLookingDelegationDataSource_cidrZero(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDGlobalConfig(0, "ip4", "unused"),
				ExpectError: regexp.MustCompile(`value must be between`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_negativeCount(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDByPoolConfig("pool-id", 26, -1, "ip4", "unused"),
				ExpectError: regexp.MustCompile(`value must be at least`),
			},
		},
	})
}

func TestAccNextAvailableForwardLookingDelegationDataSource_missingCidr(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableFLDMissingCidrConfig("block-id"),
				ExpectError: regexp.MustCompile(`(?s)Missing required argument.*cidr|cidr.*required`),
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
				Config:      testAccNextAvailableFLDBlockAndPoolConfig("block-id", "pool-id", 26),
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

func testAccNextAvailableFLDByPoolConfig(poolId string, cidr, fldCount int, protocol, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_pool_id = %q
    cidr              = %d
    fld_count         = %d
    protocol          = %q
    tags = {
        tf_test = %q
    }
}
`, poolId, cidr, fldCount, protocol, runTag)
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

func testAccNextAvailableFLDGlobalCountConfig(cidr, fldCount int, protocol, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr      = %d
    fld_count = %d
    protocol  = %q
    tags = {
        tf_test = %q
    }
}
`, cidr, fldCount, protocol, runTag)
}

func testAccNextAvailableFLDGlobalWithoutProtocolConfig(cidr int, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    cidr = %d
    tags = {
        tf_test = %q
    }
}
`, cidr, runTag)
}

func testAccNextAvailableFLDByPoolWithoutProtocolConfig(poolId string, cidr int, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_pool_id = %q
    cidr              = %d
    tags = {
        tf_test = %q
    }
}
`, poolId, cidr, runTag)
}

func testAccNextAvailableFLDByBlockNameAndCommentConfig(blockId string, cidr int, name, comment, runTag string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    cidr               = %d
    name               = %q
    comment            = %q
    tags = {
        tf_test = %q
    }
}
`, blockId, cidr, name, comment, runTag)
}

func testAccNextAvailableFLDByBlockProtocolConfig(blockId string, cidr int, protocol string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    cidr               = %d
    protocol           = %q
}
`, blockId, cidr, protocol)
}

func testAccNextAvailableFLDMissingCidrConfig(blockId string) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
}
`, blockId)
}

func testAccNextAvailableFLDBlockAndPoolConfig(blockId, poolId string, cidr int) string {
	return fmt.Sprintf(`
data "bloxone_federation_next_available_forward_looking_delegations" "test" {
    federated_block_id = %q
    federated_pool_id  = %q
    cidr               = %d
}
`, blockId, poolId, cidr)
}
