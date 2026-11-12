---
type: glossary-term
title: Docker
description: Docker is a software technology providing operating-system-level virtualization also known as containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-docker
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/docker.md
    title: 'Kubernetes website source: reference/glossary/docker.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-docker
    title: Docker | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/docker.md
---

# Docker

Docker (specifically, Docker Engine) is a software technology providing operating-system-level virtualization also known as [containers](/reference/glossary/container.md).

Docker uses the resource isolation features of the Linux kernel such as cgroups and kernel namespaces, and a union-capable file system such as OverlayFS and others to allow independent containers to run within a single Linux instance, avoiding the overhead of starting and maintaining virtual machines (VMs).
