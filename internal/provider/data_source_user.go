package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "A user in Marklogic.",

		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The user's name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description for the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"roles": {
				Description: "The roles assigned to the user.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("name").(string)

	user, err := client.GetUser(ctx, name)

	if err != nil {
		return err
	}

	d.SetId(name)
	d.Set("name", name)
	d.Set("description", user["description"].(string))
	d.Set("roles", user["role"])

	return nil
}
