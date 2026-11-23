---
type: glossary-term
title: Persistent Volume
description: API object that represents a piece of storage in the cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-persistent-volume
tags: [kubernetes, reference, glossary, core-object, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/persistent-volume.md
    title: 'Kubernetes website source: reference/glossary/persistent-volume.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-persistent-volume
    title: Persistent Volume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/persistent-volume.md
---

# Persistent Volume

An API object that represents a piece of storage in the cluster. Representation of as a general, pluggable storage
[resource](/reference/glossary/infrastructure-resource.md) that can persist beyond the lifecycle of any
individual [Pod](/reference/glossary/pod.md).

 

PersistentVolumes (PVs) provide an API that abstracts details of how storage is provided from how it is consumed.
PVs are used directly in scenarios where storage can be created ahead of time (static provisioning).
For scenarios that require on-demand storage (dynamic provisioning), PersistentVolumeClaims (PVCs) are used instead.
