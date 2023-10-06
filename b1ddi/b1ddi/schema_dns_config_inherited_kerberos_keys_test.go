package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedKerberosKeys(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedKerberosKeys
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: &models.ConfigInheritedKerberosKeys{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: []*models.ConfigKerberosKey{
					{
						Algorithm:  "unit-test-algorithm",
						Domain:     "unit-test-domain",
						Key:        swag.String("unit-test-key"),
						Principal:  "unit-test-principal",
						UploadedAt: "unit-test-uploaded-at",
						Version:    int64(0),
					},
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"action":       "inherit",
					"display_name": "unit-test-display-name",
					"source":       "unit-test-source",
					"value": []map[string]interface{}{
						{
							"algorithm":   "unit-test-algorithm",
							"domain":      "unit-test-domain",
							"key":         swag.String("unit-test-key"),
							"principal":   "unit-test-principal",
							"uploaded_at": "unit-test-uploaded-at",
							"version":     int64(0),
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedKerberosKeys(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
