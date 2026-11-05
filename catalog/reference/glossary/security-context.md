---
type: glossary-term
title: Security Context
description: The securityContext field defines privilege and access control settings for a Pod or container.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-security-context
tags: [kubernetes, reference, glossary, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/security-context.md
    title: 'Kubernetes website source: reference/glossary/security-context.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-security-context
    title: Security Context | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/security-context.md
---

# Security Context

The `securityContext` field defines privilege and access control settings for
a [Pod](/reference/glossary/pod.md) or
[container](/reference/glossary/container.md).

In a `securityContext`, you can define: the user that processes run as,
the group that processes run as, and privilege settings.
You can also configure security policies (for example: SELinux, AppArmor or seccomp).

The `PodSpec.securityContext` setting applies to all containers in a Pod.
