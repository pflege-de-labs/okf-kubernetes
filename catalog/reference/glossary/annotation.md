---
type: glossary-term
title: Annotation
description: A key-value pair that is used to attach arbitrary non-identifying metadata to objects.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-annotation
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/annotation.md
    title: 'Kubernetes website source: reference/glossary/annotation.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-annotation
    title: Annotation | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/annotation.md
---

# Annotation

A key-value pair that is used to attach arbitrary non-identifying metadata to objects.

 

The metadata in an annotation can be small or large, structured or unstructured, and can include characters not permitted by [labels](/reference/glossary/label.md). Clients such as tools and libraries can retrieve this metadata.
