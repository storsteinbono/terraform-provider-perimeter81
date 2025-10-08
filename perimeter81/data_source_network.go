package perimeter81

import (
	"context"

	perimeter81Sdk "github.com/Perimeter81-Public/perimeter-81-client-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkRead,
		Schema: map[string]*schema.Schema{
			"id":        {Type: schema.TypeString, Computed: true},
			"name":      {Type: schema.TypeString, Required: true},
			"dns":       {Type: schema.TypeString, Computed: true},
			"subnet":    {Type: schema.TypeString, Computed: true},
			"access_type": {Type: schema.TypeString, Computed: true},
			"is_default":  {Type: schema.TypeBool, Computed: true},
			"tenant_id":   {Type: schema.TypeString, Computed: true},
			"created_at":  {Type: schema.TypeString, Computed: true},
			"updated_at":  {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*perimeter81Sdk.APIClient)

	networkName := d.Get("name").(string)

	networksResp, _, err := client.NetworksApi.GetNetworks(ctx)
	if err != nil {
		return appendErrorDiags(diags, "Unable to get networks", err)
	}

	var found *perimeter81Sdk.Network
	for _, n := range networksResp {
		if n.Name == networkName {
			found = &n
			break
		}
	}

	if found == nil {
		return diag.Errorf("Network '%s' not found", networkName)
	}

	d.SetId(found.Id)
	d.Set("name", found.Name)
	d.Set("dns", found.Dns)
	d.Set("subnet", found.Subnet)
	d.Set("access_type", found.AccessType)
	d.Set("is_default", found.IsDefault)
	d.Set("tenant_id", found.TenantId)
	d.Set("created_at", found.CreatedAt)
	d.Set("updated_at", found.UpdatedAt)

	return diags
}
