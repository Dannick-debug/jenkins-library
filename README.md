# Project "Piper": ABAP Environment Pipeline

A lightweight fork of [SAP/jenkins-library](https://github.com/SAP/jenkins-library) scoped exclusively to the **ABAP Environment Pipeline** and **BTP API steps**.

This fork reduces maintenance overhead by removing everything unrelated to ABAP add-on build and delivery on SAP BTP, ABAP Environment.

## What's included

- **`abapEnvironmentPipeline`** — 11-stage pipeline for building and publishing ABAP add-on products on SAP BTP
- **ABAP steps** (`abapAddonAssemblyKit*`, `abapEnvironment*`, `abapLandscapePortalUpdateAddOnProduct`)
- **BTP API steps** (`btpCreateServiceBinding`, `btpCreateServiceInstance`, `btpDeleteServiceBinding`, `btpDeleteServiceInstance`)
- **Supporting Cloud Foundry steps** used by the pipeline (`cloudFoundryCreateServiceKey`, `cloudFoundryDeleteService`)

## What's removed

All steps, pipeline templates, documentation, and tests unrelated to the ABAP environment: general-purpose pipelines (CAP, UI5, MTA, etc.), CF application deployment steps, TMS, gCTS, Neo, and all associated Groovy utility classes.

## Documentation

See the `documentation/` folder. Build locally with:

```sh
docker run --rm -it -p 8000:8000 -v "${PWD}/documentation:/docs" squidfunk/mkdocs-material:8.5.11
```

## Upstream

This fork tracks [SAP/jenkins-library](https://github.com/SAP/jenkins-library). Upstream changes to ABAP/BTP steps are periodically merged; unrelated changes are not.
