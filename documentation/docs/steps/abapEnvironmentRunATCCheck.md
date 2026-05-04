# abapEnvironmentRunATCCheck

Runs an ATC Check


## Description

This step is for triggering an [ATC](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/d8cec788fc104ff9ad9c3757b4dd13d4.html) test run on an SAP BTP ABAP Environment system.
Please provide either of the following options:

* The host and credentials the Cloud Platform ABAP Environment system itself. The credentials must be configured for the Communication Scenario [SAP_COM_0901](https://help.sap.com/docs/BTP/65de2977205c403bbc107264b8eccf4b/d8cec788fc104ff9ad9c3757b4dd13d4.html).
* The Cloud Foundry parameters (API endpoint, organization, space), credentials, the service instance for the ABAP service and the service key for the Communication Scenario SAP_COM_0901.
* Only provide one of those options with the respective credentials. If all values are provided, the direct communication (via host) has priority.

Regardless of the option you chose, please make sure to provide the configuration the object set (e.g. with Software Components and Packages) that you want to be checked analog to the examples listed on this page.


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapEnvironmentRunATCCheck script: this
        ```

    === "Command Line"

        ```sh
        piper abapEnvironmentRunATCCheck
        ```



!!! Currently the Object Set configuration is limited to the usage of Multi Property Sets. Please note that other sets besides the Multi Property Set will not be included in the ATC runs. You can see an example of the Multi Property Sets with all configurable properties. However, we strongly reccommend to only specify packages and software components like in the first two examples of the section `ATC config file example`.

## Prerequisites

* A SAP BTP, ABAP environment system is available. On this system, a [Communication User](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/0377adea0401467f939827242c1f4014.html), a [Communication System](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/1bfe32ae08074b7186e375ab425fb114.html) and a [Communication Arrangement](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/a0771f6765f54e1c8193ad8582a32edb.html) is setup for the Communication Scenario “ABAP Test Cockpit - Test Integration (SAP_COM_0901)“. This can be done manually through the respective applications on the SAP BTP, ABAP environment system or through creating a service key for the system on Cloud Foundry with the parameters {“scenario_id”: “SAP_COM_0901", “type”: “basic”}. In a pipeline, you can do this with the step [cloudFoundryCreateServiceKey](https://sap.github.io/jenkins-library/steps/cloudFoundryCreateServiceKey/).
* You can either provide the ABAP endpoint configuration to directly trigger an ATC run on the ABAP system or optionally provide the Cloud Foundry parameters with your credentials to read a Service Key of a SAP BTP, ABAP environment system in Cloud Foundry that contains all the details of the ABAP endpoint to trigger an ATC run.
* Regardless if you chose an ABAP endpoint directly or reading a Cloud Foundry Service Key, you have to provide the configuration of the packages and software components you want to be checked in an ATC run in a .yml or .yaml file. This file must be stored in the same folder as the Jenkinsfile defining the pipeline.
* The software components and/or packages you want to be checked must be present in the configured system in order to run the check. Please make sure that you have created or pulled the respective software components and/or Packages in the SAP BTP, ABAP environment system.

Examples will be listed below.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [password](#password) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [username](#username) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapCredentialsId`](#abapcredentialsid)) |
| [atcConfig](#atcconfig) | no |  |
| [atcResultsFileName](#atcresultsfilename) | no |  |
| [certificateNames](#certificatenames) | no |  |
| [cfApiEndpoint](#cfapiendpoint) | no |  |
| [cfOrg](#cforg) | no |  |
| [cfServiceInstance](#cfserviceinstance) | no |  |
| [cfServiceKeyName](#cfservicekeyname) | no |  |
| [cfSpace](#cfspace) | no |  |
| [failOnSeverity](#failonseverity) | no |  |
| [generateHTML](#generatehtml) | no |  |
| [host](#host) | no |  |
| [idp](#idp) | no |  |
| [repositories](#repositories) | no |  |
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

#### atcConfig

Path to a YAML configuration file for the object set to be checked during ATC run

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_atcConfig` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### atcResultsFileName

Specifies output file name for the results from the ATC run. This file name will also be used for generating the HTML file

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `ATCResults.xml` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
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


#### failOnSeverity

Specifies the severity level, for which the ATC step should fail if at least one message with this severity (or "higher") level is returned by the ATC Check Run (possible values - error, warning, info). Initial value is default behavior and ATC findings of any severity do not fail the step

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_failOnSeverity` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### generateHTML

Specifies whether the ATC results should also be generated as an HTML document

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


#### password

Password for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0901

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

User for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0901

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

Jenkins credentials ID containing user and password to authenticate to the Cloud Platform ABAP Environment system or the Cloud Foundry API

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
abapEnvironmentRunATCCheck script: this
```

If you want to provide the host and credentials of the Communication Arrangement directly, the configuration could look as follows:

```yaml
steps:
  abapEnvironmentRunATCCheck:
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
    atcConfig: 'atcconfig.yml',
```

### ATC run via Cloud Foundry Service Key example in Jenkinsfile

The following example triggers an ATC run via reading the Service Key of an ABAP instance in Cloud Foundry.

You can store the credentials in Jenkins and use the cfCredentialsId parameter to authenticate to Cloud Foundry.
The username and password to authenticate to ABAP system will then be read from the Cloud Foundry service key that is bound to the ABAP instance.

This can be done accordingly:

```groovy
abapEnvironmentRunATCCheck(
    cfApiEndpoint : 'https://test.server.com',
    cfOrg : 'cfOrg',
    cfSpace: 'cfSpace',
    cfServiceInstance: 'myServiceInstance',
    cfServiceKeyName: 'myServiceKey',
    abapCredentialsId: 'cfCredentialsId',
    atcConfig: 'atcconfig.yml',
    script: this,
)
```

To trigger the ATC run an ATC config file `atcconfig.yml` will be needed. Check section 'ATC config file example' for more information.

### ATC run via direct ABAP endpoint configuration in Jenkinsfile

This  example triggers an ATC run directly on the ABAP endpoint.

In order to trigger the ATC run you have to pass the username and password for authentication to the ABAP endpoint via parameters as well as the ABAP endpoint/host. You can store the credentials in Jenkins and use the abapCredentialsId parameter to authenticate to the ABAP endpoint/host.

This must be configured as following:

```groovy
abapEnvironmentRunATCCheck(
    abapCredentialsId: 'abapCredentialsId',
    host: 'https://myABAPendpoint.com',
    atcConfig: 'atcconfig.yml',
    script: this,
)
```

To trigger the ATC run an ATC config file `atcconfig.yml` will be needed. Check section 'ATC config file example' for more information.

### ATC config file example

Providing a specifc ATC configuration is optional. If you are using a `repositories.yml` file for the `Clone` stage of the ABAP environment pipeline, a default ATC configuration will be derived if no explicit ATC configuration is available.

The following section contains an example of an `atcconfig.yml` file.
This file must be stored in the same Git folder where the `Jenkinsfile` is stored to run the pipeline. This folder must be taken as a SCM in the Jenkins pipeline to run the pipeline.

You can specify a list of packages and/or software components to be checked. This must be in the same format as below example for a `atcconfig.yml` file.
In case subpackages shall be included in the checks you can use packagetrees.
Please note that if you chose to provide both packages and software components to be checked with the `atcconfig.yml` file, the set of packages and the set of software components will be combinend by the API using a logical AND operation.
Therefore, we advise to specify either the software components or packages.
Additionally, if you don't specify a dedicated ATC check variant to be used, the `ABAP_CLOUD_DEVELOPMENT_DEFAULT` variant will be used as default. For more information on how to configure a check variant for an ATC run please check the last example on this page.

See below example for an `atcconfig.yml` file with both packages and software components to be checked:

```yaml
objectset:
  softwarecomponents:
    - name: TestComponent
    - name: TestComponent2  
  packages:
    - name: TestPackage
  packagetrees:
    - name: TestPackageWithSubpackages
```

The following example of an `atcconfig.yml` file that only contains packages and packagetrees to be checked:

```yaml
objectset:
  packages:
    - name: TestPackage
  packagetrees:
    - name: TestPackageWithSubpackages
```

The following example of an `atcconfig.yml` file that only contains software components to be checked:

```yaml
objectset:
  softwarecomponents:
    - name: TestComponent
    - name: TestComponent2
```

The following is an example of an `atcconfig.yml` file that supports the check variant and configuration ATC options and containing the software components `TestComponent` and `TestComponent2` as Objectset.

```yaml
checkvariant: "TestCheckVariant"
configuration: "TestConfiguration"
objectset:
  softwarecomponents:
    - name: TestComponent
    - name: TestComponent2
```

The following example of an `atcconfig.yml` file contains all possible properties of the Multi Property Set that can be used. Please take note that this is not the reccommended approach. If you want to check packages or software components please use the two above examples. The usage of the Multi Property Set is only reccommended for ATC runs that require these rules for the test execution. There is no official documentation on the usage of the Multi Property Set.

```yaml
checkvariant: "TestCheckVariant"
configuration: "TestConfiguration"
objectset:
  type: multiPropertySet
  multipropertyset:
    owners:
      - name: demoOwner
    softwarecomponents:
      - name: demoSoftwareComponent
    versions:
      - value: ACTIVE
    packages:
      - name: demoPackage
    packagetrees:
      - name: TestPackageWithSubpackages
    objectnamepatterns:
      - value: 'ZCL_*'
    languages:
      - value: EN
    sourcesystems:
      - name: H01
    objecttypes:
      - name: CLAS
    objecttypegroups:
      - name: CLAS
    releasestates:
      - value: RELEASED
    applicationcomponents:
      - name: demoApplicationComponent
    transportlayers:
      - name: H01
```
