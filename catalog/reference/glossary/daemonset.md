---
type: glossary-term
title: DaemonSet
description: Ensures a copy of a Pod is running across a set of nodes in a cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/daemonset.md
    title: 'Kubernetes website source: reference/glossary/daemonset.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-daemonset
    title: DaemonSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/daemonset.md
---

# DaemonSet

Ensures a copy of a [Pod](/reference/glossary/pod.md) is running across a set of nodes in a [cluster](/reference/glossary/cluster.md).

 

Used to deploy system daemons such as log collectors and monitoring agents that typically must run on every [Node](/reference/glossary/node.md).
