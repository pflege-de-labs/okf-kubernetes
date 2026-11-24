---
type: glossary-term
title: LimitRange
description: Provides constraints to limit resource consumption per Containers or Pods in a namespace.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-limitrange
tags: [kubernetes, reference, glossary, core-object, fundamental, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/limitrange.md
    title: 'Kubernetes website source: reference/glossary/limitrange.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-limitrange
    title: LimitRange | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/limitrange.md
---

# LimitRange

Constraints resource consumption per [container](/reference/glossary/container.md) or [Pod](/reference/glossary/pod.md),
specified for a particular [namespace](/reference/glossary/namespace.md).

 

A [LimitRange](/concepts/policy/limit-range.md) either limits the quantity of [API resources](/reference/glossary/api-resource.md)
that can be created (for a particular resource type),
or the amount of [infrastructure resources](/reference/glossary/infrastructure-resource.md)
that may be requested/consumed by individual containers or Pods within a namespace.
