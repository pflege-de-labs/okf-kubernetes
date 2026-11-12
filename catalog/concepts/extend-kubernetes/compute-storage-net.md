---
type: section
title: Compute, Storage, and Networking Extensions
description: This section covers extensions to your cluster that do not come as part as Kubernetes itself.
resource: https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/
tags: [kubernetes, concepts, extend-kubernetes]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/extend-kubernetes/compute-storage-net/_index.md
    title: 'Kubernetes website source: concepts/extend-kubernetes/compute-storage-net/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/
    title: Compute, Storage, and Networking Extensions | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/extend-kubernetes/compute-storage-net/_index.md
---

# Compute, Storage, and Networking Extensions

This section covers extensions to your cluster that do not come as part as Kubernetes itself.
You can use these extensions to enhance the nodes in your cluster, or to provide the network
fabric that links Pods together.

* [CSI](/concepts/storage/volumes.md#csi) and [FlexVolume](/concepts/storage/volumes.md#flexvolume) storage plugins

  [Container Storage Interface](/reference/glossary/csi.md) (CSI) plugins
  provide a way to extend Kubernetes with supports for new kinds of volumes. The volumes can
  be backed by durable external storage, or provide ephemeral storage, or they might offer a
  read-only interface to information using a filesystem paradigm.

  Kubernetes also includes support for [FlexVolume](/concepts/storage/volumes.md#flexvolume)
  plugins, which are deprecated since Kubernetes v1.23 (in favour of CSI).

  FlexVolume plugins allow users to mount volume types that aren't natively
  supported by Kubernetes. When you run a Pod that relies on FlexVolume
  storage, the kubelet calls a binary plugin to mount the volume. The archived
  [FlexVolume](https://git.k8s.io/design-proposals-archive/storage/flexvolume-deployment.md)
  design proposal has more detail on this approach.

  The [Kubernetes Volume Plugin FAQ for Storage Vendors](https://github.com/kubernetes/community/blob/main/sig-storage/volume-plugin-faq.md#kubernetes-volume-plugin-faq-for-storage-vendors)
  includes general information on storage plugins.

* [Device plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md)

  Device plugins allow a node to discover new Node facilities (in addition to the
  built-in node resources such as `cpu` and `memory`), and provide these custom node-local
  facilities to Pods that request them.

* [Network plugins](/concepts/extend-kubernetes/compute-storage-net/network-plugins.md)

  Network plugins allow Kubernetes to work with different networking topologies and technologies.
  Your Kubernetes cluster needs a _network plugin_ in order to have a working Pod network
  and to support other aspects of the Kubernetes network model.

  Kubernetes 1.37 is compatible with [CNI](/reference/glossary/cni.md)
  network plugins.
