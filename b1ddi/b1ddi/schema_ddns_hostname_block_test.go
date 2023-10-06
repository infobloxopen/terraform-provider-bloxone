package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcDDNSHostnameBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcDDNSHostnameBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcDDNSHostnameBlock{
				DdnsGenerateName:    true,
				DdnsGeneratedPrefix: "unit-test-prefix",
			},
			expect: []interface{}{
				map[string]interface{}{
					"ddns_generate_name":    true,
					"ddns_generated_prefix": "unit-test-prefix",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcDDNSHostnameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcDDNSHostnameBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcDDNSHostnameBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"ddns_generate_name":    true,
					"ddns_generated_prefix": "unit-test-prefix",
				},
			},
			expect: &models.IpamsvcDDNSHostnameBlock{
				DdnsGenerateName:    true,
				DdnsGeneratedPrefix: "unit-test-prefix",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcDDNSHostnameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
