---
type: glossary-term
title: ResourceQuota
description: Provides constraints that limit aggregate resource consumption per namespace.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resource-quota
tags: [kubernetes, reference, glossary, fundamental, operation, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/resource-quota.md
    title: 'Kubernetes website source: reference/glossary/resource-quota.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resource-quota
    title: ResourceQuota | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/resource-quota.md
---

# ResourceQuota

Object that constrains aggregate resource
consumption, per [Namespace](/reference/glossary/namespace.md).

A ResourceQuota can either limit the quantity of [API resources](/reference/glossary/api-resource.md)
that can be created in a namespace by type, or it can set a limit on the total amount of
[infrastructure resources](/reference/glossary/infrastructure-resource.md)
that may be consumed on behalf of the namespace (and the objects within it).
