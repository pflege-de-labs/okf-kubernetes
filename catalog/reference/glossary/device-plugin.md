---
type: glossary-term
title: Device Plugin
description: Software extensions to let Pods access devices that need vendor-specific initialization or setup
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-device-plugin
tags: [kubernetes, reference, glossary, fundamental, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/device-plugin.md
    title: 'Kubernetes website source: reference/glossary/device-plugin.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-device-plugin
    title: Device Plugin | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/device-plugin.md
---

# Device Plugin

Device plugins run on worker
[Nodes](/reference/glossary/node.md) and provide
[Pods](/reference/glossary/pod.md) with access to
infrastructure [resources](/reference/glossary/infrastructure-resource.md),
such as local hardware, that require vendor-specific initialization or setup
steps.

Device plugins advertise resources to the
[kubelet](/reference/glossary/kubelet.md), so that workload
Pods can access hardware features that relate to the Node where that Pod is running.
You can deploy a device plugin as a [DaemonSet](/reference/glossary/daemonset.md),
or install the device plugin software directly on each target Node.

See
[Device Plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md)
for more information.
