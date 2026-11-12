---
type: glossary-term
title: Deployment
description: Manages a replicated application on your cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-deployment
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/deployment.md
    title: 'Kubernetes website source: reference/glossary/deployment.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-deployment
    title: Deployment | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/deployment.md
---

# Deployment

An API object that manages a replicated application, typically by running Pods with no local state.

 

Each replica is represented by a [Pod](/reference/glossary/pod.md), and the Pods are distributed among the 
[nodes](/reference/glossary/node.md) of a cluster.
For workloads that do require local state, consider using a [StatefulSet](/reference/glossary/statefulset.md).
