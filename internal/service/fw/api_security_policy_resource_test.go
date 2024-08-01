package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - net_address_dfps

func TestAccSecurityPolicyResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_security_policy.test"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "default_action", "action_allow"),
					resource.TestCheckResourceAttr(resourceName, "default_redirect_name", ""),
					resource.TestCheckResourceAttr(resourceName, "ecs", "false"),
					resource.TestCheckResourceAttr(resourceName, "onprem_resolve", "false"),
					resource.TestCheckResourceAttr(resourceName, "safe_search", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPolicyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPolicyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					testAccCheckSecurityPolicyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSecurityPolicyResource_Name(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_name"
	var v fw.SecurityPolicy
	name1 := acctest.RandomNameWithPrefix("sec-policy")
	name2 := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_Description(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_description"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyDescription(name, "TEST_DESCRIPTION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyDescription(name, "TEST_DESCRIPTION_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_AccessCodes(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_access_codes"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")
	namedListName := acctest.RandomNameWithPrefix("named-list")
	accessCodeName1 := acctest.RandomNameWithPrefix("ac")
	accessCodeName2 := acctest.RandomNameWithPrefix("ac")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyAccessCodes(name, "ac_test1", namedListName, accessCodeName1, accessCodeName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "access_codes.0", "bloxone_td_access_code.ac_test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyAccessCodes(name, "ac_test2", namedListName, accessCodeName1, accessCodeName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "access_codes.0", "bloxone_td_access_code.ac_test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_DefaultAction(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_default_action"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyDefaultAction(name, "action_allow"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_action", "action_allow"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyDefaultAction(name, "action_redirect"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_action", "action_redirect"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_DefaultRedirectName(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_default_redirect_name"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")
	redirectName1 := acctest.RandomNameWithPrefix("redirect")
	redirectName2 := acctest.RandomNameWithPrefix("redirect")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyDefaultRedirectName(name, "test_a", redirectName1, redirectName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "default_redirect_name", "bloxone_td_custom_redirect.test_a", "name"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyDefaultRedirectName(name, "test_b", redirectName1, redirectName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "default_redirect_name", "bloxone_td_custom_redirect.test_b", "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_Ecs(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_ecs"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyEcs(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyEcs(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_NetworkLists(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_network_lists"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")
	networkListName1 := acctest.RandomNameWithPrefix("network-list")
	networkListName2 := acctest.RandomNameWithPrefix("network-list")
	ip1 := acctest.RandomIP() + "/32"
	ip2 := acctest.RandomIP() + "/32"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyNetworkLists(name, "nl_test1", networkListName1, ip1, networkListName2, ip2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "network_lists.0", "bloxone_td_network_list.nl_test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyNetworkLists(name, "nl_test2", networkListName1, ip1, networkListName2, ip2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "network_lists.0", "bloxone_td_network_list.nl_test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_OnpremResolve(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_onprem_resolve"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyOnpremResolve(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "onprem_resolve", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyOnpremResolve(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "onprem_resolve", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_Precedence(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_precedence"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyPrecedence(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "precedence", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyPrecedence(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "precedence", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_Rules(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_rules"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")
	namedListName1 := acctest.RandomNameWithPrefix("named-list")
	namedListName2 := acctest.RandomNameWithPrefix("named-list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyRules(name, "action_allow", "nl_test1", namedListName1, namedListName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.nl_test1", "name"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyRules(name, "action_log", "nl_test2", namedListName1, namedListName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.nl_test2", "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_SafeSearch(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_safe_search"
	var v fw.SecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicySafeSearch(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "safe_search", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicySafeSearch(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "safe_search", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPolicyResource_Tags(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_tags"
	var v fw.SecurityPolicy
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPolicyTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPolicyTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSecurityPolicyExists(ctx context.Context, resourceName string, v *fw.SecurityPolicy) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			SecurityPoliciesAPI.
			ReadSecurityPolicy(ctx, int32(id)).
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

func testAccCheckSecurityPolicyDestroy(ctx context.Context, v *fw.SecurityPolicy) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			SecurityPoliciesAPI.
			ReadSecurityPolicy(ctx, *v.Id).
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

func testAccCheckSecurityPolicyDisappears(ctx context.Context, v *fw.SecurityPolicy) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			SecurityPoliciesAPI.
			DeleteSingleSecurityPolicy(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSecurityPolicyBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
	name = %q
}
`, name)
}

func testAccSecurityPolicyName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_name" {
	name = %q
}
`, name)
}

func testAccSecurityPolicyDescription(name, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_description" {
	name = %q
	description = %q
}
`, name, description)
}

func testAccSecurityPolicyAccessCodes(name, accessCode, namedListName, accessCodeName1, accessCodeName2 string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_access_code" "ac_test1" {
	name = %[1]q
	activation = %[3]q
	expiration = %[4]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

resource "bloxone_td_access_code" "ac_test2" {
	name = %[2]q
	activation = %[3]q
	expiration = %[4]q
	rules = [
		{
			data = bloxone_td_named_list.test.name,
			type = bloxone_td_named_list.test.type
		}
	]
}

resource "bloxone_td_security_policy" "test_access_codes" {
	name = %[5]q
	access_codes = [bloxone_td_access_code.%[6]s.id]

}
`, accessCodeName1, accessCodeName2, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339), name, accessCode)
	return strings.Join([]string{testAccBaseWithNamedList(namedListName), config}, "")
}

func testAccSecurityPolicyDefaultAction(name, defaultAction string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_default_action" {
	name = %q
	default_action = %q
}
`, name, defaultAction)
}

func testAccSecurityPolicyDefaultRedirectName(name, defaultRedirect, redirectName1, redirectName2 string) string {
	return fmt.Sprintf(`
resource "bloxone_td_custom_redirect" "test_a" {
	name = %q
	data = "156.2.3.10"
}

resource "bloxone_td_custom_redirect" "test_b" {
	name = %q
	data = "192.2.3.10"
}

resource "bloxone_td_security_policy" "test_default_redirect_name" {
	name = %q
	default_action = "action_redirect"
	default_redirect_name = bloxone_td_custom_redirect.%s.name
}
`, redirectName1, redirectName2, name, defaultRedirect)
}

func testAccSecurityPolicyEcs(name, ecs string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_ecs" {
	name = %q
	ecs = %q
}
`, name, ecs)
}

func testAccSecurityPolicyNetworkLists(name, networkList, networkListName1, ip1, networkListName2, ip2 string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "nl_test1" {
	name = %q
	items = [%q]
}

resource "bloxone_td_network_list" "nl_test2" {
	name = %q
	items = [%q]
}

resource "bloxone_td_security_policy" "test_network_lists" {
	name = %q
	network_lists = [bloxone_td_network_list.%s.id]
}
`, networkListName1, ip1, networkListName2, ip2, name, networkList)
}

func testAccSecurityPolicyOnpremResolve(name, onpremResolve string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_onprem_resolve" {
	name = %q
	onprem_resolve = %q
}
`, name, onpremResolve)
}

func testAccSecurityPolicyPrecedence(name string, precedence int32) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_precedence" {
	name = %q
	precedence = %d
}
`, name, precedence)
}

func testAccSecurityPolicySafeSearch(name, safeSearch string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_safe_search" {
	name = %q
	safe_search = %q
}
`, name, safeSearch)
}

func testAccSecurityPolicyRules(name, rulesAction, rulesData, namedListName1, namedListName2 string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "nl_test1" {
	name = %q
	items_described = [
	{
		item = "tf1-domain.com"
		description = "Exaample Domain"
	}
	]
	type = "custom_list"
}

resource "bloxone_td_named_list" "nl_test2" {
	name = %q
	items_described = [
	{
		item = "tf2-domain.com"
		description = "Exaample Domain"
	}
	]
	type = "custom_list"
}

resource "bloxone_td_security_policy" "test_rules" {
	name = %q
	rules = [
		{
			action = %q
			data = bloxone_td_named_list.%s.name
			type = "custom_list"
		}
	]
}
`, namedListName1, namedListName2, name, rulesAction, rulesData)
}

func testAccSecurityPolicyTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_tags" {
    name = %q
	tags = %s
}
`, name, tagsStr)
}
