---
type: glossary-term
title: Static Pod
description: A pod managed directly by the kubelet daemon on a specific node.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-static-pod
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/static-pod.md
    title: 'Kubernetes website source: reference/glossary/static-pod.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-static-pod
    title: Static Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/static-pod.md
---

# Static Pod

A [pod](/reference/glossary/pod.md) managed directly by the [kubelet](/reference/glossary/kubelet.md)
 daemon on a specific node,

without the API server observing it.

Static Pods do not support [ephemeral containers](/reference/glossary/ephemeral-container.md).
