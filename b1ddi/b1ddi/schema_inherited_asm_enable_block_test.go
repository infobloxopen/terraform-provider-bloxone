package b1ddi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFlattenIpamsvcInheritedAsmEnableBlock(t *testing.T) {
	cases := map[string]struct {
		input  *models.IpamsvcInheritedAsmEnableBlock
		expect []interface{}
	}{
		"NilInput": {
			input:  nil,
			expect: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcInheritedAsmEnableBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcAsmEnableBlock{
					Enable:             true,
					EnableNotification: true,
					ReenableDate:       strfmt.NewDateTime(),
				},
			},
			expect: []interface{}{
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
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcInheritedAsmEnableBlock(tc.input)

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestExpandIpamsvcInheritedAsmEnableBlock(t *testing.T) {
	cases := map[string]struct {
		input       []interface{}
		expect      *models.IpamsvcInheritedAsmEnableBlock
		expectError error
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
							"enable":              true,
							"enable_notification": true,
							"reenable_date":       "1970-01-01T00:00:00.000Z",
						},
					},
				},
			},
			expect: &models.IpamsvcInheritedAsmEnableBlock{
				Action:      "inherit",
				DisplayName: "unit-test-display-name",
				Source:      "unit-test-source",
				Value: &models.IpamsvcAsmEnableBlock{
					Enable:             true,
					EnableNotification: true,
					ReenableDate:       strfmt.NewDateTime(),
				},
			},
		},
		"IncorrectReenableDate_ExpectError": {
			input: []interface{}{
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
			result, err := expandIpamsvcInheritedAsmEnableBlock(context.TODO(), tc.input)

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
