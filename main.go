package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/spaceapegames/redait-provder/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}