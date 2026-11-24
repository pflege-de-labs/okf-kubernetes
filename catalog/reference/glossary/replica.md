---
type: glossary-term
title: Replica
description: Replicas are copies of pods, ensuring availability, scalability, and fault tolerance by maintaining identical instances.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replica
tags: [kubernetes, reference, glossary, fundamental, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/replica.md
    title: 'Kubernetes website source: reference/glossary/replica.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replica
    title: Replica | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/replica.md
---

# Replica

A copy or duplicate of a [Pod](/reference/glossary/pod.md) or
a set of pods. Replicas ensure high availability, scalability, and fault tolerance
by maintaining multiple identical instances of a pod.

Replicas are commonly used in Kubernetes to achieve the desired application state and reliability.
They enable workload scaling and distribution across multiple nodes in a cluster.

By defining the number of replicas in a Deployment or ReplicaSet, Kubernetes ensures that
the specified number of instances are running, automatically adjusting the count as needed.

Replica management allows for efficient load balancing, rolling updates, and
self-healing capabilities in a Kubernetes cluster.
