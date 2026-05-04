# Extensibility

The ABAP Environment Pipeline is delivered as a GitHub Actions reusable workflow. Consumer repositories reference it by name and never maintain the pipeline logic themselves. When the pipeline is updated upstream, consumers pick up the changes by bumping the ref (e.g. `@main` or `@v1.2.0`).

## Extending individual stages

GitHub Actions does not have a built-in extension mechanism equivalent to Groovy stage extensions, but the same outcomes can be achieved with standard GitHub Actions patterns.

### Adding steps before or after a stage

Fork this repository, modify the relevant job in `.github/workflows/abapEnvironmentPipeline.yml`, and reference your fork in the consumer workflow:

```yaml
uses: my-org/jenkins-library/.github/workflows/abapEnvironmentPipeline.yml@main
```

### Adding a custom job

Add a new job to your consumer workflow that runs in parallel with or after the reusable workflow's jobs. Jobs defined in the consumer workflow have full access to repository secrets and can call any `piper` step directly:

```yaml
jobs:
  pipeline:
    uses: SAP/jenkins-library/.github/workflows/abapEnvironmentPipeline.yml@main
    with:
      addonDescriptorFileName: addon.yml
    secrets: inherit

  custom-report:
    needs: pipeline
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Download piper
        run: |
          curl -sLo piper "https://github.com/SAP/jenkins-library/releases/latest/download/piper_linux_amd64"
          chmod +x piper
      - name: Run custom step
        run: ./piper abapEnvironmentRunATCCheck
        env:
          PIPER_user: ${{ secrets.PIPER_USER }}
          PIPER_password: ${{ secrets.PIPER_PASSWORD }}
```

## Running piper steps directly

Every pipeline step is also a standalone CLI command. You can call any step outside of the pipeline for local testing or ad-hoc use:

```sh
# Run an ATC check locally against a known system
export PIPER_user=my_user
export PIPER_password=my_password
piper abapEnvironmentRunATCCheck --atcConfig atcConfig.yml
```

See the [CLI reference](cli/index.md) for all available commands and their flags.

## Pinning to a specific release

To avoid unexpected changes from upstream updates, pin the reusable workflow to a specific release tag:

```yaml
uses: SAP/jenkins-library/.github/workflows/abapEnvironmentPipeline.yml@v1.2.0
```

Releases are listed on the [GitHub releases page](https://github.com/SAP/jenkins-library/releases).
