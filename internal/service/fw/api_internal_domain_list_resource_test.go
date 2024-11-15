package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccInternalDomainListResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_internal_domain_list.test"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainListBasicConfig(name, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalDomainListResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalDomainListDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccInternalDomainListBasicConfig(name, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					testAccCheckInternalDomainListsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccInternalDomainListResource_Description(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test_description"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainDescription(name, "example.somedomain.com", "TEST_DESCRIPTION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION"),
				),
			},
			// Update and Read
			{
				Config: testAccInternalDomainDescription(name, "example.somedomain.com", "TEST_DESCRIPTION_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "TEST_DESCRIPTION_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalDomainListResource_InternalDomains(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test_internal_domain"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainListInternalDomain(name, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.0", "example.somedomain.com"),
				),
			},
			// Update and Read
			{
				Config: testAccInternalDomainListInternalDomain(name, "example.newdomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.0", "example.newdomain.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalDomainListResource_InternalDomains_Multiple(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test_internal_domain_multiple"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainListInternalDomainMultiple(name, []string{"example.somedomain.com", "dev.somedomain.com", "internal.somedomain.com", "test.somedomain.com"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.0", "example.somedomain.com"),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.1", "dev.somedomain.com"),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.2", "internal.somedomain.com"),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.3", "test.somedomain.com"),
				),
			},
			// Update and Read
			{
				Config: testAccInternalDomainListInternalDomainMultiple(name, []string{"test.newdomain.com", "internal.newdomain.com", "newdomain.com"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.0", "test.newdomain.com"),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.1", "internal.newdomain.com"),
					resource.TestCheckResourceAttr(resourceName, "internal_domains.2", "newdomain.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalDomainListResource_Name(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test_name"
	var v1, v2 fw.InternalDomains
	var name1 = acctest.RandomNameWithPrefix("td-internal_domain_list")
	var name2 = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainName(name1, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccInternalDomainName(name2, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListDestroy(context.Background(), &v1),
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalDomainListResource_Tags(t *testing.T) {
	resourceName := "bloxone_td_internal_domain_list.test_tags"
	var v fw.InternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccInternalDomainTags(name, "example.somedomain.com", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccInternalDomainTags(name, "example.somedomain.com", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInternalDomainListExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckInternalDomainListExists(ctx context.Context, resourceName string, v *fw.InternalDomains) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			InternalDomainListsAPI.
			ReadInternalDomains(ctx, int32(id)).
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

func testAccCheckInternalDomainListDestroy(ctx context.Context, v *fw.InternalDomains) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			InternalDomainListsAPI.
			ReadInternalDomains(ctx, *v.Id).
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

func testAccCheckInternalDomainListsDisappears(ctx context.Context, v *fw.InternalDomains) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			InternalDomainListsAPI.
			DeleteSingleInternalDomains(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccInternalDomainListBasicConfig(name, internalDomains string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test" {
	name = %q
	internal_domains = [%q]
}
`, name, internalDomains)
}

func testAccInternalDomainDescription(name string, internalDomains, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test_description" {
    name = %q
	internal_domains = [%q]
    description = %q
}
`, name, internalDomains, description)
}

func testAccInternalDomainListInternalDomain(name, internalDomains string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test_internal_domain" {
	name = %q
	internal_domains = [%q]
}
`, name, internalDomains)
}

func testAccInternalDomainListInternalDomainMultiple(name string, internalDomains []string) string {
	// join internalDomains into a string separated by commas, and quote each element
	internalDomainsStr := ""
	for i, domain := range internalDomains {
		if i > 0 {
			internalDomainsStr += ","
		}
		internalDomainsStr += fmt.Sprintf(`%q`, domain)
	}

	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test_internal_domain_multiple" {
	name = %q
	internal_domains = [%s]
}
`, name, internalDomainsStr)
}

func testAccInternalDomainName(name, internalDomains string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test_name" {
    name = %q
	internal_domains = [%q]
}
`, name, internalDomains)
}

func testAccInternalDomainTags(name string, internalDomains string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test_tags" {
    name = %q
	internal_domains = [%q]
    tags = %s
}
`, name, internalDomains, tagsStr)
}
