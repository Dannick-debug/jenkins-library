# abapEnvironmentBuild

Executes builds as defined with the build framework


## Description

Executes builds as defined with the build framework. Transaction overview /n/BUILD/OVERVIEW


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentBuild script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentBuild
        ```


## Outputs

| Output type | Details |
| ----------- | ------- |
| commonPipelineEnvironment | <ul><li>abap/buildValues</li></ul> |


## Prerequisites SAP BTP, ABAP environment

* A SAP BTP, ABAP environment system is available.
  * This can be created manually on Cloud Foundry.
  * In a pipeline, you can do this, for example, with the step [cloudFoundryCreateService](https://sap.github.io/jenkins-library/steps/cloudFoundryCreateService/).
* Communication Scenario [“SAP BTP, ABAP Environment - Software Assembly Integration (SAP_COM_0582)“](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/26b8df5435c649aa8ea7b3688ad5bb0a.html) is setup for this system.
  * E.g. a [Communication User](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0377adea0401467f939827242c1f4014.html), a [Communication System](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/1bfe32ae08074b7186e375ab425fb114.html) and a [Communication Arrangement](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/a0771f6765f54e1c8193ad8582a32edb.html) are configured.
  * This can be done manually through the respective applications on the SAP BTP, ABAP environment system,
  * or through creating a service key for the system on cloud foundry with the parameters {“scenario_id”: “SAP_COM_0582", “type”: “basic”}.
  * In a pipeline, you can do this, for example, with the step [cloudFoundryCreateServiceKey](https://sap.github.io/jenkins-library/steps/cloudFoundryCreateServiceKey/).
* You have following options to provide the ABAP endpoint configuration:
  * The host and credentials the SAP BTP, ABAP environment system itself. The credentials must be configured for the Communication Scenario SAP_COM_0582.
  * The Cloud Foundry parameters (API endpoint, organization, space), credentials, the service instance for the ABAP service and the service key for the Communication Scenario SAP_COM_0582.
  * Only provide one of those options with the respective credentials. If all values are provided, the direct communication (via host) has priority.

## Prerequisites On Premise

* You need to specify the host and credentials to your system
* A certificate for the system needs to be stored in .pipeline/trustStore and the name of the certificate needs to be handed over via the configuration

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [password](#password) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [phase](#phase) | **yes** |  |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [abapSourceClient](#abapsourceclient) | no |  |
| [addonDescriptor](#addondescriptor) | no |  |
| [certificateNames](#certificatenames) | no |  |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfOrg](#cforg) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServiceKeyName](#cfservicekeyname) | no |  |
| [cfSpace](#cfspace) | no |  |
| [conditionOnAddonDescriptor](#conditiononaddondescriptor) | no |  |
| [cpeValues](#cpevalues) | no |  |
| [downloadAllResultFiles](#downloadallresultfiles) | no |  |
| [downloadResultFilenames](#downloadresultfilenames) | no |  |
| [filenamePrefixForDownload](#filenameprefixfordownload) | no |  |
| [host](#host) | no |  |
| [maxRuntimeInMinutes](#maxruntimeinminutes) | no |  |
| [pollingIntervalInSeconds](#pollingintervalinseconds) | no |  |
| [publishAllDownloadedResultFiles](#publishalldownloadedresultfiles) | no |  |
| [publishResultFilenames](#publishresultfilenames) | no |  |
| [stopOnFirstError](#stoponfirsterror) | no |  |
| [subDirectoryForDownload](#subdirectoryfordownload) | no |  |
| [treatWarningsAsError](#treatwarningsaserror) | no |  |
| [useFieldsOfAddonDescriptor](#usefieldsofaddondescriptor) | no |  |
| [values](#values) | no |  |
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

#### abapSourceClient

Specifies the client of the SAP BTP ABAP Environment system, use only in combination with host

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_abapSourceClient` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### addonDescriptor

Structure in the commonPipelineEnvironment containing information about the Product Version and corresponding Software Component Versions

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_addonDescriptor` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | _commonPipelineEnvironment_:<br />&nbsp;&nbsp;reference to: `abap/addonDescriptor`<br /> |


#### certificateNames

file names of trusted (self-signed) server certificates - need to be stored in .pipeline/trustStore

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `[]string` |
| Mandatory | no |
| Default | `$PIPER_certificateNames` (if set) |
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
| Default | `$PIPER_cfApiEndpoint` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfOrg

Cloud Foundry target organization

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

Cloud Foundry Service Instance

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

Cloud Foundry Service Key

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

Cloud Foundry target space

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


#### conditionOnAddonDescriptor

normally if useFieldsOfAddonDescriptor is not initial, a build is triggered for each repository in the addonDescriptor. This can be changed by posing conditions. Please enter in the format '[{"field":"Status","operator":"==","value":"P"}]'

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_conditionOnAddonDescriptor` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
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


