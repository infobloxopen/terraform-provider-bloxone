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

func TestAccOptionCodeResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test"
	var optionSpace = "bloxone_dhcp_option_space.test"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeBasicConfig("234", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "234"),
					resource.TestCheckResourceAttr(resourceName, "name", "basic_opt_code"),
					resource.TestCheckResourceAttrPair(resourceName, "option_space", optionSpace, "id"),
					resource.TestCheckResourceAttr(resourceName, "type", "boolean"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "source"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "array", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_option_code.test"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionCodeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionCodeBasicConfig("234", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					testAccCheckOptionCodeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOptionCodeResource_Array(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_array"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeArray("234", "basic_opt_code", "boolean", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "array", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeArray("234", "basic_opt_code", "boolean", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "array", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_Code(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_code"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeCode("234", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "234"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeCode("235", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "235"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_comment"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeComment("234", "basic_opt_code", "boolean", "boolean option code type"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "boolean option code type"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeComment("234", "basic_opt_code", "boolean", "boolean option code type update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "boolean option code type update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_name"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeName("234", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "basic_opt_code"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeName("234", "basic_opt_code_1", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "basic_opt_code_1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_OptionSpace(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_option_space"
	var optSpace1 = "bloxone_dhcp_option_space.test1"
	var optSpace2 = "bloxone_dhcp_option_space.test2"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeOptionSpace("234", "basic_opt_code_1", "boolean", optSpace1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "option_space", optSpace1, "id"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeOptionSpace("234", "basic_opt_code_1", "boolean", optSpace2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "option_space", optSpace2, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionCodeResource_Type(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_code.test_type"
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionCodeType("234", "basic_opt_code", "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "boolean"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionCodeType("234", "basic_opt_code", "int16"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "int16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOptionCodeExists(ctx context.Context, resourceName string, v *ipam.IpamsvcOptionCode) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionCodeAPI.
			OptionCodeRead(ctx, rs.Primary.ID).
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

func testAccCheckOptionCodeDestroy(ctx context.Context, v *ipam.IpamsvcOptionCode) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionCodeAPI.
			OptionCodeRead(ctx, *v.Id).
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

func testAccCheckOptionCodeDisappears(ctx context.Context, v *ipam.IpamsvcOptionCode) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionCodeAPI.
			OptionCodeDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOptionCodeBasicConfig(code, name, type_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
}
`, code, name, type_)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}

func testAccOptionCodeArray(code, name, type_, array string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_array" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
    array = %q
}
`, code, name, type_, array)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}

func testAccOptionCodeCode(code, name, type_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_code" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
}
`, code, name, type_)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}

func testAccOptionCodeComment(code string, name string, type_ string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_comment" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
    comment = %q
}
`, code, name, type_, comment)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}

func testAccOptionCodeName(code string, name string, type_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_name" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
}
`, code, name, type_)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}

func testAccOptionCodeOptionSpace(code string, name string, type_, optSpaceName string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_option_space" {
    code = %q
    name = %q
    option_space = %s.id
    type = %q
}
`, code, name, optSpaceName, type_)

	return strings.Join([]string{testAccOptionSpaceMultiple("test_option_space_1", "ip4", "test_option_space_2", "ip4"), config}, "")
}

func testAccOptionCodeType(code string, name string, type_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test_type" {
    code = %q
    name = %q
    option_space = bloxone_dhcp_option_space.test.id
    type = %q
}
`, code, name, type_)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}
