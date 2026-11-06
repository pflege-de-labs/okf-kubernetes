---
type: glossary-term
title: HostAliases
description: A HostAliases is a mapping between the IP address and hostname to be injected into a Pod's hosts file.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-host-aliases
tags: [kubernetes, reference, glossary, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/host-aliases.md
    title: 'Kubernetes website source: reference/glossary/host-aliases.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-host-aliases
    title: HostAliases | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/host-aliases.md
---

# HostAliases

A HostAliases is a mapping between the IP address and hostname to be injected into a [Pod](/reference/glossary/pod.md)'s hosts file.

[HostAliases](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#hostalias-v1-core) is an optional list of hostnames and IP addresses that will be injected into the Pod's hosts file if specified. This is only valid for non-hostNetwork Pods.
