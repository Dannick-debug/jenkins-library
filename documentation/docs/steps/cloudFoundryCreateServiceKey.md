# cloudFoundryCreateServiceKey

cloudFoundryCreateServiceKey


## Description

Create CloudFoundryServiceKey


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        cloudFoundryCreateServiceKey script: this
        ```

    === "Command Line"

        ```sh
        piper cloudFoundryCreateServiceKey
        ```



## Prerequisites

* This step is for creating a Service Key for an existing Service in Cloud Foundry.
* Cloud Foundry API endpoint, Organization, Space, user and Service Instance are available
* Credentials have been configured in Jenkins with a dedicated Id
* Additionally you can set the optional `serviceKeyConfig` flag to configure the Service Key creation with your respective JSON configuration. The JSON configuration can either be a JSON or the path a dedicated JSON configuration file containing the JSON configuration. If you chose a dedicated config file, it must be stored in a file that must be referenced in the `serviceKeyConfigFile` flag. You must store the file in the same folder as your `Jenkinsfile` that starts the Pipeline in order for the Pipeline to be able to find the file. Most favourable SCM is Git.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [cfApiEndpoint](#cfapiendpoint) | **yes** |  |
| [cfOrg](#cforg) | **yes** |  |
| [cfServiceInstance](#cfserviceinstance) | **yes** |  |
| [cfServiceKeyName](#cfservicekeyname) | **yes** |  |
| [cfSpace](#cfspace) | **yes** |  |
| [password](#password) | **(yes)** |  ![Vault](https://img.shields.io/badge/-Vault-lightgrey) ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV, Vault or Jenkins credentials ([`cfCredentialsId`](#cfcredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** |  ![Vault](https://img.shields.io/badge/-Vault-lightgrey) ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV, Vault or Jenkins credentials ([`cfCredentialsId`](#cfcredentialsid)) |
| [cfAsync](#cfasync) | no |  |
| [cfServiceKeyConfig](#cfservicekeyconfig) | no |  |
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

#### cfApiEndpoint

Cloud Foundry API endpoint

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/apiEndpoint` |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_cfApiEndpoint` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfAsync

Decides if the service key creation runs asynchronously

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


#### cfOrg

CF org

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


#### cfServiceInstance

Parameter for CloudFoundry Service Instance Name

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/serviceInstance` |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_cfServiceInstance` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfServiceKeyConfig

Path to JSON config file path or JSON in-line string for Cloud Foundry Service Key creation

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | `cloudFoundry/serviceKeyConfig` |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_cfServiceKeyConfig` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### cfServiceKeyName

Parameter for Service Key name for CloudFoundry Service Key to be created

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - `cloudFoundry/serviceKey`<br />- `cloudFoundry/serviceKeyName`<br />- `cfServiceKey` |
| Type | `string` |
| Mandatory | **yes** |
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


#### password

User Password for CF User

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
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Examples

The following examples will create a Service Key named "myServiceKey" for the Service Instance "myServiceInstance" in the provided Cloud Foundry Organization and Space. For the Service Key creation in these example, the serviceKeyConfig parameter is used. It will show the different ways of passing the JSON configuration, either via a string or the path to a file containing the JSON configuration.
If you dont want to use a special configuration simply remove the parameter since it is optional.

### Create Service Key with JSON config file in Jenkinsfile

This example covers the parameters for a Jenkinsfile when using the cloudFoundryCreateServiceKey step. It uses a `serviceKeaConfig.json` file with valid JSON objects for creating a Cloud Foundry Service Key.

```groovy
cloudFoundryCreateServiceKey(
  cfApiEndpoint: 'https://test.server.com',
  cfCredentialsId: 'cfCredentialsId',
  cfOrg: 'cfOrg',
  cfSpace: 'cfSpace',
  cfServiceInstance: 'myServiceInstance',
  cfServiceKeyName: 'myServiceKey',
  cfServiceKeyConfig: 'serviceKeyConfig.json',
  script: this,
)
```

The JSON config file, e.g. `serviceKeyConfig.json` can look like this:

```json
{
  "example":"value",
  "example":"value"
}
```

### Create Service Key with JSON string in Jenkinsfile

The following example covers the creation of a Cloud Foundry Service Key in a Jenkinsfile with using a JSON string as a config for the Service Key creation. If you use a Jenkinsfile for passing the parameter values you need to escape the double quotes in the JSON config string.

```groovy
cloudFoundryCreateServiceKey(
  cfApiEndpoint: 'https://test.server.com',
  cfCredentialsId: 'cfCredentialsId',
  cfOrg: 'cfOrg',
  cfSpace: 'cfSpace',
  cfServiceInstance: 'myServiceInstance',
  cfServiceKeyName: 'myServiceKey',
  cfServiceKeyConfig: '{\"example\":\"value\",\"example\":\"value\"}',
  script: this,
)
```

### Create Service Key with JSON string as parameter in .pipeline/config.yml file

If you chose to provide a `config.yml` file you can provide the parameters including the values in this file. You only need to set the script parameter when calling the step:

```groovy
cloudFoundryCreateServiceKey(
  script: this,
)
```

The `.pipeline/config.yml` has to contain the following parameters accordingly:

```yaml
steps:
    cloudFoundryCreateServiceKey:
        cfApiEndpoint: 'https://test.server.com'
        cfOrg: 'testOrg'
        cfSpace: 'testSpace'
        cfServiceInstance: 'testInstance'
        cfServiceKeyName: 'myServiceKey'
        cfServiceKeyConfig: '{"example":"value","example":"value"}'
        cfCredentialsId: 'cfCredentialsId'
```

When using a `.pipeline/config.yml` file you don't need to escape the double quotes in the JSON config string.
You can also pass the path to a JSON config file in the `cfServiceKeyConfig` parameter. Example: `cfServiceKeyConfig: 'serviceKeyconfig.json'`
