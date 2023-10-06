package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_zone"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceConfigAuthZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigAuthZoneRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceConfigAuthZone),
				Description: "List of DNS Auth Zones matching filters. The schema of each element is identical to the b1ddi_dns_auth_zone resource schema.",
			},
		},
	}
}

func dataSourceConfigAuthZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	is := "partial"

	resp, err := c.DNSConfigurationAPI.AuthZone.AuthZoneList(&auth_zone.AuthZoneListParams{
		Filter:  swag.String(filterStr),
		Inherit: &is,
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenConfigAuthZone(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenConfigAuthZone(r *models.ConfigAuthZone) []interface{} {
	if r == nil {
		return nil
	}

	externalPrimaries := make([]map[string]interface{}, 0, len(r.ExternalPrimaries))
	for _, ep := range r.ExternalPrimaries {
		externalPrimaries = append(externalPrimaries, flattenConfigExternalPrimary(ep))
	}

	externalSecondaries := make([]map[string]interface{}, 0, len(r.ExternalSecondaries))
	for _, es := range r.ExternalSecondaries {
		externalSecondaries = append(externalSecondaries, flattenConfigExternalSecondary(es))
	}

	inheritanceAssignedHosts := make([]interface{}, 0, len(r.InheritanceAssignedHosts))
	for _, iah := range r.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritance2AssignedHost(iah))
	}

	internalSecondaries := make([]interface{}, 0, len(r.InternalSecondaries))
	for _, is := range r.InternalSecondaries {
		internalSecondaries = append(internalSecondaries, flattenConfigInternalSecondary(is))
	}

	queryACL := make([]interface{}, 0, len(r.QueryACL))
	for _, aclItem := range r.QueryACL {
		queryACL = append(queryACL, flattenConfigACLItem(aclItem))
	}

	transferACL := make([]interface{}, 0, len(r.TransferACL))
	for _, aclItem := range r.TransferACL {
		transferACL = append(transferACL, flattenConfigACLItem(aclItem))
	}

	updateACL := make([]interface{}, 0, len(r.UpdateACL))
	for _, aclItem := range r.UpdateACL {
		updateACL = append(updateACL, flattenConfigACLItem(aclItem))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                          r.ID,
			"comment":                     r.Comment,
			"created_at":                  r.CreatedAt.String(),
			"disabled":                    r.Disabled,
			"external_primaries":          externalPrimaries,
			"external_secondaries":        externalSecondaries,
			"fqdn":                        r.Fqdn,
			"gss_tsig_enabled":            r.GssTsigEnabled,
			"inheritance_assigned_hosts":  inheritanceAssignedHosts,
			"inheritance_sources":         flattenConfigAuthZoneInheritance(r.InheritanceSources),
			"initial_soa_serial":          r.InitialSoaSerial,
			"internal_secondaries":        internalSecondaries,
			"mapped_subnet":               r.MappedSubnet,
			"mapping":                     r.Mapping,
			"notify":                      r.Notify,
			"nsgs":                        r.Nsgs,
			"parent":                      r.Parent,
			"primary_type":                r.PrimaryType,
			"protocol_fqdn":               r.ProtocolFqdn,
			"query_acl":                   queryACL,
			"tags":                        r.Tags,
			"transfer_acl":                transferACL,
			"update_acl":                  updateACL,
			"updated_at":                  r.UpdatedAt.String(),
			"use_forwarders_for_subzones": r.UseForwardersForSubzones,
			"view":                        r.View,
			"zone_authority":              flattenConfigZoneAuthority(r.ZoneAuthority),
		},
	}
}
