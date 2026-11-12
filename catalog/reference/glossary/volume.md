---
type: glossary-term
title: Volume
description: A directory containing data, accessible to the containers in a pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/volume.md
    title: 'Kubernetes website source: reference/glossary/volume.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-volume
    title: Volume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/volume.md
---

# Volume

A directory containing data, accessible to the [containers](/reference/glossary/container.md) in a [Pod](/reference/glossary/pod.md).

A Kubernetes volume lives as long as the Pod that encloses it. Consequently, a volume outlives any containers that run within the Pod, and data in the volume is preserved across container restarts.

See [storage](/concepts/storage.md) for more information.
