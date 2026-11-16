---
type: glossary-term
title: Horizontal Pod Autoscaler
description: Object that automatically scales the number of pod replicas based on targeted resource utilization or custom metric targets.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-horizontal-pod-autoscaler
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/horizontal-pod-autoscaler.md
    title: 'Kubernetes website source: reference/glossary/horizontal-pod-autoscaler.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-horizontal-pod-autoscaler
    title: Horizontal Pod Autoscaler | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/horizontal-pod-autoscaler.md
---

# Horizontal Pod Autoscaler

An [object](/reference/glossary/object.md) that automatically scales the number of [Pod](/reference/glossary/pod.md) replicas,
based on targeted [resource](/reference/glossary/infrastructure-resource.md) utilization or custom metric targets.

 

HorizontalPodAutoscaler (HPA) is typically used with [Deployments](/reference/glossary/deployment.md), or [ReplicaSets](/reference/glossary/replica-set.md). It cannot be applied to objects that cannot be scaled, for example [DaemonSets](/reference/glossary/daemonset.md).
