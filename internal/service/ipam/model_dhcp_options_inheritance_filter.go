package ipam

import (
	"github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// dropInheritedOptionsSetExplicitly removes "inherit" entries from the
// inheritance_sources.dhcp_options list for options that are also set
// explicitly in dhcp_options.
//
// The inheritance_sources value is carried over from prior state. If it still
// holds an "inherit" entry for an option code that is now set explicitly with
// the same value the parent provides, the API treats the option as inherited
// and drops the explicit copy, so the apply returns fewer dhcp_options than
// planned. Entries with any other action (override, block) are kept.
func dropInheritedOptionsSetExplicitly(sources *ipam.InheritedDHCPOptionList, explicit []ipam.OptionItem) {
	if sources == nil || len(sources.Value) == 0 || len(explicit) == 0 {
		return
	}

	codes := make(map[string]struct{}, len(explicit))
	groups := make(map[string]struct{}, len(explicit))
	for _, o := range explicit {
		if o.OptionCode != nil {
			codes[*o.OptionCode] = struct{}{}
		}
		if o.Group != nil {
			groups[*o.Group] = struct{}{}
		}
	}

	kept := make([]ipam.InheritedDHCPOption, 0, len(sources.Value))
	for _, v := range sources.Value {
		if v.Action != nil && *v.Action == "inherit" && v.Value != nil && v.Value.Option != nil {
			opt := v.Value.Option
			if opt.OptionCode != nil {
				if _, ok := codes[*opt.OptionCode]; ok {
					continue
				}
			}
			if opt.Group != nil {
				if _, ok := groups[*opt.Group]; ok {
					continue
				}
			}
		}
		kept = append(kept, v)
	}
	sources.Value = kept
}
