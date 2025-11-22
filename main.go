package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/lfventura/terraform-provider-shell/shell"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shell.Provider})
}
