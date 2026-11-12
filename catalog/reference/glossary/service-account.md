---
type: glossary-term
title: ServiceAccount
description: Provides an identity for processes that run in a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-service-account
tags: [kubernetes, reference, glossary, fundamental, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/service-account.md
    title: 'Kubernetes website source: reference/glossary/service-account.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-service-account
    title: ServiceAccount | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/service-account.md
---

# ServiceAccount

Provides an identity for processes that run in a [Pod](/reference/glossary/pod.md).

 

When processes inside Pods access the cluster, they are authenticated by the API server as a particular service account, for example, `default`. When you create a Pod, if you do not specify a service account, it is automatically assigned the default service account in the same [Namespace](/reference/glossary/namespace.md).
