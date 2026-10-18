package ipamfederation

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// poolNameSuffix matches the "(ipam-pool-<id>)" suffix the API appends to the name of pools that are
// backed by a cloud provider, e.g. "my-pool(ipam-pool-0073c8bb836f02aec)".
var poolNameSuffix = regexp.MustCompile(`\(ipam-pool-[A-Za-z0-9]+\)$`)

// stripPoolNameSuffix removes the provider-added suffix from a pool name returned by the API.
func stripPoolNameSuffix(name string) string {
	return poolNameSuffix.ReplaceAllString(name, "")
}

// flattenPoolName flattens the API pool name. The suffix is only removed for AWS pools, the only ones that get it,
// so a name chosen by the user for another pool is never altered.
func flattenPoolName(name *string, provider ipamfederation.ProviderType) types.String {
	if name == nil {
		return types.StringNull()
	}
	if provider != ipamfederation.PROVIDERTYPE_AWS {
		return types.StringValue(*name)
	}
	return flex.FlattenString(stripPoolNameSuffix(*name))
}
