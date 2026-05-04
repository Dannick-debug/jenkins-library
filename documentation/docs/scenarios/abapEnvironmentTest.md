# Continuous Testing on SAP BTP, ABAP Environment

This scenario describes how to run automated ATC checks and AUnit tests against an ABAP system as part of a GitHub Actions pipeline.

## Overview

Three pipeline stages are used:

1. **Prepare System** _(optional)_ — provisions a fresh ABAP system per run
2. **Clone Repositories** — clones software components into the system
3. **ATC + AUnit** — runs checks and tests in parallel

If Prepare System is disabled, the pipeline runs against a permanent quality system. This is faster (no provisioning wait) but shares state between runs.

## Configuration

### Enable the stages

In your consumer workflow, set the relevant toggles:

```yaml
with:
  addonDescriptorFileName: addon.yml
  repositories: repositories.yml
  stage_prepareSystem: true       # false to use a permanent system
  stage_cloneRepositories: true
  stage_atc: true
  stage_aunit: true
  stage_build: false              # not needed for testing-only
  stage_confirm: false
  stage_publish: false
```

### `repositories.yml`

```yaml
repositories:
  - name: /DMO/GIT_REPOSITORY
    branch: main
```

### ATC (`atcConfig.yml`)

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/GIT_REPOSITORY
```

Configure in `.pipeline/config.yml`:

```yaml
steps:
  abapEnvironmentRunATCCheck:
    atcConfig: atcConfig.yml
```

### AUnit (`aUnitConfig.yml`)

```yaml
objectSet:
  softwarecomponent:
    - name: /DMO/GIT_REPOSITORY
```

Configure in `.pipeline/config.yml`:

```yaml
steps:
  abapEnvironmentRunAUnitTest:
    aUnitConfig: aUnitConfig.yml
```

## Scheduling nightly runs

Add a `schedule` trigger to your consumer workflow:

```yaml
on:
  workflow_dispatch:
  schedule:
    - cron: '0 3 * * *'   # every night at 03:00 UTC
```

## Sample configurations

- [Transient system (new system per run)](https://github.com/SAP-samples/abap-platform-ci-cd-samples/tree/atc-transient)
- [Permanent system (fixed quality system)](https://github.com/SAP-samples/abap-platform-ci-cd-samples/tree/atc-static)
