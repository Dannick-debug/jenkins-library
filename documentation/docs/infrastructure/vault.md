# Secrets Management

All secrets are passed to the `piper` CLI as environment variables with a `PIPER_` prefix. The `piper` binary reads them automatically — no explicit flags needed.

## GitHub repository secrets (recommended)

Store credentials as [encrypted secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) in your GitHub repository under **Settings → Secrets and variables → Actions**.

The reusable workflow forwards them to each job via the `secrets:` block in the consumer workflow:

```yaml
secrets:
  PIPER_abapAddonAssemblyKitCookie: ${{ secrets.PIPER_ABAP_ADDON_ASSEMBLY_KIT_COOKIE }}
  PIPER_user: ${{ secrets.PIPER_USER }}
  PIPER_password: ${{ secrets.PIPER_PASSWORD }}
```

## HashiCorp Vault (optional)

The `piper` binary can fetch secrets from a Vault KV engine (v1 or v2) at runtime. Configure it in `.pipeline/config.yml`:

```yaml
general:
  vaultServerUrl: 'https://your-vault-server'
  vaultNamespace: 'your-namespace'
  vaultPath: 'kv/abap-pipeline'
```

Authenticate with AppRole (recommended — supports automatic SecretID rotation):

```yaml
general:
  vaultAppRoleTokenCredentialsId: vault-approle-id
  vaultAppRoleSecretTokenCredentialsId: vault-approle-secret
```

Or with a token via the environment variable `PIPER_vaultToken`.

### Vault lookup order

Piper resolves secrets in the following order:

1. `<vaultPath>/<secretPath>`
2. `<vaultBasePath>/<vaultPipelineName>/<secretPath>`
3. `<vaultBasePath>/GROUP-SECRETS/<secretPath>`

### Options

Prevent Vault from overwriting values already set in `config.yml`:

```yaml
general:
  vaultDisableOverwrite: true
```

Skip Vault for a specific step:

```yaml
steps:
  abapEnvironmentRunATCCheck:
    skipVault: true
```

Fetch arbitrary credentials from Vault and expose them as environment variables to a step:

```yaml
steps:
  abapEnvironmentBuild:
    vaultCredentialPath: 'abap-build-credentials'
    vaultCredentialKeys: ['myUser', 'myPassword']
```

Values are available as `PIPER_VAULTCREDENTIAL_MYUSER` and `PIPER_VAULTCREDENTIAL_MYUSER_BASE64`.
