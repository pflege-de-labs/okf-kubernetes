---
type: glossary-term
title: Node
description: A node is a worker machine in Kubernetes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node
tags: [kubernetes, reference, glossary, core-object, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/node.md
    title: 'Kubernetes website source: reference/glossary/node.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-node
    title: Node | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/node.md
---

# Node

A node is a worker machine in Kubernetes.

A worker node may be a VM or physical machine, depending on the cluster. It has local daemons or services necessary to run [Pods](/reference/glossary/pod.md) and is managed by the control plane. The daemons on a node include [kubelet](/reference/glossary/kubelet.md), [kube-proxy](/reference/glossary/kube-proxy.md), and a container runtime implementing the [CRI](/reference/glossary/cri.md) such as [Docker](/reference/glossary/docker.md).

In early Kubernetes versions, Nodes were called "Minions".
