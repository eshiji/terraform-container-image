package terraform_tools

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

// TODO add go version check after fix multiarch build
// const go_version = "1.18.2"
const terragrunt_version = "v0.86.2"
const tfswitch_version = "v1.5.0"
const terraform_docs_version = "v0.20.0"
const tfsec_version = "v1.28.14"
const image_tag = "terraform-docker-terratest"

func buldImage(t *testing.T, tag string, arch []string) {
	buildOptions := &docker.BuildOptions{
		Tags:          []string{tag},
		Architectures: arch,
	}
	docker.Build(t, "../", buildOptions)
}

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
	// Given
	fmt.Print(runtime.GOROOT(), runtime.GOARCH, runtime.GOOS)
	current_arch := fmt.Sprintf("linux/%s", runtime.GOARCH)
	// multi_arch := []string{"linux/arm64"}
	// build_archs := []string{current_arch}
	buldImage(t, image_tag, []string{current_arch})

	// When
	for name, test := range versionTests {
		t.Run(name, func(t *testing.T) {
			opts := &docker.RunOptions{
				Command:  test.command,
				Platform: current_arch,
				Remove:   true,
			}
			result := docker.Run(t, image_tag, opts)
			assert.Contains(t, result, test.expected)
		})

	}
}
