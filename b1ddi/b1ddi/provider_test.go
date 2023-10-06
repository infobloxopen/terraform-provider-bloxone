package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testAccProvider          *schema.Provider
	testAccProviderFactories map[string]func() (*schema.Provider, error)
)

func init() {
	testAccProvider = Provider()

	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"b1ddi": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

func testAccPreCheck(t *testing.T) {
	if host := os.Getenv("B1DDI_HOST"); host == "" {
		t.Fatal("B1DDI_HOST must be set for acceptance tests")
	}

	if token := os.Getenv("B1DDI_API_KEY"); token == "" {
		t.Fatal("B1DDI_API_KEY must be set for acceptance tests")
	}

	err := testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
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
			assert.Equal(t, tc.expectedResult, filterFromMap(tc.inputMap))
		})
	}
}
