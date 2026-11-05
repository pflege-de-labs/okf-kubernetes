---
type: glossary-term
title: Condition
description: A condition represents the current state of a Kubernetes resource, providing information about whether certain aspects of the resource are true.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-condition
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/condition.md
    title: 'Kubernetes website source: reference/glossary/condition.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-condition
    title: Condition | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/condition.md
---

# Condition

A condition is a field in a Kubernetes resource's status that describes the current state of that resource.

Conditions provide a standardized way for Kubernetes components to communicate the status of resources. Each condition has a `type`, a `status` (True, False, or Unknown), and optional fields like `reason` and `message` that provide additional details. For example, a Pod might have conditions like `Ready`, `ContainersReady`, or `PodScheduled`.
