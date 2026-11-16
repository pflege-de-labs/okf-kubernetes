---
type: glossary-term
title: Pod
description: A Pod represents a set of running containers in your cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod
tags: [kubernetes, reference, glossary, core-object, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/pod.md
    title: 'Kubernetes website source: reference/glossary/pod.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod
    title: Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/pod.md
---

# Pod

The smallest and simplest Kubernetes object. A Pod represents a set of running [containers](/reference/glossary/container.md) on your cluster.

 

A Pod is typically set up to run a single primary container. It can also run optional sidecar containers that add supplementary features like logging. Pods are commonly managed by a [Deployment](/reference/glossary/deployment.md).
