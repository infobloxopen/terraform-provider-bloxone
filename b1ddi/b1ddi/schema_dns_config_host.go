package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigHost Host
//
// A DNS Host (_dns/host_) object associates DNS configuraton with hosts.
//
// Automatically created and destroyed based on the hosts known to the platform.
//
// swagger:model configHost
func schemaConfigHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The resource identifier.
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// Host FQDN.
			// Read Only: true
			"absolute_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host FQDN.",
			},

			// Host's primary IP Address.
			// Read Only: true
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host's primary IP Address.",
			},

			// Anycast address configured to the host. Order is not significant.
			// Read Only: true
			"anycast_addresses": {
				Type:        schema.TypeList,
				Elem:        schema.TypeString,
				Computed:    true,
				Description: "Anycast address configured to the host. Order is not significant.",
			},

			// Host associated server configuration.
			"associated_server": {
				Type:        schema.TypeList,
				Elem:        schemaConfigHostAssociatedServer(),
				MaxItems:    1,
				Optional:    true,
				Description: "Host associated server configuration.",
			},

			// Host description.
			// Read Only: true
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host description.",
			},

			// Host current version.
			// Read Only: true
			"current_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host current version.",
			},

			// Below _dfp_ field is deprecated and not supported anymore.
			// The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.
			// Read Only: true
			"dfp": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Below _dfp_ field is deprecated and not supported anymore.\nThe indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.",
			},

			// DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host.
			// If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP.
			// Allowed values:
			//  * _unavailable_: BloxOne TD DFP application is not available,
			//  * _enabled_: BloxOne TD DFP application is available and enabled,
			//  * _disabled_: BloxOne TD DFP application is available but disabled.
			// Read Only: true
			"dfp_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. \nIf so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP.\nAllowed values:\n * _unavailable_: BloxOne TD DFP application is not available,\n * _enabled_: BloxOne TD DFP application is available and enabled,\n * _disabled_: BloxOne TD DFP application is available but disabled.",
			},

			// Optional. Inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaConfigHostInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Inheritance configuration.",
			},

			// Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.
			//
			// Defaults to empty.
			"kerberos_keys": {
				Type:        schema.TypeList,
				Elem:        schemaConfigKerberosKey(),
				Optional:    true,
				Description: "Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.\n\nDefaults to empty.",
			},

			// Host display name.
			// Read Only: true
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host display name.",
			},

			// On-Prem Host ID.
			// Read Only: true
			"ophid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "On-Prem Host ID.",
			},

			// Host FQDN in punycode.
			// Read Only: true
			"protocol_absolute_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host FQDN in punycode.",
			},

			// The resource identifier.
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Host site ID.
			// Read Only: true
			"site_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host site ID.",
			},

			// Host tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Host tagging specifics.",
			},
		},
	}
}

func flattenConfigHost(r *models.ConfigHost) []interface{} {
	if r == nil {
		return nil
	}

	kerberosKeys := make([]map[string]interface{}, 0, len(r.KerberosKeys))
	for _, k := range r.KerberosKeys {
		kerberosKeys = append(kerberosKeys, flattenConfigKerberosKey(k))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                     r.ID,
			"absolute_name":          r.AbsoluteName,
			"address":                r.Address,
			"anycast_addresses":      r.AnycastAddresses,
			"associated_server":      flattenConfigHostAssociatedServer(r.AssociatedServer),
			"comment":                r.Comment,
			"current_version":        r.CurrentVersion,
			"dfp":                    r.Dfp,
			"dfp_service":            r.DfpService,
			"inheritance_sources":    flattenConfigHostInheritance(r.InheritanceSources),
			"kerberos_keys":          kerberosKeys,
			"name":                   r.Name,
			"ophid":                  r.Ophid,
			"protocol_absolute_name": r.ProtocolAbsoluteName,
			"server":                 r.Server,
			"site_id":                r.SiteID,
			"tags":                   r.Tags,
		},
	}
}
