---
type: section
title: Kubernetes API
description: Kubernetes' API is the HTTP API that lets users, cluster components, and external components communicate with one another and query or manipulate the state of Kubernetes objects.
resource: https://kubernetes.io/docs/reference/kubernetes-api/
tags: [kubernetes, reference]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/_index.md
    title: 'Kubernetes website source: reference/kubernetes-api/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/
    title: Kubernetes API | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/_index.md
---

# Kubernetes API

Kubernetes' API is the HTTP API that lets users, cluster components, and external components communicate with one another and query or manipulate the state of Kubernetes objects.

 

Kubernetes resources and "records of intent" are all stored as API objects, and modified via RESTful calls to the API. The API allows configuration to be managed in a declarative way. Users can interact with the Kubernetes API directly, or via tools like `kubectl`. The core Kubernetes API is flexible and can also be extended to support custom resources.
