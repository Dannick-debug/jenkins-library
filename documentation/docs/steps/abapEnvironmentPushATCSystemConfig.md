# abapEnvironmentPushATCSystemConfig

Create/Update ATC System Configuration


## Description

This step is for creating/updating an [ATC](https://help.sap.com/products/BTP/65de2977205c403bbc107264b8eccf4b/657285a09f7148d894c27bb8e17827cf.html?version=Cloud) system configurationon on an SAP BTP, ABAP Environment system.
Please provide either of the following options:

* The host and credentials the SAP BTP, ABAP Environment system itself. The credentials must be configured for the Communication Scenario [SAP_COM_0763](https://help.sap.com/products/BTP/65de2977205c403bbc107264b8eccf4b/657285a09f7148d894c27bb8e17827cf.html?version=Cloud).
* The Cloud Foundry parameters (API endpoint, organization, space), credentials, the service instance for the ABAP service and the service key for the Communication Scenario SAP_COM_0763.
* Only provide one of those options with the respective credentials. If all values are provided, the direct communication (via host) has priority.


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentPushATCSystemConfig script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentPushATCSystemConfig
        ```



## Prerequisites

* A SAP BTP, ABAP environment system is available. On this system, a [Communication User](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0377adea0401467f939827242c1f4014.html), a [Communication System](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/1bfe32ae08074b7186e375ab425fb114.html) and a [Communication Arrangement](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/a0771f6765f54e1c8193ad8582a32edb.html) is setup for the Communication Scenario “ABAP Test Cockpit Configuration Integration (SAP_COM_0763)“. This can be done manually through the respective applications on the SAP BTP, ABAP environment system or through creating a service key for the system on Cloud Foundry with the parameters {“scenario_id”: “SAP_COM_0763", “type”: “basic”}. In a pipeline, you can do this with the step [cloudFoundryCreateServiceKey](https://sap.github.io/jenkins-library/steps/cloudFoundryCreateServiceKey/).
* You can either provide the ABAP endpoint configuration to directly trigger an ATC run on the ABAP system or optionally provide the Cloud Foundry parameters with your credentials to read a Service Key of a SAP BTP, ABAP environment system in Cloud Foundry that contains all the details of the ABAP endpoint to trigger an ATC run.
* Regardless if you chose an ABAP endpoint directly or reading a Cloud Foundry Service Key, you have to provide the configuration of the packages and software components you want to be checked in an ATC run in a .yml or .yaml file. This file must be stored in the same folder as the Jenkinsfile defining the pipeline.
* The software components and/or packages you want to be checked must be present in the configured system in order to run the check. Please make sure that you have created or pulled the respective software components and/or Packages in the SAP BTP, ABAP environment system.

Examples will be listed below.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [atcSystemConfigFilePath](#atcsystemconfigfilepath) | **yes** |  |
| [password](#password) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfOrg](#cforg) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServiceKeyName](#cfservicekeyname) | no |  |
| [cfSpace](#cfspace) | no |  |
| [host](#host) | no |  |
| [idp](#idp) | no |  |
| [patchIfExisting](#patchifexisting) | no |  |
| [serviceBindingName](#servicebindingname) | no |  |
| [serviceInstanceName](#serviceinstancename) | no |  |
| [subaccount](#subaccount) | no |  |
| [subdomain](#subdomain) | no |  |
| [url](#url) | no |  |
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

#### atcSystemConfigFilePath

Path to a JSON file with ATC System Configuration

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_atcSystemConfigFilePath` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfApiEndpoint

Cloud Foundry API endpoint

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/apiEndpoint` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfApiEndpoint` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfOrg

CF org

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/org` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfOrg` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfServiceInstance

Parameter of ServiceInstance Name to delete CloudFoundry Service

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


#### cfServiceKeyName

Parameter of CloudFoundry Service Key to be created

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - `cloudFoundry/serviceKey`<br />- `cloudFoundry/serviceKeyName`<br />- `cfServiceKey` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfServiceKeyName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfSpace

CF Space

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/space` |
| Type | `string` |
| Mandatory | no |
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


#### host

Specifies the host address of the SAP SAP BTP, ABAP Environment system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_host` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### idp

BTP Identity Provider

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/idp` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_idp` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### password

Password for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0763

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_password` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`abapCredentialsId`](#abapcredentialsid)<br />&nbsp;&nbsp;reference to: `password`<br /> |


#### patchIfExisting

In case an configuration under the given name already exists in the system. Should the step update/patch the existing ATC Systm Configuration from the provided ATC System Configuration file?

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `true` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


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


#### serviceBindingName

BTP service binding name

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/bindingName` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_serviceBindingName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### serviceInstanceName

BTP service instance name

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/instanceName` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_serviceInstanceName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### subaccount

BTP Subaccount name

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/subaccount` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_subaccount` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### subdomain

BTP Global Account subdomain

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/subdomain` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_subdomain` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### url

BTP CLI API endpoint

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `btp/url` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_url` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### username

User for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0763

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_username` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`abapCredentialsId`](#abapcredentialsid)<br />&nbsp;&nbsp;reference to: `username`<br /> |


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


#### abapCredentialsId

Jenkins credentials ID containing user and password to authenticate to the SAP BTP, ABAP Environment system or the Cloud Foundry API

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cfCredentialsId` |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Examples

### Configuration in the config.yml

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapEnvironmentPushATCSystemConfig script: this
```

If you want to provide the host and credentials of the Communication Arrangement directly, the configuration could look as follows:

```yaml
steps:
  abapEnvironmentPushATCSystemConfig:
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
    atcSystemConfigFilePath: 'atcSystemConfig.json',
```

To trigger a create/update ATC System Configuration step an ATC System configuration file `atcSystemConfig.json` will be needed. Check section 'ATC System Configuration file example' for more information.

### Create/Update an ATC System Configuration via Cloud Foundry Service Key example in Jenkinsfile

The following example triggers a Create/Update operation on an ATC System Configuration via reading the Service Key of an ABAP instance in Cloud Foundry.

You can store the credentials in Jenkins and use the cfCredentialsId parameter to authenticate to Cloud Foundry.
The username and password to authenticate to ABAP system will then be read from the Cloud Foundry service key that is bound to the ABAP instance.

This can be done accordingly:

```groovy
abapEnvironmentPushATCSystemConfig(
    cfApiEndpoint : 'https://test.server.com',
    cfOrg : 'cfOrg',
    cfSpace: 'cfSpace',
    cfServiceInstance: 'myServiceInstance',
    cfServiceKeyName: 'myServiceKey',
    abapCredentialsId: 'cfCredentialsId',
    atcSystemConfigFilePath: 'atcSystemConfig.json',
    script: this,
)
```

### Create/Update an ATC System Configuration via direct ABAP endpoint configuration in Jenkinsfile

This example triggers a create/update operation on an ATC System Configuration run directly on the ABAP endpoint.

In order to trigger the create/update operation on an ATC System Configuration you have to pass the username and password for authentication to the ABAP endpoint via parameters as well as the ABAP endpoint/host. You can store the credentials in Jenkins and use the abapCredentialsId parameter to authenticate to the ABAP endpoint/host.

This must be configured as following:

```groovy
abapEnvironmentPushATCSystemConfig(
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
    atcSystemConfigFilePath: 'atcSystemConfig.json',
    script: this,
)
```

To create/update an ATC System Configuration a file `atcSystemConfig.json` will be needed. Check section 'ATC System configuration file example' for more information.

### ATC System configuration file example

The step always performs a check first, if an ATC System Configuration with the same name provided in the file `atcSystemConfig.json` with the attribute conf_name.
This file contains an JSON Representation of an ATC System Configuration. Some json file examples can be found below.

In case an ATC System Configuration with this name already exists, by default, the step would perform an update of this ATC System Configuration with the ATC System Configuration information provided in file `atcSystemConfig.json`.
If this is not desired, an update could be supressed by using the parameter patchIfExisting in the configuration yaml the following way:

```yaml
steps:
  abapEnvironmentPushATCSystemConfig:
    atcSystemConfigFilePath: atcSystemConfig.json,
    patchIfExisting: false,
```

In this case the step skips further processing after existence check and returns with a Warning.

Providing a specifc System configuration file `atcSystemConfig.json` is mandatory.

The following section contains an example of an `atcSystemConfig.json` file.

This file must be stored in the same Git folder where the `Jenkinsfile` is stored to run the pipeline. This folder must be taken as a SCM in the Jenkins pipeline to run the pipeline.

See below an example for an `atcSystemConfig.json` file for creating/updating an ATC System Configuration with the name myATCSystemConfigurationName including a change of one priority.

```json
{
  "conf_name": "myATCSystemConfigurationName",
  "checkvariant": "SAP_CLOUD_PLATFORM_ATC_DEFAULT",
  "block_findings": "0",
  "inform_findings": "1",
  "_priorities": [
    {
      "test": "CL_CI_TEST_AMDP_HDB_MIGRATION",
      "message_id": "FAIL_ABAP",
      "priority": 2
    }
  ]
}
```

See below an example for an `atcSystemConfig.json` file for creating/updating an ATC System Configuration with the name myATCSystemConfigurationName.

```json
{
  "conf_name": "myATCSystemConfigurationName",
  "checkvariant": "SAP_CLOUD_PLATFORM_ATC_DEFAULT",
  "block_findings": "0",
  "inform_findings": "1"
}
```
