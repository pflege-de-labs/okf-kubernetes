---
type: glossary-term
title: Proxy
description: An application acting as an intermediary between clients and servers
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-proxy
tags: [kubernetes, reference, glossary, networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/proxy.md
    title: 'Kubernetes website source: reference/glossary/proxy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-proxy
    title: Proxy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/proxy.md
---

# Proxy

In computing, a proxy is a server that acts as an intermediary for a remote
service.

A client interacts with the proxy; the proxy copies the client's data to the
actual server; the actual server replies to the proxy; the proxy sends the
actual server's reply to the client.

[kube-proxy](/reference/command-line-tools-reference/kube-proxy.md) is a
network proxy that runs on each node in your cluster, implementing part of
the Kubernetes [Service](/reference/glossary/service.md) concept.

You can run kube-proxy as a plain userland proxy service. If your operating
system supports it, you can instead run kube-proxy in a hybrid mode that
achieves the same overall effect using less system resources.
