package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/spf13/cobra"
)

//go:embed templates/consumer_workflow.yml.tmpl
var consumerWorkflowTemplate string

//go:embed templates/pipeline_config.yml.tmpl
var pipelineConfigTemplate string

type initGithubActionsOptions struct {
	AddonDescriptorFileName string
	Repositories            string
	ConfirmEnvironment      string
	CfApiEndpoint           string
	CfOrg                   string
	CfSpace                 string
	BtpSubaccountId         string
	BtpServicePlan          string
}

// InitGithubActionsCommand returns the cobra command for scaffolding a GitHub Actions consumer workflow.
func InitGithubActionsCommand() *cobra.Command {
	opts := &initGithubActionsOptions{}

	cmd := &cobra.Command{
		Use:   "initGithubActions",
		Short: "Scaffold a GitHub Actions workflow for the ABAP Environment Pipeline",
		Long: `Generates the files needed to run the ABAP Environment Pipeline on GitHub Actions:

  .github/workflows/abap-pipeline.yml  – consumer workflow that calls the reusable workflow
  .pipeline/config.yml                 – step configuration skeleton (if not already present)
  repositories.yml                     – repositories skeleton (if not already present)

After running this command, follow the checklist printed below to finish setup.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runInitGithubActions(opts)
		},
	}

	cmd.Flags().StringVar(&opts.AddonDescriptorFileName, "addon-descriptor", "addon.yml", "Path to the Add-on Descriptor file")
	cmd.Flags().StringVar(&opts.Repositories, "repositories", "repositories.yml", "Path to the repositories file")
	cmd.Flags().StringVar(&opts.ConfirmEnvironment, "confirm-environment", "production", "GitHub Environment name for the manual Confirm gate")
	cmd.Flags().StringVar(&opts.CfApiEndpoint, "cf-api-endpoint", "", "Cloud Foundry API endpoint (CF provisioning path)")
	cmd.Flags().StringVar(&opts.CfOrg, "cf-org", "", "Cloud Foundry organisation (CF provisioning path)")
	cmd.Flags().StringVar(&opts.CfSpace, "cf-space", "", "Cloud Foundry space (CF provisioning path)")
	cmd.Flags().StringVar(&opts.BtpSubaccountId, "btp-subaccount-id", "", "BTP subaccount ID (BTP provisioning path)")
	cmd.Flags().StringVar(&opts.BtpServicePlan, "btp-service-plan", "abap/standard", "BTP ABAP service plan (BTP provisioning path)")

	return cmd
}

func runInitGithubActions(opts *initGithubActionsOptions) error {
	if err := writeFromTemplate(".github/workflows/abap-pipeline.yml", consumerWorkflowTemplate, opts); err != nil {
		return err
	}
	if err := writeFromTemplateIfAbsent(".pipeline/config.yml", pipelineConfigTemplate, opts); err != nil {
		return err
	}
	if err := writeFileIfAbsent(opts.Repositories, defaultRepositoriesYml()); err != nil {
		return err
	}

	printChecklist(opts)
	return nil
}

func writeFromTemplate(destPath, tmplContent string, data any) error {
	tmpl, err := template.New("").Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("parsing template for %s: %w", destPath, err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("rendering template for %s: %w", destPath, err)
	}
	return writeFile(destPath, buf.Bytes())
}

func writeFromTemplateIfAbsent(destPath, tmplContent string, data any) error {
	if _, err := os.Stat(destPath); err == nil {
		log.Entry().Infof("Skipping %s — already exists", destPath)
		return nil
	}
	return writeFromTemplate(destPath, tmplContent, data)
}

func writeFile(destPath string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("creating directory for %s: %w", destPath, err)
	}
	if err := os.WriteFile(destPath, content, 0644); err != nil {
		return fmt.Errorf("writing %s: %w", destPath, err)
	}
	log.Entry().Infof("Created %s", destPath)
	return nil
}

func writeFileIfAbsent(destPath string, content []byte) error {
	if _, err := os.Stat(destPath); err == nil {
		log.Entry().Infof("Skipping %s — already exists", destPath)
		return nil
	}
	return writeFile(destPath, content)
}

func defaultRepositoriesYml() []byte {
	return []byte(`repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
    version: v1.0.0
    commitID: ""
`)
}

func printChecklist(opts *initGithubActionsOptions) {
	fmt.Println()
	fmt.Println("✔  Files generated successfully. Complete the following steps:")
	fmt.Println()
	fmt.Println("1. Add the following secrets to your GitHub repository")
	fmt.Println("   (Settings → Secrets and variables → Actions → New repository secret):")
	fmt.Println()
	fmt.Println("   PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE  – AAKaaS authentication cookie")
	fmt.Println("   PIPER_USER                             – ABAP system user")
	fmt.Println("   PIPER_PASSWORD                         – ABAP system password")
	if opts.CfApiEndpoint != "" {
		fmt.Println("   PIPER_CF_USER                          – Cloud Foundry user")
		fmt.Println("   PIPER_CF_PASSWORD                      – Cloud Foundry password")
	}
	if opts.BtpSubaccountId != "" {
		fmt.Println("   PIPER_BTP_API_CREDENTIALS_ID           – BTP API credentials")
	}
	fmt.Println()
	fmt.Printf("2. Create a GitHub Environment named %q:\n", opts.ConfirmEnvironment)
	fmt.Println("   Settings → Environments → New environment")
	fmt.Println("   Add required reviewers to enforce manual approval before Publish.")
	fmt.Println()
	fmt.Println("3. Edit addon.yml with your Add-on Product and Component Version details.")
	fmt.Println()
	fmt.Printf("4. Edit %s with your ABAP repository names and target branches.\n", opts.Repositories)
	fmt.Println()
	fmt.Println("5. Push the generated files and trigger the workflow via:")
	fmt.Println("   Actions → ABAP Environment Pipeline → Run workflow")
}
