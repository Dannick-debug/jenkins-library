# Custom Jenkins Setup

## Requirements

- Java Runtime Environment 8 or higher
- Jenkins 2.60.3 or higher running on Linux
- A Jenkins user with administration rights
- Docker installed on the Jenkins agent
- Network access to [github.com](https://github.com) and SAP BTP

## Docker

Most ABAP Environment Pipeline steps run inside Docker containers pulled automatically from Docker Hub. Install Docker on the Jenkins agent if not already present:

```sh
# Example for Debian/Ubuntu
sudo apt-get update
sudo apt-get install docker-ce
```

If Jenkins itself runs as a Docker container, expose the Docker socket so steps can launch sibling containers:

```
docker run ... -v /var/run/docker.sock:/var/run/docker.sock ...
```

## Shared Library

Add the Piper shared library to your Jenkins system configuration:

1. Go to **Manage Jenkins > Configure System > Global Pipeline Libraries** and click **Add**.
2. Set **Library Name** to `piper-lib-os`.
3. Set **Default Version** to the branch or tag you want to consume (e.g. `master`).
4. Set **Retrieval Method** to **Modern SCM** and **Source Code Management** to **Git**.
5. Set **Project Repository** to `https://github.com/SAP/jenkins-library`.
6. Save.

Your `Jenkinsfile` can then load the library:

```groovy
@Library('piper-lib-os') _
```

## Credentials

Configure the following credentials in Jenkins (**Manage Jenkins > Credentials**):

| Type | ID (example) | Used for |
| ---- | ------------ | -------- |
| Username with password | `cfCredentialsId` | Cloud Foundry / ABAP system provisioning |
| Username with password | `abapCredentialsId` | ABAP communication user for all `abapEnvironment*` steps |
| Secret text | `btpServiceAccountCredentialsId` | BTP API steps |

Reference these credential IDs in your `.pipeline/config.yml`:

```yaml
steps:
  abapEnvironmentCreateSystem:
    cfCredentialsId: cfCredentialsId
  abapEnvironmentCloneGitRepo:
    abapCredentialsId: abapCredentialsId
```

## User Permission Issue

If your Jenkins service user ID differs from `1000` (the default in the official Jenkins Docker image), you may encounter file permission issues when Docker containers bind-mount the workspace. To resolve this, either change the Jenkins service user ID to `1000`, or build custom Docker images that relax the file system restrictions.
