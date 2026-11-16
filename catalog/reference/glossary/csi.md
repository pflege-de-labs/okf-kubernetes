---
type: glossary-term
title: Container Storage Interface (CSI)
description: The Container Storage Interface (CSI) defines a standard interface to expose storage systems to containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-csi
tags: [kubernetes, reference, glossary, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/csi.md
    title: 'Kubernetes website source: reference/glossary/csi.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-csi
    title: Container Storage Interface (CSI) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/csi.md
---

# Container Storage Interface (CSI)

The Container Storage Interface (CSI) defines a standard interface to expose storage systems to containers.

 

CSI allows vendors to create custom storage plugins for Kubernetes without adding them to the Kubernetes repository (out-of-tree plugins). To use a CSI driver from a storage provider, you must first [deploy it to your cluster](https://kubernetes-csi.github.io/docs/deploying.html). You will then be able to create a [Storage Class](/reference/glossary/storage-class.md) that uses that CSI driver.

* [CSI in the Kubernetes documentation](/concepts/storage/volumes.md#csi)
* [List of available CSI drivers](https://kubernetes-csi.github.io/docs/drivers.html)
