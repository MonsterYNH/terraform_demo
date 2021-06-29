package monster

import (
	"os"

	"github.com/MonsterYNH/terraform_demo/mockapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("MONSTER_ACCOUNT", nil),
			},
			"passowrd": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("MONSTER_ACCOUNT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"monster_vdc": resourceMonsterVpc(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"monster_data_resource_vdc": dataResourceMonsterVpc(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	datas := map[string]string{}
	if inter, ok := data.GetOk("account"); ok {
		account, exist := inter.(string)
		if exist {
			datas["account"] = account
		}
	} else {
		account := os.Getenv("MONSTER_ACCOUNT")
		datas["account"] = account
	}
	if inter, ok := data.GetOk("password"); ok {
		password, exist := inter.(string)
		if exist {
			datas["password"] = password
		}
	} else {
		password := os.Getenv("MONSTER_PASSWORD")
		datas["password"] = password
	}

	client := &mockapi.MockapiClient{
		Account:  datas["account"],
		Password: datas["password"],
		BaseURL:  "http://localhost:8080",
	}
	return client, nil
}
