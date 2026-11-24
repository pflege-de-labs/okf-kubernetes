---
type: glossary-term
title: Namespace
description: An abstraction used by Kubernetes to support isolation of groups of resources within a single cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-namespace
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/namespace.md
    title: 'Kubernetes website source: reference/glossary/namespace.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-namespace
    title: Namespace | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/namespace.md
---

# Namespace

An abstraction used by Kubernetes to support isolation of groups of [API resources](/reference/glossary/api-resource.md)
within a single [cluster](/reference/glossary/cluster.md).

 

Namespaces are used to organize objects in a cluster and provide a way to divide cluster resources. Names of resources need to be unique within a namespace, but not across namespaces. Namespace-based scoping is applicable only for namespaced resources _(for example: Pods, Deployments, Services)_ and not for cluster-wide resources _(for example: StorageClasses, Nodes, PersistentVolumes)_.
