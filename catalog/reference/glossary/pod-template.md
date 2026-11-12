---
type: glossary-term
title: PodTemplate
description: A template for creating Pods.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-template
tags: [kubernetes, reference, glossary, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/pod-template.md
    title: 'Kubernetes website source: reference/glossary/pod-template.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-template
    title: PodTemplate | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/pod-template.md
---

# PodTemplate

An API object that defines a template for creating [Pods](/reference/glossary/pod.md).
The PodTemplate API is also embedded in API definitions for workload management, such as 
[Deployment](/reference/glossary/deployment.md) or
[StatefulSets](/reference/glossary/statefulset.md).

 

Pod templates allow you to define common metadata (such as labels, or a template for the name of a
new Pod) as well as to specify a pod's desired state.
[Workload management](/concepts/workloads/controllers.md) controllers use Pod templates
(embedded into another object, such as a Deployment or StatefulSet)
to define and manage one or more [Pods](/reference/glossary/pod.md).
When there can be multiple Pods based on the same template, these are called
[replicas](/reference/glossary/replica.md).
Although you can create a PodTemplate object directly, you rarely need to do so.
