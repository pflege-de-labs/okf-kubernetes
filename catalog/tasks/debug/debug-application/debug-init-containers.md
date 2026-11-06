---
type: task
title: Debug Init Containers
description: This page shows how to investigate problems related to the execution of Init Containers.
resource: https://kubernetes.io/docs/tasks/debug/debug-application/debug-init-containers/
tags: [kubernetes, tasks, debug]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/debug/debug-application/debug-init-containers.md
    title: 'Kubernetes website source: tasks/debug/debug-application/debug-init-containers.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/debug/debug-application/debug-init-containers/
    title: Debug Init Containers | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/debug/debug-application/debug-init-containers.md
---

# Debug Init Containers

This page shows how to investigate problems related to the execution of
Init Containers. The example command lines below refer to the Pod as
`<pod-name>` and the Init Containers as `<init-container-1>` and
`<init-container-2>`.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

* You should be familiar with the basics of
  [Init Containers](/concepts/workloads/pods/init-containers.md).
* You should have [Configured an Init Container](/tasks/configure-pod-container/configure-pod-initialization.md#create-a-pod-that-has-an-init-container).

## Checking the status of Init Containers

Display the status of your pod:

```shell
kubectl get pod <pod-name>
```

For example, a status of `Init:1/2` indicates that one of two Init Containers
has completed successfully:

```
NAME         READY     STATUS     RESTARTS   AGE
<pod-name>   0/1       Init:1/2   0          7s
```

See [Understanding Pod status](#understanding-pod-status) for more examples of
status values and their meanings.

## Getting details about Init Containers

View more detailed information about Init Container execution:

```shell
kubectl describe pod <pod-name>
```

For example, a Pod with two Init Containers might show the following:

```
Init Containers:
  <init-container-1>:
    Container ID:    ...
    ...
    State:           Terminated
      Reason:        Completed
      Exit Code:     0
      Started:       ...
      Finished:      ...
    Ready:           True
    Restart Count:   0
    ...
  <init-container-2>:
    Container ID:    ...
    ...
    State:           Waiting
      Reason:        CrashLoopBackOff
    Last State:      Terminated
      Reason:        Error
      Exit Code:     1
      Started:       ...
      Finished:      ...
    Ready:           False
    Restart Count:   3
    ...
```

You can also access the Init Container statuses programmatically by reading the
`status.initContainerStatuses` field on the Pod Spec:

```shell
kubectl get pod <pod-name> --template '{{.status.initContainerStatuses}}'
```

This command will return the same information as above, formatted using a [Go template](https://pkg.go.dev/text/template).

## Accessing logs from Init Containers

Pass the Init Container name along with the Pod name
to access its logs.

```shell
kubectl logs <pod-name> -c <init-container-2>
```

Init Containers that run a shell script print
commands as they're executed. For example, you can do this in Bash by running
`set -x` at the beginning of the script.

## Understanding Pod status

A Pod status beginning with `Init:` summarizes the status of Init Container
execution. The table below describes some example status values that you might
see while debugging Init Containers.

Status | Meaning
------ | -------
`Init:N/M` | The Pod has `M` Init Containers, and `N` have completed so far.
`Init:Error` | An Init Container has failed to execute.
`Init:CrashLoopBackOff` | An Init Container has failed repeatedly.
`Pending` | The Pod has not yet begun executing Init Containers.
`PodInitializing` or `Running` | The Pod has already finished executing Init Containers.
