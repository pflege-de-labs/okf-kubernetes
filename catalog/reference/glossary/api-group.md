---
type: glossary-term
title: API Group
description: A set of related paths in the Kubernetes API.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-group
tags: [kubernetes, reference, glossary, fundamental, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/api-group.md
    title: 'Kubernetes website source: reference/glossary/api-group.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-api-group
    title: API Group | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/api-group.md
---

# API Group

A set of related paths in Kubernetes API.

You can enable or disable each API group by changing the configuration of your API server. You can also disable or enable paths to specific
[resources](/reference/glossary/api-resource.md). An API group makes it easier to extend the Kubernetes API.
The API group is specified in a REST path and in the `apiVersion` field of a serialized [object](/reference/glossary/object.md).

* Read [API Group](/concepts/overview/kubernetes-api.md#api-groups-and-versioning) for more information.
