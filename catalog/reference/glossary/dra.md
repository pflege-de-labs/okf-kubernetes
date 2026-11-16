---
type: glossary-term
title: Dynamic Resource Allocation
description: A Kubernetes feature for requesting and sharing resources, like hardware accelerators, among Pods.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dra
tags: [kubernetes, reference, glossary, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/dra.md
    title: 'Kubernetes website source: reference/glossary/dra.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dra
    title: Dynamic Resource Allocation | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/dra.md
---

# Dynamic Resource Allocation

A Kubernetes feature that lets you request and share resources among Pods.
These resources are often attached
[devices](/reference/glossary/device.md) like hardware
accelerators.

With DRA, device drivers and cluster admins define device _classes_ that are
available to _claim_ in workloads. Kubernetes allocates matching devices to
specific claims and places the corresponding Pods on nodes that can access the
allocated devices.
