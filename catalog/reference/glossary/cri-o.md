---
type: glossary-term
title: CRI-O
description: A lightweight container runtime specifically for Kubernetes
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri-o
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/cri-o.md
    title: 'Kubernetes website source: reference/glossary/cri-o.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri-o
    title: CRI-O | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/cri-o.md
---

# CRI-O

A tool that lets you use OCI container runtimes with Kubernetes CRI.

CRI-O is an implementation of the [Container Runtime Interface (CRI)](/reference/glossary/cri.md)
to enable using [container](/reference/glossary/container.md)
runtimes that are compatible with the Open Container Initiative (OCI)
[runtime spec](https://www.github.com/opencontainers/runtime-spec).

Deploying CRI-O allows Kubernetes to use any OCI-compliant runtime as the container
runtime for running [Pods](/reference/glossary/pod.md), and to fetch
OCI container images from remote registries.
