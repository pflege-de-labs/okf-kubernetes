---
type: section
title: Extending the Kubernetes API
description: Custom resources are extensions of the Kubernetes API.
resource: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/
tags: [kubernetes, concepts, extend-kubernetes]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/extend-kubernetes/api-extension/_index.md
    title: 'Kubernetes website source: concepts/extend-kubernetes/api-extension/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/
    title: Extending the Kubernetes API | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/extend-kubernetes/api-extension/_index.md
---

# Extending the Kubernetes API

Custom resources are extensions of the Kubernetes API. Kubernetes provides two ways to add custom resources to your cluster:

- The [CustomResourceDefinition](/concepts/extend-kubernetes/api-extension/custom-resources.md)
  (CRD) mechanism allows you to declaratively define a new custom API with an API group, kind, and
  schema that you specify.
  The Kubernetes control plane serves and handles the storage of your custom resource. CRDs allow you to
  create new types of resources for your cluster without writing and running a custom API server. 
- The [aggregation layer](/concepts/extend-kubernetes/api-extension/apiserver-aggregation.md)
  sits behind the primary API server, which acts as a proxy.
  This arrangement is called API Aggregation (AA), which allows you to provide
  specialized implementations for your custom resources by writing and
  deploying your own API server.
  The main API server delegates requests to your API server for the custom APIs that you specify,
  making them available to all of its clients.
