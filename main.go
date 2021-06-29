package main

import (
	"github.com/MonsterYNH/terraform_demo/monster"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: monster.Provider})
}
