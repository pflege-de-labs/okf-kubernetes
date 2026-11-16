---
type: glossary-term
title: PodGroup
description: A PodGroup represents a set of Pods with common scheduling policy and constraints.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-podgroup
tags: [kubernetes, reference, glossary, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/podgroup.md
    title: 'Kubernetes website source: reference/glossary/podgroup.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-podgroup
    title: PodGroup | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/podgroup.md
---

# PodGroup

A PodGroup is a runtime object that represents a group of Pods scheduled
together as a single unit. While the
[Workload API](/concepts/workloads/workload-api.md) defines scheduling policy
templates, PodGroups are the runtime counterparts that carry both the policy and
the scheduling status for a specific instance of that group.
