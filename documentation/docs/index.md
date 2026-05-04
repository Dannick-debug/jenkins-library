# ABAP Environment Pipeline

A CI/CD toolkit for building and delivering ABAP add-on products on SAP BTP, ABAP Environment (Steampunk).

## What's included

- A **`piper` CLI binary** that implements every pipeline step as a sub-command
- A **GitHub Actions reusable workflow** that orchestrates those steps into the full pipeline
- A **`piper initGithubActions`** scaffold command that wires everything up in your repository in seconds

## How it works

The `piper` binary is the single execution unit. Each pipeline step — ATC checks, package assembly, target vector publishing — is one `piper <step>` call. The GitHub Actions reusable workflow sequences those calls and handles:

- Artifact-based state propagation between jobs (common pipeline environment)
- Parallel ATC and AUnit execution
- A manual Confirm gate backed by GitHub Environments with required reviewers
- System cleanup in a Post job that always runs

Consumer repos reference the reusable workflow with two lines of YAML and never need to maintain pipeline logic themselves.

## Quick start

```sh
# Download the CLI
curl -sLo piper "https://github.com/SAP/jenkins-library/releases/latest/download/piper_linux_amd64"
chmod +x piper && sudo mv piper /usr/local/bin/piper

# Scaffold GitHub Actions files in your ABAP project repo
cd my-abap-addon-repo
piper initGithubActions --btp-subaccount-id <YOUR_SUBACCOUNT_ID>
```

Then follow the checklist printed by the command — add secrets, create a GitHub Environment, and push.

See the [Getting Started guide](guidedtour.md) for the full walkthrough.

## Pipeline stages

| Stage | GitHub Actions job |
|-------|--------------------|
| Init | `init` |
| Initial Checks | `initial-checks` |
| Prepare System | `prepare-system` |
| Clone Repositories | `clone-repositories` |
| ATC Check _(parallel)_ | `atc` |
| AUnit Tests _(parallel)_ | `aunit` |
| Build | `build` |
| Integration Tests _(optional)_ | `integration-tests` |
| Confirm _(manual gate)_ | `confirm` |
| Publish | `publish` |
| Post _(always runs)_ | `post` |

## Scenarios

- [Build and Publish Add-on Products on SAP BTP, ABAP Environment](scenarios/abapEnvironmentAddons.md)
- [Continuous Testing on SAP BTP, ABAP Environment](scenarios/abapEnvironmentTest.md)
