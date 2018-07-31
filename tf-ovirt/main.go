package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ovirt/ovirt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovirt.Provider})
}
