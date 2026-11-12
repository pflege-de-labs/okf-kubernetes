---
type: glossary-term
title: Name
description: A client-provided string that refers to an object in a resource URL, such as `/api/v1/pods/some-name`.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-name
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/name.md
    title: 'Kubernetes website source: reference/glossary/name.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-name
    title: Name | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/name.md
---

# Name

A client-provided string that refers to an object in a [resource](/reference/glossary/api-resource.md)
URL, such as `/api/v1/pods/some-name`.

 

Only one object of a given kind can have a given name at a time. However, if you delete the object, you can make a new object with the same name.
