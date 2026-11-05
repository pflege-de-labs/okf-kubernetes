---
type: glossary-term
title: Feature gate
description: A way to control whether or not a particular Kubernetes feature is enabled.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-feature-gates
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/feature-gates.md
    title: 'Kubernetes website source: reference/glossary/feature-gates.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-feature-gates
    title: Feature gate | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/feature-gates.md
---

# Feature gate

Feature gates are a set of keys (opaque string values) that you can use to control which
Kubernetes features are enabled in your cluster.

You can turn these features on or off using the `--feature-gates` command line flag on each Kubernetes component.
Each Kubernetes component lets you enable or disable a set of feature gates that are relevant to that component.
The Kubernetes documentation lists all current 
[feature gates](/reference/command-line-tools-reference/feature-gates.md) and what they control.
