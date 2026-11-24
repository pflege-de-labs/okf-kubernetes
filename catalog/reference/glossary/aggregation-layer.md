---
type: glossary-term
title: Aggregation Layer
description: The aggregation layer lets you install additional Kubernetes-style APIs in your cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-aggregation-layer
tags: [kubernetes, reference, glossary, architecture, extension, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/aggregation-layer.md
    title: 'Kubernetes website source: reference/glossary/aggregation-layer.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-aggregation-layer
    title: Aggregation Layer | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/aggregation-layer.md
---

# Aggregation Layer

The aggregation layer lets you install additional Kubernetes-style APIs in your cluster.

When you've configured the [Kubernetes API Server](/reference/glossary/kube-apiserver.md) to [support additional APIs](/tasks/extend-kubernetes/configure-aggregation-layer.md), you can add `APIService` objects to "claim" a URL path in the Kubernetes API.
