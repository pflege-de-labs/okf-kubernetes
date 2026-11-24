---
type: glossary-term
title: Common Expression Language
description: An expression language that's designed to be safe for executing user code.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cel
tags: [kubernetes, reference, glossary, extension, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/cel.md
    title: 'Kubernetes website source: reference/glossary/cel.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cel
    title: Common Expression Language | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/cel.md
---

# Common Expression Language

A general-purpose expression language that's designed to be fast, portable, and
safe to execute.

In Kubernetes, CEL can be used to run queries and perform fine-grained
filtering. For example, you can use CEL expressions with
[dynamic admission control](/reference/access-authn-authz/extensible-admission-controllers.md)
to filter for specific fields in requests, and with
[dynamic resource allocation (DRA)](/concepts/resource-management/dynamic-resource-allocation.md)
to select resources based on specific attributes.
