# Piper CLI

The Piper CLI is the Go binary that implements all ABAP Environment Pipeline and BTP API steps. Jenkins calls it under the hood for each step, but you can also run it directly on the command line for local testing.

## Download

The latest released version can be downloaded via:

```sh
wget https://github.com/SAP/jenkins-library/releases/latest/download/piper
chmod +x piper
```

Specific versions are available on the [GitHub releases](https://github.com/SAP/jenkins-library/releases) page.

Once available in `$PATH`, it is ready to use.

## Verify and explore

```sh
piper version        # print the version
piper help           # list all available commands
piper <step> --help  # show parameters for a specific step
```

## Available steps

All ABAP Environment Pipeline steps are available as sub-commands:

| Command | Description |
| ------- | ----------- |
| `abapEnvironmentBuild` | Executes an ABAP build on SAP BTP |
| `abapEnvironmentCheckoutBranch` | Checks out a branch in an ABAP Git repository |
| `abapEnvironmentCloneGitRepo` | Clones a Git repository into an ABAP system |
| `abapEnvironmentCreateSystem` | Creates a new ABAP environment system |
| `abapEnvironmentCreateTag` | Creates a Git tag in an ABAP system |
| `abapEnvironmentPullGitRepo` | Pulls a Git repository in an ABAP system |
| `abapEnvironmentPushATCSystemConfig` | Pushes an ATC system configuration |
| `abapEnvironmentRunATCCheck` | Runs ATC checks |
| `abapEnvironmentRunAUnitTest` | Runs AUnit tests |
| `abapEnvironmentAssemblePackages` | Assembles add-on packages |
| `abapEnvironmentAssembleConfirm` | Confirms assembly of add-on packages |
| `abapAddonAssemblyKitCheck` | Checks the add-on descriptor |
| `abapAddonAssemblyKitCheckCVs` | Checks component versions |
| `abapAddonAssemblyKitCheckPV` | Checks the product version |
| `abapAddonAssemblyKitCreateTargetVector` | Creates the target vector |
| `abapAddonAssemblyKitPublishTargetVector` | Publishes the target vector |
| `abapAddonAssemblyKitRegisterPackages` | Registers add-on packages |
| `abapAddonAssemblyKitReleasePackages` | Releases add-on packages |
| `abapAddonAssemblyKitReserveNextPackages` | Reserves the next add-on packages |
| `abapLandscapePortalUpdateAddOnProduct` | Updates an add-on product via Landscape Portal |
| `btpCreateServiceInstance` | Creates a BTP service instance |
| `btpCreateServiceBinding` | Creates a BTP service binding |
| `btpDeleteServiceInstance` | Deletes a BTP service instance |
| `btpDeleteServiceBinding` | Deletes a BTP service binding |
| `cloudFoundryCreateServiceKey` | Creates a Cloud Foundry service key |
| `cloudFoundryDeleteService` | Deletes a Cloud Foundry service |

## Shell completion

For interactive use, set up shell completion:

```sh
piper completion --help
```

!!! note "Linux only"
    The binary is compiled for Linux. On macOS or Windows, run it inside Docker.
