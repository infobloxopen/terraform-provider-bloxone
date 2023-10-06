package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedDHCPOptionList(t *testing.T) {
	cases := map[string]struct {
		input    *models.IpamsvcInheritedDHCPOptionList
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedDHCPOptionList{
				Action: "inherit",
				Value: []*models.IpamsvcInheritedDHCPOption{
					{
						Action:      "inherit",
						DisplayName: "unit-test-display-name",
						Source:      "unit-test-source",
						Value: &models.IpamsvcOptionItem{
							Group:       "unit-test-group",
							OptionCode:  "unit-test-option-code",
							OptionValue: "unit-test-option-value",
							Type:        "option",
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"action": "inherit",
					"value": []map[string]interface{}{
						{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"group":        "unit-test-group",
									"option_code":  "unit-test-option-code",
									"option_value": "unit-test-option-value",
									"type":         "option",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedDHCPOptionList(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandIpamsvcInheritedDHCPOptionList(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.IpamsvcInheritedDHCPOptionList
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"action": "inherit",
					"value": []map[string]interface{}{
						{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"group":        "unit-test-group",
									"option_code":  "unit-test-option-code",
									"option_value": "unit-test-option-value",
									"type":         "option",
								},
							},
						},
					},
				},
			},
			expected: &models.IpamsvcInheritedDHCPOptionList{
				Action: "inherit",
				Value: []*models.IpamsvcInheritedDHCPOption{
					{
						Action:      "inherit",
						DisplayName: "unit-test-display-name",
						Source:      "unit-test-source",
						Value: &models.IpamsvcOptionItem{
							Group:       "unit-test-group",
							OptionCode:  "unit-test-option-code",
							OptionValue: "unit-test-option-value",
							Type:        "option",
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedDHCPOptionList(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
