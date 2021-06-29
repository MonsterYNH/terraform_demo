package monster

import (
	"errors"

	"github.com/MonsterYNH/terraform_demo/mockapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMonsterVpc() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"specs": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VDC specs",
			},
		},
		Create: createMonsterVpc,
		Update: updateMonsterVpc,
		Read:   readMonsterVpc,
		Delete: deleteMOnsterVpc,
	}
}

func createMonsterVpc(data *schema.ResourceData, meta interface{}) error {
	datas := map[string]string{}
	if inter, ok := data.GetOk("specs"); ok {
		if specs, exist := inter.(string); exist {
			datas["specs"] = specs
		}
	}

	client, ok := meta.(*mockapi.MockapiClient)
	if !ok {
		return errors.New("client not ready")
	}

	vdc, err := client.CreateVdc(datas["specs"])
	if err != nil {
		return err
	}

	data.SetId(vdc.ID)
	return readMonsterVpc(data, meta)
}

func readMonsterVpc(data *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(*mockapi.MockapiClient)
	if !ok {
		return errors.New("client not ready")
	}

	_, err := client.ReadVdc(data.Id())
	if err != nil {
		return err
	}

	return nil
}

func updateMonsterVpc(data *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(*mockapi.MockapiClient)
	if !ok {
		return errors.New("client not ready")
	}

	datas := map[string]string{}
	if inter, ok := data.GetOk("specs"); ok {
		specs, exist := inter.(string)
		if exist {
			datas["specs"] = specs
		}
	}

	_, err := client.UpdateVdc(data.Id(), datas["specs"])
	if err != nil {
		return err
	}
	return readMonsterVpc(data, meta)
}

func deleteMOnsterVpc(data *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(*mockapi.MockapiClient)
	if !ok {
		return errors.New("client not ready")
	}

	if err := client.DeleteVdc(data.Id()); err != nil {
		return err
	}
	return readMonsterVpc(data, meta)
}
