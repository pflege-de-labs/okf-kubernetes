---
type: glossary-term
title: Init Container
description: One or more initialization containers that must run to completion before any app containers run.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-init-container
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/init-container.md
    title: 'Kubernetes website source: reference/glossary/init-container.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-init-container
    title: Init Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/init-container.md
---

# Init Container

One or more initialization [containers](/reference/glossary/container.md) that must run to completion before any app containers run.

 

Initialization (init) containers are like regular app containers, with one difference: init containers must run to completion before any app containers can start. Init containers run in series: each init container must run to completion before the next init container begins.

Unlike [sidecar containers](/reference/glossary/sidecar-container.md), init containers do not remain running after Pod startup.

For more information, read [init containers](/concepts/workloads/pods/init-containers.md).
