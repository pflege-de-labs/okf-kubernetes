---
type: glossary-term
title: Disruption
description: An event that leads to Pod(s) going out of service
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-disruption
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/disruption.md
    title: 'Kubernetes website source: reference/glossary/disruption.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-disruption
    title: Disruption | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/disruption.md
---

# Disruption

Disruptions are events that lead to one or more
[Pods](/reference/glossary/pod.md) going out of service.
A disruption has consequences for workload management [resources](/reference/glossary/api-resource.md),
such as [Deployment](/reference/glossary/deployment.md), that rely on the affected
Pods.

If you, as cluster operator, destroy a Pod that belongs to an application,
Kubernetes terms that a _voluntary disruption_. If a Pod goes offline
because of a Node failure, or an outage affecting a wider failure zone,
Kubernetes terms that an _involuntary disruption_.

See [Disruptions](/concepts/workloads/pods/disruptions.md) for more information.
