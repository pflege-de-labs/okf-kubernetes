---
type: glossary-term
title: kube-scheduler
description: Control plane component that watches for newly created pods with no assigned node, and selects a node for them to run on.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-scheduler
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/kube-scheduler.md
    title: 'Kubernetes website source: reference/glossary/kube-scheduler.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-scheduler
    title: kube-scheduler | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
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
