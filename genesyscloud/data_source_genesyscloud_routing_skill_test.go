package genesyscloud

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRoutingSkill(t *testing.T) {
	var (
		skillResource   = "routing-skill"
		skillDataSource = "routing-skill-data"
		skillName       = "Terraform Skill-" + uuid.NewString()
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { TestAccPreCheck(t) },
		ProviderFactories: GetProviderFactories(providerResources, providerDataSources),
		Steps: []resource.TestStep{
			{
				Config: GenerateRoutingSkillResource(
					skillResource,
					skillName,
				) + generateRoutingSkillDataSource(skillDataSource, "genesyscloud_routing_skill."+skillResource+".name", "genesyscloud_routing_skill."+skillResource),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.genesyscloud_routing_skill."+skillDataSource, "id", "genesyscloud_routing_skill."+skillResource, "id"),
				),
			},
		},
	})
}

func generateRoutingSkillDataSource(
	resourceID string,
	name string,
	// Must explicitly use depends_on in terraform v0.13 when a data source references a resource
	// Fixed in v0.14 https://github.com/hashicorp/terraform/pull/26284
	dependsOnResource string) string {
	return fmt.Sprintf(`data "genesyscloud_routing_skill" "%s" {
		name = %s
        depends_on=[%s]
	}
	`, resourceID, name, dependsOnResource)
}
