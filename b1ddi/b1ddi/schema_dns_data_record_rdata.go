package b1ddi

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var (
	ParseError = "unable to parse key '%s': %v"
)

/*
		updateDataRecordRData helps convert rDATA record fields of type integer from string value
	    Introduced to fix issues where terraform converts integer values to string in rendered config
*/
func updateDataRecordRData(d interface{}, recordType string) (interface{}, diag.Diagnostics) {
	if d == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	in := d.(map[string]interface{})
	switch recordType {
	case "CAA":
		inPlaceUpdater(in, "flags", &diags)

	case "MX":
		inPlaceUpdater(in, "preference", &diags)

	case "NAPTR":
		inPlaceUpdater(in, "order", &diags)
		inPlaceUpdater(in, "preference", &diags)

	case "SOA":
		inPlaceUpdater(in, "serial", &diags)

	case "SRV":
		inPlaceUpdater(in, "port", &diags)
		inPlaceUpdater(in, "priority", &diags)
		inPlaceUpdater(in, "weight", &diags)

	default:
		return d, nil
	}

	return in, diags
}

func inPlaceUpdater(in map[string]interface{}, key string, diags *diag.Diagnostics) {
	if val, ok := in[key]; ok {
		i, err := strconv.Atoi(val.(string))
		if err != nil {
			*diags = append(*diags, diag.Errorf(ParseError, key, err)...)
		} else {
			in[key] = i
		}
	}
}
