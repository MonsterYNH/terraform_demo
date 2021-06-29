package monster

import (
	"errors"

	"github.com/MonsterYNH/terraform_demo/mockapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataResourceMonsterVpc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Result output file",
			},
		},
		Read: dataResourceMonsterVdcRead,
	}
}

func dataResourceMonsterVdcRead(data *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(*mockapi.MockapiClient)
	if !ok {
		return errors.New("client not ready")
	}

	vdc, err := client.GetVdcList()
	if err != nil {
		return err
	}

	if inter, ok := data.GetOk("result_output_file"); ok {
		if outputFile, exist := inter.(string); exist {
			if err := writeTofile(vdc, outputFile); err != nil {
				return err
			}
		}
	}
	return nil
}
