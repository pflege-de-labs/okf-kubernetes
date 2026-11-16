---
type: glossary-term
title: user namespace
description: A Linux kernel feature to emulate superuser privilege for unprivileged users.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-userns
tags: [kubernetes, reference, glossary, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/userns.md
    title: 'Kubernetes website source: reference/glossary/userns.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-userns
    title: user namespace | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/userns.md
---

# user namespace

A kernel feature to emulate root. Used for "rootless containers".

User namespaces are a Linux kernel feature that allows a non-root user to
emulate superuser ("root") privileges,
for example in order to run containers without being a superuser outside the container.

User namespace is effective for mitigating damage of potential container break-out attacks.

In the context of user namespaces, the namespace is a Linux kernel feature, and not a
[namespace](/reference/glossary/namespace.md) in the Kubernetes sense
of the term.
