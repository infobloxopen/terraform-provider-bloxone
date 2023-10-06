package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcDDNSUpdateBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcDDNSUpdateBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcDDNSUpdateBlock{
				DdnsDomain:      "unit-test-ddns-domain",
				DdnsSendUpdates: true,
			},
			expect: []interface{}{
				map[string]interface{}{
					"ddns_domain":       "unit-test-ddns-domain",
					"ddns_send_updates": true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcDDNSUpdateBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcDDNSUpdateBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcDDNSUpdateBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"ddns_domain":       "unit-test-ddns-domain",
					"ddns_send_updates": true,
				},
			},
			expect: &models.IpamsvcDDNSUpdateBlock{
				DdnsDomain:      "unit-test-ddns-domain",
				DdnsSendUpdates: true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcDDNSUpdateBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
