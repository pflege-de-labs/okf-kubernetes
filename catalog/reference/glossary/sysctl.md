---
type: glossary-term
title: sysctl
description: An interface for getting and setting Unix kernel parameters
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-sysctl
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/sysctl.md
    title: 'Kubernetes website source: reference/glossary/sysctl.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-sysctl
    title: sysctl | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/sysctl.md
---

# sysctl

`sysctl` is a semi-standardized interface for reading or changing the
 attributes of the running Unix kernel.

On Unix-like systems, `sysctl` is both the name of the tool that administrators
use to view and modify these settings, and also the system call that the tool
uses.

[Container](/reference/glossary/container.md) runtimes and
network plugins may rely on `sysctl` values being set a certain way.
