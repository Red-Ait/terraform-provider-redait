package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/Red-Ait/terraform-provider-redait/provider/api/client"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider,
	})
}