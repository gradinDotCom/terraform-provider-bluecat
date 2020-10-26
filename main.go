package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/gradinDotCom/terraform-provider-bluecat/bluecat"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bluecat.Provider})
}
