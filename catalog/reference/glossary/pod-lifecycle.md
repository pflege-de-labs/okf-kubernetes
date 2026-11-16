---
type: glossary-term
title: Pod Lifecycle
description: The sequence of states through which a Pod passes during its lifetime.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-lifecycle
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/pod-lifecycle.md
    title: 'Kubernetes website source: reference/glossary/pod-lifecycle.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-lifecycle
    title: Pod Lifecycle | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/pod-lifecycle.md
---

# Pod Lifecycle

The sequence of states through which a Pod passes during its lifetime.

 

The [Pod Lifecycle](/concepts/workloads/pods/pod-lifecycle.md) is defined by the states or phases of a Pod. There are five possible Pod phases: Pending, Running, Succeeded, Failed, and Unknown. A high-level description of the Pod state is summarized in the [PodStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#podstatus-v1-core) `phase` field.
