# Running the ABAP Environment Pipeline on GitHub Actions

This guide explains how to run the ABAP Environment Pipeline using the `piper` CLI and GitHub Actions instead of Jenkins.

## Prerequisites

- A GitHub repository containing your `addon.yml` and ABAP repository configuration
- The `piper` CLI binary (downloaded automatically by the workflow)
- Access to SAP BTP, ABAP Environment (either via BTP service provisioning or Cloud Foundry)
- An AAKaaS authentication cookie

## Quick Start

Run the scaffold command in your repository root:

```sh
piper initGithubActions \
  --addon-descriptor addon.yml \
  --btp-subaccount-id <YOUR_BTP_SUBACCOUNT_ID>
```

For the Cloud Foundry provisioning path:

```sh
piper initGithubActions \
  --addon-descriptor addon.yml \
  --cf-api-endpoint https://api.cf.<region>.hana.ondemand.com \
  --cf-org <ORG> \
  --cf-space <SPACE>
```

This generates:

| File | Purpose |
|------|---------|
| `.github/workflows/abap-pipeline.yml` | Consumer workflow that calls the reusable pipeline |
| `.pipeline/config.yml` | Step configuration skeleton |
| `repositories.yml` | ABAP repository list skeleton |

## Configuration

### `.pipeline/config.yml`

Non-secret step parameters go here. Example:

```yaml
general:
  abapAddonAssemblyKitEndpoint: "https://aakaas.example.com"

steps:
  abapEnvironmentRunATCCheck:
    atcConfig: atcConfig.yml
  abapEnvironmentRunAUnitTest:
    aUnitConfig: aUnitConfig.yml
```

### `repositories.yml`

List the ABAP Git repositories to clone. Example:

```yaml
repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
    version: v1.0.0
```

## Setting Up GitHub Secrets

Go to **Settings → Secrets and variables → Actions** and add:

| Secret name | Description |
|-------------|-------------|
| `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE` | AAKaaS authentication cookie |
| `PIPER_USER` | ABAP system user |
| `PIPER_PASSWORD` | ABAP system password |
| `PIPER_CF_USER` | Cloud Foundry user (CF path only) |
| `PIPER_CF_PASSWORD` | Cloud Foundry password (CF path only) |
| `PIPER_BTP_API_CREDENTIALS_ID` | BTP API credentials (BTP path only) |

## Setting Up the Confirm Gate

The pipeline pauses before publishing and requires manual approval. This is implemented using a [GitHub Environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment) with required reviewers.

1. Go to **Settings → Environments → New environment**
2. Name it `production` (or the value you passed to `--confirm-environment`)
3. Under **Deployment protection rules**, add **Required reviewers**

## Pipeline Stages

The reusable workflow defines the following jobs:

| Job | Condition |
|-----|-----------|
| `init` | Always |
| `initial-checks` | `stage_initialChecks: true` (default) |
| `prepare-system` | `stage_prepareSystem: true` (default) |
| `clone-repositories` | `stage_cloneRepositories: true` (default) |
| `atc` | `stage_atc: true` (default) — parallel with `aunit` |
| `aunit` | `stage_aunit: true` (default) — parallel with `atc` |
| `build` | `stage_build: true` (default) |
| `integration-tests` | `stage_integrationTest: true` (default: false) |
| `confirm` | `stage_confirm: true` (default) — manual gate |
| `publish` | `stage_publish: true` (default) |
| `post` | Always (`if: always()`) — cleanup |

Toggle any stage off by setting the corresponding input to `false` in your consumer workflow.

## Passing State Between Jobs (CPE)

The `piper` CLI persists pipeline state in the `.pipeline/` directory (the *common pipeline environment*, or CPE). Because GitHub Actions jobs run in isolated runners, the workflow uploads `.pipeline/` as an artifact after each job and downloads it at the start of the next.

This is handled automatically by the reusable workflow — no action needed on your part.

## Troubleshooting

**`pipeline-env` artifact not found on first job**

The `init` job creates the artifact. Subsequent jobs use `continue-on-error: true` when downloading, so missing artifacts on the first run do not cause failures.

**Confirm gate not pausing**

Ensure the GitHub Environment named `production` (or your custom name) has at least one required reviewer configured. Without required reviewers the environment approval is skipped.

**`PIPER_user` vs `PIPER_username`**

BTP steps (`btpCreateServiceInstance`, etc.) use the `user` parameter. Cloud Foundry steps (`cloudFoundryCreateServiceKey`, etc.) use the `username` parameter. If you use both provisioning paths, set both `PIPER_USER` (maps to `user`) and check your CF step config.
