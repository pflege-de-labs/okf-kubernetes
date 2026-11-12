---
type: glossary-term
title: ReplicaSet
description: ReplicaSet ensures that a specified number of Pod replicas are running at one time
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replica-set
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/replica-set.md
    title: 'Kubernetes website source: reference/glossary/replica-set.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replica-set
    title: ReplicaSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/replica-set.md
---

# ReplicaSet

A ReplicaSet (aims to) maintain a set of replica Pods running at any given time.

Workload objects such as [Deployment](/reference/glossary/deployment.md) make use of ReplicaSets
to ensure that the configured number of [Pods](/reference/glossary/pod.md) are
running in your cluster, based on the spec of that ReplicaSet.
