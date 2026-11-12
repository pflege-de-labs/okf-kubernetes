---
type: glossary-term
title: EndpointSlice
description: EndpointSlices track the IP addresses of Pods for Services.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoint-slice
tags: [kubernetes, reference, glossary, networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/endpoint-slice.md
    title: 'Kubernetes website source: reference/glossary/endpoint-slice.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-endpoint-slice
    title: EndpointSlice | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/endpoint-slice.md
---

# EndpointSlice

EndpointSlices track the IP addresses of backend endpoints.
EndpointSlices are normally associated with a
[Service](/reference/glossary/service.md) and the backend endpoints typically represent
[Pods](/reference/glossary/pod.md).

One Service can be backed by multiple Pods. Kubernetes represents the backing endpoints of a Service
with a set of EndpointSlices that are associated with that Service.
The backing endpoints are usually, but not always, pods running in the cluster.

The control plane usually manages EndpointSlices for you automatically. However,
EndpointSlices can be defined manually for [Services](/reference/glossary/service.md) without
[selectors](/reference/glossary/selector.md) specified.
