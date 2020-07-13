package myracloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"myracloud_domain":     resourceMyracloudDomain(),
			"myracloud_dns_record": resourceMyracloudDNSRecord(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}

	return provider
}
