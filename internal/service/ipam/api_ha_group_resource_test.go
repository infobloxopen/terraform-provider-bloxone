package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHaGroupResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test"
	var v ipam.IpamsvcHAGroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupBasicConfig("active", "active", "test-ha", "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "hosts.0.host"),
					resource.TestCheckResourceAttr(resourceName, "hosts.0.role", "active"),
					resource.TestCheckResourceAttrSet(resourceName, "hosts.1.host"),
					resource.TestCheckResourceAttr(resourceName, "hosts.1.role", "active"),
					resource.TestCheckResourceAttr(resourceName, "name", "test-ha"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_ha_group.test"
	var v ipam.IpamsvcHAGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupBasicConfig("active", "active", "test-ha", "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					testAccCheckHaGroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHaGroupResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_comment"
	var v ipam.IpamsvcHAGroup
	name := acctest.RandomNameWithPrefix("test-ha")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupComment("active", "active", name, "active-active", "HA Group created with Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "HA Group created with Terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupComment("active", "active", name, "active-active", "HA Group was created with Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "HA Group was created with Terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Hosts(t *testing.T) {
	var (
		v               ipam.IpamsvcHAGroup
		resourceName    = "bloxone_ipam_ha_group.test_hosts"
		name            = acctest.RandomNameWithPrefix("test-ha")
		dataSourceHost1 = "data.bloxone_dhcp_hosts.test_01"
		dataSourceHost2 = "data.bloxone_dhcp_hosts.test_02"
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupHosts(dataSourceHost1, "active", dataSourceHost2, "passive", name, "active-passive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.0.host", dataSourceHost1, "results.0.id"),
					resource.TestCheckResourceAttr(resourceName, "hosts.0.role", "active"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.1.host", dataSourceHost2, "results.0.id"),
					resource.TestCheckResourceAttr(resourceName, "hosts.1.role", "passive"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupHosts(dataSourceHost2, "active", dataSourceHost1, "passive", name, "active-passive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.0.host", dataSourceHost2, "results.0.id"),
					resource.TestCheckResourceAttr(resourceName, "hosts.0.role", "active"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.1.host", dataSourceHost1, "results.0.id"),
					resource.TestCheckResourceAttr(resourceName, "hosts.1.role", "passive"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Mode(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_mode"
	var v ipam.IpamsvcHAGroup
	name := acctest.RandomNameWithPrefix("test-ha")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupMode("active", "active", name, "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "active-active"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupMode("active", "passive", name, "active-passive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "active-passive"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Name(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_name"
	var v ipam.IpamsvcHAGroup
	name := acctest.RandomNameWithPrefix("test-ha")
	updateName := acctest.RandomNameWithPrefix("test-ha-new")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupName("active", "active", name, "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupName("active", "active", updateName, "active-active"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHaGroupResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_ha_group.test_tags"
	var v ipam.IpamsvcHAGroup
	name := acctest.RandomNameWithPrefix("test-ha")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHaGroupTags("active", "passive", name, "active-passive", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccHaGroupTags("active", "passive", name, "active-passive", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHaGroupExists(ctx context.Context, resourceName string, v *ipam.IpamsvcHAGroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			HaGroupRead(ctx, rs.Primary.ID).
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

func testAccCheckHaGroupDestroy(ctx context.Context, v *ipam.IpamsvcHAGroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			HaGroupRead(ctx, *v.Id).
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

func testAccCheckHaGroupDisappears(ctx context.Context, v *ipam.IpamsvcHAGroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			HaGroupAPI.
			HaGroupDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccHaGroupBasicConfig(role1, role2, name, mode string) string {
	// TODO: create basic resource with required fields
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test" {
	hosts = [
		{
			host = data.bloxone_dhcp_hosts.test_01.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test_02.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
}
`, role1, role2, name, mode)

	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupComment(role1, role2, name, mode, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_comment" {
	hosts = [
		{
			host = data.bloxone_dhcp_hosts.test_01.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test_02.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
	comment = %q
}
`, role1, role2, name, mode, comment)

	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupHosts(host1, role1, host2, role2, name, mode string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_hosts" {
	hosts = [
		{
			host = %s.results.0.id
			role = %q
		},
		{
			host = %s.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
}
`, host1, role1, host2, role2, name, mode)
	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupMode(role1, role2, name, mode string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_mode" {
	hosts = [
		{
			host = data.bloxone_dhcp_hosts.test_01.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test_02.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
}
`, role1, role2, name, mode)
	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupName(role1, role2, name, mode string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_name" {
	hosts = [
		{
			host = data.bloxone_dhcp_hosts.test_01.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test_02.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
}
`, role1, role2, name, mode)
	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupTags(role1, role2, name, mode string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test_tags" {
    hosts = [
		{
			host = data.bloxone_dhcp_hosts.test_01.results.0.id
			role = %q
		},
		{
			host = data.bloxone_dhcp_hosts.test_02.results.0.id
			role = %q
		}
	]
	name = %q
	mode = %q
	tags = %s
}
`, role1, role2, name, mode, tagsStr)

	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}
