---
type: glossary-term
title: DaemonSet
description: Ensures a copy of a Pod is running across a set of nodes in a cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/daemonset.md
    title: 'Kubernetes website source: reference/glossary/daemonset.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
    title: DaemonSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/daemonset.md
---

# DaemonSet

Ensures a copy of a [Pod](/reference/glossary/pod.md) is running across a set of nodes in a [cluster](/reference/glossary/cluster.md).

 

Used to deploy system daemons such as log collectors and monitoring agents that typically must run on every [Node](/reference/glossary/node.md).
