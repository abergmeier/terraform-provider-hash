package provider

import (
	"github.com/abergmeier/terraform-provider-hash/internal/datasources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"hash_filesha256s": datasources.FileSha256s(),
		},
	}
}
