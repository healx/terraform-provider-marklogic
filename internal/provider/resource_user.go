package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "A user in MarkLogic.",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The user's name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "A description for the user.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Created by terraform",
			},
			"password": {
				Description: "The user's password.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"roles": {
				Description: "The roles to assign to the user.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	user := map[string]interface{}{
		"user-name":   d.Get("name"),
		"description": d.Get("description"),
		"role":        d.Get("roles"),
		"password":    d.Get("password"),
	}

	err := client.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	// use the meta value to retrieve your client from the provider configure method
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

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	user := map[string]interface{}{
		"user-name":   d.Get("name"),
		"description": d.Get("description"),
		"role":        d.Get("roles"),
		"password":    d.Get("password"),
	}

	err := client.UpdateUser(ctx, user)

	if err != nil {
		return err
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	name := d.Get("name").(string)

	err := client.DeleteUser(ctx, name)

	if err != nil {
		return err
	}

	return nil
}
