---
type: glossary-term
title: Container Runtime
description: The container runtime is the software that is responsible for running containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-runtime
tags: [kubernetes, reference, glossary, fundamental, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/container-runtime.md
    title: 'Kubernetes website source: reference/glossary/container-runtime.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-runtime
    title: Container Runtime | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/container-runtime.md
---

# Container Runtime

A fundamental component that empowers Kubernetes to run containers effectively.
 It is responsible for managing the execution and lifecycle of containers within the Kubernetes environment.

Kubernetes supports container runtimes such as
[containerd](/reference/glossary/containerd.md), [CRI-O](/reference/glossary/cri-o.md),
and any other implementation of the [Kubernetes CRI (Container Runtime
Interface)](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-node/container-runtime-interface.md).
