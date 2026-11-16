---
type: glossary-term
title: Resource (infrastructure)
description: A defined amount of infrastructure available for consumption (CPU, memory, etc).
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-infrastructure-resource
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/infrastructure-resource.md
    title: 'Kubernetes website source: reference/glossary/infrastructure-resource.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-infrastructure-resource
    title: Resource (infrastructure) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/infrastructure-resource.md
---

# Resource (infrastructure)

Capabilities provided to one or more [nodes](/reference/glossary/node.md) (CPU, memory, GPUs, etc), and made available for consumption by
[Pods](/reference/glossary/pod.md) running on those nodes.

Kubernetes also uses the term _resource_ to describe an [API resource](/reference/glossary/api-resource.md).

Computers provide fundamental hardware facilities: processing power, storage memory, network, etc.
These resources have finite capacity, measured in a unit applicable to that resource (number of CPUs, bytes of memory, etc).
Kubernetes abstracts common [resources](/concepts/configuration/manage-resources-containers.md)
for allocation to workloads and utilizes operating system primitives (for example, Linux [cgroups](/reference/glossary/cgroup.md)) to manage consumption by [workloads](/reference/glossary/workload.md)).

You can also use [dynamic resource allocation](/concepts/resource-management/dynamic-resource-allocation.md) to
manage complex resource allocations automatically.
