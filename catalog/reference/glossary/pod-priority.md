---
type: glossary-term
title: Pod Priority
description: Pod Priority indicates the importance of a Pod relative to other Pods.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-priority
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/pod-priority.md
    title: 'Kubernetes website source: reference/glossary/pod-priority.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-priority
    title: Pod Priority | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/pod-priority.md
---

# Pod Priority

Pod Priority indicates the importance of a [Pod](/reference/glossary/pod.md) relative to other Pods.

[Pod Priority](/concepts/scheduling-eviction/pod-priority-preemption.md#pod-priority) gives the ability to set scheduling priority of a Pod to be higher and lower than other Pods — an important feature for production clusters workload.
