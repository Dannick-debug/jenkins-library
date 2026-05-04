# abapEnvironmentCreateSystem

Creates a SAP BTP ABAP Environment system (aka Steampunk system)


## Description

This step creates a SAP BTP ABAP Environment system (aka Steampunk system) via the cloud foundry command line interface (cf CLI). This can be done by providing a service manifest as a configuration file (parameter `serviceManifest`) or by passing the configuration values directly via the other parameters of this step.


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentCreateSystem script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentCreateSystem
        ```



## Prerequisites

- On SAP Business Technology Platform (SAP BTP), Cloud Foundry needs to be enabled on subaccount level. This can be done on the Subaccount Overview page. The subaccount is then mapped to a “Cloud Foundry Organization”, for which you must provide a suitable name during the creation. Have a look at the [documentation](https://help.sap.com/viewer/a96b1df8525f41f79484717368e30626/Cloud/en-US/dc18bac42270468d84b6c030a668e003.html) for more details.
- A (technical) user is required to access the SAP BTP via the cf CLI. The user needs to be a [member of the global account](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/4a0491330a164f5a873fa630c7f45f06.html) and has to have the [Space Developer](https://help.sap.com/viewer/a96b1df8525f41f79484717368e30626/Cloud/en-US/967fc4e2b1314cf7afc7d7043b53e566.html) role. The user and password need to be stored in the Jenkins Credentials Store.
- Please make sure, that there are enough entitlements in the subaccount for the [Service Plan](https://help.sap.com/viewer/a96b1df8525f41f79484717368e30626/Cloud/en-US/c40cb18aeaa343389036fdcdd03c41d0.html), which you want to use for this step.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [cfOrg](#cforg) | **yes** |  |
| [cfSpace](#cfspace) | **yes** |  |
| [password](#password) | **(yes)** |  ![Vault](https://img.shields.io/badge/-Vault-lightgrey) ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV, Vault or Jenkins credentials ([`cfCredentialsId`](#cfcredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** |  ![Vault](https://img.shields.io/badge/-Vault-lightgrey) ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV, Vault or Jenkins credentials ([`cfCredentialsId`](#cfcredentialsid)) |
| [abapSystemAdminEmail](#abapsystemadminemail) | no |  |
| [abapSystemDescription](#abapsystemdescription) | no |  |
| [abapSystemID](#abapsystemid) | no |  |
| [abapSystemIsDevelopmentAllowed](#abapsystemisdevelopmentallowed) | no |  |
| [abapSystemSizeOfPersistence](#abapsystemsizeofpersistence) | no |  |
| [abapSystemSizeOfRuntime](#abapsystemsizeofruntime) | no |  |
| [addonDescriptorFileName](#addondescriptorfilename) | no |  |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfService](#cfservice) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServicePlan](#cfserviceplan) | no |  |
| [includeAddon](#includeaddon) | no |  |
| [serviceManifest](#servicemanifest) | no |  |
| [verbose](#verbose) | no | activates debug output |

### Overview - Execution Environment

!!! note "Orchestrator-specific only"

    These parameters are relevant for orchestrator usage and not considered when using the command line option.

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [containerCommand](#containercommand) | no | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) |
| [containerShell](#containershell) | no | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) |
| [dockerEnvVars](#dockerenvvars) | no |  |
| [dockerImage](#dockerimage) | no |  |
| [dockerName](#dockername) | no |  |
| [dockerOptions](#dockeroptions) | no |  |
| [dockerPullImage](#dockerpullimage) | no |  |
| [dockerVolumeBind](#dockervolumebind) | no | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) |
| [dockerWorkspace](#dockerworkspace) | no | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) |

### Details

#### abapSystemAdminEmail

Admin E-Mail address for the initial administrator of the system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_abapSystemAdminEmail` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapSystemDescription

Description for the ABAP Environment system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `Test system created by an automated pipeline` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapSystemID

The three character name of the system - maps to 'sapSystemName'

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `H02` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapSystemIsDevelopmentAllowed

This parameter determines, if development is allowed on the system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `true` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapSystemSizeOfPersistence

The size of the persistence

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `int` |
| Mandatory | no |
| Default | `0` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapSystemSizeOfRuntime

The size of the runtime

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `int` |
| Mandatory | no |
| Default | `0` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### addonDescriptorFileName

The file name of the addonDescriptor

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_addonDescriptorFileName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfApiEndpoint

Cloud Foundry API endpoint

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/apiEndpoint` |
| Type | `string` |
| Mandatory | no |
| Default | `https://api.cf.eu10.hana.ondemand.com` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfOrg

Cloud Foundry org

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/org` |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_cfOrg` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfService

Parameter for Cloud Foundry Service to be used for creating Cloud Foundry Service

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/service` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfService` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfServiceInstance

Parameter for naming the Service Instance when creating a Cloud Foundry Service

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/serviceInstance` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfServiceInstance` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfServicePlan

Parameter for Cloud Foundry Service Plan to be used when creating a Cloud Foundry Service

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/servicePlan` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfServicePlan` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfSpace

Cloud Foundry Space

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/space` |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_cfSpace` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### containerCommand

**Jenkins-specific:** Used for proper environment setup.

