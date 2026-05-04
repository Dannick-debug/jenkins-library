# Infrastructure Overview

This section covers the infrastructure required to run the ABAP Environment Pipeline.

## Jenkins

The ABAP Environment Pipeline runs on Jenkins. You need a Jenkins instance with:

- The [Piper shared library](https://github.com/SAP/jenkins-library) configured as a global pipeline library named `piper-lib-os`
- Docker installed on the Jenkins agent (most steps execute in Docker containers)
- Credentials configured for:
  - SAP BTP service account (for BTP API steps)
  - Cloud Foundry credentials (for CF-based system provisioning)
  - ABAP communication user credentials

See the [Custom Jenkins Setup](customjenkins.md) guide for detailed installation instructions.

## Credentials

The pipeline requires credentials stored in Jenkins or fetched from Vault:

| Credential | Used by |
| ---------- | ------- |
| CF username/password | `abapEnvironmentCreateSystem`, `cloudFoundryCreateServiceKey`, `cloudFoundryDeleteService` |
| ABAP communication user | All `abapEnvironment*` steps communicating with the ABAP system |
| BTP service account | `btpCreate/DeleteServiceInstance/Binding` |

## Vault (optional)

Secrets can be fetched directly from [HashiCorp Vault](https://www.hashicorp.com/products/vault) instead of Jenkins credentials. See the [Vault for Pipeline Secrets](vault.md) guide.

## Docker

Most steps pull their required tools from Docker Hub as needed — no manual installation required. If you hit Docker Hub rate limits, see the [Docker Rate Limit](docker-rate-limit.md) guide.
