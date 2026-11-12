---
type: glossary-term
title: Persistent Volume Claim
description: Claims storage resources defined in a PersistentVolume so that it can be mounted as a volume in a container.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-persistent-volume-claim
tags: [kubernetes, reference, glossary, core-object, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/persistent-volume-claim.md
    title: 'Kubernetes website source: reference/glossary/persistent-volume-claim.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-persistent-volume-claim
    title: Persistent Volume Claim | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/persistent-volume-claim.md
---

# Persistent Volume Claim

Claims storage [resources](/reference/glossary/infrastructure-resource.md) defined in a
[PersistentVolume](/reference/glossary/persistent-volume.md), so that the storage can be mounted as
a volume in a [container](/reference/glossary/container.md).

 

Specifies the amount of storage, how the storage will be accessed (read-only, read-write and/or exclusive) and how it is reclaimed (retained, recycled or deleted). Details of the storage itself are described in the PersistentVolume object.
