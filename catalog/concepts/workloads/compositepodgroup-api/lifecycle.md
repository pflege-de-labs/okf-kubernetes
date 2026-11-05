---
type: concept
title: CompositePodGroup Lifecycle
description: A CompositePodGroup represents a non-leaf node in a multi-level `PodGroup` hierarchy.
resource: https://kubernetes.io/docs/concepts/workloads/compositepodgroup-api/lifecycle/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/compositepodgroup-api/lifecycle.md
    title: 'Kubernetes website source: concepts/workloads/compositepodgroup-api/lifecycle.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/compositepodgroup-api/lifecycle/
    title: CompositePodGroup Lifecycle | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/workloads/compositepodgroup-api/lifecycle.md
---

# CompositePodGroup Lifecycle

**Feature state:** Kubernetes v1.37 [alpha] (feature gate `CompositePodGroup`, disabled by default)

A [CompositePodGroup](/concepts/workloads/compositepodgroup-api.md) represents a non-leaf node
in a multi-level `PodGroup` hierarchy. Unlike `PodGroup` resources, `CompositePodGroup` resources do
not directly contain Pods. Instead, they maintain a hierarchy of descendant `CompositePodGroup` and
`PodGroup` objects and carry scheduling policies that apply to its children groups.

## Ownership and garbage collection

A `CompositePodGroup` object, together with its descendant `CompositePodGroup` and `PodGroup`
resources, is owned by the workload controller that created it via Kubernetes `ownerReferences`.
When the owning workload object gets deleted, cascading garbage collection automatically deletes the
associated group hierarchy.

`CompositePodGroup` names must be unique within a namespace and must be valid
[DNS subdomains](/concepts/overview/working-with-objects/names.md#dns-subdomain-names).

## Creation ordering

To ensure proper hierarchy resolution and scheduling, workload controllers create resources in a
top-down order:

1. **`Workload`**: Defines static templates (`CompositePodGroupTemplates` and `PodGroupTemplates`).
2. **Root `CompositePodGroup`**: Created with `spec.workloadRef` pointing to the root template in
   the `Workload`.
3. **Descendant `CompositePodGroups` and `PodGroups`**: Created top-down. Each child group specifies
   its parent by using `spec.parentCompositePodGroupName` and its template using `spec.workloadRef`.
4. **`Pods`**: Created with `spec.schedulingGroup.podGroupName` pointing to their leaf `PodGroup`.

If a group references a parent `CompositePodGroup` that does not exist, or if a Pod references
a `PodGroup` that has not yet been created, the scheduler holds off scheduling until all parent
resources in the hierarchy exist.

## Limitations and validation rules

- **Consistent scheduler name**: All Pods across an entire `CompositePodGroup` hierarchy must use
  the same `spec.schedulerName`. If a mismatch is detected, the scheduler rejects the hierarchy as
  unschedulable.
- **Consistent priority**: All Pods across an entire `CompositePodGroup` hierarchy must specify the
  same value of `spec.priority` which must be equal to the priority specified by the root group. If
  a mismatch is detected, the scheduler rejects the hierarchy as unschedulable.
- **Consistent preemption policy**: All Pods across an entire `CompositePodGroup` hierarchy must use
  the same `spec.preemptionPolicy`. In addition, when the
  [PodGroupPreemptionPolicy](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/podgroup-preemption-policy/)
  feature gate is enabled, the root group's preemption policy must be equal to the one specified by
  the Pods. If a mismatch is detected, the scheduler rejects the hierarchy as unschedulable.
- **Maximum nesting depth**: The group-template hierarchy supports a maximum depth of 4 levels.
- **List item limit**: The maximum number of child `CompositePodGroupTemplates` and
  `PodGroupTemplates` at any level of a `Workload` is 8.
- **Immutable gang group count**: The `spec.schedulingPolicy.gang.minGroupCount` field on a
  `CompositePodGroup` is immutable after creation.
- **Immutable hierarchy references**: `spec.parentCompositePodGroupName` on groups and
  `spec.schedulingGroup` on Pods are immutable once set.

## What's next

* Read the [CompositePodGroup API](/concepts/workloads/compositepodgroup-api.md) overview.
* Learn about the [Workload API](/concepts/workloads/workload-api.md) and template definitions.
* See how leaf groups are structured in the [PodGroup API](/concepts/workloads/podgroup-api.md).
* Read about [PodGroup scheduling policies](/concepts/workloads/workload-api/policies.md).
* Learn about [Workload-aware preemption](/concepts/scheduling-eviction/workload-aware-preemption.md).
