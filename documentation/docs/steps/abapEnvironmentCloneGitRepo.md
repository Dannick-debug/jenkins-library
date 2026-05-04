# abapEnvironmentCloneGitRepo

Clones a git repository to a SAP BTP ABAP Environment system


## Description

Clones a git repository (Software Component) to a SAP BTP ABAP Environment system. If the repository is already cloned, the step will checkout the configured branch and pull the specified commit, instead.
Please provide either of the following options:

* The host and credentials the BTP ABAP Environment system itself. The credentials must be configured for the Communication Scenario [SAP_COM_0948](https://help.sap.com/docs/sap-btp-abap-environment/abap-environment/api-for-managing-software-components-61f4d47af1394b1c8ad684b71d3ad6a0?locale=en-US).
* The Cloud Foundry parameters (API endpoint, organization, space), credentials, the service instance for the ABAP service and the service key for the Communication Scenario SAP_COM_0948.
* Only provide one of those options with the respective credentials. If all values are provided, the direct communication (via host) has priority.


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentCloneGitRepo script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentCloneGitRepo
        ```



## Prerequisites

A SAP BTP, ABAP environment system is available.
On this system, a [Communication User](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0377adea0401467f939827242c1f4014.html), a [Communication System](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/1bfe32ae08074b7186e375ab425fb114.html) and a [Communication Arrangement](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/a0771f6765f54e1c8193ad8582a32edb.html) is setup for the Communication Scenario "Software Component Management Integration (SAP_COM_0948)". This can be done manually through the respective applications on the SAP BTP, ABAP environment system or through creating a service key for the system on Cloud Foundry with the parameters {"scenario_id": "SAP_COM_0948", "type": "basic"}. In a pipeline, you can do this with the step [cloudFoundryCreateServiceKey](https://sap.github.io/jenkins-library/steps/cloudFoundryCreateServiceKey/).

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [byogPassword](#byogpassword) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`byogCredentialsId`](#byogcredentialsid)) |
| [byogUsername](#byogusername) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`byogCredentialsId`](#byogcredentialsid)) |
| [password](#password) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [branchName](#branchname) | no |  |
| [byogAuthMethod](#byogauthmethod) | no |  |
| [certificateNames](#certificatenames) | no |  |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfOrg](#cforg) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServiceKeyName](#cfservicekeyname) | no |  |
| [cfSpace](#cfspace) | no |  |
| [host](#host) | no |  |
| [idp](#idp) | no |  |
| [logOutput](#logoutput) | no |  |
| [repositories](#repositories) | no |  |
| [repositoryName](#repositoryname) | no |  |
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

#### branchName

Specifies a branch of a repository (Software Components) on the SAP BTP ABAP Environment system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_branchName` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### byogAuthMethod

Specifies which authentication method is used for bring your own git (BYOG) repositories

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `TOKEN` |
| Possible values | - `TOKEN`<br />- `BASIC` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### byogPassword

Password for bring your own git (BYOG) authentication

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_byogPassword` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`byogCredentialsId`](#byogcredentialsid)<br />&nbsp;&nbsp;reference to: `password`<br /> |


#### byogUsername

Username for bring your own git (BYOG) authentication

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_byogUsername` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`byogCredentialsId`](#byogcredentialsid)<br />&nbsp;&nbsp;reference to: `username`<br /> |


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

Cloud Foundry API Enpoint

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


#### logOutput

Specifies how the clone logs from the Manage Software Components App are displayed or saved

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `STANDARD` |
| Possible values | - `ZIP`<br />- `STANDARD` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### password

Password for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0948

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


#### repositories

Specifies a YAML file containing the repositories configuration

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_repositories` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### repositoryName

Specifies a repository (Software Components) on the SAP BTP ABAP Environment system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_repositoryName` (if set) |
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

User for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0948

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

Jenkins credentials ID containing user and password to authenticate to the BTP ABAP Environment system or the Cloud Foundry API

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - `cfCredentialsId`<br />- `credentialsId` |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |


#### byogCredentialsId

Jenkins credentials ID containing ByogUsername and ByogPassword to authenticate to a software component which is used in a BYOG scenario. (https://help.sap.com/docs/btp/sap-business-technology-platform/cloning-software-components-to-abap-environment-system-383ce2f9e2eb40f1b8ad538ddf79e656)

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Example: Configuration in the config.yml

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapEnvironmentCheckoutBranch script: this
```

If you want to provide the host and credentials of the Communication Arrangement directly, the configuration could look as follows:

```yaml
steps:
  abapEnvironmentCloneGitRepo:
    repositoryName: '/DMO/GIT_REPOSITORY'
    branchName: 'my-demo-branch'
    abapCredentialsId: 'abapCredentialsId'
    host: '1234-abcd-5678-efgh-ijk.abap.eu10.hana.ondemand.com'
```

Please note that the branchName parameter specifies the target branch you want to clone. Also keep in mind that the repositoryName parameter must define a single repository.

Another option is to read the host and credentials from the cloud foundry service key of the respective instance. Furthermore, if you want to clone multiple repositories, they can be specified in a configuration file.

With this approach the `config.yml` would look like this:

```yaml
steps:
  abapEnvironmentCloneGitRepo:
    repositories: 'repositories.yml'
    cfCredentialsId: 'cfCredentialsId'
    cfApiEndpoint: 'https://test.server.com'
    cfOrg: 'cfOrg'
    cfSpace: 'cfSpace'
    cfServiceInstance: 'cfServiceInstance'
    cfServiceKeyName: 'cfServiceKeyName'
```

and the configuration file `repositories.yml` would look like this:

```yaml
repositories:
  - name: '/DMO/REPO'
    branch: 'main'
  - name: '/DMO/REPO_COMMIT'
    branch: 'feature'
    commitID: 'cd87a3cac2bc946b7629580e58598c3db56a26f8'
  - name: '/DMO/REPO_TAG'
    branch: 'release'
    tag: 'myTag'
```

Using such a configuration file is the recommended approach. Please note that you need to use the YAML data structure as in the example above when using the `repositories.yml` config file.
If you want to clone a specific commit, either a `commitID` or a `tag` can be specified. If both are specified, the `tag` will be ignored.

## Example: Configuration in the Jenkinsfile

It is also possible to call the steps - including all parameters - directly in the Jenkinsfile.
In the first example, the host and the credentialsId of the Communication Arrangement are directly provided.

```groovy
abapEnvironmentCloneGitRepo (
  script: this,
  repositoryName: '/DMO/GIT_REPOSITORY',
  branchName: 'my-demo-branch',
  abapCredentialsId: 'abapCredentialsId',
  host: '1234-abcd-5678-efgh-ijk.abap.eu10.hana.ondemand.com'
)
```

In the second example, the host and credentialsId will be read from the provided cloud foundry service key of the specified service instance.

```groovy
abapEnvironmentCloneGitRepo (
  script: this,
  repositoryName: '/DMO/GIT_REPOSITORY',
  branchName: 'my-demo-branch'
  abapCredentialsId: 'cfCredentialsId',
  cfApiEndpoint: 'https://test.server.com',
  cfOrg: 'cfOrg',
  cfSpace: 'cfSpace',
  cfServiceInstance: 'cfServiceInstance',
  cfServiceKeyName: 'cfServiceKeyName'
)
```

## Example: Cloning a Bring Your Own Git (BYOG) repository

> Feature will be available in November 2024.

Since a ByoG repository is an external repository, you must be authenticated to clone it.
For this, the corresponding credentials must be stored in Jenkins as a username and password/token.

<strong> Store the credentials: </strong> <br>
A new credential with the type username and password must be stored.<br>
`Jenkins Dashboard > Manage Jenkins > Credentials` <br>
These credentials are used to clone the ByoG repository.
More information on configuring the credentials can be found [here](https://www.jenkins.io/doc/book/using/using-credentials/).

The config.yaml should look like this:

```yaml
steps:
  abapEnvironmentCloneGitRepo:
    repositories: 'repos.yaml'
    byogCredentialsId: 'byogCredentialsId'
    abapCredentialsId: 'abapCredentialsId'
    host: '1234-abcd-5678-efgh-ijk.abap.eu10.hana.ondemand.com'
```

`byogCredentialsId: 'byogCredentialsId'` is the reference to the defined credential in Jenkins. So take care that this matches with your setup.

After that, the ByoG repository that is to be cloned must be specified in the repos.yaml:

```yaml
repositories:
  - name: '/DMO/REPO_BYOG'
    branch: 'main'
```

After the pipeline has run through, the repository has been cloned.
