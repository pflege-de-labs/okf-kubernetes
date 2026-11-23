---
type: glossary-term
title: Volume
description: A directory containing data, accessible to the containers in a pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/volume.md
    title: 'Kubernetes website source: reference/glossary/volume.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume
    title: Volume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/volume.md
---

# Volume

A directory containing data, accessible to the [containers](/reference/glossary/container.md) in a [Pod](/reference/glossary/pod.md).

A Kubernetes volume lives as long as the Pod that encloses it. Consequently, a volume outlives any containers that run within the Pod, and data in the volume is preserved across container restarts.

See [storage](/concepts/storage.md) for more information.
