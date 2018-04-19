# jenknis-jobs-configmaps

Some spikes (and maybe eventually a project) to try and manage jenkins jobs in a CM rather than with xml files. This would allow for a much smoother integration into Helm charts...

Current State: Very minimal yaml to xml parser (template based)

## Try it out

Clone and modify `jobs.yaml`. Run `go run main.go` and check `jobs/` folder for output.
