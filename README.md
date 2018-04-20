# jenkins-jobs-configmaps

Some spikes (and maybe eventually a project) to try and manage jenkins jobs in a CM rather than with xml files. This would allow for a much smoother integration into Helm charts...

Current State: Very minimal yaml to xml parser (template based)

## Try it out

Clone and modify `jobs.yaml`. Run `go run main.go` and check `jobs/` folder for output.

## Possible Roadmap (subject to change)

* Run this tool in an initContainer alongside Jenkins to configure jobs at startup
* Run as sidecar container to constantly watch updates and make adding of new jobs possible without having to restart Jenkins
  * At this point it probably makes sense to save the configuration in CRDs rather than config maps

## Sample `jobs.yaml`:

```yaml
---
jobs:
  # A minimal job
  - name: sample-job
    git:
      enabled: true
      remoteName: "origin"
      remoteURL: "git@github.com/bar/foo.git"
      jenkinsCredentialsID: "foobar"
      branches: "*/master"

  # A more advanced job
  - name: advanced-job
    displayName: A bigger job
    description: Possibly the best job in the world
    git:
      enabled: true
      remoteName: "origin"
      remoteURL: "git@github.com/bar/foo.git"
      jenkinsCredentialsID: "foobar"
      branches: "*/master"
    triggers:
      scmPolling:
        enabled: true
        schedule: "H * * * *"
        ingorePostCommitHooks: false
      Timer:
        enabled: true
        schedule: "H 2 * * *"
```
