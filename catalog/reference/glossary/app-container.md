---
type: glossary-term
title: App Container
description: A container used to run part of a workload. Compare with init container.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-app-container
tags: [kubernetes, reference, glossary, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/app-container.md
    title: 'Kubernetes website source: reference/glossary/app-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-app-container
    title: App Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/app-container.md
---

# App Container

Application containers (or app containers) are the [containers](/reference/glossary/container.md) in a [pod](/reference/glossary/pod.md) that are started after any [init containers](/reference/glossary/init-container.md) have completed.

An init container lets you separate initialization details that are important for the overall 
[workload](/reference/glossary/workload.md), and that don't need to keep running
once the application container has started.
If a pod doesn't have any init containers configured, all the containers in that pod are app containers.
