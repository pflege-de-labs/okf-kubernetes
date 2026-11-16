---
type: glossary-term
title: ResourceSlice
description: Represents one or more infrastructure resources, like devices, in a pool of similar resources.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resourceslice
tags: [kubernetes, reference, glossary, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/resourceslice.md
    title: 'Kubernetes website source: reference/glossary/resourceslice.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resourceslice
    title: ResourceSlice | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/resourceslice.md
---

# ResourceSlice

Represents one or more infrastructure resources, such as
[devices](/reference/glossary/device.md), that are attached to
nodes. Drivers create and manage ResourceSlices in the cluster. ResourceSlices
are used for
[dynamic resource allocation (DRA)](/concepts/resource-management/dynamic-resource-allocation.md).

When a [ResourceClaim](/reference/glossary/resourceclaim.md) is
created, Kubernetes uses ResourceSlices to find nodes that have access to
resources that can satisfy the claim. Kubernetes allocates resources to the
ResourceClaim and schedules the Pod onto a node that can access the resources.
