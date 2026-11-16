---
type: glossary-term
title: Storage Class
description: A StorageClass provides a way for administrators to describe different available storage types.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-storage-class
tags: [kubernetes, reference, glossary, core-object, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/storage-class.md
    title: 'Kubernetes website source: reference/glossary/storage-class.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-storage-class
    title: Storage Class | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/storage-class.md
---

# Storage Class

A StorageClass provides a way for administrators to describe different available storage types.

 

StorageClasses can map to quality-of-service levels, backup policies, or to arbitrary policies determined by cluster administrators. Each StorageClass contains the fields `provisioner`, `parameters`, and `reclaimPolicy`, which are used when a [Persistent Volume](/reference/glossary/persistent-volume.md) belonging to the class needs to be dynamically provisioned. Users can request a particular class using the name of a StorageClass object.
