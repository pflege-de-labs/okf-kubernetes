---
type: glossary-term
title: Watch
description: A verb that is used to track changes to an object in Kubernetes as a stream.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-watch
tags: [kubernetes, reference, glossary, API verb, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/watch.md
    title: 'Kubernetes website source: reference/glossary/watch.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-watch
    title: Watch | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/watch.md
---

# Watch

A verb that is used to track changes to an object in Kubernetes as a stream.
It is used for the efficient detection of changes.

A verb that is used to track changes to an object in Kubernetes as a stream. Watches allow
efficient detection of changes; for example, a
[controller](/reference/glossary/controller.md) that needs to know whenever a
ConfigMap has changed can use a watch rather than polling.

See [Efficient Detection of Changes in API Concepts](/reference/using-api/api-concepts.md#efficient-detection-of-changes) for more information.
