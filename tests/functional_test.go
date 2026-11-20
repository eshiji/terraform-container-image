package terraform_tools

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fixtureMountPath   = "/fixture"
	fixtureWorkspace   = "/tmp/workspace"
	expectedPlanOutput = "hello from terraform-tools"
)

// The fixture is mounted read-only and copied, because init writes .terraform/
// and the container user may not own the host directory.
var terraformWorkflow = []string{
	"cp -r " + fixtureMountPath + " " + fixtureWorkspace,
	"cd " + fixtureWorkspace,
	"terraform init -input=false -no-color",
	"terraform validate -no-color",
	"terraform plan -input=false -no-color",
}

func TestTerraformWorkflowRunsInsideImage(t *testing.T) {
	fixture, err := filepath.Abs("fixtures/terraform")
	require.NoError(t, err)
	opts := &docker.RunOptions{
		Command:  []string{"sh", "-ec", strings.Join(terraformWorkflow, " && ")},
		Volumes:  []string{fixture + ":" + fixtureMountPath + ":ro"},
		Platform: currentPlatform(),
		Remove:   true,
	}

	output := docker.Run(t, builtImage(t), opts)

	assert.Contains(t, output, "Terraform has been successfully initialized")
	assert.Contains(t, output, "The configuration is valid")
	assert.Contains(t, output, expectedPlanOutput)
}
