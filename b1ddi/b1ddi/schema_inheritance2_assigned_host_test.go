package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenInheritance2AssignedHost(t *testing.T) {
	cases := map[string]struct {
		input  *models.Inheritance2AssignedHost
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.Inheritance2AssignedHost{
				DisplayName: "unit-test-display-name",
				Host:        "unit-test-host",
				Ophid:       "unit-test-ophid",
			},
			expect: []interface{}{
				map[string]interface{}{
					"display_name": "unit-test-display-name",
					"host":         "unit-test-host",
					"ophid":        "unit-test-ophid",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenInheritance2AssignedHost(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
