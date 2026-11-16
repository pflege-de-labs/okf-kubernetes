---
type: glossary-term
title: kube-scheduler
description: Control plane component that watches for newly created pods with no assigned node, and selects a node for them to run on.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-scheduler
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/kube-scheduler.md
    title: 'Kubernetes website source: reference/glossary/kube-scheduler.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-scheduler
    title: kube-scheduler | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/kube-scheduler.md
---

# kube-scheduler

Control plane component that watches for newly created
[Pods](/reference/glossary/pod.md) with no assigned
[node](/reference/glossary/node.md), and selects a node for them
to run on.

Factors taken into account for scheduling decisions include:
individual and collective [resource](/reference/glossary/infrastructure-resource.md)
requirements, hardware/software/policy constraints, affinity and anti-affinity specifications,
data locality, inter-workload interference, and deadlines.
