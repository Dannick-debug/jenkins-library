# ABAP Environment Pipeline

![ABAP Environment Pipeline](../../images/abapPipelineOverviewFull.png)

The ABAP Environment Pipeline enables Continuous Integration for SAP BTP, ABAP Environment (Steampunk). It is orchestrated by a GitHub Actions reusable workflow that calls the `piper` CLI binary for each step.

## Scenarios

### Continuous Testing

Run nightly ATC checks and AUnit tests against a permanent or transient ABAP system. See the [scenario documentation](../../scenarios/abapEnvironmentTest.md).

### Building ABAP Add-ons

Build and publish ABAP add-on products for SAP partners delivering SaaS solutions on Steampunk. See the [scenario documentation](../../scenarios/abapEnvironmentAddons.md).

## Pipeline Stages

| Stage | GitHub Actions job | Steps |
|-------|--------------------|-------|
| Init | `init` | `piper initRunStageConfiguration` |
| [Initial Checks](stages/initialChecks.md) | `initial-checks` | `abapAddonAssemblyKitCheck` |
| [Prepare System](stages/prepareSystem.md) | `prepare-system` | `btpCreateServiceInstance`, `btpCreateServiceBinding` (BTP) or `abapEnvironmentCreateSystem`, `cloudFoundryCreateServiceKey` (CF) |
| [Clone Repositories](stages/cloneRepositories.md) | `clone-repositories` | `abapEnvironmentCloneGitRepo` |
| [ATC](stages/test.md) | `atc` | `abapEnvironmentRunATCCheck` |
| [AUnit](stages/test.md) | `aunit` _(parallel with atc)_ | `abapEnvironmentRunAUnitTest` |
| [Build](stages/build.md) | `build` | `abapAddonAssemblyKitReserveNextPackages`, `abapEnvironmentAssemblePackages`, `abapEnvironmentAssembleConfirm`, `abapAddonAssemblyKitRegisterPackages`, `abapAddonAssemblyKitReleasePackages` |
| [Integration Tests](stages/integrationTest.md) | `integration-tests` | `abapEnvironmentRunATCCheck` |
| [Confirm](stages/confirm.md) | `confirm` | manual gate via GitHub Environment |
| [Publish](stages/publish.md) | `publish` | `abapAddonAssemblyKitCreateTargetVector`, `abapAddonAssemblyKitPublishTargetVector` |
| [Post](stages/post.md) | `post` _(always runs)_ | `abapLandscapePortalUpdateAddOnProduct`, `btpDeleteServiceBinding`, `btpDeleteServiceInstance` (BTP) or `cloudFoundryDeleteService` (CF) |

For configuration details see the [configuration guide](configuration.md) and the [GitHub Actions setup guide](github-actions.md).
