---
type: glossary-term
title: Eviction
description: Process of terminating one or more Pods on Nodes
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-eviction
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/eviction.md
    title: 'Kubernetes website source: reference/glossary/eviction.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-eviction
    title: Eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/eviction.md
---

# Eviction

Eviction is the process of terminating one or more Pods on Nodes.

There are two kinds of eviction:
* [Node-pressure eviction](/concepts/scheduling-eviction/node-pressure-eviction.md)
* [API-initiated eviction](/concepts/scheduling-eviction/api-eviction.md)
