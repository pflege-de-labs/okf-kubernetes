---
type: glossary-term
title: Cloud Controller Manager
description: Control plane component that integrates Kubernetes with third-party cloud providers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cloud-controller-manager
tags: [kubernetes, reference, glossary, architecture, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cloud-controller-manager.md
    title: 'Kubernetes website source: reference/glossary/cloud-controller-manager.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cloud-controller-manager
    title: Cloud Controller Manager | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/cloud-controller-manager.md
---

# Cloud Controller Manager

A Kubernetes [control plane](/reference/glossary/control-plane.md) component
that embeds cloud-specific control logic. The cloud controller manager lets you link your
cluster into your cloud provider's API, and separates out the components that interact
with that cloud platform from components that only interact with your cluster.

By decoupling the interoperability logic between Kubernetes and the underlying cloud
infrastructure, the cloud-controller-manager component enables cloud providers to release
features at a different pace compared to the main Kubernetes project.
