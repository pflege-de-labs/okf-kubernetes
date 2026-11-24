---
type: glossary-term
title: API resource
description: A Kubernetes entity, representing an endpoint on the Kubernetes API server.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-resource
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/api-resource.md
    title: 'Kubernetes website source: reference/glossary/api-resource.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-resource
    title: API resource | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/api-resource.md
---

# API resource

An entity in the Kubernetes type system, corresponding to an endpoint on the [Kubernetes API](/reference/glossary/kubernetes-api.md).
A resource typically represents an [object](/reference/glossary/object.md).
Some resources represent an operation on other objects, such as a permission check.

Each resource represents an HTTP endpoint (URI) on the Kubernetes API server, defining the schema for the objects or operations on that resource.
