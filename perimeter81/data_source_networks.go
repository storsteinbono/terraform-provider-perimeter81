package perimeter81

import (
	"context"
	"strconv"
	"time"

	perimeter81Sdk "github.com/Perimeter81-Public/perimeter-81-client-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworksRead,
		Schema: map[string]*schema.Schema{
			"networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id":         {Type: schema.TypeString, Computed: true},
						"name":       {Type: schema.TypeString, Computed: true},
						"dns":        {Type: schema.TypeString, Computed: true},
						"subnet":     {Type: schema.TypeString, Computed: true},
						"access_type": {Type: schema.TypeString, Computed: true},
						"is_default":  {Type: schema.TypeBool, Computed: true},
						"tenant_id":   {Type: schema.TypeString, Computed: true},
						"created_at":  {Type: schema.TypeString, Computed: true},
						"updated_at":  {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*perimeter81Sdk.APIClient)

	networksResp, _, err := client.NetworksApi.GetNetworks(ctx)
	if err != nil {
		return appendErrorDiags(diags, "Unable to get networks", err)
	}

	var list []map[string]interface{}
	for _, n := range networksResp {
		item := map[string]interface{}{
			"id":          n.Id,
			"name":        n.Name,
			"dns":         n.Dns,
			"subnet":      n.Subnet,
			"access_type": n.AccessType,
			"is_default":  n.IsDefault,
			"tenant_id":   n.TenantId,
			"created_at":  n.CreatedAt,
			"updated_at":  n.UpdatedAt,
		}
		list = append(list, item)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("networks", list)

	return diags
}
