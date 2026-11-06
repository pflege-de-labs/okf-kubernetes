---
type: glossary-term
title: cAdvisor
description: Tool that provides understanding of the resource usage and performance characteristics for containers
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cadvisor
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cadvisor.md
    title: 'Kubernetes website source: reference/glossary/cadvisor.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cadvisor
    title: cAdvisor | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/cadvisor.md
---

# cAdvisor

cAdvisor (Container Advisor) provides container users an understanding of the [resource](/reference/glossary/infrastructure-resource.md)
usage and performance characteristics of their running [containers](/reference/glossary/container.md).

It is a running daemon that collects, aggregates, processes, and exports information about running containers. Specifically, for each container it keeps resource isolation parameters, historical resource usage, histograms of complete historical resource usage and network statistics. This data is exported by container and machine-wide.
