package ipamfederation_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func TestAccFederatedPoolResource_basic(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName, "ip4", "NIOS_X", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "name", poolName),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
					resource.TestCheckResourceAttr(resourceName, "provider_type", "NIOS_X"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_federated_pool.test"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName, "ip4", "NIOS_X", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					testAccCheckFederatedPoolDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFederatedPoolResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_tags"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolTags(realmName, poolName, "ip4", "NIOS_X", "us-east-1", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedPoolTags(realmName, poolName, "ip4", "NIOS_X", "us-east-1", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_Description(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_description"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolDescription(realmName, poolName, "ip4", "NIOS_X", "us-east-1", "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedPoolDescription(realmName, poolName, "ip4", "NIOS_X", "us-east-1", "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_name"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolName(realmName, "ip4", "NIOS_X", "us-east-1", "test-pool-name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "test-pool-name"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedPoolName(realmName, "ip4", "NIOS_X", "us-east-1", "test-pool-name-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "test-pool-name-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_Metadata(t *testing.T) {
	t.Skip("metadata is a read-only field on the API — it cannot be set by the caller")
}

func TestAccFederatedPoolResource_NetworkCompliance(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_network_compliance"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolNetworkCompliance(realmName, poolName, "ip4", "NIOS_X", "us-east-1", 20, 17, 28),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.default_netmask_length", "20"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.minimum_netmask_length", "17"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.maximum_netmask_length", "28"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedPoolNetworkCompliance(realmName, poolName, "ip4", "NIOS_X", "us-east-1", 22, 18, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.default_netmask_length", "22"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.minimum_netmask_length", "18"),
					resource.TestCheckResourceAttr(resourceName, "network_compliance.maximum_netmask_length", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_Protocol(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test"
	var v1, v2 ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName, "ip4", "NIOS_X", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
				),
			},
			// Protocol cannot be updated on the API, so the pool is replaced
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName, "ip6", "NIOS_X", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v2),
					testAccCheckFederatedPoolRecreated(&v1, &v2),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedPoolResource_Parent(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.child"
	var v1, v2, v3, v4 ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	parentName := acctest.RandomNameWithPrefix("federated-pool")
	otherName := acctest.RandomNameWithPrefix("federated-pool")
	childName := acctest.RandomNameWithPrefix("federated-pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create child pool under the first parent
			{
				Config: testAccFederatedPoolParent(realmName, parentName, otherName, childName, "parent"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_federation_federated_pool.parent", "id"),
				),
			},
			// Changing the parent is not supported by the API, so the pool is replaced
			{
				Config: testAccFederatedPoolParent(realmName, parentName, otherName, childName, "other"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v2),
					testAccCheckFederatedPoolRecreated(&v1, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_federation_federated_pool.other", "id"),
				),
			},
			// Removing the parent also replaces the pool
			{
				Config: testAccFederatedPoolParent(realmName, parentName, otherName, childName, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v3),
					testAccCheckFederatedPoolRecreated(&v2, &v3),
					resource.TestCheckNoResourceAttr(resourceName, "parent"),
				),
			},
			// Updating other fields of a pool that has a parent must not send the parent
			{
				Config: testAccFederatedPoolParentWithDescription(realmName, parentName, otherName, childName, "parent", "updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v4),
					resource.TestCheckResourceAttr(resourceName, "description", "updated"),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_federation_federated_pool.parent", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFederatedPoolRecreated(before, after *ipamfederation.FederatedPool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if before.GetId() == after.GetId() {
			return fmt.Errorf("expected federated pool to be recreated, but id is unchanged: %s", before.GetId())
		}
		return nil
	}
}

func testAccCheckFederatedPoolExists(ctx context.Context, resourceName string, v *ipamfederation.FederatedPool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedPoolAPI.
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

func testAccCheckFederatedPoolDestroy(ctx context.Context, v *ipamfederation.FederatedPool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedPoolAPI.
			Read(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckFederatedPoolDisappears(ctx context.Context, v *ipamfederation.FederatedPool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedPoolAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFederatedPoolBasicConfig(realmName, poolName, protocol, providerType, region string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = %q
    provider_type   = %q
    region          = %q
}
`, poolName, protocol, providerType, region)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolDescription(realmName, poolName, protocol, providerType, region, description string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_description" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = %q
    provider_type   = %q
    region          = %q
    description     = %q
}
`, poolName, protocol, providerType, region, description)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolName(realmName, protocol, providerType, region, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_name" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = %q
    provider_type   = %q
    region          = %q
    name            = %q
}
`, protocol, providerType, region, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolNetworkCompliance(realmName, poolName, protocol, providerType, region string, defaultMask, minMask, maxMask int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_network_compliance" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = %q
    provider_type   = %q
    region          = %q
    network_compliance = {
        default_netmask_length = %d
        minimum_netmask_length = %d
        maximum_netmask_length = %d
    }
}
`, poolName, protocol, providerType, region, defaultMask, minMask, maxMask)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolTags(realmName, poolName, protocol, providerType, region string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
        %s = %q
`, k, v)
	}
	tagsStr += "    }"
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_tags" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = %q
    provider_type   = %q
    region          = %q
    tags            = %s
}
`, poolName, protocol, providerType, region, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolParent(realmName, parentName, otherName, childName, parentRef string) string {
	return testAccFederatedPoolParentWithDescription(realmName, parentName, otherName, childName, parentRef, "")
}

func testAccFederatedPoolParentWithDescription(realmName, parentName, otherName, childName, parentRef, description string) string {
	parentAttr := ""
	if parentRef != "" {
		parentAttr = fmt.Sprintf("parent = bloxone_federation_federated_pool.%s.id", parentRef)
	}
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = "ip4"
    region          = "us-east-1"
}

resource "bloxone_federation_federated_pool" "other" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = "ip4"
    region          = "us-east-1"
}

resource "bloxone_federation_federated_pool" "child" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = "ip4"
    region          = "us-east-1"
    description     = %q
    %s
}
`, parentName, otherName, childName, description, parentAttr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
