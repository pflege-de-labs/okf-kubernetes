---
type: task
title: Define a Command and Arguments for a Container
description: This page shows how to define commands and arguments when you run a container in a Pod.
resource: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/
tags: [kubernetes, tasks, inject-data-application]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/inject-data-application/define-command-argument-container.md
    title: 'Kubernetes website source: tasks/inject-data-application/define-command-argument-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/
    title: Define a Command and Arguments for a Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/inject-data-application/define-command-argument-container.md
---

# Define a Command and Arguments for a Container

This page shows how to define commands and arguments when you run a container
in a [Pod](/reference/glossary/pod.md).

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Define a command and arguments when you create a Pod

When you create a Pod, you can define a command and arguments for the
containers that run in the Pod. To define a command, include the `command`
field in the configuration file. To define arguments for the command, include
the `args` field in the configuration file. The command and arguments that
you define cannot be changed after the Pod is created.

The command and arguments that you define in the configuration file
override the default command and arguments provided by the container image.
If you define args, but do not define a command, the default command is used
with your new arguments.

**Note:**
The `command` field corresponds to `ENTRYPOINT`, and the `args` field corresponds to `CMD` in some container runtimes.

In this exercise, you create a Pod that runs one container. The configuration
file for the Pod defines a command and two arguments:

[`pods/commands.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/commands.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: command-demo
  labels:
    purpose: demonstrate-command
spec:
  containers:
  - name: command-demo-container
    image: debian
    command: ["printenv"]
    args: ["HOSTNAME", "KUBERNETES_PORT"]
  restartPolicy: OnFailure
```

1. Create a Pod based on the YAML configuration file:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/commands.yaml
   ```

1. List the running Pods:

   ```shell
   kubectl get pods
   ```

   The output shows that the container that ran in the command-demo Pod has
   completed.

1. To see the output of the command that ran in the container, view the logs
from the Pod:

   ```shell
   kubectl logs command-demo
   ```

   The output shows the values of the HOSTNAME and KUBERNETES_PORT environment
   variables:

   ```
   command-demo
   tcp://10.3.240.1:443
   ```

## Use environment variables to define arguments

In the preceding example, you defined the arguments directly by
providing strings. As an alternative to providing strings directly,
you can define arguments by using environment variables:

```yaml
env:
- name: MESSAGE
  value: "hello world"
command: ["/bin/echo"]
args: ["$(MESSAGE)"]
```

This means you can define an argument for a Pod using any of
the techniques available for defining environment variables, including
[ConfigMaps](/tasks/configure-pod-container/configure-pod-configmap.md)
and
[Secrets](/concepts/configuration/secret.md).

**Note:**
The environment variable appears in parentheses, `"$(VAR)"`. This is
required for the variable to be expanded in the `command` or `args` field.

## Run a command in a shell

In some cases, you need your command to run in a shell. For example, your
command might consist of several commands piped together, or it might be a shell
script. To run your command in a shell, wrap it like this:

```shell
command: ["/bin/sh"]
args: ["-c", "while true; do echo hello; sleep 10;done"]
```

## What's next

* Learn more about [configuring pods and containers](/tasks.md).
* Learn more about [running commands in a container](/tasks/debug/debug-application/get-shell-running-container.md).
* See [Container](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#container-v1-core).
