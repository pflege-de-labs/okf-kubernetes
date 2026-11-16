---
type: glossary-term
title: Dockershim
description: A component of Kubernetes v1.23 and earlier, which allows Kubernetes system components to communicate with Docker Engine.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dockershim
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/dockershim.md
    title: 'Kubernetes website source: reference/glossary/dockershim.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dockershim
    title: Dockershim | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/dockershim.md
---

# Dockershim

The dockershim is a component of Kubernetes version 1.23 and earlier. It allows the [kubelet](/reference/glossary/kubelet.md)
to communicate with [Docker Engine](/reference/glossary/docker.md).

Starting with version 1.24, dockershim has been removed from Kubernetes. For more information, see [Dockershim FAQ](https://kubernetes.io/dockershim).
