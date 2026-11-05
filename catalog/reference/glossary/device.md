---
type: glossary-term
title: Device
description: Any resource that's directly or indirectly attached your cluster's nodes, like GPUs or circuit boards.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-device
tags: [kubernetes, reference, glossary, extension, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/device.md
    title: 'Kubernetes website source: reference/glossary/device.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-device
    title: Device | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/device.md
---

# Device

One or more
[infrastructure resources](/reference/glossary/infrastructure-resource.md)
that are directly or indirectly attached to your
[nodes](/reference/glossary/node.md).

Devices might be commercial products like GPUs, or custom hardware like
[ASIC boards](https://en.wikipedia.org/wiki/Application-specific_integrated_circuit).
Attached devices usually require device drivers that let Kubernetes
[Pods](/reference/glossary/pod.md) access the devices.
