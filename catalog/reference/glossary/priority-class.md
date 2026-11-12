---
type: glossary-term
title: PriorityClass
description: A mapping from a class name to the scheduling priority that a Pod should have.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-priority-class
tags: [kubernetes, reference, glossary, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/priority-class.md
    title: 'Kubernetes website source: reference/glossary/priority-class.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-priority-class
    title: PriorityClass | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/priority-class.md
---

# PriorityClass

A PriorityClass is a named class for the scheduling priority that should be assigned to a Pod
in that class.

A [PriorityClass](/concepts/scheduling-eviction/pod-priority-preemption.md#how-to-use-priority-and-preemption)
is a non-namespaced object mapping a name to an integer priority, used for a Pod. The name is
specified in the `metadata.name` field, and the priority value in the `value` field. Priorities range from
-2147483648 to 1000000000 inclusive. Higher values indicate higher priority.
