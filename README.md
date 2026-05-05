# Piper: ABAP Environment Pipeline

Run the SAP BTP, ABAP Environment pipeline entirely on **GitHub Actions** using the `piper` CLI binary — no Jenkins, no Groovy, no CI/CD server to manage.

## Quick start

**1. Download the piper binary**

```sh
curl -sL https://github.com/SAP/jenkins-library/releases/latest/download/piper_linux_amd64 -o piper
chmod +x piper
```

**2. Scaffold your consumer repository**

```sh
./piper initGithubActions
```

This generates:
- `.github/workflows/abap-pipeline.yml` — consumer workflow referencing the reusable pipeline
- `.pipeline/config.yml` — step configuration skeleton
- `repositories.yml` — software component list skeleton

**3. Add GitHub secrets**

| Secret | Description |
|--------|-------------|
| `PIPER_USER` | ABAP system username |
| `PIPER_PASSWORD` | ABAP system password |
| `PIPER_CF_USER` | CF credentials (CF provisioning path) |
| `PIPER_CF_PASSWORD` | CF credentials (CF provisioning path) |
| `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE` | AAKaaS authentication cookie (add-on build only) |

**4. Push and trigger**

Push to your repository and run the workflow via **Actions → Run workflow**, or add a `schedule:` trigger for nightly runs.

## How it works

