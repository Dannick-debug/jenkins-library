# abapAddonAssemblyKitCheck

This step calls AAKaaS to check the validity of the Addon Product Modelling.


## Description

[The Addon Product Modelling](https://www.project-piper.io/scenarios/abapEnvironmentAddons/#add-on-descriptor-file) is read from the <b>addonDescriptorFileName</b> (e.g. addon.yml).
A connection to AAKaaS (Addon Assembly Kit as a Service) is established and the Addon Product Modelling is transfered for detailed [checks](https://www.project-piper.io/scenarios/abapEnvironmentAddons/#versioning-rules).
The semantic versions are resolved and stored into the piper commonPipelineEnviroment for usage in subsequent pipeline steps.
<br />
For logon to AAKaaS you can either provide a credential with basic authorization (username and password) or two secret text credentials containing the technical s-users certificate (see note [2805811](https://me.sap.com/notes/2805811) for download) as base64 encoded string and the password to decrypt the file
<br />
For Terminology refer to the [Scenario Description](https://www.project-piper.io/scenarios/abapEnvironmentAddons/).


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapAddonAssemblyKitCheck script: this
        ```

    === "Command Line"

        ```sh
        piper abapAddonAssemblyKitCheck
        ```


## Outputs

| Output type | Details |
| ----------- | ------- |
| commonPipelineEnvironment | <ul><li>abap/addonDescriptor</li></ul> |


### Artifacts

- addonDescriptorFile (addon.yml)
    The addonDescriptorFile as specified in parameter addonDescriptorFileName is archived as artifact. This is done as this file is the main configuration and usually changed with every run. Thus it simplifies support if the corresponding configuration file is directly accessible in the pipeline.

## Prerequisites

* The credentials to access the AAKaaS (Technical Communication User) must be stored in the Jenkins Credential Store
* The step needs an addon.yml containing information about the Product Version and corresponding Software Component Versions/Repositories

A detailed description of all prerequisites of the scenario and how to configure them can be found in the [Scenario Description](https://www.project-piper.io/scenarios/abapEnvironmentAddons/).

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [abapAddonAssemblyKitCertificateFile](#abapaddonassemblykitcertificatefile) | no | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapAddonAssemblyKitCertificateFileCredentialsId`](#abapaddonassemblykitcertificatefilecredentialsid)) |
| [abapAddonAssemblyKitCertificatePass](#abapaddonassemblykitcertificatepass) | no | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`abapAddonAssemblyKitCertificatePassCredentialsId`](#abapaddonassemblykitcertificatepasscredentialsid)) |
| [abapAddonAssemblyKitEndpoint](#abapaddonassemblykitendpoint) | no |  |
| [abapAddonAssemblyKitOriginHash](#abapaddonassemblykitoriginhash) | no | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [addonDescriptor](#addondescriptor) | no |  |
| [addonDescriptorFileName](#addondescriptorfilename) | no |  |
| [password](#password) | no | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [username](#username) | no | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials |
| [verbose](#verbose) | no | activates debug output |

### Overview - Execution Environment

!!! note "Orchestrator-specific only"

    These parameters are relevant for orchestrator usage and not considered when using the command line option.

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |

### Details

#### abapAddonAssemblyKitCertificateFile

base64 encoded certificate pfx file (PKCS12 format) see note [2805811](https://me.sap.com/notes/2805811)

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_abapAddonAssemblyKitCertificateFile` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`abapAddonAssemblyKitCertificateFileCredentialsId`](#abapaddonassemblykitcertificatefilecredentialsid)<br />&nbsp;&nbsp;reference to: `abapAddonAssemblyKitCertificateFile`<br /> |


#### abapAddonAssemblyKitCertificatePass

password to decrypt the certificate file

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_abapAddonAssemblyKitCertificatePass` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`abapAddonAssemblyKitCertificatePassCredentialsId`](#abapaddonassemblykitcertificatepasscredentialsid)<br />&nbsp;&nbsp;reference to: `abapAddonAssemblyKitCertificatePass`<br /> |


#### abapAddonAssemblyKitEndpoint

Base URL to the Addon Assembly Kit as a Service (AAKaaS) system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `https://apps.support.sap.com` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### abapAddonAssemblyKitOriginHash

Origin Hash for restricted AAKaaS scenarios

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_abapAddonAssemblyKitOriginHash` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
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


#### addonDescriptorFileName

File name of the YAML file which describes the Product Version and corresponding Software Component Versions

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `addon.yml` |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


#### password

Password for the Addon Assembly Kit as a Service (AAKaaS) system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
| Default | `$PIPER_password` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
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

User for the Addon Assembly Kit as a Service (AAKaaS) system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | no |
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


#### abapAddonAssemblyKitCredentialsId

CredentialsId stored in Jenkins for the Addon Assembly Kit as a Service (AAKaaS) system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |


#### abapAddonAssemblyKitCertificateFileCredentialsId

Jenkins secret text credential ID containing the base64 encoded certificate pfx file (PKCS12 format) see note [2805811](https://me.sap.com/notes/2805811)

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |


#### abapAddonAssemblyKitCertificatePassCredentialsId

Jenkins secret text credential ID containing the password to decrypt the certificate file stored in abapAddonAssemblyKitCertificateFileCredentialsId

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Examples

### Configuration in the config.yml

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapAddonAssemblyKitCheck script: this
```

If the step is to be configured individually the config.yml should look like this:

```yaml
steps:
  abapAddonAssemblyKitCheck:
    abapAddonAssemblyKitCredentialsId: 'abapAddonAssemblyKitCredentialsId',
    addonDescriptorFileName: 'addon.yml'
```

More convenient ways of configuration (e.g. on stage level) are described in the respective scenario/pipeline documentation.
