package b1ddi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateDataRecordOptions(t *testing.T) {
	testData := map[string]struct {
		input      map[string]interface{}
		output     map[string]interface{}
		recordType string
		errCount   int
	}{
		"Invalid Options - Bad create_ptr value": {
			input: map[string]interface{}{
				"create_ptr": "asda",
			},
			output: map[string]interface{}{
				"create_ptr": "asda",
			},
			recordType: "A",
			errCount:   1,
		},
		"Invalid Options - Bad check_rmz value": {
			input: map[string]interface{}{
				"create_ptr": "true",
				"check_rmz":  "lsdclksj",
			},
			output: map[string]interface{}{
				"create_ptr": true,
				"check_rmz":  "lsdclksj",
			},
			recordType: "A",
			errCount:   1,
		},
		"Invalid Options - Bad check_rmz and create_ptr value": {
			input: map[string]interface{}{
				"create_ptr": "true12",
				"check_rmz":  "lsdclksj",
			},
			output: map[string]interface{}{
				"create_ptr": "true12",
				"check_rmz":  "lsdclksj",
			},
			recordType: "A",
			errCount:   2,
		},
		"Valid Options": {
			input: map[string]interface{}{
				"create_ptr": "true",
			},
			output: map[string]interface{}{
				"create_ptr": true,
			},
			recordType: "A",
			errCount:   0,
		},
		"Valid Options - 1": {
			input: map[string]interface{}{
				"create_ptr": "true",
				"check_rmz":  "true",
			},
			output: map[string]interface{}{
				"create_ptr": true,
				"check_rmz":  true,
			},
			recordType: "A",
			errCount:   0,
		},
		"Valid Options - 2": {
			// Fake message w/o the optional message to test the integrity of the code
			input: map[string]interface{}{
				"address": "10.0.0.1",
			},
			output: map[string]interface{}{
				"address": "10.0.0.1",
			},
			recordType: "A",
			errCount:   0,
		},
	}

	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			options, diags := updateDataRecordOptions(tc.input, tc.recordType)
			assert.Equal(t, tc.errCount, len(diags))
			assert.Equal(t, tc.output, options)
		})
	}
}
