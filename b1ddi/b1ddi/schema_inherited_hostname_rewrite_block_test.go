package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedHostnameRewriteBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedHostnameRewriteBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedHostnameRewriteBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcHostnameRewriteBlock{
					HostnameRewriteChar:    "u",
					HostnameRewriteEnabled: true,
					HostnameRewriteRegex:   "unit-test-hostname-rewrite-regex",
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"hostname_rewrite_char":    "u",
							"hostname_rewrite_enabled": true,
							"hostname_rewrite_regex":   "unit-test-hostname-rewrite-regex",
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedHostnameRewriteBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcInheritedHostnameRewriteBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcInheritedHostnameRewriteBlock
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
							"hostname_rewrite_char":    "u",
							"hostname_rewrite_enabled": true,
							"hostname_rewrite_regex":   "unit-test-hostname-rewrite-regex",
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedHostnameRewriteBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcHostnameRewriteBlock{
					HostnameRewriteChar:    "u",
					HostnameRewriteEnabled: true,
					HostnameRewriteRegex:   "unit-test-hostname-rewrite-regex",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedHostnameRewriteBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
