package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedDDNSHostnameBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedDDNSHostnameBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedDDNSHostnameBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcDDNSHostnameBlock{
					DdnsGenerateName:    true,
					DdnsGeneratedPrefix: "unit-test-prefix",
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"ddns_generate_name":    true,
							"ddns_generated_prefix": "unit-test-prefix",
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			response := flattenIpamsvcInheritedDDNSHostnameBlock(tc.input)

			assert.Equal(t, tc.expect, response)
		})
	}
}

func TestExpandIpamsvcInheritedDDNSHostnameBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcInheritedDDNSHostnameBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"ddns_generate_name":    true,
							"ddns_generated_prefix": "unit-test-prefix",
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedDDNSHostnameBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcDDNSHostnameBlock{
					DdnsGenerateName:    true,
					DdnsGeneratedPrefix: "unit-test-prefix",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedDDNSHostnameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
