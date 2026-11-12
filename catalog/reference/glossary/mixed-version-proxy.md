---
type: glossary-term
title: Mixed Version Proxy (MVP)
description: Feature that lets a kube-apiserver proxy a resource request to a different peer API server.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-mixed-version-proxy
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/mixed-version-proxy.md
    title: 'Kubernetes website source: reference/glossary/mixed-version-proxy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-mixed-version-proxy
    title: Mixed Version Proxy (MVP) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/mixed-version-proxy.md
---

# Mixed Version Proxy (MVP)

Feature to let a kube-apiserver proxy a resource request to a different peer API server.

When a cluster has multiple API servers running different versions of Kubernetes, this
feature enables [resource](/reference/glossary/api-resource.md)
requests to be served by the correct API server.

MVP is disabled by default and can be activated by enabling
the [feature gate](/reference/command-line-tools-reference/feature-gates.md) named `UnknownVersionInteroperabilityProxy` when 
the [API Server](/reference/glossary/kube-apiserver.md) is started.
