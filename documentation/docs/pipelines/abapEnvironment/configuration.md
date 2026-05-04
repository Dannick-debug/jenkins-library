# Pipeline Configuration

The ABAP Environment Pipeline is configured through three files in your repository plus GitHub repository secrets.

For a step-by-step setup guide see the [Getting Started](../../guidedtour.md) page.

## Prerequisites

- GitHub repository with Actions enabled
- SAP BTP, ABAP Environment system (permanent or provisioned per run)
- AAKaaS access (for add-on build scenarios)
- ABAP communication user credentials

## Files to create

### `addon.yml` — Add-on Descriptor

Describes the add-on product and software component versions to build.

```yaml
addonProduct: /NAMESPC/PRODUCTX
addonVersion: "1.2.0"
repositories:
  - name: /NAMESPC/COMPONENTA
    branch: v1.2.0
    version: "1.2.0"
    commitID: "7d4516e9"
  - name: /NAMESPC/COMPONENTB
    branch: v2.0.0
    version: "2.0.0"
    commitID: "9f102ffb"
```

### `repositories.yml` — Repository List

Lists the software components to clone in the Clone Repositories stage.

```yaml
repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
    version: v1.0.0
```

### `.pipeline/config.yml` — Step Configuration

Non-secret parameters for each step.

```yaml
general:
  abapAddonAssemblyKitEndpoint: "https://aakaas.example.com"

steps:
  abapEnvironmentCloneGitRepo:
    repositories: repositories.yml

  abapEnvironmentRunATCCheck:
    atcConfig: atcConfig.yml

  abapEnvironmentRunAUnitTest:
    aUnitConfig: aUnitConfig.yml
```

## ATC configuration (`atcConfig.yml`)

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/MY_COMPONENT
```

## AUnit configuration (`aUnitConfig.yml`)

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/MY_COMPONENT
```

## GitHub secrets

| Secret | Description |
|--------|-------------|
| `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE` | AAKaaS authentication cookie |
| `PIPER_USER` / `PIPER_PASSWORD` | ABAP system credentials |
| `PIPER_CF_USER` / `PIPER_CF_PASSWORD` | CF credentials (CF path) |
| `PIPER_BTP_API_CREDENTIALS_ID` | BTP credentials (BTP path) |

## Consumer workflow

Run `piper initGithubActions` to generate the consumer workflow. To toggle stages, edit the `with:` inputs block:

```yaml
with:
  addonDescriptorFileName: addon.yml
  repositories: repositories.yml
  stage_atc: true
  stage_aunit: true
  stage_integrationTest: false
  stage_confirm: true
  confirmEnvironment: production
```

Set any stage toggle to `false` to skip that stage entirely.

## Confirm gate

The Confirm stage pauses the pipeline and requires manual approval before publishing. It uses a [GitHub Environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment) with required reviewers. Create the environment under **Settings → Environments** and add reviewers.

## Sample configurations

See the [SAP-samples/abap-platform-ci-cd-samples](https://github.com/SAP-samples/abap-platform-ci-cd-samples) repository for complete working examples.
