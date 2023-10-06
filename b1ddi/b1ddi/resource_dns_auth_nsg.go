package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_nsg"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigAuthNSG AuthNSG
//
// Authoritative DNS Server Group for authoritative zones.
//
// swagger:model configAuthNSG
func resourceConfigAuthNSG() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigAuthNSGCreate,
		ReadContext:   resourceConfigAuthNSGRead,
		UpdateContext: resourceConfigAuthNSGUpdate,
		DeleteContext: resourceConfigAuthNSGDelete,
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

			// Optional. DNS primaries external to BloxOne DDI. Order is not significant.
			"external_primaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigExternalPrimary(),
				Optional:    true,
				Description: "Optional. DNS primaries external to BloxOne DDI. Order is not significant.",
			},

			// DNS secondaries external to BloxOne DDI. Order is not significant.
			"external_secondaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigExternalSecondary(),
				Optional:    true,
				Description: "DNS secondaries external to BloxOne DDI. Order is not significant.",
			},

			// Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.
			"internal_secondaries": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInternalSecondary(),
				Optional:    true,
				Description: "Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.",
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

func resourceConfigAuthNSGCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalPrimaries := make([]*models.ConfigExternalPrimary, 0)
	for _, ep := range d.Get("external_primaries").([]interface{}) {
		externalPrimaries = append(externalPrimaries, expandConfigExternalPrimary(ep.(map[string]interface{})))
	}

	externalSecondaries := make([]*models.ConfigExternalSecondary, 0)
	for _, es := range d.Get("external_secondaries").([]interface{}) {
		externalSecondaries = append(externalSecondaries, expandConfigExternalSecondary(es.(map[string]interface{})))
	}

	internalSecondaries := make([]*models.ConfigInternalSecondary, 0)
	for _, is := range d.Get("internal_secondaries").([]interface{}) {
		internalSecondaries = append(internalSecondaries, expandConfigInternalSecondary(is.(map[string]interface{})))
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	nsg := &models.ConfigAuthNSG{
		Comment:             d.Get("comment").(string),
		ExternalPrimaries:   externalPrimaries,
		ExternalSecondaries: externalSecondaries,
		InternalSecondaries: internalSecondaries,
		Name:                swag.String(d.Get("name").(string)),
		Nsgs:                nsgs,
		Tags:                d.Get("tags"),
	}

	resp, err := c.DNSConfigurationAPI.AuthNsg.AuthNsgCreate(
		&auth_nsg.AuthNsgCreateParams{
			Body: nsg, Context: ctx,
		},
		nil,
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigAuthNSGRead(ctx, d, m)
}

func resourceConfigAuthNSGRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.DNSConfigurationAPI.AuthNsg.AuthNsgRead(
		&auth_nsg.AuthNsgReadParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	externalPrimaries := make([]map[string]interface{}, 0, len(resp.Payload.Result.ExternalPrimaries))
	for _, ep := range resp.Payload.Result.ExternalPrimaries {
		externalPrimaries = append(externalPrimaries, flattenConfigExternalPrimary(ep))
	}
	err = d.Set("external_primaries", externalPrimaries)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	externalSecondaries := make([]map[string]interface{}, 0, len(resp.Payload.Result.ExternalSecondaries))
	for _, es := range resp.Payload.Result.ExternalSecondaries {
		externalSecondaries = append(externalSecondaries, flattenConfigExternalSecondary(es))
	}
	err = d.Set("external_secondaries", externalSecondaries)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	internalSecondaries := make([]map[string]interface{}, 0, len(resp.Payload.Result.InternalSecondaries))
	for _, is := range resp.Payload.Result.InternalSecondaries {
		internalSecondaries = append(internalSecondaries, flattenConfigInternalSecondary(is))
	}
	err = d.Set("internal_secondaries", internalSecondaries)
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

func resourceConfigAuthNSGUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	externalPrimaries := make([]*models.ConfigExternalPrimary, 0)
	for _, ep := range d.Get("external_primaries").([]interface{}) {
		externalPrimaries = append(externalPrimaries, expandConfigExternalPrimary(ep.(map[string]interface{})))
	}

	externalSecondaries := make([]*models.ConfigExternalSecondary, 0)
	for _, es := range d.Get("external_secondaries").([]interface{}) {
		externalSecondaries = append(externalSecondaries, expandConfigExternalSecondary(es.(map[string]interface{})))
	}

	internalSecondaries := make([]*models.ConfigInternalSecondary, 0)
	for _, is := range d.Get("internal_secondaries").([]interface{}) {
		internalSecondaries = append(internalSecondaries, expandConfigInternalSecondary(is.(map[string]interface{})))
	}

	nsgs := make([]string, 0)
	for _, n := range d.Get("nsgs").([]interface{}) {
		if n != nil {
			nsgs = append(nsgs, n.(string))
		}
	}

	nsg := &models.ConfigAuthNSG{
		Comment:             d.Get("comment").(string),
		ExternalPrimaries:   externalPrimaries,
		ExternalSecondaries: externalSecondaries,
		InternalSecondaries: internalSecondaries,
		Name:                swag.String(d.Get("name").(string)),
		Nsgs:                nsgs,
		Tags:                d.Get("tags"),
	}

	resp, err := c.DNSConfigurationAPI.AuthNsg.AuthNsgUpdate(
		&auth_nsg.AuthNsgUpdateParams{
			ID: d.Id(), Body: nsg, Context: ctx,
		},
		nil,
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceConfigAuthNSGRead(ctx, d, m)
}

func resourceConfigAuthNSGDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)
	_, err := c.DNSConfigurationAPI.AuthNsg.AuthNsgDelete(
		&auth_nsg.AuthNsgDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
