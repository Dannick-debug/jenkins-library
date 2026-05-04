package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitGithubActions_DefaultBTP(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	require.NoError(t, os.Chdir(dir))
	defer os.Chdir(origDir) //nolint:errcheck

	opts := &initGithubActionsOptions{
		AddonDescriptorFileName: "addon.yml",
		Repositories:            "repositories.yml",
		ConfirmEnvironment:      "production",
		BtpSubaccountId:         "my-subaccount",
		BtpServicePlan:          "abap/standard",
	}

	require.NoError(t, runInitGithubActions(opts))

	// Consumer workflow created
	workflowPath := filepath.Join(".github", "workflows", "abap-pipeline.yml")
	content := readTestFile(t, workflowPath)
	assert.Contains(t, content, "uses: SAP/jenkins-library/.github/workflows/abapEnvironmentPipeline.yml@main")
	assert.Contains(t, content, "addonDescriptorFileName: addon.yml")
	assert.Contains(t, content, "btpSubaccountId: my-subaccount")
	assert.NotContains(t, content, "cfApiEndpoint")

	// Pipeline config created
	assert.FileExists(t, filepath.Join(".pipeline", "config.yml"))

	// Repositories file created
	reposContent := readTestFile(t, "repositories.yml")
	assert.Contains(t, reposContent, "repositories:")
}

func TestInitGithubActions_CF(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	require.NoError(t, os.Chdir(dir))
	defer os.Chdir(origDir) //nolint:errcheck

	opts := &initGithubActionsOptions{
		AddonDescriptorFileName: "addon.yml",
		Repositories:            "repos.yml",
		ConfirmEnvironment:      "staging",
		CfApiEndpoint:           "https://api.cf.example.com",
		CfOrg:                   "my-org",
		CfSpace:                 "my-space",
	}

	require.NoError(t, runInitGithubActions(opts))

	content := readTestFile(t, filepath.Join(".github", "workflows", "abap-pipeline.yml"))
	assert.Contains(t, content, "cfApiEndpoint: https://api.cf.example.com")
	assert.Contains(t, content, "cfOrg: my-org")
	assert.NotContains(t, content, "btpSubaccountId")
	assert.Contains(t, content, "PIPER_CF_USER")
	assert.NotContains(t, content, "PIPER_BTP_API_CREDENTIALS_ID")
}

func TestInitGithubActions_SkipsExistingFiles(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	require.NoError(t, os.Chdir(dir))
	defer os.Chdir(origDir) //nolint:errcheck

	// Pre-create the config and repos files
	require.NoError(t, os.MkdirAll(".pipeline", 0755))
	require.NoError(t, os.WriteFile(".pipeline/config.yml", []byte("existing: true\n"), 0644))
	require.NoError(t, os.WriteFile("repositories.yml", []byte("existing: true\n"), 0644))

	opts := &initGithubActionsOptions{
		AddonDescriptorFileName: "addon.yml",
		Repositories:            "repositories.yml",
		ConfirmEnvironment:      "production",
	}

	require.NoError(t, runInitGithubActions(opts))

	// Existing files must NOT be overwritten
	assert.Equal(t, "existing: true\n", readTestFile(t, ".pipeline/config.yml"))
	assert.Equal(t, "existing: true\n", readTestFile(t, "repositories.yml"))
}

func TestInitGithubActions_WorkflowOverwritesExisting(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	require.NoError(t, os.Chdir(dir))
	defer os.Chdir(origDir) //nolint:errcheck

	opts := &initGithubActionsOptions{
		AddonDescriptorFileName: "addon.yml",
		Repositories:            "repositories.yml",
		ConfirmEnvironment:      "production",
	}

	// Run twice — workflow must be regenerated on the second run
	require.NoError(t, runInitGithubActions(opts))
	require.NoError(t, runInitGithubActions(opts))

	content := readTestFile(t, filepath.Join(".github", "workflows", "abap-pipeline.yml"))
	assert.True(t, strings.Contains(content, "ABAP Environment Pipeline"))
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(data)
}
