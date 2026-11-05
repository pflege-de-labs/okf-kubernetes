---
type: glossary-term
title: Kubelet
description: An agent that runs on each node in the cluster. It makes sure that containers are running in a pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kubelet
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/kubelet.md
    title: 'Kubernetes website source: reference/glossary/kubelet.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kubelet
    title: Kubelet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/kubelet.md
---

# Kubelet

An agent that runs on each [node](/reference/glossary/node.md) in the cluster. It makes sure that [containers](/reference/glossary/container.md) are running in a [Pod](/reference/glossary/pod.md).

The [kubelet](/reference/command-line-tools-reference/kubelet.md) takes a set of PodSpecs that 
are provided through various mechanisms and ensures that the containers described in those 
PodSpecs are running and healthy. The kubelet doesn't manage containers which were not created by 
Kubernetes.
