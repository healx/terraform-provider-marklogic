package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Support markdown syntax.
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:     schema.TypeString,
					Optional: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"MARKLOGIC_USERNAME",
					}, nil),
				},
				"password": {
					Type:     schema.TypeString,
					Optional: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"MARKLOGIC_PASSWORD",
					}, nil),
				},
				"base_url": {
					Type:     schema.TypeString,
					Optional: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"MARKLOGIC_BASE_URL",
					}, nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"marklogic_user": dataSourceUser(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"marklogic_user": resourceUser(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	baseUrl   string
	userAgent string
	username  string
	password  string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return &apiClient{
			baseUrl:   d.Get("base_url").(string),
			userAgent: p.UserAgent("terraform-provider-marklogic", version),
			username:  d.Get("username").(string),
			password:  d.Get("password").(string),
		}, nil
	}
}
