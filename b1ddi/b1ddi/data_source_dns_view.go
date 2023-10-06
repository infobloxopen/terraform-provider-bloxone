package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/view"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceConfigView() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigViewRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceConfigView),
				Description: "List of DNS Views matching filters. The schema of each element is identical to the b1ddi_dns_view resource schema.",
			},
		},
	}
}

func dataSourceConfigViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.DNSConfigurationAPI.View.ViewList(&view.ViewListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenConfigView(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenConfigView(r *models.ConfigView) []interface{} {
	if r == nil {
		return nil
	}

	customRootNs := make([]map[string]interface{}, 0, len(r.CustomRootNs))
	for _, ns := range r.CustomRootNs {
		customRootNs = append(customRootNs, flattenConfigRootNS(ns))
	}

	dnssecRootKeys := make([]map[string]interface{}, 0, len(r.DnssecRootKeys))
	for _, drk := range r.DnssecRootKeys {
		dnssecRootKeys = append(dnssecRootKeys, flattenConfigTrustAnchor(drk))
	}

	dnssecTrustAnchors := make([]map[string]interface{}, 0, len(r.DnssecTrustAnchors))
	for _, dta := range r.DnssecTrustAnchors {
		dnssecTrustAnchors = append(dnssecTrustAnchors, flattenConfigTrustAnchor(dta))
	}

	ecsZones := make([]map[string]interface{}, 0, len(r.EcsZones))
	for _, ecsZone := range r.EcsZones {
		ecsZones = append(ecsZones, flattenConfigECSZone(ecsZone))
	}

	forwarders := make([]map[string]interface{}, 0, len(r.Forwarders))
	for _, f := range r.Forwarders {
		forwarders = append(forwarders, flattenConfigForwarder(f))
	}

	matchClientsACL := make([]map[string]interface{}, 0, len(r.MatchClientsACL))
	for _, aclItem := range r.MatchClientsACL {
		matchClientsACL = append(matchClientsACL, flattenConfigACLItem(aclItem))
	}

	matchDestinationsACL := make([]map[string]interface{}, 0, len(r.MatchDestinationsACL))
	for _, aclItem := range r.MatchDestinationsACL {
		matchDestinationsACL = append(matchDestinationsACL, flattenConfigACLItem(aclItem))
	}

	queryACL := make([]map[string]interface{}, 0, len(r.QueryACL))
	for _, aclItem := range r.QueryACL {
		queryACL = append(queryACL, flattenConfigACLItem(aclItem))
	}

	recursionACL := make([]map[string]interface{}, 0, len(r.RecursionACL))
	for _, aclItem := range r.RecursionACL {
		recursionACL = append(recursionACL, flattenConfigACLItem(aclItem))
	}

	transferACL := make([]map[string]interface{}, 0, len(r.TransferACL))
	for _, aclItem := range r.TransferACL {
		transferACL = append(transferACL, flattenConfigACLItem(aclItem))
	}

	updateACL := make([]map[string]interface{}, 0, len(r.UpdateACL))
	for _, aclItem := range r.UpdateACL {
		updateACL = append(updateACL, flattenConfigACLItem(aclItem))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                          r.ID,
			"comment":                     r.Comment,
			"created_at":                  r.CreatedAt.String(),
			"custom_root_ns":              customRootNs,
			"custom_root_ns_enabled":      r.CustomRootNsEnabled,
			"disabled":                    r.Disabled,
			"dnssec_enable_validation":    r.DnssecEnableValidation,
			"dnssec_enabled":              r.DnssecEnabled,
			"dnssec_root_keys":            dnssecRootKeys,
			"dnssec_trust_anchors":        dnssecTrustAnchors,
			"dnssec_validate_expiry":      r.DnssecValidateExpiry,
			"ecs_enabled":                 r.EcsEnabled,
			"ecs_forwarding":              r.EcsForwarding,
			"ecs_prefix_v4":               r.EcsPrefixV4,
			"ecs_prefix_v6":               r.EcsPrefixV6,
			"ecs_zones":                   ecsZones,
			"edns_udp_size":               r.EdnsUDPSize,
			"forwarders":                  forwarders,
			"forwarders_only":             r.ForwardersOnly,
			"gss_tsig_enabled":            r.GssTsigEnabled,
			"inheritance_sources":         flattenConfigViewInheritance(r.InheritanceSources),
			"ip_spaces":                   r.IPSpaces,
			"lame_ttl":                    r.LameTTL,
			"match_clients_acl":           matchClientsACL,
			"match_destinations_acl":      matchDestinationsACL,
			"match_recursive_only":        r.MatchRecursiveOnly,
			"max_cache_ttl":               r.MaxCacheTTL,
			"max_negative_ttl":            r.MaxNegativeTTL,
			"max_udp_size":                r.MaxUDPSize,
			"minimal_responses":           r.MinimalResponses,
			"name":                        r.Name,
			"notify":                      r.Notify,
			"query_acl":                   queryACL,
			"recursion_acl":               recursionACL,
			"recursion_enabled":           r.RecursionEnabled,
			"tags":                        r.Tags,
			"transfer_acl":                transferACL,
			"update_acl":                  updateACL,
			"updated_at":                  r.UpdatedAt.String(),
			"use_forwarders_for_subzones": r.UseForwardersForSubzones,
			"zone_authority":              flattenConfigZoneAuthority(r.ZoneAuthority),
		},
	}
}
