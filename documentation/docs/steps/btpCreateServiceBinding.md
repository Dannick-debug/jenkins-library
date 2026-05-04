# btpCreateServiceBinding

## Description

btpCreateServiceBinding

## Usage

We recommend to define values of [step parameters](#parameters) via [.pipeline/config.yml file](../configuration.md).<br />In this case, calling the step is essentially reduced to defining the step name.<br />Calling the step can be done either in an orchestrator specific way (e.g. via a Jenkins library step) or on the command line.

!!! tip ""

    === "Jenkins"

        ```groovy
        library('piper-lib-os')

        btpCreateServiceBinding script: this
        ```

    === "Command Line"

        ```sh
        piper btpCreateServiceBinding
        ```

## Parameters
