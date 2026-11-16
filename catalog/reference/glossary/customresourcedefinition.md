---
type: glossary-term
title: CustomResourceDefinition
description: Custom code that defines a resource to add to your Kubernetes API server without building a complete custom server.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-customresourcedefinition
tags: [kubernetes, reference, glossary, fundamental, operation, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/customresourcedefinition.md
    title: 'Kubernetes website source: reference/glossary/customresourcedefinition.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-customresourcedefinition
    title: CustomResourceDefinition | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/customresourcedefinition.md
---

# CustomResourceDefinition

A kind of [API object](/reference/glossary/object.md) that defines a new custom API to add
to your Kubernetes API server, without building a complete custom server.

CustomResourceDefinitions let you extend the Kubernetes API for your environment if the built-in
[API resources](/reference/glossary/api-resource.md) can't meet your needs.
