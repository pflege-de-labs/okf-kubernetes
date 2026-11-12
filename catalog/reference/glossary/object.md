---
type: glossary-term
title: Object
description: An entity in the Kubernetes system, representing part of the state of your cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-object
tags: [kubernetes, reference, glossary, architecture, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/object.md
    title: 'Kubernetes website source: reference/glossary/object.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-object
    title: Object | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/object.md
---

# Object

An entity in the Kubernetes system. An object is an
[API resource](/reference/glossary/api-resource.md) that the Kubernetes API
uses to represent the state of your cluster.

A Kubernetes object is typically a “record of intent”—once you create the object, the Kubernetes
[control plane](/reference/glossary/control-plane.md) works constantly to ensure
that the item it represents actually exists.
By creating an object, you're effectively telling the Kubernetes system what you want that part of
your cluster's workload to look like; this is your cluster's desired state.
