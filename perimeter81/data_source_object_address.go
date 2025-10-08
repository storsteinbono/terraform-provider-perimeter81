package perimeter81

import (
	"context"

	perimeter81Sdk "github.com/Perimeter81-Public/perimeter-81-client-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
dataSourceObjectAddress Query a single Object Address by name

@return &schema.Resource
*/
func dataSourceObjectAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceObjectAddressRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the address object to find",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Address object ID",
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Description of the address object",
			},
			"value_type": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Type of value (e.g., IP, range, etc.)",
			},
			"ip_version": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "IP version (IPv4 or IPv6)",
			},
			"value": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "List of values (addresses, ranges, etc.)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

/*
dataSourceObjectAddressRead Find an object address by name
  - @param ctx context.Context
  - @param d *schema.ResourceData
  - @param m interface{}

@return diag.Diagnostics
*/
func dataSourceObjectAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*perimeter81Sdk.APIClient)

	searchName := d.Get("name").(string)

	// Get all object addresses
	resp, _, err := client.ObjectsAddressesApi.GetObjectsAddresses(ctx)
	if err != nil {
		return appendErrorDiags(diags, "Unable to get Object Addresses", err)
	}

	// Find the address with matching name
	var foundAddress *perimeter81Sdk.ObjectsAddressObj
	for _, address := range resp.Data {
		if address.Name == searchName {
			foundAddress = &address
			break
		}
	}

	if foundAddress == nil {
		return diag.Errorf("Object Address with name '%s' not found", searchName)
	}

	// Set the data
	d.SetId(foundAddress.Id)
	d.Set("name", foundAddress.Name)
	d.Set("description", foundAddress.Description)
	d.Set("value_type", foundAddress.ValueType)
	d.Set("ip_version", foundAddress.IpVersion)
	d.Set("value", foundAddress.Value)

	return diags
}
