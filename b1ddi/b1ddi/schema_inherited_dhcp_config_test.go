package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenIpamsvcInheritedDHCPConfig(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedDHCPConfig
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedDHCPConfig{
				AllowUnknown: &models.InheritanceInheritedBool{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       true,
				},
				Filters: &models.InheritedDHCPConfigFilterList{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: []string{
						"unit-test-value",
					},
				},
				IgnoreList: &models.InheritedDHCPConfigIgnoreItemList{
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
				LeaseTime: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"allow_unknown": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        true,
						},
					},
					"filters": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []string{
								"unit-test-value",
							},
						},
					},
					"ignore_list": []interface{}{
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
					"lease_time": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedDHCPConfig(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcInheritedDHCPConfig(t *testing.T) {
	cases := map[string]struct {
		input  []interface{}
		expect *models.IpamsvcInheritedDHCPConfig
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"allow_unknown": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        true,
						},
					},
					"filters": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []string{
								"unit-test-value",
							},
						},
					},
					"ignore_list": []interface{}{
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
					"lease_time": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedDHCPConfig{
				AllowUnknown: &models.InheritanceInheritedBool{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       true,
				},
				Filters: &models.InheritedDHCPConfigFilterList{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: []string{
						"unit-test-value",
					},
				},
				IgnoreList: &models.InheritedDHCPConfigIgnoreItemList{
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
				LeaseTime: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := expandIpamsvcInheritedDHCPConfig(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}
