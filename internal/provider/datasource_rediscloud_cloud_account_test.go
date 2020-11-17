package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRedisCloudCloudAccount(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRedisCloudCloudAccount,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.rediscloud_cloud_account.test", "id", regexp.MustCompile("^\\d*$")),
					resource.TestCheckResourceAttr("data.rediscloud_cloud_account.test", "provider_type", "AWS"),
					resource.TestCheckResourceAttrSet("data.rediscloud_cloud_account.test", "name"),
					resource.TestCheckResourceAttrSet("data.rediscloud_cloud_account.test", "access_key_id"),
				),
			},
		},
	})
}

const testAccDataSourceRedisCloudCloudAccount = `
data "rediscloud_cloud_account" "test" {
  exclude_internal_account = true
  provider_type = "AWS" 
}
`
