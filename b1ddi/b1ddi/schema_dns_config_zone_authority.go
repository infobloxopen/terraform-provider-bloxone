package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigZoneAuthority ZoneAuthority
//
// Construct for fields: _refresh_, _retry_, _expire_, _default_ttl_, _negative_ttl_, _rname_, _protocol_rname_, _mname_, _protocol_mname_, _use_default_mname_.
//
// swagger:model configZoneAuthority
func schemaConfigZoneAuthority() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).
			//
			// Defaults to 28800.
			"default_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     28800,
				Description: "Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).\n\nDefaults to 28800.",
			},

			// Optional. ZoneAuthority expire time in seconds.
			//
			// Defaults to 2419200.
			"expire": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2419200,
				Description: "Optional. ZoneAuthority expire time in seconds.\n\nDefaults to 2419200.",
			},

			// Optional. ZoneAuthority master name server (partially qualified domain name)
			//
			// Defaults to empty.
			"mname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defaults to empty.",
			},

			// Optional. ZoneAuthority negative caching (minimum) ttl in seconds.
			//
			// Defaults to 900.
			"negative_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     900,
				Description: "Optional. ZoneAuthority negative caching (minimum) ttl in seconds.\n\nDefaults to 900.",
			},

			// Optional. ZoneAuthority master name server in punycode.
			//
			// Defaults to empty.
			// Read Only: true
			"protocol_mname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional. ZoneAuthority master name server in punycode.\n\nDefaults to empty.",
			},

			// Optional. A domain name which specifies the mailbox of the person responsible for this zone.
			//
			// Defaults to empty.
			// Read Only: true
			"protocol_rname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional. A domain name which specifies the mailbox of the person responsible for this zone.\n\nDefaults to empty.",
			},

			// Optional. ZoneAuthority refresh.
			//
			// Defaults to 10800.
			"refresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10800,
				Description: "Optional. ZoneAuthority refresh.\n\nDefaults to 10800.",
			},

			// Optional. ZoneAuthority retry.
			//
			// Defaults to 3600.
			"retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3600,
				Description: "Optional. ZoneAuthority retry.\n\nDefaults to 3600.",
			},

			// Optional. ZoneAuthority rname.
			//
			// Defaults to empty.
			"rname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. ZoneAuthority rname.\n\nDefaults to empty.",
			},

			// Optional. Use default value for master name server.
			//
			// Defaults to true.
			"use_default_mname": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Optional. Use default value for master name server.\n\nDefaults to true.",
			},
		},
	}
}

func flattenConfigZoneAuthority(r *models.ConfigZoneAuthority) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"default_ttl":       r.DefaultTTL,
			"expire":            r.Expire,
			"mname":             r.Mname,
			"negative_ttl":      r.NegativeTTL,
			"protocol_mname":    r.ProtocolMname,
			"protocol_rname":    r.ProtocolRname,
			"refresh":           r.Refresh,
			"retry":             r.Retry,
			"rname":             r.Rname,
			"use_default_mname": r.UseDefaultMname,
		},
	}
}

func expandConfigZoneAuthority(d []interface{}) *models.ConfigZoneAuthority {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	mname := in["mname"].(string)
	if in["use_default_mname"].(bool) {
		mname = ""
	}

	return &models.ConfigZoneAuthority{
		DefaultTTL:      int64(in["default_ttl"].(int)),
		Expire:          int64(in["expire"].(int)),
		Mname:           mname,
		NegativeTTL:     int64(in["negative_ttl"].(int)),
		Refresh:         int64(in["refresh"].(int)),
		Retry:           int64(in["retry"].(int)),
		Rname:           in["rname"].(string),
		UseDefaultMname: swag.Bool(in["use_default_mname"].(bool)),
	}
}
