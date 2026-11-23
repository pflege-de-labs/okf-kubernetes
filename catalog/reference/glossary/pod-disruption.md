---
type: glossary-term
title: Pod Disruption
description: The process by which Pods on Nodes are terminated either voluntarily or involuntarily.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-disruption
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/pod-disruption.md
    title: 'Kubernetes website source: reference/glossary/pod-disruption.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-disruption
    title: Pod Disruption | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/pod-disruption.md
---

# Pod Disruption

[Pod disruption](/concepts/workloads/pods/disruptions.md) is the process by which 
Pods on Nodes are terminated either voluntarily or involuntarily. 

 

Voluntary disruptions are started intentionally by application owners or cluster 
administrators. Involuntary disruptions are unintentional and can be triggered by 
unavoidable issues like Nodes running out of [resources](/reference/glossary/infrastructure-resource.md),
or by accidental deletions.
