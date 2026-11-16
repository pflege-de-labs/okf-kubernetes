---
type: glossary-term
title: Drain
description: Safely evicts Pods from a Node to prepare for maintenance or removal.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-drain
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/drain.md
    title: 'Kubernetes website source: reference/glossary/drain.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-drain
    title: Drain | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/drain.md
---

# Drain

The process of safely evicting [Pods](/reference/glossary/pod.md) from a [Node](/reference/glossary/node.md) to prepare it for maintenance or removal from a [cluster](/reference/glossary/cluster.md).

The `kubectl drain` command is used to mark a [Node](/reference/glossary/node.md) as going out of service. 
When executed, it evicts all [Pods](/reference/glossary/pod.md) from the [Node](/reference/glossary/node.md). 
If an eviction request is temporarily rejected, `kubectl drain` retries until all [Pods](/reference/glossary/pod.md) are terminated or a configurable timeout is reached.
