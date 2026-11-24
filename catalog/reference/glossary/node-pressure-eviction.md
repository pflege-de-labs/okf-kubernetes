---
type: glossary-term
title: Node-pressure eviction
description: Node-pressure eviction is the process by which the kubelet proactively fails pods to reclaim resources on nodes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node-pressure-eviction
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/node-pressure-eviction.md
    title: 'Kubernetes website source: reference/glossary/node-pressure-eviction.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node-pressure-eviction
    title: Node-pressure eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/node-pressure-eviction.md
---

# Node-pressure eviction

Node-pressure eviction is the process by which the [kubelet](/reference/glossary/kubelet.md) proactively terminates
pods to reclaim [resource](/reference/glossary/infrastructure-resource.md)
on nodes.

The kubelet monitors resources like CPU, memory, disk space, and filesystem 
inodes on your cluster's nodes. When one or more of these resources reach
specific consumption levels, the kubelet can proactively fail one or more pods
on the node to reclaim resources and prevent starvation. 

Node-pressure eviction is not the same as [API-initiated eviction](/concepts/scheduling-eviction/api-eviction.md).