#### cpeValues

Values taken from the previous step, if a value was also specified in the config file, the value from cpe will be discarded

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cpeValues` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9744; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
| Resource references | _commonPipelineEnvironment_:<br />&nbsp;&nbsp;reference to: `abap/buildValues`<br /> |


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


#### downloadAllResultFiles

If true, all build artefacts are downloaded

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### downloadResultFilenames

Only the specified files are downloaded. If downloadAllResultFiles is true, this parameter is ignored

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `[]string` |
| Mandatory | no |
| Default | `$PIPER_downloadResultFilenames` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### filenamePrefixForDownload

Filename prefix for the downloaded files, {buildID} and {taskID} can be used and will be resolved accordingly

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_filenamePrefixForDownload` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### host

Specifies the host address of the SAP BTP ABAP Environment system

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


#### maxRuntimeInMinutes

maximal runtime of the step in minutes

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `int` |
| Mandatory | no |
| Default | `360` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### password

Password

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_password` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### phase

Phase as specified in the build script in the backend system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_phase` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### pollingIntervalInSeconds

wait time in seconds till next status request in the backend system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `int` |
| Mandatory | no |
| Default | `60` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### publishAllDownloadedResultFiles

If true, it publishes all downloaded files

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### publishResultFilenames

Only the specified files get published, in case the file was not downloaded before an error occures

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `[]string` |
| Mandatory | no |
| Default | `$PIPER_publishResultFilenames` (if set) |
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


#### stopOnFirstError

If false, it does not stop if an error occured for one repository in the addonDescriptor, but continues with the next repository. However the step is marked as failed in the end if an error occured.

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### subDirectoryForDownload

Target directory to store the downloaded files, {buildID} and {taskID} can be used and will be resolved accordingly

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_subDirectoryForDownload` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### treatWarningsAsError

If a warrning occures, the step will be set to unstable

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `bool` |
| Mandatory | no |
| Default | `false` |
| Possible values | - `true`<br />- `false` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### useFieldsOfAddonDescriptor

use fields of the addonDescriptor in the cpe as input values. Please enter in the format '[{"use":"Name","renameTo":"SWC"}]'

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_useFieldsOfAddonDescriptor` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### username

User

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_username` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### values

Input values for the build framework, please enter in the format '[{"value_id":"Id1","value":"value1"},{"value_id":"Id2","value":"value2"}]'

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_values` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


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

Jenkins credentials ID containing user and password to authenticate to the Cloud Platform ABAP Environment system or the Cloud Foundry API

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - `cfCredentialsId`<br />- `credentialsId` |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Examples

### Configuration in the config.yml

If you want to use this step several time in one pipeline with different phases, the steps have to be put in different stages as it is not allowed to run the same step repeatedly in one stage.

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
stage('MyPhase') {
            steps {
                abapEnvironmentBuild script: this
            }
        }
```

If you want to provide the host and credentials of the Communication Arrangement directly or you want to run in on premise, the configuration could look as follows:

```yaml
stages:
  MyPhase:
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
```

Or by authenticating against Cloud Foundry and reading the Service Key details from there:

```yaml
stages:
  MyPhase:
    abapCredentialsId: 'cfCredentialsId',
    cfApiEndpoint : 'https://test.server.com',
    cfOrg : 'cfOrg',
    cfSpace: 'cfSpace',
    cfServiceInstance: 'myServiceInstance',
    cfServiceKeyName: 'myServiceKey',
```

One possible complete config example. Please note that the values are handed over as a string, which has inside a json structure:

```yaml
stages:
  MyPhase:
    abapCredentialsId: 'abapCredentialsId'
    host: 'https://myABAPendpoint.com'
    certificateNames: ['myCert.cer']
    phase: 'MyPhase'
    values: '[{"value_id":"ID1","value":"Value1"},{"value_id":"ID2","value":"Value2"}]'
    downloadResultFilenames: ['File1','File2']
    publishResultFilenames: ['File2']
    subDirectoryForDownload: 'MyDir'
    filenamePrefixForDownload: 'MyPrefix'
    treatWarningsAsError: true
    maxRuntimeInMinutes: 360
    pollingIntervallInSeconds: 15
```
