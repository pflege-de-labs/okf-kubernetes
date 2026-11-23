---
type: glossary-term
title: DeviceClass
description: A category of devices in the cluster. Users can claim specific devices in a DeviceClass.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-deviceclass
tags: [kubernetes, reference, glossary, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/deviceclass.md
    title: 'Kubernetes website source: reference/glossary/deviceclass.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-deviceclass
    title: DeviceClass | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/deviceclass.md
---

# DeviceClass

A category of [devices](/reference/glossary/device.md) in the
 cluster that can be used with dynamic resource allocation (DRA).

Administrators or device owners use DeviceClasses to define a set of devices
that can be claimed and used in workloads. Devices are claimed by creating
[ResourceClaims](/reference/glossary/resourceclaim.md)
that filter for specific device parameters in a DeviceClass.

For more information, see
[Dynamic Resource Allocation](/concepts/resource-management/dynamic-resource-allocation/dra-api.md#deviceclass)
