package opentelekomcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAntiDdosV1_importBasic(t *testing.T) {
	resourceName := "opentelekomcloud_antiddos_v1.antiddos_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAntiDdosV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAntiDdosV1_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
