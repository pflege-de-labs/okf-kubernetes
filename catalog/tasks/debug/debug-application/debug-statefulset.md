---
type: task
title: Debug a StatefulSet
description: This task shows you how to debug a StatefulSet.
resource: https://kubernetes.io/docs/tasks/debug/debug-application/debug-statefulset/
tags: [kubernetes, tasks, debug]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/debug/debug-application/debug-statefulset.md
    title: 'Kubernetes website source: tasks/debug/debug-application/debug-statefulset.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/debug/debug-application/debug-statefulset/
    title: Debug a StatefulSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/debug/debug-application/debug-statefulset.md
---

# Debug a StatefulSet

This task shows you how to debug a StatefulSet.

## Before you begin

* You need to have a Kubernetes cluster, and the kubectl command-line tool must be configured to communicate with your cluster.
* You should have a StatefulSet running that you want to investigate.

## Debugging a StatefulSet

In order to list all the pods which belong to a StatefulSet, which have a label `app.kubernetes.io/name=MyApp` set on them,
you can use the following:

```shell
kubectl get pods -l app.kubernetes.io/name=MyApp
```

If you find that any Pods listed are in `Unknown` or `Terminating` state for an extended period of time,
refer to the [Deleting StatefulSet Pods](/tasks/run-application/delete-stateful-set.md) task for
instructions on how to deal with them.
You can debug individual Pods in a StatefulSet using the
[Debugging Pods](/tasks/debug/debug-application/debug-pods.md) guide.

## What's next

Learn more about [debugging an init-container](/tasks/debug/debug-application/debug-init-containers.md).
