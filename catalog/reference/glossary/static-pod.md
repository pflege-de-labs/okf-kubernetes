---
type: glossary-term
title: Static Pod
description: A pod managed directly by the kubelet daemon on a specific node.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-static-pod
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/static-pod.md
    title: 'Kubernetes website source: reference/glossary/static-pod.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-static-pod
    title: Static Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/static-pod.md
---

# Static Pod

A [pod](/reference/glossary/pod.md) managed directly by the [kubelet](/reference/glossary/kubelet.md)
 daemon on a specific node,

without the API server observing it.

Static Pods do not support [ephemeral containers](/reference/glossary/ephemeral-container.md).
