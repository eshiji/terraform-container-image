package terraform_tools

import (
	"encoding/json"
	"os/exec"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nana reads the versions the Dockerfile declares, so they are written in one place.
// renovate: datasource=go depName=github.com/eshiji/nana
const nanaModule = "github.com/eshiji/nana/cmd/nana@v0.1.0"

const versionArgSuffix = "_VERSION"

// toolChecks maps each declared version ARG to the command that reports the installed version.
var toolChecks = map[string][]string{
	"TERRAFORM_VERSION":      {"terraform", "version"},
	"TERRAGRUNT_VERSION":     {"terragrunt", "--version"},
	"TERRAFORM_DOCS_VERSION": {"terraform-docs", "--version"},
	"TFSEC_VERSION":          {"tfsec", "--version"},
	"TFSWITCH_VERSION":       {"tfswitch", "--version"},
	"GOLANG_VERSION":         {"go", "version"},
}

// notCheckedYet lists declared versions without a check, each with the reason.
var notCheckedYet = map[string]string{}

type dockerfileArg struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

func declaredVersions(t *testing.T) map[string]string {
	t.Helper()
	out, err := exec.Command("go", "run", nanaModule, "dockerfile", "args", "-f", "../Dockerfile", "-o", "json").Output()
	require.NoError(t, err, "nana dockerfile args")
	var args []dockerfileArg
	require.NoError(t, json.Unmarshal(out, &args))
	versions := map[string]string{}
	for _, arg := range args {
		if strings.HasSuffix(arg.Name, versionArgSuffix) {
			versions[arg.Name] = arg.Default
		}
	}
	return versions
}

func TestEveryDeclaredVersionIsChecked(t *testing.T) {
	for name := range declaredVersions(t) {
		_, checked := toolChecks[name]
		_, excused := notCheckedYet[name]
		assert.True(t, checked || excused, "%s has no version check and no reason to skip it", name)
	}
}

func TestInstalledToolsMatchDeclaredVersions(t *testing.T) {
	versions := declaredVersions(t)
	image := builtImage(t)
	for name, command := range toolChecks {
		t.Run(command[0], func(t *testing.T) {
			declared, ok := versions[name]
			require.True(t, ok, "Dockerfile no longer declares %s", name)
			opts := &docker.RunOptions{Command: command, Platform: currentPlatform(), Remove: true}

			output := docker.Run(t, image, opts)

			assert.Contains(t, output, strings.TrimPrefix(declared, "v"))
		})
	}
}
