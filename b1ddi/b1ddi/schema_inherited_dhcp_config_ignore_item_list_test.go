package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritedDHCPConfigIgnoreItemList(t *testing.T) {
	cases := map[string]struct {
		input    *models.InheritedDHCPConfigIgnoreItemList
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput": {
			input: &models.InheritedDHCPConfigIgnoreItemList{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []*models.IpamsvcIgnoreItem{
					{
						Type:  swag.String("hardware"),
						Value: swag.String("unit-test-value"),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []map[string]interface{}{
						{
							"type":  swag.String("hardware"),
							"value": swag.String("unit-test-value"),
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritedDHCPConfigIgnoreItemList(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandInheritedDHCPConfigIgnoreItemList(t *testing.T) {
	cases := map[string]struct {
		input    []interface{}
		expected *models.InheritedDHCPConfigIgnoreItemList
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
							"type":  "hardware",
							"value": "unit-test-value",
						},
					},
				},
			},
			expected: &models.InheritedDHCPConfigIgnoreItemList{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []*models.IpamsvcIgnoreItem{
					{
						Type:  swag.String("hardware"),
						Value: swag.String("unit-test-value"),
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandInheritedDHCPConfigIgnoreItemList(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
