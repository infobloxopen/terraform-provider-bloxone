package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigKerberosKey(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigKerberosKey
		expect map[string]interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: map[string]interface{}(nil),
		},
		"FullConfig": {
			input: &models.ConfigKerberosKey{
				Algorithm:  "unit-test-algorithm",
				Domain:     "unit-test-domain",
				Key:        swag.String("unit-test-key"),
				Principal:  "unit-test-principal",
				UploadedAt: "unit-test-uploaded-at",
				Version:    int64(0),
			},
			expect: map[string]interface{}{
				"algorithm":   "unit-test-algorithm",
				"domain":      "unit-test-domain",
				"key":         swag.String("unit-test-key"),
				"principal":   "unit-test-principal",
				"uploaded_at": "unit-test-uploaded-at",
				"version":     int64(0),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigKerberosKey(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
