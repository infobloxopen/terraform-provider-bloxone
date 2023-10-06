package b1ddi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFlattenIpamsvcInheritedASMConfig(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedASMConfig
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedASMConfig{
				AsmEnableBlock: &models.IpamsvcInheritedAsmEnableBlock{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: &models.IpamsvcAsmEnableBlock{
						Enable:             true,
						EnableNotification: true,
						ReenableDate:       strfmt.NewDateTime(),
					},
				},
				AsmGrowthBlock: &models.IpamsvcInheritedAsmGrowthBlock{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: &models.IpamsvcAsmGrowthBlock{
						GrowthFactor: 20,
						GrowthType:   "percent",
					},
				},
				AsmThreshold: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				ForecastPeriod: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				History: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MinTotal: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MinUnused: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
			},
			expect: []interface{}{
				map[string]interface{}{
					"asm_enable_block": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"enable":              true,
									"enable_notification": true,
									"reenable_date":       "1970-01-01T00:00:00.000Z",
								},
							},
						},
					},
					"asm_growth_block": []interface{}{
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
					"asm_threshold": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"forecast_period": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"history": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"min_total": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        int64(20),
						},
					},
					"min_unused": []interface{}{
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
			result := flattenIpamsvcInheritedASMConfig(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcInheritedASMConfig(t *testing.T) {
	cases := map[string]struct {
		input       []interface{}
		expect      *models.IpamsvcInheritedASMConfig
		expectError error
	}{
		"NilInput": {
			input:  nil,
			expect: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"asm_enable_block": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"enable":              true,
									"enable_notification": true,
									"reenable_date":       "1970-01-01T00:00:00.000Z",
								},
							},
						},
					},
					"asm_growth_block": []interface{}{
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
					"asm_threshold": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"forecast_period": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"history": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"min_total": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"min_unused": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedASMConfig{
				AsmEnableBlock: &models.IpamsvcInheritedAsmEnableBlock{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: &models.IpamsvcAsmEnableBlock{
						Enable:             true,
						EnableNotification: true,
						ReenableDate:       strfmt.NewDateTime(),
					},
				},
				AsmGrowthBlock: &models.IpamsvcInheritedAsmGrowthBlock{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value: &models.IpamsvcAsmGrowthBlock{
						GrowthFactor: 20,
						GrowthType:   "percent",
					},
				},
				AsmThreshold: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				ForecastPeriod: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				History: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MinTotal: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
				MinUnused: &models.InheritanceInheritedUInt32{
					Action:      "inherit",
					DisplayName: "unit-test-display-name",
					Source:      "unit-test-source",
					Value:       int64(20),
				},
			},
		},
		"IncorrectAsmEnableBlock.Value.ReenableDate_ExpectError": {
			input: []interface{}{
				map[string]interface{}{
					"asm_enable_block": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value": []interface{}{
								map[string]interface{}{
									"enable":              true,
									"enable_notification": true,
									"reenable_date":       "incorrect-reenable-date",
								},
							},
						},
					},
					"asm_growth_block": []interface{}{
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
					"asm_threshold": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"forecast_period": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"history": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"min_total": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
					"min_unused": []interface{}{
						map[string]interface{}{
							"action":       "inherit",
							"display_name": "unit-test-display-name",
							"source":       "unit-test-source",
							"value":        20,
						},
					},
				},
			},
			expectError: &time.ParseError{
				Layout:     "2006-01-02 15:04:05",
				LayoutElem: "2006",
				Value:      "incorrect-reenable-date",
				ValueElem:  "incorrect-reenable-date",
				Message:    "",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := expandIpamsvcInheritedASMConfig(context.TODO(), tc.input)

			if err != nil {
				if tc.expectError != nil {
					assert.Equal(t, tc.expectError, err)
				} else {
					t.Fatal(err)
				}
			} else {
				assert.Equal(t, tc.expect, result)
			}
		})
	}
}
