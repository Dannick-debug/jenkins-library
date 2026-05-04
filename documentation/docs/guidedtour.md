# Getting Started

This guide walks you through setting up the ABAP Environment Pipeline in your GitHub repository using the `piper` CLI and GitHub Actions.

## Prerequisites

- A GitHub repository containing your ABAP add-on sources (`addon.yml`, `repositories.yml`)
- Access to SAP BTP, ABAP Environment — either via BTP service provisioning or Cloud Foundry
- An AAKaaS authentication cookie from SAP

## 1. Download the piper CLI

```sh
curl -sLo piper "https://github.com/SAP/jenkins-library/releases/latest/download/piper_linux_amd64"
chmod +x piper
sudo mv piper /usr/local/bin/piper
piper version
```

Specific releases are available on the [GitHub releases page](https://github.com/SAP/jenkins-library/releases).

## 2. Scaffold your pipeline files

Run `piper initGithubActions` in the root of your ABAP project repository.

**BTP provisioning path:**

```sh
piper initGithubActions \
  --addon-descriptor addon.yml \
  --btp-subaccount-id <YOUR_BTP_SUBACCOUNT_ID>
```

**Cloud Foundry provisioning path:**

```sh
piper initGithubActions \
  --addon-descriptor addon.yml \
  --cf-api-endpoint https://api.cf.<region>.hana.ondemand.com \
  --cf-org <ORG> \
  --cf-space <SPACE>
```

This generates three files:

| File | Purpose |
|------|---------|
| `.github/workflows/abap-pipeline.yml` | Consumer workflow that calls the reusable pipeline |
| `.pipeline/config.yml` | Step configuration (created only if absent) |
| `repositories.yml` | ABAP repository list (created only if absent) |

## 3. Add GitHub secrets

Go to **Settings → Secrets and variables → Actions** in your repository and add:

| Secret name | Description |
|-------------|-------------|
| `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE` | AAKaaS authentication cookie |
| `PIPER_USER` | ABAP system user |
| `PIPER_PASSWORD` | ABAP system password |
| `PIPER_CF_USER` | Cloud Foundry user _(CF path only)_ |
| `PIPER_CF_PASSWORD` | Cloud Foundry password _(CF path only)_ |
| `PIPER_BTP_API_CREDENTIALS_ID` | BTP API credentials _(BTP path only)_ |

## 4. Create the Confirm gate environment

The pipeline pauses before publishing and requires manual approval via a [GitHub Environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment).

1. Go to **Settings → Environments → New environment**
2. Name it `production` (the default, or whatever you passed to `--confirm-environment`)
3. Under **Deployment protection rules**, add **Required reviewers**

## 5. Edit your configuration files

**`addon.yml`** — describe your ABAP add-on product and component versions. See the [add-on descriptor reference](steps/abapAddonAssemblyKitCheck.md).

**`repositories.yml`** — list the ABAP Git repositories to clone:

```yaml
repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
    version: v1.0.0
```

**`.pipeline/config.yml`** — configure step behaviour (non-secret parameters):

```yaml
general:
  abapAddonAssemblyKitEndpoint: "https://aakaas.example.com"

steps:
  abapEnvironmentRunATCCheck:
    atcConfig: atcConfig.yml
  abapEnvironmentRunAUnitTest:
    aUnitConfig: aUnitConfig.yml
```

## 6. Push and trigger

```sh
git add .github/workflows/abap-pipeline.yml .pipeline/config.yml repositories.yml addon.yml
git commit -m "feat: add ABAP Environment Pipeline"
git push
```

Then go to **Actions → ABAP Environment Pipeline → Run workflow**.

## What's next

- Read the full [pipeline configuration reference](pipelines/abapEnvironment/configuration.md)
- Understand each stage in the [pipeline introduction](pipelines/abapEnvironment/introduction.md)
- Learn about the [piper CLI commands](cli/index.md)
