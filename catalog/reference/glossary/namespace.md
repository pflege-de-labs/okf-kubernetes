---
type: glossary-term
title: Namespace
description: An abstraction used by Kubernetes to support isolation of groups of resources within a single cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-namespace
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/namespace.md
    title: 'Kubernetes website source: reference/glossary/namespace.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-namespace
    title: Namespace | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/namespace.md
---

# Namespace

An abstraction used by Kubernetes to support isolation of groups of [API resources](/reference/glossary/api-resource.md)
within a single [cluster](/reference/glossary/cluster.md).

 

Namespaces are used to organize objects in a cluster and provide a way to divide cluster resources. Names of resources need to be unique within a namespace, but not across namespaces. Namespace-based scoping is applicable only for namespaced resources _(for example: Pods, Deployments, Services)_ and not for cluster-wide resources _(for example: StorageClasses, Nodes, PersistentVolumes)_.
