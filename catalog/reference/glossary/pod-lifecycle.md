---
type: glossary-term
title: Pod Lifecycle
description: The sequence of states through which a Pod passes during its lifetime.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-lifecycle
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/pod-lifecycle.md
    title: 'Kubernetes website source: reference/glossary/pod-lifecycle.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-lifecycle
    title: Pod Lifecycle | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/pod-lifecycle.md
---

# Pod Lifecycle

The sequence of states through which a Pod passes during its lifetime.

 

The [Pod Lifecycle](/concepts/workloads/pods/pod-lifecycle.md) is defined by the states or phases of a Pod. There are five possible Pod phases: Pending, Running, Succeeded, Failed, and Unknown. A high-level description of the Pod state is summarized in the [PodStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#podstatus-v1-core) `phase` field.
