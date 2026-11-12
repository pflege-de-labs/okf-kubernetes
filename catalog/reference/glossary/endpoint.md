---
type: glossary-term
title: Endpoints
description: (Deprecated) API representing endpoints of a Service
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoint
tags: [kubernetes, reference, glossary, networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/endpoint.md
    title: 'Kubernetes website source: reference/glossary/endpoint.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoint
    title: Endpoints | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/endpoint.md
---

# Endpoints

A deprecated API that represents the set of all endpoints for a
[Service](/reference/glossary/service.md).

Since v1.21, Kubernetes uses 
[EndpointSlices](/reference/glossary/endpoint-slice.md)
rather than Endpoints; the original Endpoints API was deprecated due to
concerns about scalability.

To learn more about Endpoints, read [Endpoints](/concepts/services-networking/service.md#endpoints).
