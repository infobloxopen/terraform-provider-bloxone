package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcHostnameRewriteBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcHostnameRewriteBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcHostnameRewriteBlock{
				HostnameRewriteChar:    "u",
				HostnameRewriteEnabled: true,
				HostnameRewriteRegex:   "unit-test-hostname-rewrite-regex",
			},
			expect: []interface{}{
				map[string]interface{}{
					"hostname_rewrite_char":    "u",
					"hostname_rewrite_enabled": true,
					"hostname_rewrite_regex":   "unit-test-hostname-rewrite-regex",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcHostnameRewriteBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcHostnameRewriteBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcHostnameRewriteBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"hostname_rewrite_char":    "u",
					"hostname_rewrite_enabled": true,
					"hostname_rewrite_regex":   "unit-test-hostname-rewrite-regex",
				},
			},
			expect: &models.IpamsvcHostnameRewriteBlock{
				HostnameRewriteChar:    "u",
				HostnameRewriteEnabled: true,
				HostnameRewriteRegex:   "unit-test-hostname-rewrite-regex",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcHostnameRewriteBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
