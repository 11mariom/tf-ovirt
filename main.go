package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/11mariom/tf-ovirt/ovirt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovirt.Provider})
}
