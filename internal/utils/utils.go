package utils

import (
	"fmt"

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ptr is a helper routine that returns a pointer to given value.
func Ptr[T any](t T) *T {
	return &t
}

// DataSourceAttributeMap converts a map of resource schema attributes to data source schema attributes
func DataSourceAttributeMap(r map[string]resourceschema.Attribute, diags *diag.Diagnostics) map[string]datasourceschema.Attribute {
	d := map[string]datasourceschema.Attribute{}
	for k, v := range r {
		d[k] = DataSourceAttribute(k, v, diags)
	}
	return d
}

// DataSourceNestedAttributeObject converts a resource schema nested attribute object to data source schema nested attribute object
func DataSourceNestedAttributeObject(r resourceschema.NestedAttributeObject, diags *diag.Diagnostics) datasourceschema.NestedAttributeObject {
	return datasourceschema.NestedAttributeObject{
		Attributes: DataSourceAttributeMap(r.Attributes, diags),
		CustomType: r.CustomType,
		Validators: r.Validators,
	}
}

// DataSourceAttribute converts a resource schema attribute to data source schema attribute
func DataSourceAttribute(name string, val resourceschema.Attribute, diags *diag.Diagnostics) datasourceschema.Attribute {
	switch a := val.(type) {
	case resourceschema.BoolAttribute:
		return datasourceschema.BoolAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.StringAttribute:
		return datasourceschema.StringAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Int64Attribute:
		return datasourceschema.Int64Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.Float64Attribute:
		return datasourceschema.Float64Attribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.NumberAttribute:
		return datasourceschema.NumberAttribute{
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ObjectAttribute:
		return datasourceschema.ObjectAttribute{
			AttributeTypes:      a.AttributeTypes,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ListAttribute:
		return datasourceschema.ListAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.ListNestedAttribute:
		return datasourceschema.ListNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.MapAttribute:
		return datasourceschema.MapAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.MapNestedAttribute:
		return datasourceschema.MapNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SetAttribute:
		return datasourceschema.SetAttribute{
			ElementType:         a.ElementType,
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SetNestedAttribute:
		return datasourceschema.SetNestedAttribute{
			NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	case resourceschema.SingleNestedAttribute:
		return datasourceschema.SingleNestedAttribute{
			Attributes:          DataSourceAttributeMap(a.Attributes, diags),
			CustomType:          a.CustomType,
			Required:            a.Required,
			Optional:            a.Optional,
			Computed:            a.Computed,
			Sensitive:           a.Sensitive,
			Description:         a.Description,
			MarkdownDescription: a.MarkdownDescription,
			DeprecationMessage:  a.DeprecationMessage,
			Validators:          a.Validators,
		}
	}
	diags.AddError("Provider error",
		fmt.Sprintf("Failed to convert schema attribute of type '%T' for '%s'", val, name))
	return nil
}
