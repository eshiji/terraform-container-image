package terraform_tools

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
)

const imageTag = "terraform-docker-terratest"

var buildOnce sync.Once

func currentPlatform() string {
	return fmt.Sprintf("linux/%s", runtime.GOARCH)
}

// builtImage builds the image once per test run and returns its tag.
func builtImage(t *testing.T) string {
	buildOnce.Do(func() {
		docker.Build(t, "../", &docker.BuildOptions{
			Tags:          []string{imageTag},
			Architectures: []string{currentPlatform()},
		})
	})
	return imageTag
}
