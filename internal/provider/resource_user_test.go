package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"marklogic_user.test-user", "name", "test-user"),
					resource.TestCheckResourceAttr(
						"marklogic_user.test-user", "password", "foobar1234"),
					resource.TestCheckResourceAttr(
						"marklogic_user.test-user", "description", "test description"),
				),
			},
		},
	})
}

const testAccResourceUser = `
resource "marklogic_user" "test-user" {
  name        = "test-user"
  description = "test description"
  roles       = ["manage-user", "rest-reader"]
  password    = "foobar1234"
}
`
