package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - dfps
// - net_address_dfps
// - roaming_device_groups
// - user_groups

func TestAccSecurityPoliciesResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_security_policy.test"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
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

func TestAccSecurityPoliciesResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPoliciesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPoliciesBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					testAccCheckSecurityPoliciesDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSecurityPoliciesResource_Name(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_name"
	var v fw.AtcfwSecurityPolicy
	name1 := acctest.RandomNameWithPrefix("sec-policy")
	name2 := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Description(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_description"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesDescription(name, "TEST_DESCRIPTION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesDescription(name, "TEST_DESCRIPTION_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_AccessCodes(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_access_codes"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesAccessCodes(name, "ac_test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "access_codes.0", "bloxone_td_access_code.ac_test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesAccessCodes(name, "ac_test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "access_codes.0", "bloxone_td_access_code.ac_test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_DefaultAction(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_default_action"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesDefaultAction(name, "action_allow"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_action", "action_allow"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesDefaultAction(name, "action_redirect"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_action", "action_redirect"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_DefaultRedirectName(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_default_redirect_name"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesDefaultRedirectName(name, "redirect_a"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_redirect_name", "redirect_a"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesDefaultRedirectName(name, "redirect_b"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_redirect_name", "redirect_b"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Ecs(t *testing.T) {
	resourceName := "bloxone_td_security_policy.test_ecs"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesEcs(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesEcs(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_NetworkLists(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_network_lists"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesNetworkLists(name, "nl_test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "network_lists.0", "bloxone_td_network_list.nl_test1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesNetworkLists(name, "nl_test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "network_lists.0", "bloxone_td_network_list.nl_test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_OnpremResolve(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_onprem_resolve"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesOnpremResolve(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "onprem_resolve", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesOnpremResolve(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "onprem_resolve", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Precedence(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_precedence"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesPrecedence(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "precedence", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesPrecedence(name, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "precedence", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Rules(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_rules"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesRules(name, "action_allow", "nl_test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.nl_test1", "name"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesRules(name, "action_log", "nl_test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "rules.0.data", "bloxone_td_named_list.nl_test2", "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_SafeSearch(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_safe_search"
	var v fw.AtcfwSecurityPolicy
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesSafeSearch(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "safe_search", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesSafeSearch(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "safe_search", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSecurityPoliciesResource_Tags(t *testing.T) {
	t.Skip()
	resourceName := "bloxone_td_security_policy.test_tags"
	var v fw.AtcfwSecurityPolicy
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSecurityPoliciesTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccSecurityPoliciesTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSecurityPoliciesExists(ctx context.Context, resourceName string, v *fw.AtcfwSecurityPolicy) resource.TestCheckFunc {
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
			SecurityPoliciesReadSecurityPolicy(ctx, int32(id)).
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

func testAccCheckSecurityPoliciesDestroy(ctx context.Context, v *fw.AtcfwSecurityPolicy) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			SecurityPoliciesAPI.
			SecurityPoliciesReadSecurityPolicy(ctx, *v.Id).
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

func testAccCheckSecurityPoliciesDisappears(ctx context.Context, v *fw.AtcfwSecurityPolicy) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			SecurityPoliciesAPI.
			SecurityPoliciesDeleteSingleSecurityPolicy(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSecurityPoliciesBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
	name=%q
}
`, name)
}

func testAccSecurityPoliciesName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_name" {
	name=%q
}
`, name)
}

func testAccSecurityPoliciesDescription(name, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_description" {
	name = %q
	description = %q
}
`, name, description)
}

func testAccSecurityPoliciesAccessCodes(name, accessCode string) string {
	act := time.Now().UTC().Format(time.RFC3339)
	exp := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "ac_test1" {
	name = "terraform-test-ac1"
	activation = %[1]q
	expiration = %[2]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

resource "bloxone_td_access_code" "ac_test2" {
	name = "terraform-test-ac2"
	activation = %[1]q
	expiration = %[2]q
	rules = [
		{
			data = "terraform_test",
			type = "custom_list"
		}
	]
}

resource "bloxone_td_security_policy" "test_access_codes" {
	name = %[3]q
	access_codes = [bloxone_td_access_code.%[4]s.id]

}
`, act, exp, name, accessCode)
}

func testAccSecurityPoliciesDefaultAction(name, defaultAction string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_default_action" {
	name = %q
	default_action = %q
}
`, name, defaultAction)
}

func testAccSecurityPoliciesDefaultRedirectName(name, defaultRedirectName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_default_redirect_name" {
	name = %q
	default_action = "action_redirect"
	default_redirect_name = %q
}
`, name, defaultRedirectName)
}

func testAccSecurityPoliciesEcs(name, ecs string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_ecs" {
	name = %q
	ecs = %q
}
`, name, ecs)
}

func testAccSecurityPoliciesNetworkLists(name, accessCode string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "nl_test1" {
	name = "terraform-test-nl1"
	items = ["156.2.3.0/24"]
}

resource "bloxone_td_network_list" "nl_test2" {
	name = "terraform-test-nl2"
	items = ["176.2.3.0/24"]
}

resource "bloxone_td_security_policy" "test_network_lists" {
	name = %q
	network_lists = [bloxone_td_network_list.%s.id]
}
`, name, accessCode)
}

func testAccSecurityPoliciesOnpremResolve(name, onpremResolve string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_onprem_resolve" {
	name = %q
	onprem_resolve = %q
}
`, name, onpremResolve)
}

func testAccSecurityPoliciesPrecedence(name string, precedence int32) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_precedence" {
	name = %q
	precedence = %d
}
`, name, precedence)
}

func testAccSecurityPoliciesSafeSearch(name, safeSearch string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_safe_search" {
	name = %q
	safe_search = %q
}
`, name, safeSearch)
}

func testAccSecurityPoliciesRules(name, rulesAction, rulesData string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "nl_test1" {
	name = "terraform-test-nl1"
	items_described = [
	{
		item = "tf1-domain.com"
		description = "Exaample Domain"
	}
	]
	type = "custom_list"
}

resource "bloxone_td_named_list" "nl_test2" {
	name = "terraform-test-nl2"
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
`, name, rulesAction, rulesData)
}

func testAccSecurityPoliciesRulesTags(name, rulesAction string, rulesTags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range rulesTags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test_rules" {
	name = %q
	rules = [
		{
			action = %q
			tag_scope = "custom_list"
			tags = %s
			type = "rule_tag"
		}
	]
}
`, name, rulesAction, tagsStr)
}

func testAccSecurityPoliciesTags(name string, tags map[string]string) string {
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
