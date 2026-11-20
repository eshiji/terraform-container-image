package terraform_tools

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

// TODO add go version check after fix multiarch build
// const go_version = "1.18.2"
const terragrunt_version = "v0.86.2"
const tfswitch_version = "v1.5.1"
const terraform_docs_version = "v0.20.0"
const tfsec_version = "v1.28.14"

func TestDocker(t *testing.T) {
	var versionTests = map[string]struct {
		command  []string
		expected string
	}{
		"tfswitch": {
			command:  []string{"tfswitch", "--version"},
			expected: tfswitch_version,
		},
		"terragrunt": {
			command:  []string{"terragrunt", "--version"},
			expected: terragrunt_version,
		},
		"terraform-docs": {
			command:  []string{"terraform-docs", "--version"},
			expected: terraform_docs_version,
		},
		"tfsec": {
			command:  []string{"tfsec", "--version"},
			expected: tfsec_version,
		},
		// "go": {
		// 	command:  []string{"go", "version"},
		// 	expected: go_version,
		// },
	}
	image := builtImage(t)

	// When
	for name, test := range versionTests {
		t.Run(name, func(t *testing.T) {
			opts := &docker.RunOptions{
				Command:  test.command,
				Platform: currentPlatform(),
				Remove:   true,
			}
			result := docker.Run(t, image, opts)
			assert.Contains(t, result, test.expected)
		})

	}
}
