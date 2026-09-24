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
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	poolName := acctest.RandomNameWithPrefix("federated_pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "name", poolName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
				),
			},
		},
	})
}

func TestAccFederatedPoolResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_federated_pool.test"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	poolName := acctest.RandomNameWithPrefix("federated_pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolBasicConfig(realmName, poolName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					testAccCheckFederatedPoolDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFederatedPoolResource_Description(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_description"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	poolName := acctest.RandomNameWithPrefix("federated_pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolDescription(realmName, poolName, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			{
				Config: testAccFederatedPoolDescription(realmName, poolName, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
		},
	})
}

func TestAccFederatedPoolResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_name"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	name1 := acctest.RandomNameWithPrefix("federated_pool")
	name2 := acctest.RandomNameWithPrefix("federated_pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolName(realmName, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			{
				Config: testAccFederatedPoolName(realmName, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
		},
	})
}

func TestAccFederatedPoolResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_federated_pool.test_tags"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	poolName := acctest.RandomNameWithPrefix("federated_pool")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolTags(realmName, poolName, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			{
				Config: testAccFederatedPoolTags(realmName, poolName, map[string]string{"tag1": "value1updated"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1updated"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
		},
	})
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

func testAccFederatedPoolBasicConfig(realmName, poolName string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
}
`, poolName)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolDescription(realmName, poolName, description string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_description" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
    description     = %q
}
`, poolName, description)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolName(realmName, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_name" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
}
`, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolTags(realmName, poolName string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("        %s = %q\n", k, v)
	}
	tagsStr += "    }"
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test_tags" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
    tags = %s
}
`, poolName, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
