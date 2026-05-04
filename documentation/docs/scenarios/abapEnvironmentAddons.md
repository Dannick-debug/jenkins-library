# Build and Publish Add-on Products on SAP BTP, ABAP Environment

This scenario describes how SAP partners build and publish ABAP add-on products for SaaS delivery on SAP BTP, ABAP Environment (Steampunk).

## Introduction

Add-on products bundle one or more ABAP software component versions into a shippable unit. The ABAP Environment Pipeline automates the build process: reserving packages, assembling them on the ABAP system, registering them with AAKaaS, publishing the target vector, and deprovisioning the build system.

Once published, the add-on is available for installation into customer ABAP systems provisioned as multitenant applications.

## The Add-on Descriptor (`addon.yml`)

The build is driven by `addon.yml` in your repository root:

```yaml
addonProduct: /NAMESPC/PRODUCTX
addonVersion: "1.2.0"
repositories:
  - name: /NAMESPC/COMPONENTA
    branch: v1.2.0
    version: "1.2.0"
    commitID: "7d4516e9"
    languages:
      - DE
      - EN
  - name: /NAMESPC/COMPONENTB
    branch: v2.0.0
    version: "2.0.0"
    commitID: "9f102ffb"
    languages:
      - DE
      - EN
      - FR
```

**Field reference:**

| Field | Description |
|-------|-------------|
| `addonProduct` | Technical name: `/<namespace>/<product>` |
| `addonVersion` | `<release>.<support package stack>.<patch level>` |
| `name` | Software component technical name |
| `branch` | Git branch in the ABAP system |
| `version` | `<release>.<support package>.<patch>` |
| `commitID` | Git commit that defines the exact state to build |
| `languages` | ISO-639 language codes to deliver |

### Versioning rules

Version numbers must increase without gaps:

- `1.0.0 → 2.0.0` ✓
- `1.1.2 → 2.0.0` ✓
- `1.0.0 → 3.0.0` ✗ (version 2.0.0 missing)
- `2.1.0 → 2.3.0` ✗ (version 2.2.0 missing)

## Build Pipeline Stages

| Stage | What happens |
|-------|-------------|
| Initial Checks | AAKaaS validates product and component versions |
| Prepare System | Provisions a transient ABAP build system |
| Clone Repositories | Clones software components at the specified commit |
| ATC | Static code checks (blocks the build on findings) |
| AUnit | Unit test execution |
| Build | Reserves package slots, assembles packages, registers and releases them with AAKaaS |
| Integration Tests | Optional: installs the add-on on a test system and runs verification checks |
| Confirm | Manual approval before publishing (GitHub Environment gate) |
| Publish | Publishes the target vector — add-on becomes available for installation |
| Post | Deprovisions the build system and reports to Landscape Portal |

## Prerequisites

### Register your add-on product

Before the first build, register your add-on product for your SAP BTP global account using the [Landscape Portal](https://help.sap.com/docs/BTP/65de2977205c403bbc107264b8eccf4b/dc15fb4ebab5453fa4641b98190b1f85.html). This is a one-time manual step.

### Reserve a namespace

Development objects must live in a reserved namespace. See [SAP Note 105132](https://launchpad.support.sap.com/#/notes/105132).

### AAKaaS access

Obtain an AAKaaS technical communication user as described in [SAP Note 2174416](https://launchpad.support.sap.com/#/notes/2174416). Store the authentication cookie as the GitHub secret `PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE`.

### ABAP assembly system

Use a permanent `abap/standard` system provisioned with `is_development_allowed = false`. See [assembly system recommendations](introduction.md).

## ATC quality gate

!!! caution "Block on ATC findings"
    Configure your ATC check variant to block any error or warning findings. Delivery packages are final — fixing issues requires a new `addonVersion`. Resolve ATC findings during development before triggering a build.

    Include all software components from `addon.yml` in `atcConfig.yml`.

## Troubleshooting

| Error | Cause | Resolution |
|-------|-------|------------|
| `Quota is not sufficient` (Prepare System) | Missing `abap/standard` entitlement | Assign entitlements to the subaccount |
| `Branch checkout is currently performed` | Parallel actions on same software component | Wait and re-run |
| `Package was already built but with commit X` | `commitID` changed without bumping `version` | Increment the software component `version` |
| `CommitID of package is the same as predecessor` | Same commit for a new patch level | Use a new `commitID` |
| `Product not registered for productive development` | Add-on not registered | Follow [Register Add-on Product](https://help.sap.com/docs/BTP/65de2977205c403bbc107264b8eccf4b/dc15fb4ebab5453fa4641b98190b1f85.html) |

## Support components

| Stage | Support component |
|-------|------------------|
| Initial Checks, Build (AAKaaS steps) | BC-UPG-OCS |
| Prepare System, Clone Repositories | BC-CP-ABA |
| ATC / AUnit | BC-DWB-TOO-ATF |
| Build (assembly steps) | BC-UPG-ADDON |

## Sample configuration

See [SAP-samples/abap-platform-ci-cd-samples (addon-build)](https://github.com/SAP-samples/abap-platform-ci-cd-samples/tree/addon-build) for a complete working example.
