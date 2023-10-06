package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedACLItems(t *testing.T) {
	cases := map[string]struct {
		input    *models.ConfigInheritedACLItems
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedACLItems{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []*models.ConfigACLItem{
					{
						Access:  swag.String("allow"),
						ACL:     "unit-test-acl-id",
						Address: "unit-test-address",
						Element: swag.String("ip"),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"access":   swag.String("allow"),
							"acl":      "unit-test-acl-id",
							"address":  "unit-test-address",
							"element":  swag.String("ip"),
							"tsig_key": []interface{}{},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedACLItems(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandConfigInheritedACLItems(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.ConfigInheritedACLItems
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"access":   "allow",
							"acl":      "unit-test-acl-id",
							"address":  "unit-test-address",
							"element":  "ip",
							"tsig_key": []interface{}{},
						},
					},
				},
			},
			expected: &models.ConfigInheritedACLItems{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []*models.ConfigACLItem{
					{
						Access:  swag.String("allow"),
						ACL:     "unit-test-acl-id",
						Address: "unit-test-address",
						Element: swag.String("ip"),
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedACLItems(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
