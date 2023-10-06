package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// DataRecordInheritance RecordInheritance
//
// The inheritance configuration specifies how the _Record_ object inherits the _ttl_ field.
//
// swagger:model dataRecordInheritance
func schemaDataRecordInheritance() *schema.Resource {
	return &schema.Resource{

		Schema: map[string]*schema.Schema{

			// The field config for the _ttl_ field from the _Record_ object
			"ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "The field config for the _ttl_ field from the _Record_ object",
			},
		},
	}
}

func flattenDataRecordInheritance(r *models.DataRecordInheritance) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"ttl": flattenInheritanceInheritedUInt32(r.TTL),
		},
	}
}

func expandDataRecordInheritance(d []interface{}) *models.DataRecordInheritance {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})
	return &models.DataRecordInheritance{
		TTL: expandInheritance2InheritedUInt32(in["ttl"].([]interface{})),
	}
}
