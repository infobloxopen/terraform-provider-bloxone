package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigECSBlock ECSBlock
//
// Block for fields: _ecs_enabled_, _ecs_forwarding_, _ecs_prefix_v4_, _ecs_prefix_v6_, _ecs_zones_.
//
// swagger:model configECSBlock
func schemaConfigECSBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _ecs_enabled_ field.
			"ecs_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _ecs_enabled_ field.",
			},

			// Optional. Field config for _ecs_forwarding_ field.
			"ecs_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _ecs_forwarding_ field.",
			},

			// Optional. Field config for _ecs_prefix_v4_ field.
			"ecs_prefix_v4": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional. Field config for _ecs_prefix_v4_ field.",
			},

			// Optional. Field config for _ecs_prefix_v6_ field.
			"ecs_prefix_v6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional. Field config for _ecs_prefix_v6_ field.",
			},

			// Optional. Field config for _ecs_zones_ field.
			"ecs_zones": {
				Type:        schema.TypeList,
				Elem:        schemaConfigECSZone(),
				Optional:    true,
				Description: "Optional. Field config for _ecs_zones_ field.",
			},
		},
	}
}

func flattenConfigECSBlock(r *models.ConfigECSBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	ecsZones := make([]interface{}, 0, len(r.EcsZones))
	for _, ecsZone := range r.EcsZones {
		ecsZones = append(ecsZones, flattenConfigECSZone(ecsZone))
	}

	return []interface{}{
		map[string]interface{}{
			"ecs_enabled":    r.EcsEnabled,
			"ecs_forwarding": r.EcsForwarding,
			"ecs_prefix_v4":  r.EcsPrefixV4,
			"ecs_prefix_v6":  r.EcsPrefixV6,
			"ecs_zones":      ecsZones,
		},
	}
}
