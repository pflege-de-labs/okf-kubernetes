---
type: concept
title: Scheduling Group
description: You can link a `Pod` to a PodGroup to indicate that the `Pod` belongs to a group of `Pods` scheduled together.
resource: https://kubernetes.io/docs/concepts/workloads/pods/scheduling-group/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/pods/scheduling-group.md
    title: 'Kubernetes website source: concepts/workloads/pods/scheduling-group.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/pods/scheduling-group/
    title: Scheduling Group | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/workloads/pods/scheduling-group.md
---

# Scheduling Group

**Feature state:** Kubernetes v1.37 [beta] (feature gate `GenericWorkload`, disabled by default)

You can link a `Pod` to a [PodGroup](/concepts/workloads/podgroup-api.md) to indicate
that the `Pod` belongs to a group of `Pods` scheduled together. This enables the scheduler
to apply group-level policies such as gang scheduling rather than treating each `Pod` independently.

## Specifying a scheduling group

When the [`GenericWorkload`](/reference/command-line-tools-reference/feature-gates.md#GenericWorkload)
feature gate is enabled,
you can set the `spec.schedulingGroup` field in your `Pod` manifest. This field establishes a link to a specific `PodGroup` object in the same namespace by name.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: worker-0
  namespace: some-ns
spec:
  schedulingGroup:
    podGroupName: training-worker-0
  containers:
  - name: ml-worker
    image: training:v1
```

The `schedulingGroup` field is immutable. Once set, a `Pod` cannot be moved to a
different `PodGroup`.

## Behavior

When you set `spec.schedulingGroup`, the scheduler looks up the referenced
[PodGroup](/concepts/workloads/podgroup-api.md) and applies the
[scheduling policy](/concepts/workloads/workload-api/policies.md) defined in it:

* If the `PodGroup` uses the `basic` policy, each `Pod` is scheduled independently using
  standard Kubernetes behavior. The grouping is used as group-level label.
* If the `PodGroup` uses the `gang` policy, the `Pod` enters an "all-or-nothing" scheduling
  lifecycle. The scheduler tries to place at least `minCount` `Pods` in the group
  simultaneously; none of them bind to nodes unless the minimum is met.

**Feature state:** Kubernetes v1.37 [alpha] (feature gate `CompositePodGroup`, disabled by default)

When the [`CompositePodGroup`](/reference/command-line-tools-reference/feature-gates.md#CompositePodGroup)
feature gate is enabled, a `PodGroup` may also specify a parent `CompositePodGroup`. In a hierarchical
workload, scheduling is governed by policies defined across the entire group tree (such as multi-level
gang scheduling or topology constraints).

## Missing group references

If a `Pod` references a `PodGroup` that does not yet exist, the `Pod` remains pending.
Similarly, if the referenced `PodGroup` specifies a parent `CompositePodGroup` (via
`spec.parentCompositePodGroupName`) that has not been created yet, scheduling does not start and
the `Pod` remains pending until the entire group hierarchy exists in the cluster.

The scheduler automatically reconsiders the `Pod` once all required `PodGroup` and
`CompositePodGroup` resources exist.

## What's next

* Learn about the [PodGroup API](/concepts/workloads/podgroup-api.md) and its lifecycle.
* Read about the [CompositePodGroup API](/concepts/workloads/compositepodgroup-api.md).
* Read about [PodGroup scheduling policies](/concepts/workloads/workload-api/policies.md).
* Understand the [gang scheduling](/concepts/scheduling-eviction/gang-scheduling.md) algorithm.
