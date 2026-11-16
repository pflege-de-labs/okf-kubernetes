---
type: glossary-term
title: ResourceQuota
description: Provides constraints that limit aggregate resource consumption per namespace.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resource-quota
tags: [kubernetes, reference, glossary, fundamental, operation, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/resource-quota.md
    title: 'Kubernetes website source: reference/glossary/resource-quota.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resource-quota
    title: ResourceQuota | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/resource-quota.md
---

# ResourceQuota

Object that constrains aggregate resource
consumption, per [Namespace](/reference/glossary/namespace.md).

A ResourceQuota can either limit the quantity of [API resources](/reference/glossary/api-resource.md)
that can be created in a namespace by type, or it can set a limit on the total amount of
[infrastructure resources](/reference/glossary/infrastructure-resource.md)
that may be consumed on behalf of the namespace (and the objects within it).
