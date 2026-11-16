---
type: glossary-term
title: API server
description: Control plane component that serves the Kubernetes API.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-apiserver
tags: [kubernetes, reference, glossary, architecture, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/kube-apiserver.md
    title: 'Kubernetes website source: reference/glossary/kube-apiserver.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-apiserver
    title: API server | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/kube-apiserver.md
---

# API server

The API server is a component of the Kubernetes
[control plane](/reference/glossary/control-plane.md) that exposes the Kubernetes API.
The API server is the front end for the Kubernetes control plane.

The main implementation of a Kubernetes API server is [kube-apiserver](https://kubernetes.io/docs/reference/generated/kube-apiserver/).
kube-apiserver is designed to scale horizontally&mdash;that is, it scales by deploying more instances.
You can run several instances of kube-apiserver and balance traffic between those instances.
