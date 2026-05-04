# abapEnvironmentAssemblePackages

Assembly of installation, support package or patch in SAP BTP ABAP Environment system


## Description

This step runs the assembly of a list of provided [installations, support packages or patches](https://help.sap.com/viewer/9043aa5d2f834ad385e1cdfdadc06b6f/LATEST/en-US/9a81f55473568c77e10000000a174cb4.html) in SAP BTP ABAP Environment system and saves the corresponding [SAR archive](https://launchpad.support.sap.com/#/notes/212876) to the filesystem.
<br />
Among others a semantic version [API snapshot](https://help.sap.com/docs/btp/sap-business-technology-platform/creating-api-snapshots?version=Cloud) will be
searched and marked as check-relevant in the assembly system, ATC checks for [consistency of software component dependencies](https://help.sap.com/docs/abap-cloud/abap-development-tools-user-guide/software-component-relations) as well as [API compatibility](https://help.sap.com/docs/abap-cloud/abap-development-tools-user-guide/checking-compatibility-of-released-apis) will run, and a new semantic version API snapshot will be created as well.
<br />
Refer to [Software Assembly Integration (SAP_COM_0582)](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/26b8df5435c649aa8ea7b3688ad5bb0a.html).


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentAssemblePackages script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentAssemblePackages
        ```


## Outputs

| Output type | Details |
| ----------- | ------- |
| commonPipelineEnvironment | <ul><li>abap/addonDescriptor</li></ul> |


### Artifacts

- package logs ({packagename}.zip)
    This archive contains all relevant transport logs per assembled package which might be needed for detailed analysis in case of support requests or for audit purpose. For productive builds it might be advisable to store this file as well as the overall pipeline run logs in a revision proof manner. For every assembled package an respective zip archive with its related logs are created and archived as artifact.

## Prerequisites

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
* The step needs information about the packages which should be assembled present in the CommonPipelineEnvironment.
  * For each repository/component version it needs the name of the repository, the version, splevel, patchlevel, namespace, packagename, package type, the status of the package, and optional the predecessor commit id.
  * To upload this information to the CommonPipelineEnvironment run prior to this step the steps:
    * [abapAddonAssemblyKitCheckCVs](https://sap.github.io/jenkins-library/steps/abapAddonAssemblyKitCheckCVs/),
    * [abapAddonAssemblyKitReserveNextPackages](https://sap.github.io/jenkins-library/steps/abapAddonAssemblyKitCheckPV/).
  * If one of the package is already in status released, the assembly for this package will not be executed.
* The Software Components for which packages are to be assembled need to be present in the system.
  * This can be done manually through the respective applications on the SAP BTP, ABAP environment system.
  * In a pipeline, you can do this, for example, with the step [abapEnvironmentPullGitRepo](https://sap.github.io/jenkins-library/steps/abapEnvironmentPullGitRepo/).
  * In case multiple software component are used, any dependencies between the components need to be defined in [Software Component Relations](https://help.sap.com/docs/abap-cloud/abap-development-tools-user-guide/software-component-relations).
* The packages to be assembled need to be reserved in AAKaaS and the corresponding information needs to be present in CommonPipelineEnvironment. To do so run step [abapAddonAssemblyKitReserveNextPackages](https://sap.github.io/jenkins-library/steps/abapAddonAssemblyKitReserveNextPackages/) prior this step.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [addonDescriptor](#addondescriptor) | **yes** |  |
| [password](#password) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [alternativePhaseName](#alternativephasename) | no |  |
| [certificateNames](#certificatenames) | no |  |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfOrg](#cforg) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServiceKeyName](#cfservicekeyname) | no |  |
| [cfSpace](#cfspace) | no |  |
| [host](#host) | no |  |
| [maxRuntimeInMinutes](#maxruntimeinminutes) | no |  |
| [pollIntervalsInMilliseconds](#pollintervalsinmilliseconds) | no |  |
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

#### addonDescriptor

Structure in the commonPipelineEnvironment containing information about the Product Version and corresponding Software Component Versions

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_addonDescriptor` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | _commonPipelineEnvironment_:<br />&nbsp;&nbsp;reference to: `abap/addonDescriptor`<br /> |


#### alternativePhaseName

overrides default phase name

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_alternativePhaseName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


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

Password for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0582

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


#### pollIntervalsInMilliseconds

wait time in milliseconds till next status request in the backend system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `int` |
| Mandatory | no |
| Default | `60000` |
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


#### username

User for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0582

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

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapEnvironmentAssemblePackages script: this
```

If you want to provide the host and credentials of the Communication Arrangement directly, the configuration could look as follows:

```yaml
steps:
  abapEnvironmentAssemblePackages:
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
```

Or by authenticating against Cloud Foundry and reading the Service Key details from there:

```yaml
steps:
  abapEnvironmentAssemblePackages:
    abapCredentialsId: 'cfCredentialsId',
    cfApiEndpoint : 'https://test.server.com',
    cfOrg : 'cfOrg',
    cfSpace: 'cfSpace',
    cfServiceInstance: 'myServiceInstance',
    cfServiceKeyName: 'myServiceKey',
```

### Input via the CommonPipelineEnvironment

```json
{"addonProduct":"",
"addonVersion":"",
"addonVersionAAK":"",
"addonUniqueID":"",
"customerID":"",
"AddonSpsLevel":"",
"AddonPatchLevel":"",
"TargetVectorID":"",
"repositories":[
  {
    "name":"/DMO/REPO_A",
    "tag":"",
    "branch":"",
    "version":"",
    "versionAAK":"0001",
    "PackageName":"SAPK001001REPOA",
    "PackageType":"CPK",
    "SpLevel":"0000",
    "PatchLevel":"0001",
    "PredecessorCommitID":"cbb834e9e03cde177d2f109a6676901972983fbc",
    "Status":"P",
    "Namespace":"/DMO/",
    "SarXMLFilePath":""
  },
  {
    "name":"/DMO/REPO_B",
    "tag":"",
    "branch":"",
    "version":"",
    "versionAAK":"0002",
    "PackageName":"SAPK002001REPOB",
    "PackageType":"CPK",
    "SpLevel":"0001",
    "PatchLevel":"0001",
    "PredecessorCommitID":"2f7d43923c041a07a76c8adc859c737ad772ef26",
    "Status":"P",
    "Namespace":"/DMO/",
    "SarXMLFilePath":""
  }
]}
```
