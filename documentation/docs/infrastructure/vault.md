# Vault for Pipeline Secrets

The Piper ABAP Environment Pipeline supports fetching secrets from [HashiCorp Vault](https://www.hashicorp.com/products/vault) (KV engine v1 and v2) instead of storing them as Jenkins credentials.

Parameters that support Vault are marked with the Vault label in the step documentation.

## Authentication

### AppRole (recommended)

1. Enable AppRole authentication in your Vault instance.
2. Create an AppRole role for Piper and assign the necessary policies.
3. Store the **AppRole ID** and **AppRole Secret ID** as Jenkins `Secret Text` credentials.
4. Reference them in your pipeline configuration via `vaultAppRoleTokenCredentialsId` and `vaultAppRoleSecretTokenCredentialsId`.

Piper will rotate the SecretID automatically, which is not possible with token authentication.

### Token

Store the Vault token as a Jenkins `Secret Text` credential and reference it via `vaultTokenCredentialsId`.

Alternatively, pass the token directly as the environment variable `PIPER_vaultToken`.

## Pipeline Configuration

Add the following to your `.pipeline/config.yml`:

```yaml
general:
  vaultServerUrl: 'https://your-vault-server'
  vaultNamespace: 'your-namespace'   # omit if not using namespaces
  vaultPath: 'kv/abap-pipeline'      # path where secrets are stored
```

Piper looks up secrets in this order:

1. `<vaultPath>/<secretPath>`
2. `<vaultBasePath>/<vaultPipelineName>/<secretPath>`
3. `<vaultBasePath>/GROUP-SECRETS/<secretPath>`

## Controlling Secret Lookup

### Prevent overwriting explicit config values

By default, Vault values overwrite parameters set in `config.yml`. To disable this:

```yaml
general:
  vaultDisableOverwrite: true
```

### Skip Vault for specific steps

```yaml
steps:
  abapEnvironmentRunATCCheck:
    skipVault: true
```

## Fetching General Purpose Credentials

Vault can supply arbitrary credentials to any step, for example for custom extensions:

```yaml
steps:
  abapEnvironmentBuild:
    vaultCredentialPath: 'abap-build-credentials'
    vaultCredentialKeys: ['myUser', 'myPassword']
```

The values are exposed as environment variables prefixed by `PIPER_VAULTCREDENTIAL_` (e.g. `PIPER_VAULTCREDENTIAL_MYUSER`). A Base64-encoded variant is also provided as `PIPER_VAULTCREDENTIAL_MYUSER_BASE64`.

To use a custom prefix:

```yaml
    vaultCredentialEnvPrefix: 'ABAP_CRED_'
```

Enable verbose logging for Vault lookups with `verbose: true`.
