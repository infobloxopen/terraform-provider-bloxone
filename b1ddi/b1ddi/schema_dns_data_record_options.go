package b1ddi

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

/*
	updateDataRecordOptions helps convert string options(supposed to be boolean) values into boolean
    Introduced to fix issues where terraform converts boolean values to string in rendered config
*/
func updateDataRecordOptions(d interface{}, recordType string) (interface{}, diag.Diagnostics) {
	if d == nil {
		return nil, nil
	}
	in := d.(map[string]interface{})
	var diags diag.Diagnostics
	switch recordType {
	case "A", "AAAA":
		if val, ok := in["create_ptr"]; ok {
			b, err := strconv.ParseBool(val.(string))
			if err != nil {
				diags = append(diags, diag.Errorf(ParseError, "create_ptr", err)...)
			} else {
				in["create_ptr"] = b
			}
		}
		if val, ok := in["check_rmz"]; ok {
			b, err := strconv.ParseBool(val.(string))
			if err != nil {
				diags = append(diags, diag.Errorf(ParseError, "check_rmz", err)...)
			} else {
				in["check_rmz"] = b
			}
		}
	}

	return in, diags
}
