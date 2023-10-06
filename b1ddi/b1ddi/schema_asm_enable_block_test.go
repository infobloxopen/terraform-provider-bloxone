package b1ddi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFlattenIpamsvcAsmEnableBlock(t *testing.T) {
	cases := map[string]struct {
		input    *models.IpamsvcAsmEnableBlock
		expected []interface{}
	}{
		"NilInput": {
			input:    nil,
			expected: []interface{}{},
		},
		"FullInput": {
			input: &models.IpamsvcAsmEnableBlock{
				Enable:             true,
				EnableNotification: true,
				ReenableDate:       strfmt.NewDateTime(),
			},
			expected: []interface{}{
				map[string]interface{}{
					"enable":              true,
					"enable_notification": true,
					"reenable_date":       "1970-01-01T00:00:00.000Z",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := flattenIpamsvcAsmEnableBlock(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExpandIpamsvcAsmEnableBlock(t *testing.T) {
	cases := map[string]struct {
		input         []interface{}
		expected      *models.IpamsvcAsmEnableBlock
		expectedError error
	}{
		"NilInput": {
			input:    nil,
			expected: nil,
		},
		"FullInput": {
			input: []interface{}{
				map[string]interface{}{
					"enable":              true,
					"enable_notification": true,
					"reenable_date":       "1970-01-01T00:00:00.000Z",
				},
			},
			expected: &models.IpamsvcAsmEnableBlock{
				Enable:             true,
				EnableNotification: true,
				ReenableDate:       strfmt.NewDateTime(),
			},
		},
		"IncorrectReenableDate_ExpectError": {
			input: []interface{}{
				map[string]interface{}{
					"reenable_date": "incorrect-reenable-date",
				},
			},
			expectedError: &time.ParseError{
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
			result, err := expandIpamsvcAsmEnableBlock(context.TODO(), tc.input)
			if err != nil {
				if tc.expectedError != nil {
					assert.Equal(t, tc.expectedError, err)
				} else {
					t.Fatal(err)
				}
			} else {
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
