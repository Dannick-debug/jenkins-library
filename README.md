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
