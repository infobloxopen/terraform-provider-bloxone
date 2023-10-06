package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedDDNSUpdateBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedDDNSUpdateBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedDDNSUpdateBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcDDNSUpdateBlock{
					DdnsDomain:      "unit-test-ddns-domain",
					DdnsSendUpdates: true,
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"ddns_domain":       "unit-test-ddns-domain",
							"ddns_send_updates": true,
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedDDNSUpdateBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})

	}
}

func TestExpandIpamsvcInheritedDDNSUpdateBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcInheritedDDNSUpdateBlock
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
							"ddns_domain":       "unit-test-ddns-domain",
							"ddns_send_updates": true,
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedDDNSUpdateBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcDDNSUpdateBlock{
					DdnsDomain:      "unit-test-ddns-domain",
					DdnsSendUpdates: true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedDDNSUpdateBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
