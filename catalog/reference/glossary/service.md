---
type: glossary-term
title: Service
description: A way to expose an application running on a set of Pods as a network service.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-service
tags: [kubernetes, reference, glossary, fundamental, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/service.md
    title: 'Kubernetes website source: reference/glossary/service.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-service
    title: Service | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/service.md
---

# Service

A method for exposing a network application that is running as one or more
[Pods](/reference/glossary/pod.md) in your cluster.

The set of Pods targeted by a Service is (usually) determined by a
[selector](/reference/glossary/selector.md). If more Pods are added or removed,
the set of Pods matching the selector will change. The Service makes sure that network traffic
can be directed to the current set of Pods for the workload.

Kubernetes Services either use IP networking (IPv4, IPv6, or both), or reference an external name in
the Domain Name System (DNS).

The Service abstraction enables other mechanisms, such as Ingress and Gateway.
