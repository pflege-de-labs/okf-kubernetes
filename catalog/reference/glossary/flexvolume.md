---
type: glossary-term
title: FlexVolume
description: FlexVolume is a deprecated interface for creating out-of-tree volume plugins. The Container Storage Interface is a newer interface that addresses several problems with FlexVolume.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-flexvolume
tags: [kubernetes, reference, glossary, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/flexvolume.md
    title: 'Kubernetes website source: reference/glossary/flexvolume.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-flexvolume
    title: FlexVolume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/flexvolume.md
---

# FlexVolume

FlexVolume is a deprecated interface for creating out-of-tree volume plugins. The [Container Storage Interface](/reference/glossary/csi.md) is a newer interface that addresses several problems with FlexVolume.

 

FlexVolumes enable users to write their own drivers and add support for their volumes in Kubernetes. FlexVolume driver binaries and dependencies must be installed on host machines. This requires root access. The Storage SIG suggests implementing a [CSI](/reference/glossary/csi.md) driver if possible since it addresses the limitations with FlexVolumes.

* [FlexVolume in the Kubernetes documentation](/concepts/storage/volumes.md#flexvolume)
* [More information on FlexVolumes](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-storage/flexvolume.md)
* [Volume Plugin FAQ for Storage Vendors](https://github.com/kubernetes/community/blob/main/sig-storage/volume-plugin-faq.md)
