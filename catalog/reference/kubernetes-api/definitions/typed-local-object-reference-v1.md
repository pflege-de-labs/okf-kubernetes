---
type: api-reference
title: TypedLocalObjectReference
description: TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/typed-local-object-reference-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1/
    title: TypedLocalObjectReference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: TypedLocalObjectReference
---

# TypedLocalObjectReference

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## TypedLocalObjectReference {#TypedLocalObjectReference}

TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiGroup</code><br/><em>string</em></td>
      <td>APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.</td>
    </tr>
    <tr>
      <td><code>kind</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>Kind is the type of resource being referenced</td>
    </tr>
    <tr>
      <td><code>name</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>Name is the name of resource being referenced</td>
    </tr>
  </tbody>
</table>
