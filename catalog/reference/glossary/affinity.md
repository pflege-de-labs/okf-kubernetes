---
type: glossary-term
title: Affinity
description: Rules used by the scheduler to determine where to place pods
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-affinity
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/affinity.md
    title: 'Kubernetes website source: reference/glossary/affinity.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-affinity
    title: Affinity | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/affinity.md
---

# Affinity

In Kubernetes, _affinity_ is a set of rules that give hints to the scheduler about where to place pods.

There are two kinds of affinity:
* [node affinity](/concepts/scheduling-eviction/assign-pod-node.md#node-affinity)
* [pod-to-pod affinity](/concepts/scheduling-eviction/assign-pod-node.md#inter-pod-affinity-and-anti-affinity)

The rules are defined using the Kubernetes [labels](/reference/glossary/label.md),
and [selectors](/reference/glossary/selector.md) specified in [pods](/reference/glossary/pod.md), 
and they can be either required or preferred, depending on how strictly you want the scheduler to enforce them.
