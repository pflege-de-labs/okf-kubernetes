---
type: glossary-term
title: Horizontal Pod Autoscaler
description: Object that automatically scales the number of pod replicas based on targeted resource utilization or custom metric targets.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-horizontal-pod-autoscaler
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/horizontal-pod-autoscaler.md
    title: 'Kubernetes website source: reference/glossary/horizontal-pod-autoscaler.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-horizontal-pod-autoscaler
    title: Horizontal Pod Autoscaler | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/horizontal-pod-autoscaler.md
---

# Horizontal Pod Autoscaler

An [object](/reference/glossary/object.md) that automatically scales the number of [Pod](/reference/glossary/pod.md) replicas,
based on targeted [resource](/reference/glossary/infrastructure-resource.md) utilization or custom metric targets.

 

HorizontalPodAutoscaler (HPA) is typically used with [Deployments](/reference/glossary/deployment.md), or [ReplicaSets](/reference/glossary/replica-set.md). It cannot be applied to objects that cannot be scaled, for example [DaemonSets](/reference/glossary/daemonset.md).
