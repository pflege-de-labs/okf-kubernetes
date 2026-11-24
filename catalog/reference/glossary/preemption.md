---
type: glossary-term
title: Preemption
description: Preemption logic in Kubernetes helps a pending Pod to find a suitable Node by evicting low priority Pods existing on that Node.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-preemption
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/preemption.md
    title: 'Kubernetes website source: reference/glossary/preemption.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-preemption
    title: Preemption | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/preemption.md
---

# Preemption

Preemption logic in Kubernetes helps a pending [Pod](/reference/glossary/pod.md) to find a suitable [Node](/reference/glossary/node.md) by evicting low priority Pods existing on that Node.

If a Pod cannot be scheduled, the scheduler tries to [preempt](/concepts/scheduling-eviction/pod-priority-preemption.md#preemption) lower priority Pods to make scheduling of the pending Pod possible.
