---
type: glossary-term
title: CRI-O
description: A lightweight container runtime specifically for Kubernetes
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri-o
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/cri-o.md
    title: 'Kubernetes website source: reference/glossary/cri-o.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cri-o
    title: CRI-O | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
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
