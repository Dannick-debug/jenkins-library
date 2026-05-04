# Infrastructure Overview

The ABAP Environment Pipeline runs entirely on **GitHub Actions**. No Jenkins instance, no Docker daemon on a build server, and no Groovy runtime are required.

## What you need

| Requirement | Details |
|-------------|---------|
| GitHub repository | Hosts your `addon.yml`, `repositories.yml`, and `.pipeline/config.yml` |
| GitHub Actions | Enabled on the repository (free for public repos, included in GitHub plans for private) |
| GitHub Environment | Named `production` (or custom) with required reviewers for the manual Confirm gate |
| SAP BTP, ABAP Environment | Target system — provisioned via BTP service API or Cloud Foundry |
| AAKaaS access | Authentication cookie for the ABAP Add-on Assembly Kit as a Service |
| ABAP communication user | User/password for the ABAP system steps |

## Credentials

All credentials are stored as **GitHub repository secrets** and passed to the `piper` binary via `PIPER_` prefixed environment variables. No secrets are written to disk.

| Secret | Used by |
|--------|---------|
| `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE` | All `abapAddonAssemblyKit*` steps |
| `PIPER_USER` / `PIPER_PASSWORD` | All `abapEnvironment*` steps |
| `PIPER_CF_USER` / `PIPER_CF_PASSWORD` | CF provisioning path |
| `PIPER_BTP_API_CREDENTIALS_ID` | BTP provisioning path |

Add secrets at **Settings → Secrets and variables → Actions**.

## Confirm gate (GitHub Environment)

The pipeline pauses before publishing and waits for manual approval. This is implemented using a [GitHub Environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment) with required reviewers.

1. Go to **Settings → Environments → New environment**
2. Name it `production` (or whatever you set as `confirmEnvironment` in the workflow)
3. Add **Required reviewers** under deployment protection rules

Reviewers receive a notification and can approve or reject directly from the GitHub UI.

## Common pipeline environment (CPE)

The `piper` binary persists state between steps in the `.pipeline/` directory. Because each GitHub Actions job runs on an isolated runner, the reusable workflow uploads `.pipeline/` as an artifact after every job and downloads it at the start of the next. This is handled automatically — no configuration required.

## The piper binary

The `piper` Linux binary is downloaded at the start of each job directly from the GitHub releases page. No pre-installation on the runner is needed. See the [CLI reference](../cli/index.md) for available commands.
