---
type: glossary-term
title: Container Runtime
description: The container runtime is the software that is responsible for running containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-runtime
tags: [kubernetes, reference, glossary, fundamental, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/container-runtime.md
    title: 'Kubernetes website source: reference/glossary/container-runtime.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-runtime
    title: Container Runtime | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/container-runtime.md
---

# Container Runtime

A fundamental component that empowers Kubernetes to run containers effectively.
 It is responsible for managing the execution and lifecycle of containers within the Kubernetes environment.

Kubernetes supports container runtimes such as
[containerd](/reference/glossary/containerd.md), [CRI-O](/reference/glossary/cri-o.md),
and any other implementation of the [Kubernetes CRI (Container Runtime
Interface)](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-node/container-runtime-interface.md).
