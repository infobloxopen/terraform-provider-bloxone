package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_nsg"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigForwardNSG ForwardNSG
//
// Forward DNS Server Group for forward zones.
//
// swagger:model configForwardNSG
func resourceConfigForwardNSG() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigForwardNSGCreate,
		ReadContext:   resourceConfigForwardNSGRead,
		UpdateContext: resourceConfigForwardNSGUpdate,
		DeleteContext: resourceConfigForwardNSGDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// Optional. Comment for the object.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Comment for the object.",
			},

			// Optional. External DNS servers to forward to. Order is not significant.
			"external_forwarders": {
				Type:        schema.TypeList,
				Elem:        schemaConfigForwarder(),
				Optional:    true,
				Description: "Optional. External DNS servers to forward to. Order is not significant.",
			},

			// Optional. _true_ to only forward.
			"forwarders_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. _true_ to only forward.",
			},

			// The resource identifier.
			"hosts": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"internal_forwarders": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Name of the object.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the object.",
			},

			// The resource identifier.
			"nsgs": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Tagging specifics.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tagging specifics.",
			},
		},
	}
}

func resourceConfigForwardNSGCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalForwarders := make([]*models.ConfigForwarder, 0)
	for _, ef := range d.Get("external_forwarders").([]interface{}) {
		externalForwarders = append(externalForwarders, expandConfigForwarder(ef.(map[string]interface{})))
	}

	hosts := make([]string, 0)
	for _, h := range d.Get("hosts").([]interface{}) {
		if h != nil {
			hosts = append(hosts, h.(string))
		}
	}

	internalForwarders := make([]string, 0)
	for _, ifwd := range d.Get("internal_forwarders").([]interface{}) {
		if ifwd != nil {
			internalForwarders = append(internalForwarders, ifwd.(string))
		}
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	nsg := &models.ConfigForwardNSG{
		Comment:            d.Get("comment").(string),
		ExternalForwarders: externalForwarders,
		ForwardersOnly:     d.Get("forwarders_only").(bool),
		Hosts:              hosts,
		InternalForwarders: internalForwarders,
		Name:               swag.String(d.Get("name").(string)),
		Nsgs:               nsgs,
		Tags:               d.Get("tags"),
	}

	resp, err := c.DNSConfigurationAPI.ForwardNsg.ForwardNsgCreate(
		&forward_nsg.ForwardNsgCreateParams{
			Body: nsg, Context: ctx,
		},
		nil,
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigForwardNSGRead(ctx, d, m)
}

func resourceConfigForwardNSGRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.ForwardNsg.ForwardNsgRead(
		&forward_nsg.ForwardNsgReadParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	externalForwarders := make([]map[string]interface{}, 0, len(resp.Payload.Result.ExternalForwarders))
	for _, ef := range resp.Payload.Result.ExternalForwarders {
		externalForwarders = append(externalForwarders, flattenConfigForwarder(ef))
	}
	err = d.Set("external_forwarders", externalForwarders)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("forwarders_only", resp.Payload.Result.ForwardersOnly)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hosts", resp.Payload.Result.Hosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("internal_forwarders", resp.Payload.Result.InternalForwarders)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("nsgs", resp.Payload.Result.Nsgs)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceConfigForwardNSGUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalForwarders := make([]*models.ConfigForwarder, 0)
	for _, ef := range d.Get("external_forwarders").([]interface{}) {
		externalForwarders = append(externalForwarders, expandConfigForwarder(ef.(map[string]interface{})))
	}

	hosts := make([]string, 0)
	for _, h := range d.Get("hosts").([]interface{}) {
		if h != nil {
			hosts = append(hosts, h.(string))
		}
	}

	internalForwarders := make([]string, 0)
	for _, ifwd := range d.Get("internal_forwarders").([]interface{}) {
		if ifwd != nil {
			internalForwarders = append(internalForwarders, ifwd.(string))
		}
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	nsg := &models.ConfigForwardNSG{
		Comment:            d.Get("comment").(string),
		ExternalForwarders: externalForwarders,
		ForwardersOnly:     d.Get("forwarders_only").(bool),
		Hosts:              hosts,
		InternalForwarders: internalForwarders,
		Name:               swag.String(d.Get("name").(string)),
		Nsgs:               nsgs,
		Tags:               d.Get("tags"),
	}

	resp, err := c.DNSConfigurationAPI.ForwardNsg.ForwardNsgUpdate(
		&forward_nsg.ForwardNsgUpdateParams{
			ID: d.Id(), Body: nsg, Context: ctx,
		},
		nil,
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigForwardNSGRead(ctx, d, m)
}

func resourceConfigForwardNSGDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.DNSConfigurationAPI.ForwardNsg.ForwardNsgDelete(
		&forward_nsg.ForwardNsgDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
