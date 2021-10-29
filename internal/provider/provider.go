package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"files": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The files to include in hash",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hash_filesha256s": dataSourceFileSha256s(),
		},
	}
}
