package perimeter81

import (
	"context"

	perimeter81Sdk "github.com/Perimeter81-Public/perimeter-81-client-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceObjectService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceObjectServiceRead,
		Schema: map[string]*schema.Schema{
			"name":        {Type: schema.TypeString, Required: true},
			"id":          {Type: schema.TypeString, Computed: true},
			"description": {Type: schema.TypeString, Computed: true},
			"protocols": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceObjectServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*perimeter81Sdk.APIClient)

	searchName := d.Get("name").(string)

	resp, _, err := client.ObjectsServicesApi.GetObjectsServices(ctx)
	if err != nil {
		return appendErrorDiags(diags, "Unable to get object services", err)
	}

	var foundService *perimeter81Sdk.ObjectsServicesResponseObj
	for _, service := range resp.Data {
		if service.Name == searchName {
			foundService = &service
			break
		}
	}

	if foundService == nil {
		return diag.Errorf("Object Service '%s' not found", searchName)
	}

	d.SetId(foundService.Id)
	d.Set("name", foundService.Name)
	d.Set("description", foundService.Description)

	var protocols []map[string]interface{}
	for _, p := range foundService.Protocols {
		protocols = append(protocols, map[string]interface{}{
			"protocol": p.Protocol,
		})
	}
	d.Set("protocols", protocols)

	return diags
}
