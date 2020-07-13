package myracloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMyracloudDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceMyracloudDomainCreate,
		Read:   resourceMyracloudDomainRead,
		Update: resourceMyracloudDomainUpdate,
		Delete: resourceMyracloudDomainDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Computed: true,
				Type:     schema.TypeInt,
			},
			"modified": &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			},
			"name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"auto_update": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"auto_dns": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
		},
	}
}

func resourceMyracloudDomainCreate(d *schema.ResourceData, m interface{}) error {
	return resourceMyracloudDomainRead(d, m)
}

func resourceMyracloudDomainRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMyracloudDomainUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMyracloudDomainRead(d, m)
}

func resourceMyracloudDomainDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
