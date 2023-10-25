package provider

import (
	"os"
	"testing"
)

func testAccPreCheck(t *testing.T) {
	if host := os.Getenv("BLOXONE_HOST"); host == "" {
		t.Fatal("BLOXONE_HOST must be set for acceptance tests")
	}

	if token := os.Getenv("BLOXONE_API_KEY"); token == "" {
		t.Fatal("BLOXONE_API_KEY must be set for acceptance tests")
	}
}

func TestFilterFromMap(t *testing.T) {
	testCases := []struct {
		expectedResult string
		inputMap       map[string]interface{}
	}{
		{
			expectedResult: "name=='test_name'",
			inputMap: map[string]interface{}{
				"name": "test_name",
			},
		},
		{
			expectedResult: "int_val==15",
			inputMap: map[string]interface{}{
				"int_val": "15",
			},
		},
		{
			expectedResult: "mac_addr=='00:00:00:00:00:00'",
			inputMap: map[string]interface{}{
				"mac_addr": "00:00:00:00:00:00",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expectedResult, func(t *testing.T) {
			testAccPreCheck(t)
		})
	}
}
