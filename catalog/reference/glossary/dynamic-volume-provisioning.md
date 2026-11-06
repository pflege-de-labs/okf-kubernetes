---
type: glossary-term
title: Dynamic Volume Provisioning
description: Allows users to request automatic creation of storage Volumes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dynamic-volume-provisioning
tags: [kubernetes, reference, glossary, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/dynamic-volume-provisioning.md
    title: 'Kubernetes website source: reference/glossary/dynamic-volume-provisioning.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dynamic-volume-provisioning
    title: Dynamic Volume Provisioning | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/dynamic-volume-provisioning.md
---

# Dynamic Volume Provisioning

Allows users to request automatic creation of storage  [Volumes](/reference/glossary/volume.md).

 

Dynamic provisioning eliminates the need for cluster administrators to pre-provision storage. Instead, it automatically provisions storage by user request. Dynamic volume provisioning is based on an API object, [StorageClass](/reference/glossary/storage-class.md), referring to a [Volume Plugin](/reference/glossary/volume-plugin.md) that provisions a [Volume](/reference/glossary/volume.md) and the set of parameters to pass to the Volume Plugin.
