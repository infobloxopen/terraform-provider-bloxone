package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedAsmGrowthBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedAsmGrowthBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedAsmGrowthBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcAsmGrowthBlock{
					GrowthFactor: 20,
					GrowthType:   "percent",
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []interface{}{
						map[string]interface{}{
							"growth_factor": int64(20),
							"growth_type":   "percent",
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedAsmGrowthBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcInheritedAsmGrowthBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcInheritedAsmGrowthBlock
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
							"growth_factor": 20,
							"growth_type":   "percent",
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedAsmGrowthBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcAsmGrowthBlock{
					GrowthFactor: 20,
					GrowthType:   "percent",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedAsmGrowthBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
