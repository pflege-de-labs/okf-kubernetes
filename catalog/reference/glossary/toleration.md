---
type: glossary-term
title: Toleration
description: 'A core object consisting of three required properties: key, value, and effect. Tolerations enable the scheduling of pods on nodes or node groups that have a matching taint.'
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-toleration
tags: [kubernetes, reference, glossary, core-object, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/toleration.md
    title: 'Kubernetes website source: reference/glossary/toleration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-toleration
    title: Toleration | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/toleration.md
---

# Toleration

A core object consisting of three required properties: key, value, and effect. Tolerations enable the scheduling of pods on nodes or node groups that have matching [taints](/reference/glossary/taint.md).

Tolerations and [taints](/reference/glossary/taint.md) work together to ensure that pods are not scheduled onto inappropriate nodes. One or more tolerations are applied to a [pod](/reference/glossary/pod.md). A toleration indicates that the [pod](/reference/glossary/pod.md) is allowed (but not required) to be scheduled on nodes or node groups with matching [taints](/reference/glossary/taint.md).
