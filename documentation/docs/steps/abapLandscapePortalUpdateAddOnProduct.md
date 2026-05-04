# abapLandscapePortalUpdateAddOnProduct

Update the AddOn product in SAP BTP ABAP Environment system of Landscape Portal


## Description

This step describes the AddOn product update in SAP BTP ABAP Environment system of Landscape Portal


## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        abapLandscapePortalUpdateAddOnProduct script: this
        ```

    === "Command Line"

        ```sh
        piper abapLandscapePortalUpdateAddOnProduct
        ```



## Prerequisites

- Please make sure, that you are under Embedded Steampunk environment.
- Please make sure, that the service landscape-portal-api-for-s4hc with plan api was assigned as entitlement to the subaccount, where you are about to deploy addon product.
- Please make sure, that before deploying addon product, an instance of landscape-portal-api-for-s4hc (plan api) was created, and a service key with x509 authentication mechanism was created for the instance. The service key needs to be stored in the Jenkins Credentials Store.
- Please make sure, that the system to deploy addon product is active, and the descriptor file with deployment information is available.

## Parameters

### Overview - Step

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |
| [abapSystemNumber](#abapsystemnumber) | **yes** |  |
| [landscapePortalAPIServiceKey](#landscapeportalapiservicekey) | **(yes)** | ![Secret](https://img.shields.io/badge/-Secret-yellowgreen) pass via ENV or Jenkins credentials ([`landscapePortalAPICredentialsId`](#landscapeportalapicredentialsid)) |
| [script](#script) | **(yes)** | ![Jenkins only](https://img.shields.io/badge/-Jenkins%20only-yellowgreen) reference to Jenkins main pipeline script |
| [addonDescriptorFileName](#addondescriptorfilename) | no |  |
| [verbose](#verbose) | no | activates debug output |

### Overview - Execution Environment

!!! note "Orchestrator-specific only"

    These parameters are relevant for orchestrator usage and not considered when using the command line option.

| Name | Mandatory | Additional information |
| ---- | --------- | ---------------------- |

### Details

#### abapSystemNumber

System Number of the abap integration test system

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_abapSystemNumber` (if set) |
| Secret | no |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |
| Resource references | none |


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


#### landscapePortalAPIServiceKey

Service key JSON string to access the Landscape Portal Access API

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Mandatory | **yes** |
| Default | `$PIPER_landscapePortalAPIServiceKey` (if set) |
| Secret | **yes** |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9744; general</li><li>&#9744; steps</li><li>&#9744; stages</li></ul> |
| Resource references | Jenkins credential id:<br />&nbsp;&nbsp;id: [`landscapePortalAPICredentialsId`](#landscapeportalapicredentialsid)<br />&nbsp;&nbsp;reference to: `landscapePortalAPIServiceKey`<br /> |


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


#### landscapePortalAPICredentialsId

Jenkins secret text credential ID containing the service key to access the Landscape Portal Access API

[back to overview](#parameters)

| Scope | Details |
| ---- | --------- |
| Aliases | - |
| Type | `string` |
| Configuration scope | <ul><li>&#9746; parameter</li><li>&#9746; general</li><li>&#9746; steps</li><li>&#9746; stages</li></ul> |








## Example: Configuration in the config.yml

The recommended way to configure your pipeline is via the config.yml file. In this case, calling the step in the Jenkinsfile is reduced to one line:

```groovy
abapLandscapePortalUpdateAddOnProduct script: this
```

The configuration values for the addon update can be passed through the `config.yml` file:

```yaml
steps:
  abapLandscapePortalUpdateAddOnProduct:
    landscapePortalAPICredentialsId: 'landscapePortalAPICredentialsId'
    abapSystemNumber: 'abapSystemNumber'
    addonDescriptorFileName: 'addon.yml'
    addonDescriptor: 'addonDescriptor'
```

## Example: Configuration in the Jenkinsfile

The step, including all parameters, can also be called directly from the Jenkinsfile. In the following example, a configuration file is used.

```groovy
abapLandscapePortalUpdateAddOnProduct (
  script: this,
  landscapePortalAPICredentialsId: 'landscapePortalAPICredentialsId'
  abapSystemNumber: 'abapSystemNumber'
  addonDescriptorFileName: 'addon.yml'
  addonDescriptor: 'addonDescriptor'
)
```

The file `addon.yml` would look like this:

```yaml
addonProduct: some-addon-product
addonVersion: some-addon-version
```
