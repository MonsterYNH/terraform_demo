package monster

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider := Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"monster": testAccProvider,
	}
}

func TestAccInstanceDataSource_basic(t *testing.T) {
	os.Setenv("TF_ACC", "asdasd")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMonsterVdcDataSource_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccInstanceExists("monster_vdc.my_vdc"),
					resource.TestCheckResourceAttr("monster_vdc.my_vdc", "specs", "1C2G"),
				),
			},
		},
	})
}

func testAccMonsterVdcDataSource_basic() string {
	return `
	resource monster_vdc my_vdc {
		specs = "1C2G"
	}
	data monster_data_resource_vdc my_data {
		result_output_file = "vdc.json"
	}
	`
}

func testAccInstanceExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//logId := getLogId(contextNil)
		//ctx := context.WithValue(context.TODO(), "logId", logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		fmt.Println(rs.Primary.ID)

		return nil
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MONSTER_ACCOUNT"); v == "" {
		t.Fatal("CDS_SECRET_ID must be set for acceptance tests")
	}
	if v := os.Getenv("MONSTER_PASSWORD"); v == "" {
		t.Fatal("CDS_SECRET_KEY must be set for acceptance tests")
	}
}

func testAccInstanceDestroy(s *terraform.State) error {
	return nil
}
