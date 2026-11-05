---
type: glossary-term
title: Taint
description: 'A core object consisting of three required properties: key, value, and effect. Taints prevent the scheduling of pods on nodes or node groups.'
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-taint
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/taint.md
    title: 'Kubernetes website source: reference/glossary/taint.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-taint
    title: Taint | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/taint.md
---

# Taint

A core object consisting of three required properties: key, value, and effect. Taints prevent the scheduling of [Pods](/reference/glossary/pod.md) on [nodes](/reference/glossary/node.md) or node groups.

Taints and [tolerations](/reference/glossary/toleration.md) work together to ensure that pods are not scheduled onto inappropriate nodes. One or more taints are applied to a node. A node should only schedule a Pod with the matching tolerations for the configured taints.
