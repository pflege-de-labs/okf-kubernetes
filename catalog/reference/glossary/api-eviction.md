---
type: glossary-term
title: API-initiated eviction
description: API-initiated eviction is the process by which you use the Eviction API to create an Eviction object that triggers graceful pod termination.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-eviction
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/api-eviction.md
    title: 'Kubernetes website source: reference/glossary/api-eviction.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-eviction
    title: API-initiated eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/api-eviction.md
---

# API-initiated eviction

API-initiated eviction is the process by which you use the [Eviction API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#create-eviction-pod-v1-core)
to create an `Eviction` object that triggers graceful pod termination.

You can request eviction either by directly calling the Eviction API 
using a client of the kube-apiserver, like the `kubectl drain` command. 
When an `Eviction` object is created, the API server terminates the Pod. 

API-initiated evictions respect your configured [`PodDisruptionBudgets`](/tasks/run-application/configure-pdb.md)
and [`terminationGracePeriodSeconds`](/concepts/workloads/pods/pod-lifecycle.md#pod-termination).

API-initiated eviction is not the same as [node-pressure eviction](/concepts/scheduling-eviction/node-pressure-eviction.md).

* See [API-initiated eviction](/concepts/scheduling-eviction/api-eviction.md) for more information.
