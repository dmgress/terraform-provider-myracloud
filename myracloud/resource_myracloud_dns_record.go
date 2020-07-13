package myracloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceMyracloudDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceMyracloudDNSRecordCreate,
		Read:   resourceMyracloudDNSRecordRead,
		Update: resourceMyracloudDNSRecordUpdate,
		Delete: resourceMyracloudDNSRecordDelete,

		Schema: map[string]*schema.Schema{

			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"record_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"A", "AAAA", "MX", "CNAME", "TXT", "NS", "SRV", "PTR"}, false),
			},
			"ssl_cert_template": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"protection": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"id": &schema.Schema{
				Computed: true,
				Type:     schema.TypeInt,
			},
			"modified": &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			// 			id Number UD Id of the object
			// modified Date UD Date of last modification
			// created Date - Date of creation
			// name String CU Subdomain name of a DNS record
			// ttl String CU Time to live
			// recordType String CU A recordType to identify the type of a record. Valid types
			// are: A, AAAA, MX, CNAME, TXT, NS, SRV and PTR
			// sslCertTemplate String C A FQDN in the same domain
			// alternativeCname String - The alternative CNAME that points to the record.
			// active Boolean CU Define wether this subdomain should be protected by
			// Myra or not
			// value String CU Depends on the record type. Typically an IPv4/6 address
			// or a domain entry
			// priority Number CU Priority of MX records
			// port Number CU Port for SRV records
		},
	}
}

func resourceMyracloudDNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	return resourceMyracloudDNSRecordRead(d, m)
}

func resourceMyracloudDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMyracloudDNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMyracloudDNSRecordRead(d, m)
}

func resourceMyracloudDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
