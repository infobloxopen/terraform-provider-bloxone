package ipam

import (
	"context"
	"regexp"
	
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	
	"github.com/infobloxopen/bloxone-go-client/ipam"
	
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcHostAddressModel struct {
	Address         types.String `tfsdk:"address"`
	Ref             types.String `tfsdk:"ref"`
	Space           types.String `tfsdk:"space"`
	NextAvailableId types.String `tfsdk:"next_available_id"`
}

var IpamsvcHostAddressAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"ref":               types.StringType,
	"space":             types.StringType,
	"next_available_id": types.StringType,
}

var IpamsvcHostAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("address"), path.MatchRelative().AtParent().AtName("next_available_id")),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: `Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.`,
	},
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"space": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("address")),
		},
		MarkdownDescription: `The resource identifier.`,
	},
	"next_available_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier for the network container where the next available address should be generated for the host",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("address"), path.MatchRelative().AtParent().AtName("next_available_id")),
			stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(subnet|range)/[0-9a-f-].*$`), "Should be the resource identifier of the network container."),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
	},
}

func ExpandIpamsvcHostAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHostAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcHostAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcHostAddressModel) Expand(_ context.Context, _ *diag.Diagnostics) *ipam.IpamsvcHostAddress {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcHostAddress{}
	
	if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() {
		naip := flex.ExpandString(m.NextAvailableId) + "/nextavailableip"
		to.Ref = &naip
	} else {
		to.Address = flex.ExpandStringPointer(m.Address)
		to.Space = flex.ExpandStringPointer(m.Space)
	}
	return to
}

func FlattenIpamsvcHostAddress(ctx context.Context, from []ipam.IpamsvcHostAddress, diags *diag.Diagnostics, data []IpamsvcHostAddressModel) types.List {
	if len(from) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: IpamsvcHostAddressAttrTypes})
	}
	
	tfData := ApplyToAll(from, data, func(t ipam.IpamsvcHostAddress, e IpamsvcHostAddressModel) types.Object {
		m := IpamsvcHostAddressModel{}
		m.Flatten(ctx, &t, diags)
		m.NextAvailableId = e.NextAvailableId
		o, d := types.ObjectValueFrom(ctx, IpamsvcHostAddressAttrTypes, m)
		diags.Append(d...)
		return o
	})
	
	tfList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: IpamsvcHostAddressAttrTypes}, tfData)
	diags.Append(d...)
	
	return tfList
}

func (m *IpamsvcHostAddressModel) Flatten(_ context.Context, from *ipam.IpamsvcHostAddress, _ *diag.Diagnostics) {
	if from == nil {
		return
	}
	
	if m == nil {
		*m = IpamsvcHostAddressModel{}
	}
	
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Space = flex.FlattenStringPointer(from.Space)
	
}

func (m *IpamsvcHostAddressModel) FromTerraform5Value(val tftypes.Value) error {
	if m == nil {
		return nil
	}
	
	v := map[string]tftypes.Value{}
	err := val.As(&v)
	if err != nil {
		return err
	}
	
	if !m.Address.IsNull() && !m.Address.IsUnknown() {
		err = v["address"].As(&m.Address)
		if err != nil {
			return err
		}
	}
	
	if !m.Ref.IsNull() && !m.Ref.IsUnknown() {
		err = v["ref"].As(&m.Ref)
		if err != nil {
			return err
		}
	}
	if !m.Space.IsNull() && !m.Space.IsUnknown() {
		err = v["space"].As(&m.Space)
		if err != nil {
			return err
		}
	}
	
	if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() {
		err = v["next_available_id"].As(&m.NextAvailableId)
		if err != nil {
			return err
		}
	}
	return nil
}

// ApplyToAll is returns a new slice containing the results of applying the function `f` to each element of the original slice of Ipam Host address and Ipam host address model.
// It is an extension of the ApplyToAll package found in the flex package
func ApplyToAll(from []ipam.IpamsvcHostAddress, d []IpamsvcHostAddressModel, f func(t ipam.IpamsvcHostAddress, e IpamsvcHostAddressModel) types.Object) []types.Object {
	v := make([]types.Object, len(from))
	
	for i, ha := range from {
		if i <= len(d)-1 {
			v[i] = f(ha, d[i])
		} else {
			v[i] = f(ha, IpamsvcHostAddressModel{})
		}
	}
	
	return v
}
