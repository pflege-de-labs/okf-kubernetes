---
type: glossary-term
title: Garbage Collection
description: A collective term for the various mechanisms Kubernetes uses to clean up cluster resources.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-garbage-collection
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/garbage-collection.md
    title: 'Kubernetes website source: reference/glossary/garbage-collection.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-garbage-collection
    title: Garbage Collection | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/garbage-collection.md
---

# Garbage Collection

Garbage collection is a collective term for the various mechanisms Kubernetes uses to clean up
cluster resources. 

Kubernetes uses garbage collection to clean up resources like
[unused containers and images](/concepts/architecture/garbage-collection.md#containers-images),
[failed Pods](/concepts/workloads/pods/pod-lifecycle.md#pod-garbage-collection),
[objects owned by the targeted resource](/concepts/overview/working-with-objects/owners-dependents.md),
[completed Jobs](/concepts/workloads/controllers/ttlafterfinished.md), and resources
that have expired or failed.
