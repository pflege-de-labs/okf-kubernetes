---
type: glossary-term
title: DaemonSet
description: Ensures a copy of a Pod is running across a set of nodes in a cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/daemonset.md
    title: 'Kubernetes website source: reference/glossary/daemonset.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
    title: DaemonSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/daemonset.md
---

# DaemonSet

Ensures a copy of a [Pod](/reference/glossary/pod.md) is running across a set of nodes in a [cluster](/reference/glossary/cluster.md).

 

Used to deploy system daemons such as log collectors and monitoring agents that typically must run on every [Node](/reference/glossary/node.md).
