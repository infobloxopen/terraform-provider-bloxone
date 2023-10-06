package b1ddi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateDataRecordRData(t *testing.T) {
	testData := map[string]struct {
		input      map[string]interface{}
		output     map[string]interface{}
		recordType string
		errCount   int
	}{
		"Valid Rdata - MX Record": {
			input: map[string]interface{}{
				"preference": "1",
				"exchange":   "mail.infoblox.com",
			},
			output: map[string]interface{}{
				"preference": 1,
				"exchange":   "mail.infoblox.com",
			},
			recordType: "MX",
			errCount:   0,
		},
		"Invalid Rdata - MX Record": {
			input: map[string]interface{}{
				"preference": "1qw",
				"exchange":   "mail.infoblox.com",
			},
			output: map[string]interface{}{
				"preference": "1qw",
				"exchange":   "mail.infoblox.com",
			},
			recordType: "MX",
			errCount:   1,
		},
		"Valid Rdata - CAA Record": {
			input: map[string]interface{}{
				"flags": "0",
				"tag":   "issue",
				"value": "infoblox",
			},
			output: map[string]interface{}{
				"flags": 0,
				"tag":   "issue",
				"value": "infoblox",
			},
			recordType: "CAA",
			errCount:   0,
		},
		"Invalid Rdata - CAA Record": {
			input: map[string]interface{}{
				"flags": "0qq",
				"tag":   "issue",
				"value": "infoblox",
			},
			output: map[string]interface{}{
				"flags": "0qq",
				"tag":   "issue",
				"value": "infoblox",
			},
			recordType: "CAA",
			errCount:   1,
		},
		"Valid Rdata - SRV Record": {
			input: map[string]interface{}{
				"port":     "1234",
				"priority": "1",
				"target":   "infoblox",
				"weight":   "100",
			},
			output: map[string]interface{}{
				"port":     1234,
				"priority": 1,
				"target":   "infoblox",
				"weight":   100,
			},
			recordType: "SRV",
			errCount:   0,
		},
		"Invalid Rdata - SRV Record": {
			input: map[string]interface{}{
				"port":     "1234qq",
				"priority": "1wwe",
				"target":   "infoblox",
				"weight":   "1wqq00",
			},
			output: map[string]interface{}{
				"port":     "1234qq",
				"priority": "1wwe",
				"target":   "infoblox",
				"weight":   "1wqq00",
			},
			recordType: "SRV",
			errCount:   3,
		},
		"Valid Rdata - SOA Record": {
			input: map[string]interface{}{
				"serial": "100",
				"mname":  "infoblox.com",
			},
			output: map[string]interface{}{
				"serial": 100,
				"mname":  "infoblox.com",
			},
			recordType: "SOA",
			errCount:   0,
		},
		"Invalid Rdata - SOA Record": {
			input: map[string]interface{}{
				"serial": "1234qq",
			},
			output: map[string]interface{}{
				"serial": "1234qq",
			},
			recordType: "SOA",
			errCount:   1,
		},
	}

	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			output, diagErr := updateDataRecordRData(tc.input, tc.recordType)
			assert.Equal(t, tc.errCount, len(diagErr))
			assert.Equal(t, tc.output, output)
		})
	}
}
