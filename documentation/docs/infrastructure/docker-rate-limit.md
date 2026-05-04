# Fixing Docker Rate Limit Errors

You may see the following error in pipeline runs:

```
Error response from daemon: toomanyrequests: You have reached your pull rate limit.
```

Docker Hub introduced pull rate limits in November 2020. Anonymous pulls are limited to 100/6h; authenticated free accounts to 200/6h.

## Options

### Company-internal Docker Hub mirror

If your company runs an Artifactory or similar registry with Docker Hub mirroring enabled, configure it in `.pipeline/config.yml`:

```yaml
steps:
  dockerExecute:
    dockerRegistryUrl: 'https://my.internal.registry:1234'
```

### Authenticated pulls

Configure a Docker Hub account in Jenkins and pass the credential ID to the step:

```yaml
steps:
  dockerExecute:
    dockerRegistryCredentialsId: 'dockerHubCredentialsId'
```

### GitHub Container Registry

Piper's Docker images are also published to [GitHub Container Registry](https://github.com/orgs/SAP/packages). Reference the `ghcr.io` image URL in your step configuration if Docker Hub pulls fail consistently.

### Hyperscaler mirror

Cloud providers (AWS, Azure, GCP) often offer a Docker Hub mirror for workloads running on their platforms. Check your provider's documentation for the mirror URL.
