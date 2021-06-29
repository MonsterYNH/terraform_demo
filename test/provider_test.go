package test

import (
	"github.com/MonsterYNH/terraform_demo/monster"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var testAccProvider *schema.Provider

func init() {
	testAccProvider = monster.Provider().(*schema.Provider)
}
