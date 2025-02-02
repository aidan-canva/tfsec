package compute

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
		Service:   "compute",
		ShortCode: "disk-encryption-customer-key",
		Documentation: rule.RuleDocumentation{
			Summary:     "Disks should be encrypted with Customer Supplied Encryption Keys",
			Explanation: `Using unmanaged keys makes rotation and general management difficult.`,
			Impact:      "Using unmanaged keys does not allow for proper management",
			Resolution:  "Use managed keys ",
			BadExample: []string{  `
resource "google_compute_disk" "bad_example" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}
`},
			GoodExample: []string{ `
resource "google_compute_disk" "good_example" {
  name  = "test-disk"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-9-stretch-v20200805"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
  disk_encryption_key {
    kms_key_self_link = "something"
  }
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_disk#kms_key_self_link",
			},
		},
		RequiredTypes:  []string{ 
			"resource",
		},
		RequiredLabels: []string{ 
			"google_compute_disk",
		},
		DefaultSeverity: severity.Low, 
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context){
			if kmsKeySelfLinkAttr := resourceBlock.GetBlock("disk_encryption_key").GetAttribute("kms_key_self_link"); kmsKeySelfLinkAttr.IsNil() { // alert on use of default value
				set.AddResult().
					WithDescription("Resource '%s' uses default value for disk_encryption_key.kms_key_self_link", resourceBlock.FullName())
			} else if kmsKeySelfLinkAttr.IsNotResolvable() {
				set.AddResult().
					WithDescription("Resource '%s' does not set disk_encryption_key.kms_key_self_link", resourceBlock.FullName()).
					WithAttribute(kmsKeySelfLinkAttr)
			}
		},
	})
}
