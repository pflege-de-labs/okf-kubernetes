---
type: glossary-term
title: Dockershim
description: A component of Kubernetes v1.23 and earlier, which allows Kubernetes system components to communicate with Docker Engine.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dockershim
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/dockershim.md
    title: 'Kubernetes website source: reference/glossary/dockershim.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-dockershim
    title: Dockershim | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/dockershim.md
---

# Dockershim

The dockershim is a component of Kubernetes version 1.23 and earlier. It allows the [kubelet](/reference/glossary/kubelet.md)
to communicate with [Docker Engine](/reference/glossary/docker.md).

Starting with version 1.24, dockershim has been removed from Kubernetes. For more information, see [Dockershim FAQ](https://kubernetes.io/dockershim).
