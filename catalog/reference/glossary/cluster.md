---
type: glossary-term
title: Cluster
description: A set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cluster
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cluster.md
    title: 'Kubernetes website source: reference/glossary/cluster.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cluster
    title: Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
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
