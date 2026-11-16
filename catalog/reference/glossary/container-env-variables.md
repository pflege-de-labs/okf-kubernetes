---
type: glossary-term
title: Container Environment Variables
description: Container environment variables are name=value pairs that provide useful information into containers running in a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-env-variables
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/container-env-variables.md
    title: 'Kubernetes website source: reference/glossary/container-env-variables.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-env-variables
    title: Container Environment Variables | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/container-env-variables.md
---

# Container Environment Variables

Container environment variables are name=value pairs that provide useful information into containers running in a [pod](/reference/glossary/pod.md)

Container environment variables provide information that is required by the running containerized applications along with information about important related details to the [containers](/reference/glossary/container.md). For example, file system details, information about the container itself, and other cluster resources such as service endpoints.
