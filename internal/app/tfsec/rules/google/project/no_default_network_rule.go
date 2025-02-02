package project

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

import (
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/provider"
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/aquasecurity/tfsec/pkg/severity"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		Provider:       provider.GoogleProvider,
		Service:   "project",
		ShortCode: "no-default-network",
		Documentation: rule.RuleDocumentation{
			Summary:     "Default network should not be created at project level",
			Explanation: `The default network which is provided for a project contains multiple insecure firewall rules which allow ingress to the project's infrastructure. Creation of this network should therefore be disabled.`,
			Impact:      "Exposure of internal infrastructure/services to public internet",
			Resolution:  "Disable automatic default network creation",
			BadExample: []string{  `
resource "google_project" "bad_example" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
  auto_create_network = true
}
`},
			GoodExample: []string{ `
resource "google_project" "good_example" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
  auto_create_network = false
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project#auto_create_network",
			},
		},
		RequiredTypes:  []string{ 
			"resource",
		},
		RequiredLabels: []string{ 
			"google_project",
		},
		DefaultSeverity: severity.High, 
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context){
			if autoCreateNetworkAttr := resourceBlock.GetAttribute("auto_create_network"); autoCreateNetworkAttr.IsNil() { // alert on use of default value
				set.AddResult().
					WithDescription("Resource '%s' uses default value for auto_create_network", resourceBlock.FullName())
			} else if autoCreateNetworkAttr.IsTrue() {
				set.AddResult().
					WithDescription("Resource '%s' does not have auto_create_network set to false", resourceBlock.FullName()).
					WithAttribute(autoCreateNetworkAttr)
			}
		},
	})
}
