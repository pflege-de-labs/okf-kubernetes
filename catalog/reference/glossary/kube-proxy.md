---
type: glossary-term
title: kube-proxy
description: '`kube-proxy` is a network proxy that runs on each node in the cluster.'
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-proxy
tags: [kubernetes, reference, glossary, fundamental, networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/kube-proxy.md
    title: 'Kubernetes website source: reference/glossary/kube-proxy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-proxy
    title: kube-proxy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/kube-proxy.md
---

# kube-proxy

kube-proxy is a network proxy that runs on each
[node](/reference/glossary/node.md) in your cluster,
implementing part of the Kubernetes
[Service](/reference/glossary/service.md) concept.

[kube-proxy](/reference/command-line-tools-reference/kube-proxy.md)
maintains network rules on nodes. These network rules allow network
communication to your Pods from network sessions inside or outside of
your cluster.

kube-proxy uses the operating system packet filtering layer if there is one
and it's available. Otherwise, kube-proxy forwards the traffic itself.
