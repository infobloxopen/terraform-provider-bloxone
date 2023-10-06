package b1ddi

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcRange Range
//
// A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
//
// swagger:model ipamsvcRange
func resourceIpamsvcRange() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcRangeCreate,
		ReadContext:   resourceIpamsvcRangeRead,
		UpdateContext: resourceIpamsvcRangeUpdate,
		DeleteContext: resourceIpamsvcRangeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the range. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The resource identifier.
			"dhcp_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The list of DHCP options. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options. May be either a specific option or a group of options.",
			},

			// The end IP address of the range.
			// Required: true
			"end": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The end IP address of the range.",
			},

			// The list of all exclusion ranges in the scope of the range.
			"exclusion_ranges": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcExclusionRange(),
				Optional:    true,
				Description: "The list of all exclusion ranges in the scope of the range.",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			"inheritance_assigned_hosts": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceAssignedHost(),
				Computed:    true,
				Description: "The list of the inheritance assigned hosts of the object.",
			},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The DHCP inheritance configuration for the range.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPOptionsInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP inheritance configuration for the range.",
			},

			// The name of the range. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the range. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The type of protocol (_ipv4_ or _ipv6_).
			// Read Only: true
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of protocol (_ipv4_ or _ipv6_).",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The start IP address of the range.
			// Required: true
			"start": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The start IP address of the range.",
			},

			// The tags for the range in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the range in JSON format.",
			},

			// The utilization threshold settings for the range.
			"threshold": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilizationThreshold(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The utilization threshold settings for the range.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The utilization statistics for the range.
			// Read Only: true
			"utilization": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilization(),
				Computed:    true,
				Description: "The utilization statistics for the range.",
			},
		},
	}
}

func resourceIpamsvcRangeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	exclusionRanges := make([]*models.IpamsvcExclusionRange, 0)
	for _, er := range d.Get("exclusion_ranges").([]interface{}) {
		if er != nil {
			exclusionRanges = append(exclusionRanges, expandIpamsvcExclusionRange(er.(map[string]interface{})))
		}
	}

	r := &models.IpamsvcRange{
		Comment:            d.Get("comment").(string),
		DhcpHost:           d.Get("dhcp_host").(string),
		DhcpOptions:        dhcpOptions,
		End:                swag.String(d.Get("end").(string)),
		ExclusionRanges:    exclusionRanges,
		InheritanceParent:  d.Get("inheritance_parent").(string),
		InheritanceSources: expandIpamsvcDHCPOptionsInheritance(d.Get("inheritance_sources").([]interface{})),
		Name:               d.Get("name").(string),
		Parent:             d.Get("parent").(string),
		Space:              swag.String(d.Get("space").(string)),
		Start:              swag.String(d.Get("start").(string)),
		Tags:               d.Get("tags"),
		Threshold:          expandIpamsvcUtilizationThreshold(d.Get("threshold").([]interface{})),
	}

	resp, err := c.IPAddressManagementAPI.RangeOperations.RangeCreate(&range_operations.RangeCreateParams{Body: r, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcRangeRead(ctx, d, m)
}

func resourceIpamsvcRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.RangeOperations.RangeRead(&range_operations.RangeReadParams{
		ID:      d.Id(),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("created_at", resp.Payload.Result.CreatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dhcp_host", resp.Payload.Result.DhcpHost)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	dhcpOptions := make([]map[string]interface{}, 0, len(resp.Payload.Result.DhcpOptions))
	for _, dhcpOption := range resp.Payload.Result.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}
	err = d.Set("dhcp_options", dhcpOptions)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("end", resp.Payload.Result.End)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	exclusionRanges := make([]interface{}, 0, len(resp.Payload.Result.ExclusionRanges))
	for _, er := range resp.Payload.Result.ExclusionRanges {
		exclusionRanges = append(exclusionRanges, flattenIpamsvcExclusionRange(er))
	}
	err = d.Set("exclusion_ranges", exclusionRanges)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	inheritanceAssignedHosts := make([]interface{}, 0, len(resp.Payload.Result.InheritanceAssignedHosts))
	for _, inheritanceAssignedHost := range resp.Payload.Result.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritanceAssignedHost(inheritanceAssignedHost))
	}
	err = d.Set("inheritance_assigned_hosts", inheritanceAssignedHosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_parent", resp.Payload.Result.InheritanceParent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_sources", flattenIpamsvcDHCPOptionsInheritance(resp.Payload.Result.InheritanceSources))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("parent", resp.Payload.Result.Parent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("protocol", resp.Payload.Result.Protocol)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("space", resp.Payload.Result.Space)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("start", resp.Payload.Result.Start)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("threshold", flattenIpamsvcUtilizationThreshold(resp.Payload.Result.Threshold))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("updated_at", resp.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("utilization", flattenIpamsvcUtilization(resp.Payload.Result.Utilization))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamsvcRangeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	if d.HasChange("space") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'space' field is not allowed"))
	}

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	exclusionRanges := make([]*models.IpamsvcExclusionRange, 0)
	for _, er := range d.Get("exclusion_ranges").([]interface{}) {
		if er != nil {
			exclusionRanges = append(exclusionRanges, expandIpamsvcExclusionRange(er.(map[string]interface{})))
		}
	}

	body := &models.IpamsvcRange{
		Comment:            d.Get("comment").(string),
		DhcpHost:           d.Get("dhcp_host").(string),
		DhcpOptions:        dhcpOptions,
		End:                swag.String(d.Get("end").(string)),
		ExclusionRanges:    exclusionRanges,
		InheritanceSources: expandIpamsvcDHCPOptionsInheritance(d.Get("inheritance_sources").([]interface{})),
		Name:               d.Get("name").(string),
		Start:              swag.String(d.Get("start").(string)),
		Tags:               d.Get("tags"),
		Threshold:          expandIpamsvcUtilizationThreshold(d.Get("threshold").([]interface{})),
	}

	resp, err := c.IPAddressManagementAPI.RangeOperations.RangeUpdate(
		&range_operations.RangeUpdateParams{ID: d.Id(), Body: body, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcRangeRead(ctx, d, m)
}

func resourceIpamsvcRangeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	_, err := c.IPAddressManagementAPI.RangeOperations.RangeDelete(&range_operations.RangeDeleteParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
