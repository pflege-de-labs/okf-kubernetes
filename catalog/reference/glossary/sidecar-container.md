---
type: glossary-term
title: Sidecar Container
description: An auxiliary container that stays running throughout the lifecycle of a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-sidecar-container
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/sidecar-container.md
    title: 'Kubernetes website source: reference/glossary/sidecar-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-sidecar-container
    title: Sidecar Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/sidecar-container.md
---

# Sidecar Container

One or more [containers](/reference/glossary/container.md) that are typically started before any app containers run.

 

Sidecar containers are like regular app containers, but with a different purpose: the sidecar provides a Pod-local service to the main app container.
Unlike [init containers](/reference/glossary/init-container.md), sidecar containers
continue running after Pod startup.

Read [Sidecar containers](/concepts/workloads/pods/sidecar-containers.md) for more information.
