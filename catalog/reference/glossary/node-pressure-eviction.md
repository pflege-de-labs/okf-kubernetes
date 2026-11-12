---
type: glossary-term
title: Node-pressure eviction
description: Node-pressure eviction is the process by which the kubelet proactively fails pods to reclaim resources on nodes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node-pressure-eviction
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/node-pressure-eviction.md
    title: 'Kubernetes website source: reference/glossary/node-pressure-eviction.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node-pressure-eviction
    title: Node-pressure eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
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
