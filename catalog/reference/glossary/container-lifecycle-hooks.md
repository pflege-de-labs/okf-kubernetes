---
type: glossary-term
title: Container Lifecycle Hooks
description: The lifecycle hooks expose events in the container management lifecycle and let the user run code when the events occur.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-lifecycle-hooks
tags: [kubernetes, reference, glossary, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/container-lifecycle-hooks.md
    title: 'Kubernetes website source: reference/glossary/container-lifecycle-hooks.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-lifecycle-hooks
    title: Container Lifecycle Hooks | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/container-lifecycle-hooks.md
---

# Container Lifecycle Hooks

The lifecycle hooks expose events in the [Container](/reference/glossary/container.md) management lifecycle and let the user run code when the events occur.

Two hooks are exposed to Containers: PostStart which executes immediately after a container is created and PreStop which is blocking and is called immediately before a container is terminated.
