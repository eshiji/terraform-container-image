package terraform_tools;

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDocker(t *testing.T) {
	// website::tag::1:: Configure the tag to use on the Docker image.
	tag := "test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// website::tag::2:: Build the Docker image.
	docker.Build(t, "../", buildOptions)

	// website::tag::3:: Run the Docker image, read the text file from it, and make sure it contains the expected output.
	go_version := &docker.RunOptions{Command: []string{"go", "version"}}
	terragrunt_version := &docker.RunOptions{Command: []string{"terragrunt", "-v"}}

	output_go := docker.Run(t, tag, go_version)
	output_terragrunt := docker.Run(t, tag, terragrunt_version)
	assert.Equal(t, "go version go1.18.2 linux/amd64", output_go)
	assert.Equal(t, "terragrunt version v0.36.10", output_terragrunt)
}