---
type: glossary-term
title: Manifest
description: A serialized specification of one or more Kubernetes API objects.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-manifest
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/manifest.md
    title: 'Kubernetes website source: reference/glossary/manifest.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-manifest
    title: Manifest | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/manifest.md
---

# Manifest

Specification of a Kubernetes API object in [JSON](https://www.json.org/json-en.html)
or [YAML](https://yaml.org/) format.

A manifest specifies the desired state of an object that Kubernetes will maintain when you apply the manifest.
For YAML format, each file can contain multiple manifests.
