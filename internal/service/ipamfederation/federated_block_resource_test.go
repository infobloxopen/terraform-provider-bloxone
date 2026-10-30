package ipamfederation_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func TestAccFederatedBlockResource_basic(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "name", ""),
					resource.TestCheckResourceAttr(resourceName, "region", ""),
					resource.TestCheckResourceAttr(resourceName, "state", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_IPv6(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	var idBefore string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("2001:db8:c1::", 48, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "2001:db8:c1::"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "48"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip6"),
					func(*terraform.State) error { idBefore = *v.Id; return nil },
				),
			},
			// Same address in different letter case must update in place, not replace
			{
				Config: testAccFederatedBlockBasicConfig("2001:DB8:C1::", 48, realmName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					func(*terraform.State) error {
						if *v.Id != idBefore {
							return fmt.Errorf("expected resource to be updated in place, but id changed from %s to %s", idBefore, *v.Id)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccFederatedBlockResource_AddressWithPrefixInvalid(t *testing.T) {
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFederatedBlockBasicConfig("10.12.0.0/16", 16, realmName),
				ExpectError: regexp.MustCompile(`Invalid IP Address`),
			},
		},
	})
}

func TestAccFederatedBlockResource_AddressRequired(t *testing.T) {
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// address is Optional/Computed to support next_available_id, so omitting both is
				// caught by the ExactlyOneOf validator rather than a plain "required" error.
				Config:      testAccFederatedBlockWithoutAddress(16, realmName),
				ExpectError: regexp.MustCompile(`one \(and only one\) of \[address,next_available_id\]\s*\n?\s*is required`),
			},
		},
	})
}

func TestAccFederatedBlockResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_federated_block.test"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					testAccCheckFederatedBlockDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFederatedBlockResource_Address(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test"
	var v1 ipamfederation.FederatedBlock
	var v2 ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.11.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockDestroy(context.Background(), &v1),
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", "10.11.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 18, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "18"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_comment"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockComment("10.10.0.0", 16, realmName, "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockComment("10.10.0.0", 16, realmName, "COMMENT_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_FederatedPoolId(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_federated_pool_id"
	var v1, v2 ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName1 := acctest.RandomNameWithPrefix("federated-pool")
	poolName2 := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockFederatedPoolId(realmName, poolName1, poolName2, "bloxone_federation_federated_pool.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "federated_pool_id", "bloxone_federation_federated_pool.one", "id"),
				),
			},
			// A block cannot be moved to another pool, so it is replaced
			{
				Config: testAccFederatedBlockFederatedPoolId(realmName, poolName1, poolName2, "bloxone_federation_federated_pool.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockDestroy(context.Background(), &v1),
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "federated_pool_id", "bloxone_federation_federated_pool.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_NextAvailableId(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_next_available_id"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockNextAvailableId(realmName, "10.40.0.0", 16, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "next_available_id", "bloxone_federation_federated_block.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "address"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_NextAvailableIdWithFederatedPoolId(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_next_available_id_with_pool"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockNextAvailableIdWithFederatedPoolId(realmName, poolName, "10.41.0.0", 16, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "next_available_id", "bloxone_federation_federated_block.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_pool_id", "bloxone_federation_federated_pool.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "address"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_NextAvailableIdWithNetworkCompliance(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_next_available_id_with_network_compliance"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockNextAvailableIdWithNetworkCompliance(realmName, "10.42.0.0", 16, 24, 28, 25, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "next_available_id", "bloxone_federation_federated_block.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "address"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.default_netmask_length", "28"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.minimum_netmask_length", "25"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.maximum_netmask_length", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_NextAvailableIdFromPool(t *testing.T) {
	t.Skip("CreateNextAvailableBlocks is only supported for AWS-backed Federated Pools; requires an AWS-linked Federated Realm, not available in this test environment")

	var resourceName = "bloxone_federation_federated_block.test_next_available_id_from_pool"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockNextAvailableIdFromPool(realmName, poolName, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "next_available_id", "bloxone_federation_federated_pool.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "address"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_FederatedRealm(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_federated_realm"
	var v ipamfederation.FederatedBlock
	realmName1 := acctest.RandomNameWithPrefix("federated-realm")
	realmName2 := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_name"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockName("10.0.0.0", 24, realmName, "NAME_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockName("10.0.0.0", 24, realmName, "NAME_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_NetworkCompliance(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_network_compliance"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockNetworkCompliance("10.0.0.0", 16, realmName, 20, 17, 28),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.default_netmask_length", "20"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.minimum_netmask_length", "17"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.maximum_netmask_length", "28"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockNetworkCompliance("10.0.0.0", 16, realmName, 22, 18, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.default_netmask_length", "22"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.minimum_netmask_length", "18"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.maximum_netmask_length", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_federated_block.test_tags"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockTags("10.0.0.0", realmName, 24, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockTags("10.0.0.0", realmName, 24, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFederatedBlockExists(ctx context.Context, resourceName string, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
			Read(ctx, rs.Primary.ID).
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

func testAccCheckFederatedBlockDestroy(ctx context.Context, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
			Read(ctx, *v.Id).
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

func testAccCheckFederatedBlockDisappears(ctx context.Context, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithFederatedRealm(name string) string {
	return fmt.Sprintf(`
resource "bloxone_federation_federated_realm" "test" {
    name = %q
}
`, name)
}

func testAccFederatedBlockBasicConfig(address string, cidr int, federatedRealm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockWithoutAddress(cidr int, federatedRealm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
}
`, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockComment(address string, cidr int, federatedRealm string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_comment" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    comment = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockFederatedRealm(federatedRealm1, federatedRealm2, realm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_federated_realm" {
	address = "10.0.0.0"
	cidr = 16
	federated_realm = %s.id
}
`, realm)
	return strings.Join([]string{testAccBaseWithTwoFederatedRealm(federatedRealm1, federatedRealm2), config}, "")
}

func testAccFederatedBlockFederatedPoolId(realmName, poolName1, poolName2, pool string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "one" {
	federated_realm = bloxone_federation_federated_realm.test.id
	name            = %q
	protocol        = "ip4"
}

resource "bloxone_federation_federated_pool" "two" {
	federated_realm = bloxone_federation_federated_realm.test.id
	name            = %q
	protocol        = "ip4"
}

resource "bloxone_federation_federated_block" "test_federated_pool_id" {
	address           = "10.0.0.0"
	cidr              = 16
	federated_realm   = bloxone_federation_federated_realm.test.id
	federated_pool_id = %s.id
}
`, poolName1, poolName2, pool)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedBlockNextAvailableId(federatedRealm string, parentAddress string, parentCidr int, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_next_available_id" {
    federated_realm   = bloxone_federation_federated_realm.test.id
    next_available_id = bloxone_federation_federated_block.test.id
    cidr = %d
}
`, cidr)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccFederatedBlockNextAvailableIdWithFederatedPoolId(federatedRealm, poolName string, parentAddress string, parentCidr int, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = "ip4"
}

resource "bloxone_federation_federated_block" "test_next_available_id_with_pool" {
    federated_realm   = bloxone_federation_federated_realm.test.id
    next_available_id = bloxone_federation_federated_block.test.id
    federated_pool_id = bloxone_federation_federated_pool.test.id
    cidr = %d
}
`, poolName, cidr)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccFederatedBlockNextAvailableIdWithNetworkCompliance(federatedRealm string, parentAddress string, parentCidr int, cidr int, defaultMask, minMask, maxMask int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_next_available_id_with_network_compliance" {
    federated_realm   = bloxone_federation_federated_realm.test.id
    next_available_id = bloxone_federation_federated_block.test.id
    cidr = %d
    network_compliance = {
        default_netmask_length = %d
        minimum_netmask_length = %d
        maximum_netmask_length = %d
    }
}
`, cidr, defaultMask, minMask, maxMask)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccFederatedBlockNextAvailableIdFromPool(federatedRealm, poolName string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = "ip4"
}

resource "bloxone_federation_federated_block" "test_next_available_id_from_pool" {
    federated_realm   = bloxone_federation_federated_realm.test.id
    next_available_id = bloxone_federation_federated_pool.test.id
    cidr = %d
}
`, poolName, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccBaseWithTwoFederatedRealm(name1, name2 string) string {
	return fmt.Sprintf(`
resource "bloxone_federation_federated_realm" "one" {
	name = %q
}
resource "bloxone_federation_federated_realm" "two" {
	name = %q
}`, name1, name2)
}

func testAccFederatedBlockName(address string, cidr int, federatedRealm string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_name" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    name = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockNetworkCompliance(address string, cidr int, federatedRealm string, defaultMask, minMask, maxMask int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_network_compliance" {
	address = %q
	cidr = %d
	federated_realm = bloxone_federation_federated_realm.test.id
	network_compliance = {
		default_netmask_length = %d
		minimum_netmask_length = %d
		maximum_netmask_length = %d
	}
}
`, address, cidr, defaultMask, minMask, maxMask)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockTags(address string, federatedRealm string, cidr int, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test_tags" {
    address = %q
    federated_realm = bloxone_federation_federated_realm.test.id
    cidr = %d
    tags = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}
