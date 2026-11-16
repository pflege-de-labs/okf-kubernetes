---
type: glossary-term
title: Operator pattern
description: A specialized controller used to manage a custom resource
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-operator-pattern
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/operator-pattern.md
    title: 'Kubernetes website source: reference/glossary/operator-pattern.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-operator-pattern
    title: Operator pattern | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/operator-pattern.md
---

# Operator pattern

The [operator pattern](/concepts/extend-kubernetes/operator.md) is a system
design that links a [Controller](/reference/glossary/controller.md) to one or more custom
resources.

You can extend Kubernetes by adding controllers to your cluster, beyond the built-in
controllers that come as part of Kubernetes itself.

If a running application acts as a controller and has API access to carry out tasks
against a custom resource that's defined in the control plane, that's an example of
the Operator pattern.