Kubernetes only: Allows to specify start command for container created with dockerImage parameter to overwrite Piper default (/usr/bin/tail -f /dev/null).

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### containerShell

**Jenkins-specific:** Used for proper environment setup.

Allows to specify the shell to be executed for container with containerName.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerEnvVars

Environment variables to set in the container, e.g. [http_proxy: "proxy:8080"].

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `map[string]string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerImage

Name of the docker image that should be used. If empty, Docker is not used and the command is executed directly on the Jenkins system.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `ppiper/cf-cli:latest` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerName

Kubernetes only: Name of the container launching dockerImage. SideCar only: Name of the container in local network.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `cf` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerOptions

Docker options to be set when starting the container.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `[]string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerPullImage

Set this to 'false' to bypass a docker image pull. Useful during development process. Allows testing of images which are available in the local registry only.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `true` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerVolumeBind

**Jenkins-specific:** Used for proper environment setup.

Volumes that should be mounted into the docker container.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `map[string]string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### dockerWorkspace

**Jenkins-specific:** Used for proper environment setup.

Kubernetes only: Specifies a dedicated user home directory for the container which will be passed as value for environment variable `HOME`.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### includeAddon

Must be set to true to install the addon provided via 'addonDescriptorFileName'

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### password

Password for Cloud Foundry User

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_password` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`cfCredentialsId`](#cfcredentialsid)<br />&nbsp;&nbsp;reference to: `password`<br /><br/>Vault resource:<br />&nbsp;&nbsp;name: `cloudfoundryVaultSecretName`<br />&nbsp;&nbsp;default value: `cloudfoundry-$(org)-$(space)`<br /><br/>Vault paths: <br /><ul><li>`$(vaultPath)/cloudfoundry-$(org)-$(space)`</li><li>`$(vaultBasePath)/$(vaultPipelineName)/cloudfoundry-$(org)-$(space)`</li><li>`$(vaultBasePath)/GROUP-SECRETS/cloudfoundry-$(org)-$(space)`</li></ul> |


#### script

The common script environment of the Jenkinsfile running. Typically the reference to the script calling the pipeline step is provided with the `this` parameter, as in `script: this`. This allows the function to access the `commonPipelineEnvironment` for retrieving, e.g. configuration parameters.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `Jenkins Script` |
| Mandatory | **yes** |
| Default |  |
| Secret | no |
| Configuration scope | <ul><li>&#9744; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
| Resource references | none |


#### serviceManifest

Path to Cloud Foundry Service Manifest in YAML format for multiple service creations that are being passed to a Create-Service-Push Cloud Foundry cli plugin

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - `cloudFoundry/serviceManifest`<br />- `cfServiceManifest` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_serviceManifest` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### username

User or E-Mail for CF

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_username` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`cfCredentialsId`](#cfcredentialsid)<br />&nbsp;&nbsp;reference to: `username`<br /><br/>Vault resource:<br />&nbsp;&nbsp;name: `cloudfoundryVaultSecretName`<br />&nbsp;&nbsp;default value: `cloudfoundry-$(org)-$(space)`<br /><br/>Vault paths: <br /><ul><li>`$(vaultPath)/cloudfoundry-$(org)-$(space)`</li><li>`$(vaultBasePath)/$(vaultPipelineName)/cloudfoundry-$(org)-$(space)`</li><li>`$(vaultBasePath)/GROUP-SECRETS/cloudfoundry-$(org)-$(space)`</li></ul> |


#### verbose

verbose output

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfCredentialsId

Jenkins 'Username with password' credentials ID containing user and password to authenticate to the Cloud Foundry API.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/credentialsId` |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Example: Configuration in the config.yml

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapEnvironmentCreateSystem script: this
```

The configuration values for the system can be passed through the `config.yml` file:

```yaml
steps:
  abapEnvironmentCreateSystem:
    cfCredentialsId: 'cfCredentialsId'
    cfApiEndpoint: 'https://test.server.com'
    cfOrg: 'cfOrg'
    cfSpace: 'cfSpace'
    cfServiceInstance: 'H02_Q_system'
    cfService: 'abap'
    cfServicePlan: 'standard'
    abapSystemAdminEmail: 'user@example.com'
    abapSystemDescription: 'ABAP Environment Q System'
    abapSystemIsDevelopmentAllowed: true
    abapSystemID: 'H02'
    abapSystemSizeOfPersistence: 4
    abapSystemSizeOfRuntime: 1
```

## Example: Configuration in the Jenkinsfile

The step, including all parameters, can also be called directly from the Jenkinsfile. In the following example, a configuration file is used.

```groovy
abapEnvironmentCreateSystem (
  script: this,
  cfCredentialsId: 'cfCredentialsId',
  cfApiEndpoint: 'https://test.server.com',
  cfOrg: 'cfOrg',
  cfSpace: 'cfSpace',
  cfServiceManifest: 'manifest.yml'
)
```

The file `manifest.yml` would look like this:

```yaml
---
create-services:
- name:   "H02_Q_system"
  broker: "abap"
  plan:   "standard"
  parameters: "{ \"admin_email\" : \"user@example.com\", \"description\" : \"ABAP Environment Q System\", \"is_development_allowed\" : true, \"sapsystemname\" : \"H02\", \"size_of_persistence\" : 4, \"size_of_runtime\" : 1 }"
```
