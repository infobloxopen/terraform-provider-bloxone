package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigZoneAuthorityMNameBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigZoneAuthorityMNameBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigZoneAuthorityMNameBlock{
				Mname:           "unit-test-mname",
				ProtocolMname:   "unit-test-mname",
				UseDefaultMname: true,
			},
			expect: []interface{}{
				map[string]interface{}{
					"mname":             "unit-test-mname",
					"protocol_mname":    "unit-test-mname",
					"use_default_mname": true,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigZoneAuthorityMNameBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
