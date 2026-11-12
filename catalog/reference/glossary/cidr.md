---
type: glossary-term
title: CIDR
description: CIDR is a notation for describing blocks of IP addresses and is used heavily in various networking configurations.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cidr
tags: [kubernetes, reference, glossary, networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/cidr.md
    title: 'Kubernetes website source: reference/glossary/cidr.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cidr
    title: CIDR | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/cidr.md
---

# CIDR

CIDR (Classless Inter-Domain Routing) is a notation for describing blocks of IP addresses and is used heavily in various networking configurations.

In the context of Kubernetes, each [Node](/reference/glossary/node.md) is assigned a range of IP addresses through the start address and a subnet mask using CIDR. This allows Nodes to assign each [Pod](/reference/glossary/pod.md) a unique IP address. Although originally a concept for IPv4, CIDR has also been expanded to include IPv6.
