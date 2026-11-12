---
type: glossary-term
title: Pod Priority
description: Pod Priority indicates the importance of a Pod relative to other Pods.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-priority
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/pod-priority.md
    title: 'Kubernetes website source: reference/glossary/pod-priority.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-priority
    title: Pod Priority | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/pod-priority.md
---

# Pod Priority

Pod Priority indicates the importance of a [Pod](/reference/glossary/pod.md) relative to other Pods.

[Pod Priority](/concepts/scheduling-eviction/pod-priority-preemption.md#pod-priority) gives the ability to set scheduling priority of a Pod to be higher and lower than other Pods — an important feature for production clusters workload.
