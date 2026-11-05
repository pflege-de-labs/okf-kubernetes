---
type: glossary-term
title: Ephemeral Container
description: A type of container that you can temporarily run inside a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-ephemeral-container
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/ephemeral-container.md
    title: 'Kubernetes website source: reference/glossary/ephemeral-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-ephemeral-container
    title: Ephemeral Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/ephemeral-container.md
---

# Ephemeral Container

A [Container](/reference/glossary/container.md) type that you can temporarily run inside a [Pod](/reference/glossary/pod.md).

If you want to investigate a Pod that's running with problems, you can add an ephemeral container to that Pod and carry out diagnostics.
Ephemeral containers have no [resource](/reference/glossary/infrastructure-resource.md) or scheduling guarantees,
and you should not use them to run any part of the workload itself.

Ephemeral containers are not supported by [static pods](/reference/glossary/static-pod.md).
