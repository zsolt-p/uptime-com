package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/zsolt-p/uptime-com/uptime"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func () terraform.ResourceProvider {
			return uptime.Provider()
		},
	})
}
