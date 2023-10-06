package b1ddi

import (
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenConfigInheritedZoneAuthority(t *testing.T) {
	cases := map[string]struct {
		input  *models.ConfigInheritedZoneAuthority
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.ConfigInheritedZoneAuthority{
				DefaultTTL: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Expire: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MnameBlock: &models.ConfigInheritedZoneAuthorityMNameBlock{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: &models.ConfigZoneAuthorityMNameBlock{
						Mname:           "unit-test-mname",
						ProtocolMname:   "unit-test-mname",
						UseDefaultMname: true,
					},
				},
				NegativeTTL: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				ProtocolRname: &models.Inheritance2InheritedString{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       "unit-test-value",
				},
				Refresh: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Retry: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Rname: &models.Inheritance2InheritedString{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       "unit-test-value",
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"default_ttl": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"expire": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"mname_block": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"mname":             "unit-test-mname",
									"protocol_mname":    "unit-test-mname",
									"use_default_mname": true,
								},
							},
						},
					},
					"negative_ttl": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"protocol_rname": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        "unit-test-value",
						},
					},
					"refresh": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"retry": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"rname": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        "unit-test-value",
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenConfigInheritedZoneAuthority(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandConfigInheritedZoneAuthority(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.ConfigInheritedZoneAuthority
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"default_ttl": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"expire": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"mname_block": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"mname":             "unit-test-mname",
									"protocol_mname":    "unit-test-mname",
									"use_default_mname": true,
								},
							},
						},
					},
					"negative_ttl": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"protocol_rname": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        "unit-test-value",
						},
					},
					"refresh": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"retry": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"rname": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        "unit-test-value",
						},
					},
				},
			},
			expect: &models.ConfigInheritedZoneAuthority{
				DefaultTTL: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Expire: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MnameBlock: &models.ConfigInheritedZoneAuthorityMNameBlock{
					Action: "inherit",
					Source: "unit-test-source",
				},
				NegativeTTL: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				ProtocolRname: &models.Inheritance2InheritedString{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       "unit-test-value",
				},
				Refresh: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Retry: &models.Inheritance2InheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				Rname: &models.Inheritance2InheritedString{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       "unit-test-value",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandConfigInheritedZoneAuthority(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
