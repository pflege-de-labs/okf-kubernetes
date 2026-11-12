---
type: glossary-term
title: Pod Disruption Budget
description: An object that limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-disruption-budget
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/pod-disruption-budget.md
    title: 'Kubernetes website source: reference/glossary/pod-disruption-budget.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-disruption-budget
    title: Pod Disruption Budget | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/pod-disruption-budget.md
---

# Pod Disruption Budget

A [Pod Disruption Budget](/concepts/workloads/pods/disruptions.md) allows an 
 application owner to create an object for a replicated application, that ensures 
 a certain number or percentage of [Pods](/reference/glossary/pod.md)
 with an assigned label will not be voluntarily evicted at any point in time.

 

Involuntary disruptions cannot be prevented by PDBs; however they 
do count against the budget.
