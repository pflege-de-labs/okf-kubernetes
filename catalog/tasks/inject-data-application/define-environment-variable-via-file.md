---
type: task
title: Define Environment Variable Values Using An Init Container
description: This page show how to configure environment variables for containers in a Pod via file.
resource: https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-via-file/
tags: [kubernetes, tasks, inject-data-application]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/inject-data-application/define-environment-variable-via-file.md
    title: 'Kubernetes website source: tasks/inject-data-application/define-environment-variable-via-file.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-via-file/
    title: Define Environment Variable Values Using An Init Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/inject-data-application/define-environment-variable-via-file.md
min_kubernetes_server_version: v1.34
---

# Define Environment Variable Values Using An Init Container

**Feature state:** Kubernetes v1.35 [beta] (feature gate `EnvFiles`, enabled by default)

This page show how to configure environment variables for containers in a Pod via file.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

Your Kubernetes server must be at or later than version v1.34. To check the version, enter `kubectl version`.

## How the design works

In this exercise, you will create a Pod that sources environment variables from files, 
projecting these values into the running container.

[`pods/inject/envars-file-container.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/inject/envars-file-container.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: envfile-test-pod
spec:
  initContainers:
    - name: setup-envfile
      image:  nginx
      command: ['sh', '-c', "echo \"DB_ADDRESS=\'address\'\nREST_ENDPOINT=\'endpoint\'\" > /data/config.env"]
      volumeMounts:
        - name: config
          mountPath: /data
  containers:
    - name: use-envfile
      image: nginx
      command: [ "/bin/sh", "-c", "env" ]
      env:
        - name: DB_ADDRESS
          valueFrom:
            fileKeyRef:
              path: config.env
              volumeName: config
              key: DB_ADDRESS
              optional: false
  restartPolicy: Never
  volumes:
    - name: config
      emptyDir: {}
```

In this manifest, you can see the `initContainer` mounts an `emptyDir` volume and writes environment variables to a file within it,
and the regular containers reference both the file and the environment variable key 
through the `fileKeyRef` field without needing to mount the volume. 
When `optional` field is set to false, the specified `key` in `fileKeyRef` must exist in the environment variables file.

The volume will only be mounted to the container that writes to the file
(`initContainer`), while the consumer container that consumes the environment variable will not have the volume mounted.

The env file format adheres to the [kubernetes env file standard](/tasks/inject-data-application/define-environment-variable-via-file.md#env-file-syntax).

During container initialization, the kubelet retrieves environment variables 
from specified files in the `emptyDir` volume and exposes them to the container.

**Note:**
All container types (initContainers, regular containers, sidecars containers,
and ephemeral containers) support environment variable loading from files.

While these environment variables can store sensitive information, 
`emptyDir` volumes don't provide the same protection mechanisms as
dedicated Secret objects. Therefore, exposing confidential environment variables 
to containers through this feature is not considered a security best practice.

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/inject/envars-file-container.yaml
```

Verify that the container in the Pod is running:

```shell
# If the new Pod isn't yet healthy, rerun this command a few times.
kubectl get pods
```

Check container logs for environment variables:

```shell
kubectl logs envfile-test-pod -c use-envfile | grep DB_ADDRESS
```

The output shows the values of selected environment variables:

```
DB_ADDRESS=address
```

## Env file syntax {#env-file-syntax}

The env file format used by Kubernetes is a well-defined subset of the environment variable semantics for POSIX-compliant bash. Any env file supported by Kubernetes will produce the same environment variables as when interpreted by a POSIX-compliant bash. However, POSIX-compliant bash supports some additional formats that Kubernetes does not accept.

Example:

```
MY_VAR='my-literal-value'
```

### Rules

* Variable declaration: Use the form `VAR='value'`. Spaces surrounding `=` are ignored; leading spaces on a line are ignored; blank lines are ignored.
* Quoted values: Values must be enclosed in single quotes (`'`).
  * The content inside single quotes is preserved literally. No escape-sequence processing, whitespace folding, or character interpretation is applied.
  * Newlines inside single quotes are preserved (multi-line values are supported).
* Comments: Lines that begin with `#` are treated as comments and ignored. A `#` character inside a single-quoted value is not a comment.

Examples:

```
# comment
DB_ADDRESS='address'

MULTI='line1
line2'
```

### Unsupported forms

* Unquoted values are **prohibited**:
  * `VAR=value` — not supported.
* Double-quoted values are **prohibited**:
  * `VAR="value"` — not supported.
* Multiple adjacent quoted strings are **not** supported:
  * `VAR='val1''val2'` — not supported.
* Any form of interpolation, expansion, or concatenation is **not** supported:
  * `VAR='a'$OTHER` or `VAR=${OTHER}` — not supported.

The strict single-quote requirement ensures the value is taken literally by the kubelet when loading environment variables from files.

## What's next

* Learn more about [environment variables](/tasks/inject-data-application/environment-variable-expose-pod-information.md).
* Read [Defining Environment Variables for a Container](/tasks/inject-data-application/define-environment-variable-container.md)
* Read [Expose Pod Information to Containers Through Environment Variables](/tasks/inject-data-application/environment-variable-expose-pod-information.md)
