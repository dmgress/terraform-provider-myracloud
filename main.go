package main

import (
	"github.com/dmgress/terraform-provider-myracloud/myracloud"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: myracloud.Provider})
}
