---
type: glossary-term
title: ReplicationController
description: A (deprecated) API object that manages a replicated application.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replication-controller
tags: [kubernetes, reference, glossary, workload, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/replication-controller.md
    title: 'Kubernetes website source: reference/glossary/replication-controller.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-replication-controller
    title: ReplicationController | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/replication-controller.md
---

# ReplicationController

A workload management [object](/reference/glossary/object.md)
that manages a replicated application, ensuring that
a specific number of instances of a [Pod](/reference/glossary/pod.md) are running.

The control plane ensures that the defined number of Pods are running, even if some
Pods fail, if you delete Pods manually, or if too many are started by mistake.

**Note:**
ReplicationController is deprecated. See
[Deployment](/reference/glossary/deployment.md), which is similar.
