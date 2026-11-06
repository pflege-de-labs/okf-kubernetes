---
type: glossary-term
title: Manifest
description: A serialized specification of one or more Kubernetes API objects.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-manifest
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/manifest.md
    title: 'Kubernetes website source: reference/glossary/manifest.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-manifest
    title: Manifest | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/manifest.md
---

# Manifest

Specification of a Kubernetes API object in [JSON](https://www.json.org/json-en.html)
or [YAML](https://yaml.org/) format.

A manifest specifies the desired state of an object that Kubernetes will maintain when you apply the manifest.
For YAML format, each file can contain multiple manifests.
