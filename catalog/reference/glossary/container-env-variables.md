---
type: glossary-term
title: Container Environment Variables
description: Container environment variables are name=value pairs that provide useful information into containers running in a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-env-variables
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/container-env-variables.md
    title: 'Kubernetes website source: reference/glossary/container-env-variables.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-container-env-variables
    title: Container Environment Variables | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/container-env-variables.md
---

# Container Environment Variables

Container environment variables are name=value pairs that provide useful information into containers running in a [pod](/reference/glossary/pod.md)

Container environment variables provide information that is required by the running containerized applications along with information about important related details to the [containers](/reference/glossary/container.md). For example, file system details, information about the container itself, and other cluster resources such as service endpoints.
