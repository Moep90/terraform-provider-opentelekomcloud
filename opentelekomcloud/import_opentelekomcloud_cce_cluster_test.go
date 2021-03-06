package opentelekomcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCCEClusterV3_importBasic(t *testing.T) {
	resourceName := "opentelekomcloud_cce_cluster_v3.cluster_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCCEClusterV3Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCCEClusterV3_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
