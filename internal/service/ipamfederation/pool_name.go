package ipamfederation

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// poolNameSuffix matches the "(ipam-pool-<id>)" suffix the API appends to the name of pools that are
// backed by a cloud provider, e.g. "my-pool(ipam-pool-0073c8bb836f02aec)".
var poolNameSuffix = regexp.MustCompile(`\(ipam-pool-[A-Za-z0-9]+\)$`)

// stripPoolNameSuffix removes the provider-added suffix from a pool name returned by the API.
func stripPoolNameSuffix(name string) string {
	return poolNameSuffix.ReplaceAllString(name, "")
}

// flattenPoolName flattens the API pool name without the provider-added suffix.
func flattenPoolName(name *string) types.String {
	if name == nil {
		return types.StringNull()
	}
	return flex.FlattenString(stripPoolNameSuffix(*name))
}
