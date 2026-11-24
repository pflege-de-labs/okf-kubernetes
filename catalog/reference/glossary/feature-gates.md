---
type: glossary-term
title: Feature gate
description: A way to control whether or not a particular Kubernetes feature is enabled.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-feature-gates
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/feature-gates.md
    title: 'Kubernetes website source: reference/glossary/feature-gates.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-feature-gates
    title: Feature gate | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/feature-gates.md
---

# Feature gate

Feature gates are a set of keys (opaque string values) that you can use to control which
Kubernetes features are enabled in your cluster.

You can turn these features on or off using the `--feature-gates` command line flag on each Kubernetes component.
Each Kubernetes component lets you enable or disable a set of feature gates that are relevant to that component.
The Kubernetes documentation lists all current 
[feature gates](/reference/command-line-tools-reference/feature-gates.md) and what they control.
