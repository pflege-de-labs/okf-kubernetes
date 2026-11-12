---
type: glossary-term
title: Secret
description: Stores sensitive information, such as passwords, OAuth tokens, and ssh keys.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-secret
tags: [kubernetes, reference, glossary, core-object, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/secret.md
    title: 'Kubernetes website source: reference/glossary/secret.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-secret
    title: Secret | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/secret.md
---

# Secret

Stores sensitive information, such as passwords, OAuth tokens, and SSH keys.

Secrets give you more control over how sensitive information is used and reduces
the risk of accidental exposure. Secret values are encoded as base64 strings and
are stored unencrypted by default, but can be configured to be
[encrypted at rest](/tasks/administer-cluster/encrypt-data.md#ensure-all-secrets-are-encrypted).

A [Pod](/reference/glossary/pod.md) can reference the Secret in
a variety of ways, such as in a volume mount or as an environment variable.
Secrets are designed for confidential data and
[ConfigMaps](/tasks/configure-pod-container/configure-pod-configmap.md) are
designed for non-confidential data.
