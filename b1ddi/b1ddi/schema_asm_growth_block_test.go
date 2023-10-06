package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcAsmGrowthBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcAsmGrowthBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcAsmGrowthBlock{
				GrowthFactor: 20,
				GrowthType:   "percent",
			},
			expect: []interface{}{
				map[string]interface{}{
					"growth_factor": int64(20),
					"growth_type":   "percent",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcAsmGrowthBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcAsmGrowthBlock(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcAsmGrowthBlock
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"growth_factor": 20,
					"growth_type":   "percent",
				},
			},
			expect: &models.IpamsvcAsmGrowthBlock{
				GrowthFactor: 20,
				GrowthType:   "percent",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcAsmGrowthBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
