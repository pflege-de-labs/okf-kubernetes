---
type: task
title: Scale a StatefulSet
description: This task shows how to scale a StatefulSet.
resource: https://kubernetes.io/docs/tasks/run-application/scale-stateful-set/
tags: [kubernetes, tasks, run-application]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/run-application/scale-stateful-set.md
    title: 'Kubernetes website source: tasks/run-application/scale-stateful-set.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/run-application/scale-stateful-set/
    title: Scale a StatefulSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/run-application/scale-stateful-set.md
---

# Scale a StatefulSet

This task shows how to scale a StatefulSet. Scaling a StatefulSet refers to
increasing or decreasing the number of replicas.

## Before you begin

- StatefulSets are only available in Kubernetes version 1.5 or later.
  To check your version of Kubernetes, run `kubectl version`.

- Not all stateful applications scale nicely. If you are unsure about whether
  to scale your StatefulSets, see [StatefulSet concepts](/concepts/workloads/controllers/statefulset.md)
  or [StatefulSet tutorial](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/) for further information.

- You should perform scaling only when you are confident that your stateful application
  cluster is completely healthy.

## Scaling StatefulSets

### Use kubectl to scale StatefulSets

First, find the StatefulSet you want to scale.

```shell
kubectl get statefulsets <stateful-set-name>
```

Change the number of replicas of your StatefulSet:

```shell
kubectl scale statefulsets <stateful-set-name> --replicas=<new-replicas>
```

### Make in-place updates on your StatefulSets

Alternatively, you can do
[in-place updates](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#in-place-updates-of-resources)
on your StatefulSets.

If your StatefulSet was initially created with `kubectl apply`,
update `.spec.replicas` of the StatefulSet manifests, and then do a `kubectl apply`:

```shell
kubectl apply -f <stateful-set-file-updated>
```

Otherwise, edit that field with `kubectl edit`:

```shell
kubectl edit statefulsets <stateful-set-name>
```

Or use `kubectl patch`:

```shell
kubectl patch statefulsets <stateful-set-name> -p '{"spec":{"replicas":<new-replicas>}}'
```

## Troubleshooting

### Scaling down does not work right

You cannot scale down a StatefulSet when any of the stateful Pods it manages is
unhealthy. Scaling down only takes place after those stateful Pods become running and ready.

If spec.replicas > 1, Kubernetes cannot determine the reason for an unhealthy Pod.
It might be the result of a permanent fault or of a transient fault. A transient
fault can be caused by a restart required by upgrading or maintenance.

If the Pod is unhealthy due to a permanent fault, scaling
without correcting the fault may lead to a state where the StatefulSet membership
drops below a certain minimum number of replicas that are needed to function
correctly. This may cause your StatefulSet to become unavailable.

If the Pod is unhealthy due to a transient fault and the Pod might become available again,
the transient error may interfere with your scale-up or scale-down operation. Some distributed
databases have issues when nodes join and leave at the same time. It is better
to reason about scaling operations at the application level in these cases, and
perform scaling only when you are sure that your stateful application cluster is
completely healthy.

## What's next

- Learn more about [deleting a StatefulSet](/tasks/run-application/delete-stateful-set.md).
