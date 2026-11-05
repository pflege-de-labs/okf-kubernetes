---
type: glossary-term
title: Container Runtime Interface (CRI)
description: Protocol for communication between the kubelet and the local container runtime.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cri.md
    title: 'Kubernetes website source: reference/glossary/cri.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri
    title: Container Runtime Interface (CRI) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/cri.md
---

# Container Runtime Interface (CRI)

The main protocol for the communication between the [kubelet](/reference/glossary/kubelet.md) and Container Runtime.

The Kubernetes Container Runtime Interface (CRI) defines the main
[gRPC](https://grpc.io) protocol for the communication between the
[node components](/concepts/architecture.md#node-components)
[kubelet](/reference/glossary/kubelet.md) and
[container runtime](/reference/glossary/container-runtime.md).
