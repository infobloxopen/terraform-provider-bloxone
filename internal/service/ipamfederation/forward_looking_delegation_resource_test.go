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

func TestAccForwardLookingDelegationResource_basic(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.10.0.0"
	cidr := 16

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationBasicConfig(realmName, address, cidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "cidr", fmt.Sprintf("%d", cidr)),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.0", "bloxone_federation_federated_realm.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "name", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_forward_looking_delegation.test"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.20.0.0"
	cidr := 16

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardLookingDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardLookingDelegationBasicConfig(realmName, address, cidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					testAccCheckForwardLookingDelegationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccForwardLookingDelegationResource_Address(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_address"
	var v1 ipamfederation.ForwardLookingDelegation
	var v2 ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address1 := "10.30.0.0"
	address2 := "10.31.0.0"
	cidr := 16

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationAddress(realmName, address1, cidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", address1),
					resource.TestCheckResourceAttr(resourceName, "cidr", fmt.Sprintf("%d", cidr)),
				),
			},
			// Update triggers replace (RequiresReplaceIfConfigured)
			{
				Config: testAccForwardLookingDelegationAddress(realmName, address2, cidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", address2),
					resource.TestCheckResourceAttr(resourceName, "cidr", fmt.Sprintf("%d", cidr)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_cidr"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.40.0.0"
	cidr1 := 16
	cidr2 := 17

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationCidr(realmName, address, cidr1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "cidr", fmt.Sprintf("%d", cidr1)),
				),
			},
			// Update and Read
			{
				Config: testAccForwardLookingDelegationCidr(realmName, address, cidr2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "cidr", fmt.Sprintf("%d", cidr2)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_comment"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.50.0.0"
	cidr := 16
	comment1 := "test comment"
	comment2 := "test comment updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationComment(realmName, address, cidr, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccForwardLookingDelegationComment(realmName, address, cidr, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_FederatedPoolId(t *testing.T) {
	t.Skip("FederatedPoolId requires a valid federated pool resource ID; to be added when a pool fixture is available")
}

func TestAccForwardLookingDelegationResource_FederatedRealms(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_federated_realms"
	var v ipamfederation.ForwardLookingDelegation
	realmName1 := acctest.RandomNameWithPrefix("fld-realm")
	realmName2 := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.60.0.0"
	cidr := 16

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read with one realm
			{
				Config: testAccForwardLookingDelegationFederatedRealms(realmName1, realmName2, address, cidr, "[bloxone_federation_federated_realm.one.id]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.0", "bloxone_federation_federated_realm.one", "id"),
				),
			},
			// Update to include two realms
			{
				Config: testAccForwardLookingDelegationFederatedRealms(realmName1, realmName2, address, cidr, "[bloxone_federation_federated_realm.one.id, bloxone_federation_federated_realm.two.id]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms.#", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_name"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.70.0.0"
	cidr := 16
	name1 := "test-fld-name"
	name2 := "test-fld-name-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationName(realmName, address, cidr, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccForwardLookingDelegationName(realmName, address, cidr, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardLookingDelegationResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_forward_looking_delegation.test_tags"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.80.0.0"
	cidr := 16

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardLookingDelegationTags(realmName, address, cidr, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardLookingDelegationTags(realmName, address, cidr, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckForwardLookingDelegationExists(ctx context.Context, resourceName string, v *ipamfederation.ForwardLookingDelegation) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			ForwardLookingDelegationAPI.
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

func testAccCheckForwardLookingDelegationDestroy(ctx context.Context, v *ipamfederation.ForwardLookingDelegation) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			ForwardLookingDelegationAPI.
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

func testAccCheckForwardLookingDelegationDisappears(ctx context.Context, v *ipamfederation.ForwardLookingDelegation) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			ForwardLookingDelegationAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccForwardLookingDelegationBasicConfig(realmName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationAddress(realmName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_address" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationCidr(realmName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_cidr" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationComment(realmName, address string, cidr int, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_comment" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
    comment          = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationFederatedRealms(realmName1, realmName2, address string, cidr int, realmsExpr string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_federated_realms" {
    address          = %q
    cidr             = %d
    federated_realms = %s
}
`, address, cidr, realmsExpr)
	return strings.Join([]string{testAccBaseWithTwoFederatedRealm(realmName1, realmName2), config}, "")
}

func testAccForwardLookingDelegationName(realmName, address string, cidr int, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_name" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
    name             = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationTags(realmName, address string, cidr int, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
        %s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test_tags" {
    address          = %q
    cidr             = %d
    federated_realms = [bloxone_federation_federated_realm.test.id]
    tags             = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
