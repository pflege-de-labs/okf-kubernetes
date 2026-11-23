---
type: glossary-term
title: kube-controller-manager
description: Control Plane component that runs controller processes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-controller-manager
tags: [kubernetes, reference, glossary, architecture, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/kube-controller-manager.md
    title: 'Kubernetes website source: reference/glossary/kube-controller-manager.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kube-controller-manager
    title: kube-controller-manager | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/kube-controller-manager.md
---

# kube-controller-manager

Control plane component that runs [controller](/reference/glossary/controller.md) processes.

Logically, each [controller](/reference/glossary/controller.md) is a separate process, but to reduce complexity, they are all compiled into a single binary and run in a single process.
