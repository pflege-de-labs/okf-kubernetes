---
type: glossary-term
title: Cluster
description: A set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cluster
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/cluster.md
    title: 'Kubernetes website source: reference/glossary/cluster.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cluster
    title: Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/cluster.md
---

# Cluster

A set of worker machines, called [nodes](/reference/glossary/node.md),
that run containerized applications. Every cluster has at least one worker node.

The worker node(s) host the [Pods](/reference/glossary/pod.md) that are
the components of the application workload. The
[control plane](/reference/glossary/control-plane.md) manages the worker
nodes and the Pods in the cluster. In production environments, the control plane usually
runs across multiple computers and a cluster usually runs multiple nodes, providing
fault-tolerance and high availability.
