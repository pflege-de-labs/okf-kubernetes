---
type: api-reference
title: Binding
description: Binding ties one object to another; for example, a pod is bound to a node by a scheduler.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/binding-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/binding-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/binding-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/binding-v1/
    title: Binding | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/binding-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: Binding
---

# Binding

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## Binding {#Binding}

Binding ties one object to another; for example, a pod is bound to a node by a scheduler.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>metadata</code><br/><em><a href="/reference/kubernetes-api/definitions/object-meta-v1-meta.md#ObjectMeta">ObjectMeta</a></em></td>
      <td>Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata</td>
    </tr>
    <tr>
      <td><code>target</code>&nbsp;<strong>*</strong><br/><em><a href="/reference/kubernetes-api/definitions/object-reference-v1.md#ObjectReference">ObjectReference</a></em></td>
      <td>The target object that you want to bind to the standard object.</td>
    </tr>
  </tbody>
</table>
