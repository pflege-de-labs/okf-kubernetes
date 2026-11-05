---
type: glossary-term
title: Group Version Resource
description: The API group, API version and name of a Kubernetes API.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-group-version-resource
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/group-version-resource.md
    title: 'Kubernetes website source: reference/glossary/group-version-resource.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-group-version-resource
    title: Group Version Resource | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/group-version-resource.md
---

# Group Version Resource

Means of representing specific Kubernetes APIs uniquely.

Group Version Resources (GVRs) specify the API group, API version, and _resource_ (name for the object kind as it appears in the URI) associated with accessing a particular id of object in Kubernetes.
GVRs let you define and distinguish different Kubernetes objects, and to specify a way of accessing
objects that is stable even as APIs change.

In this usage, _resource_ refers to an HTTP resource. Because some APIs are namespaced, a GVR may
not refer to a specific [API resource](/reference/glossary/api-resource.md).