The pipeline runs as a [reusable GitHub Actions workflow](https://docs.github.com/en/actions/using-workflows/reusing-workflows) defined in `.github/workflows/abapEnvironmentPipeline.yml`. Your consumer repo references it with a single `uses:` line — you never maintain pipeline logic yourself.

Each stage is a separate GitHub Actions job. `piper` steps communicate through a pipeline environment artifact (`.pipeline/`) uploaded and downloaded between jobs.

```
init → initial-checks → prepare-system → clone-repositories
     → atc ──┐
             ├── build → integration-tests → confirm → publish → post
     → aunit ┘
```

## Pipeline stages

| Stage | What it does |
|-------|-------------|
| `init` | Reads configuration, validates `addon.yml` |
| `initial-checks` | AAKaaS validates product and component versions |
| `prepare-system` | Provisions a transient ABAP build system (BTP or CF path) |
| `clone-repositories` | Clones software components at the specified commit |
| `atc` | Static code checks — blocks build on findings |
| `aunit` | ABAP unit test execution |
| `build` | Reserves packages, assembles, registers, and releases with AAKaaS |
| `integration-tests` | Optional post-build verification on a test system |
| `confirm` | Manual approval gate (GitHub Environment with required reviewers) |
| `publish` | Publishes the target vector — add-on available for installation |
| `post` | Deprovisions build system, reports to Landscape Portal |

## Configuration

Three files drive the pipeline:

- **`addon.yml`** — add-on product descriptor (product name, version, software component list)
- **`repositories.yml`** — software components to clone for testing-only scenarios
- **`.pipeline/config.yml`** — non-secret step parameters (ATC config path, endpoints, etc.)

See the [Pipeline Configuration](documentation/docs/pipelines/abapEnvironment/configuration.md) page for the full reference.

## Real-world example

The files below are a direct translation of a real Jenkins `.pipeline/config` into the GitHub Actions equivalent. The system is provisioned on the BTP **canary** landscape with a CF-based service (`abap/internal`), BYOG credentials are read from GitHub secrets, and ATC + AUnit run against two software components.

### Repository layout

```
my-addon-repo/
├── .github/
│   └── workflows/
│       └── abap-pipeline.yml        # consumer workflow
├── .pipeline/
│   └── config.yml                   # step parameters (non-secret)
├── addon.yml                        # add-on product descriptor
├── atcConfig.yml                    # ATC object set
└── aUnitConfig.yml                  # AUnit object set
```

### `.github/workflows/abap-pipeline.yml`

```yaml
name: ABAP Environment Pipeline

on:
  workflow_dispatch:
  schedule:
    - cron: '0 3 * * 1-5'   # nightly Mon–Fri at 03:00 UTC

jobs:
  pipeline:
    uses: SAP/jenkins-library/.github/workflows/abapEnvironmentPipeline.yml@main
    with:
      addonDescriptorFileName: addon.yml
      repositories: repos.yml

      # CF provisioning path — matches old stages.Prepare System
      cfApiEndpoint: https://api.cf.sap.hana.ondemand.com
      cfOrg:   my-cf-org
      cfSpace: my-cf-space

      # Stage toggles
      stage_prepareSystem:    true
      stage_cloneRepositories: true
      stage_atc:              true
      stage_aunit:            true
      stage_build:            false   # testing-only, no add-on build
      stage_confirm:          false
      stage_publish:          false

    secrets:
      PIPER_abapAddonAssemblyKitCookie: ${{ secrets.PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE }}
      PIPER_user:        ${{ secrets.PIPER_USER }}
      PIPER_password:    ${{ secrets.PIPER_PASSWORD }}
      PIPER_cfUser:      ${{ secrets.PIPER_CF_USER }}
      PIPER_cfPassword:  ${{ secrets.PIPER_CF_PASSWORD }}
```

### `.pipeline/config.yml`

Maps directly from the old `general:` and per-stage keys. Secrets stay in GitHub — only non-secret values go here.

```yaml
general:
  verbose: true

  # BTP service API endpoint (canary landscape)
  url: https://canary.cli.btp.int.sap

  # BTP subaccount and CF space identifiers
  subdomain:  2fcd6ac6-b8e0-40e8-aa71-a357aa99585e
  subaccount: f57f211e-2733-4cc6-b645-74f02d034a58

  # Persistent ABAP system used when stage_prepareSystem: false
  serviceInstanceName: demo_system_btp_steps_x

steps:
  # Clone Repositories — matches old stages.Clone Repositories
  abapEnvironmentCloneGitRepo:
    repositories: repos.yml
    strategy: Clone

  # ATC — matches old stages.ATC
  abapEnvironmentRunATCCheck:
    atcConfig: atcConfig.yml

  # AUnit — matches old stages.AUnit
  abapEnvironmentRunAUnitTest:
    aUnitConfig: aUnitConfig.yml

  # Prepare System (CF path) — matches old stages.Prepare System
  abapEnvironmentCreateSystem:
    cfServiceName:    abap
    cfServicePlan:    internal
    cfServiceInstance: demo_system_btp_steps_x
    additionalParameters: >-
      {
        "admin_email": "dannick.arnold.kwengang.tankeu@sap.com",
        "description": "ABAP Environment Test System",
        "is_development_allowed": false,
        "sapsystemname": "DK3",
        "size_of_persistence": 4,
        "size_of_runtime": 1
      }

  # Post — matches old stages.Post
  abapLandscapePortalUpdateAddOnProduct:
    deleteServiceBindings: true
```

> **BYOG credentials** (`byogCredentialsId: byog` in the old config) become the GitHub secrets `PIPER_USER` / `PIPER_PASSWORD`. Store them under **Settings → Secrets and variables → Actions** and reference them in the `secrets:` block above.

### `repos.yml`

```yaml
repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
```

### `atcConfig.yml`

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/GIT_REPOSITORY
```

### `aUnitConfig.yml`

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/GIT_REPOSITORY
```

### Mapping: Jenkins config → GitHub Actions

| Jenkins `.pipeline/config` key | GitHub Actions equivalent |
|-------------------------------|--------------------------|
| `general.url` | `.pipeline/config.yml` → `general.url` |
| `general.subdomain` / `subaccount` | `.pipeline/config.yml` → `general.subdomain` / `subaccount` |
| `general.abapCredentialsId` | GitHub secrets `PIPER_USER` + `PIPER_PASSWORD` |
| `general.byogCredentialsId` | GitHub secrets `PIPER_USER` + `PIPER_PASSWORD` |
| `general.serviceInstanceName` | `.pipeline/config.yml` → `general.serviceInstanceName` |
| `stages.Prepare System.offeringName` / `planName` | `.pipeline/config.yml` → `abapEnvironmentCreateSystem.cfServiceName/Plan` |
| `stages.Prepare System.parameters` | `.pipeline/config.yml` → `abapEnvironmentCreateSystem.additionalParameters` |
| `stages.Clone Repositories.repositories` | Consumer workflow `with.repositories:` |
| `stages.Clone Repositories.strategy` | `.pipeline/config.yml` → `abapEnvironmentCloneGitRepo.strategy` |
| `stages.ATC.atcConfig` | `.pipeline/config.yml` → `abapEnvironmentRunATCCheck.atcConfig` |
| `stages.AUnit.aUnitConfig` | `.pipeline/config.yml` → `abapEnvironmentRunAUnitTest.aUnitConfig` |
| `stages.Post.deleteServiceBindings` | `.pipeline/config.yml` → `abapLandscapePortalUpdateAddOnProduct.deleteServiceBindings` |

## Scenarios

- **[Build and publish add-on products](documentation/docs/scenarios/abapEnvironmentAddons.md)** — full build pipeline for SaaS delivery
- **[Continuous testing](documentation/docs/scenarios/abapEnvironmentTest.md)** — ATC checks and AUnit tests against a permanent or transient system

## Documentation

Full documentation is in the `documentation/` folder. Build and browse locally:

```sh
docker run --rm -it -p 8000:8000 -v "${PWD}/documentation:/docs" squidfunk/mkdocs-material:8.5.11
```

## Steps included

- `abapAddonAssemblyKit*` — AAKaaS interaction (reserve packages, register, release, publish)
- `abapEnvironment*` — ABAP system operations (clone, ATC, AUnit, build, assembly)
- `abapLandscapePortalUpdateAddOnProduct` — Landscape Portal reporting
- `btpCreateServiceBinding`, `btpCreateServiceInstance`, `btpDeleteServiceBinding`, `btpDeleteServiceInstance` — BTP service provisioning
- `cloudFoundryCreateServiceKey`, `cloudFoundryDeleteService` — CF provisioning path
