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

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - dhcp_host
// - dhcp_options
// - filters
// - inheritance_sources - Currently inheritance sources is always nil

func TestAccRangeResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBasicConfig("10.0.0.8", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end", "10.0.0.20"),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "start", "10.0.0.8"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "utilization.%"),
					resource.TestCheckResourceAttrSet(resourceName, "utilization_v6.%"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_range.test"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeBasicConfig("10.0.0.8", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					testAccCheckRangeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRangeResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_comment"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeComment("10.0.0.8", "10.0.0.20", "this range is created by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this range is created by terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeComment("10.0.0.8", "10.0.0.20", "this range was created by terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this range was created by terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DisableDhcp(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_disable_dhcp"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDisableDhcp("10.0.0.8", "10.0.0.20", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDisableDhcp("10.0.0.8", "10.0.0.20", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_End(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_end"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnd("10.0.0.8", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end", "10.0.0.20"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnd("10.0.0.8", "10.0.0.29"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end", "10.0.0.29"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_ExclusionRanges(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_exclusion_ranges"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeExclusionRanges("10.0.0.8", "10.0.0.20", "10.0.0.16", "10.0.0.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclusion_ranges.0.start", "10.0.0.12"),
					resource.TestCheckResourceAttr(resourceName, "exclusion_ranges.0.end", "10.0.0.16"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeExclusionRanges("10.0.0.8", "10.0.0.20", "10.0.0.16", "10.0.0.14"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclusion_ranges.0.start", "10.0.0.14"),
					resource.TestCheckResourceAttr(resourceName, "exclusion_ranges.0.end", "10.0.0.16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Name(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_name"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeName("10.0.0.8", "10.0.0.20", "range-test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "range-test"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeName("10.0.0.8", "10.0.0.20", "range-test-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "range-test-1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Space(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_space"
	var v1, v2 ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSpace("10.0.0.8", "10.0.0.20", "bloxone_ipam_ip_space.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSpace("10.0.0.8", "10.0.0.20", "bloxone_ipam_ip_space.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeDestroy(context.Background(), &v1),
					testAccCheckRangeExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.one", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Start(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_start"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeStart("10.0.0.8", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start", "10.0.0.8"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeStart("10.0.0.12", "10.0.0.20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start", "10.0.0.12"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_range.test_tags"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeTags("10.0.0.8", "10.0.0.20", map[string]string{
					"site": "NA",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.site", "NA"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeTags("10.0.0.8", "10.0.0.20", map[string]string{
					"site": "CA",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.site", "CA"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRangeExists(ctx context.Context, resourceName string, v *ipam.IpamsvcRange) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			RangeAPI.
			RangeRead(ctx, rs.Primary.ID).
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

func testAccCheckRangeDestroy(ctx context.Context, v *ipam.IpamsvcRange) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			RangeAPI.
			RangeRead(ctx, *v.Id).
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

func testAccCheckRangeDisappears(ctx context.Context, v *ipam.IpamsvcRange) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			RangeAPI.
			RangeDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRangeBasicConfig(start, end string) string {
	// TODO: create basic resource with required fields
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeComment(start, end, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_comment" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    comment = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end, comment)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeDisableDhcp(start, end, disableDhcp string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_disable_dhcp" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    disable_dhcp = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end, disableDhcp)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeEnd(start, end string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_end" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeExclusionRanges(start, end string, exclusionEnd, exclusionStart string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_exclusion_ranges" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    exclusion_ranges = [
      {
        end = %q
        start = %q
      }
    ]
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end, exclusionEnd, exclusionStart)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeName(start, end string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_name" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    name = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end, name)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeSpace(start, end, space string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = "10.0.0.0"
    cidr = 8
    space = %s.id
}

resource "bloxone_ipam_range" "test_space" {
    space = %s.id
    start = %q
    end = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, space, space, start, end)
	return strings.Join([]string{testAccBaseWithTwoIPSpace(), config}, "")
}

func testAccRangeStart(start, end string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_start" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeTags(start, end string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test_tags" {
    space = bloxone_ipam_ip_space.test.id
    start = %q
    end = %q
    tags = %s
    depends_on = [bloxone_ipam_subnet.test]
}
`, start, end, tagsStr)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}
