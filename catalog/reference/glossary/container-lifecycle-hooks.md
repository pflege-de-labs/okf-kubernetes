---
type: glossary-term
title: Container Lifecycle Hooks
description: The lifecycle hooks expose events in the container management lifecycle and let the user run code when the events occur.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-lifecycle-hooks
tags: [kubernetes, reference, glossary, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/container-lifecycle-hooks.md
    title: 'Kubernetes website source: reference/glossary/container-lifecycle-hooks.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-lifecycle-hooks
    title: Container Lifecycle Hooks | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/container-lifecycle-hooks.md
---

# Container Lifecycle Hooks

The lifecycle hooks expose events in the [Container](/reference/glossary/container.md) management lifecycle and let the user run code when the events occur.

Two hooks are exposed to Containers: PostStart which executes immediately after a container is created and PreStop which is blocking and is called immediately before a container is terminated.
