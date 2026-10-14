package ipam

import (
	"testing"

	"github.com/infobloxopen/universal-ddi-go-client/ipam"
)

func str(s string) *string { return &s }

func inheritedOption(action, code, value string) ipam.InheritedDHCPOption {
	return ipam.InheritedDHCPOption{
		Action: str(action),
		Value: &ipam.InheritedDHCPOptionItem{
			Option: &ipam.OptionItem{OptionCode: str(code), OptionValue: str(value), Type: str("option")},
		},
	}
}

func TestDropInheritedOptionsSetExplicitly(t *testing.T) {
	explicit := []ipam.OptionItem{
		{OptionCode: str("code/6"), OptionValue: str("192.0.2.53")},
		{OptionCode: str("code/15"), OptionValue: str("test.example")},
	}

	t.Run("drops inherit entries for explicit codes only", func(t *testing.T) {
		src := &ipam.InheritedDHCPOptionList{
			Action: str("inherit"),
			Value: []ipam.InheritedDHCPOption{
				inheritedOption("override", "code/6", "192.0.2.53,192.0.2.54"),
				inheritedOption("inherit", "code/15", "test.example"),
				inheritedOption("inherit", "code/3", "198.51.100.1"),
			},
		}
		dropInheritedOptionsSetExplicitly(src, explicit)
		if len(src.Value) != 2 {
			t.Fatalf("expected 2 entries, got %d", len(src.Value))
		}
		if *src.Value[0].Value.Option.OptionCode != "code/6" || *src.Value[1].Value.Option.OptionCode != "code/3" {
			t.Fatalf("unexpected entries kept: %+v", src.Value)
		}
		if *src.Action != "inherit" {
			t.Fatalf("list action must be unchanged")
		}
	})

	t.Run("keeps block entries", func(t *testing.T) {
		src := &ipam.InheritedDHCPOptionList{Value: []ipam.InheritedDHCPOption{inheritedOption("block", "code/15", "test.example")}}
		dropInheritedOptionsSetExplicitly(src, explicit)
		if len(src.Value) != 1 {
			t.Fatalf("block entry must be kept")
		}
	})

	t.Run("drops inherit entries for explicit groups", func(t *testing.T) {
		e := inheritedOption("inherit", "", "")
		e.Value.Option = &ipam.OptionItem{Group: str("group/1")}
		src := &ipam.InheritedDHCPOptionList{Value: []ipam.InheritedDHCPOption{e}}
		dropInheritedOptionsSetExplicitly(src, []ipam.OptionItem{{Group: str("group/1")}})
		if len(src.Value) != 0 {
			t.Fatalf("group entry must be dropped")
		}
	})

	t.Run("nil and empty inputs", func(t *testing.T) {
		dropInheritedOptionsSetExplicitly(nil, explicit)
		src := &ipam.InheritedDHCPOptionList{Value: []ipam.InheritedDHCPOption{inheritedOption("inherit", "code/15", "x")}}
		dropInheritedOptionsSetExplicitly(src, nil)
		if len(src.Value) != 1 {
			t.Fatalf("no explicit options: nothing should be dropped")
		}
		dropInheritedOptionsSetExplicitly(&ipam.InheritedDHCPOptionList{}, explicit)
	})
}
