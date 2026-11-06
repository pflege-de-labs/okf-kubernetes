---
type: glossary-term
title: Volume Plugin
description: A Volume Plugin enables integration of storage within a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume-plugin
tags: [kubernetes, reference, glossary, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/volume-plugin.md
    title: 'Kubernetes website source: reference/glossary/volume-plugin.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume-plugin
    title: Volume Plugin | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/volume-plugin.md
---

# Volume Plugin

A Volume Plugin enables integration of storage within a [Pod](/reference/glossary/pod.md).

 

A Volume Plugin lets you attach and mount storage volumes for use by a [Pod](/reference/glossary/pod.md). Volume plugins can be _in tree_ or _out of tree_. _In tree_ plugins are part of the Kubernetes code repository and follow its release cycle. _Out of tree_ plugins are developed independently.
