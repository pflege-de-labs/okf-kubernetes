---
type: glossary-term
title: Image
description: Stored instance of a container that holds a set of software needed to run an application.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-image
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/image.md
    title: 'Kubernetes website source: reference/glossary/image.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-image
    title: Image | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/image.md
---

# Image

Stored instance of a [Container](/reference/glossary/container.md) that holds a set of software needed to run an application.

A way of packaging software that allows it to be stored in a container registry, pulled to a local system, and run as an application. Meta data is included in the image that can indicate what executable to run, who built it, and other information.
