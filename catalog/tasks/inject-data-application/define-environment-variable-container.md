---
type: task
title: Define Environment Variables for a Container
description: This page shows how to define environment variables for a container in a Kubernetes Pod.
resource: https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
tags: [kubernetes, tasks, inject-data-application]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/inject-data-application/define-environment-variable-container.md
    title: 'Kubernetes website source: tasks/inject-data-application/define-environment-variable-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
    title: Define Environment Variables for a Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/inject-data-application/define-environment-variable-container.md
---

# Define Environment Variables for a Container

This page shows how to define environment variables for a container
in a Kubernetes Pod.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

## Define an environment variable for a container

When you create a Pod, you can set environment variables for the containers
that run in the Pod. To set environment variables, include the `env` or
`envFrom` field in the configuration file.

The `env` and `envFrom` fields have different effects.

`env`
: allows you to set environment variables for a container, specifying a value directly for each variable that you name.

`envFrom`
: allows you to set environment variables for a container by referencing either a ConfigMap or a Secret.
 When you use `envFrom`, all the key-value pairs in the referenced ConfigMap or Secret
 are set as environment variables for the container.
 You can also specify a common prefix string.

You can read more about [ConfigMap](/tasks/configure-pod-container/configure-pod-configmap.md#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables)
and [Secret](/tasks/inject-data-application/distribute-credentials-secure.md#configure-all-key-value-pairs-in-a-secret-as-container-environment-variables).

This page explains how to use `env`.

In this exercise, you create a Pod that runs one container. The configuration
file for the Pod defines an environment variable with name `DEMO_GREETING` and
value `"Hello from the environment"`. Here is the configuration manifest for the
Pod:

[`pods/inject/envars.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/inject/envars.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: envar-demo
  labels:
    purpose: demonstrate-envars
spec:
  containers:
  - name: envar-demo-container
    image: gcr.io/google-samples/hello-app:2.0
    env:
    - name: DEMO_GREETING
      value: "Hello from the environment"
    - name: DEMO_FAREWELL
      value: "Such a sweet sorrow"
```

1. Create a Pod based on that manifest:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/inject/envars.yaml
   ```

1. List the running Pods:

   ```shell
   kubectl get pods -l purpose=demonstrate-envars
   ```

   The output is similar to:

   ```
   NAME            READY     STATUS    RESTARTS   AGE
   envar-demo      1/1       Running   0          9s
   ```

1. List the Pod's container environment variables:

   ```shell
   kubectl exec envar-demo -- printenv
   ```

   The output is similar to this:

   ```
   NODE_VERSION=4.4.2
   EXAMPLE_SERVICE_PORT_8080_TCP_ADDR=10.3.245.237
   HOSTNAME=envar-demo
   ...
   DEMO_GREETING=Hello from the environment
   DEMO_FAREWELL=Such a sweet sorrow
   ```

**Note:**
The environment variables set using the `env` or `envFrom` field
override any environment variables specified in the container image.

**Note:**
Environment variables may reference each other, however ordering is important.
Variables making use of others defined in the same context must come later in
the list. Similarly, avoid circular references.

## Using environment variables inside of your config

Environment variables that you define in a Pod's configuration under 
`.spec.containers[*].env[*]` can be used elsewhere in the configuration, for 
example in commands and arguments that you set for the Pod's containers.
In the example configuration below, the `GREETING`, `HONORIFIC`, and
`NAME` environment variables are set to `Warm greetings to`, `The Most
Honorable`, and `Kubernetes`, respectively. The environment variable 
`MESSAGE` combines the set of all these environment variables and then uses it 
as a CLI argument passed to the `env-print-demo` container.

Environment variable names may consist of any [printable ASCII characters](https://www.ascii-code.com/characters/printable-characters) except '='.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: print-greeting
spec:
  containers:
  - name: env-print-demo
    image: bash
    env:
    - name: GREETING
      value: "Warm greetings to"
    - name: HONORIFIC
      value: "The Most Honorable"
    - name: NAME
      value: "Kubernetes"
    - name: MESSAGE
      value: "$(GREETING) $(HONORIFIC) $(NAME)"
    command: ["echo"]
    args: ["$(MESSAGE)"]
```

Upon creation, the command `echo Warm greetings to The Most Honorable Kubernetes` is run on the container.

## What's next

* Learn more about [environment variables](/tasks/inject-data-application/environment-variable-expose-pod-information.md).
* Learn about [using secrets as environment variables](/concepts/configuration/secret.md#using-secrets-as-environment-variables).
* See [EnvVarSource](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#envvarsource-v1-core).
