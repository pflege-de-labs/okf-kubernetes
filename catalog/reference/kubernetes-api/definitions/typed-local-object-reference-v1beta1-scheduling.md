---
type: api-reference
title: TypedLocalObjectReference
description: TypedLocalObjectReference allows to reference typed object inside the same namespace.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1beta1-scheduling/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1beta1-scheduling.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/typed-local-object-reference-v1beta1-scheduling.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1beta1-scheduling/
    title: TypedLocalObjectReference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/typed-local-object-reference-v1beta1-scheduling.md
api_metadata:
  apiVersion: scheduling.k8s.io/v1beta1
  import: k8s.io/api/scheduling/v1beta1
  kind: TypedLocalObjectReference
---

# TypedLocalObjectReference

`apiVersion: scheduling.k8s.io/v1beta1`

`import "k8s.io/api/scheduling/v1beta1"`

## TypedLocalObjectReference {#TypedLocalObjectReference}

TypedLocalObjectReference allows to reference typed object inside the same namespace.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiGroup</code><br/><em>string</em></td>
      <td>apiGroup is the group for the resource being referenced. If apiGroup is empty, the specified Kind must be in the core API group. For any other third-party types, setting apiGroup is required. It must be a DNS subdomain.</td>
    </tr>
    <tr>
      <td><code>kind</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>kind is the type of resource being referenced. It must be a path segment name.</td>
    </tr>
    <tr>
      <td><code>name</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>name is the name of resource being referenced. It must be a path segment name.</td>
    </tr>
  </tbody>
</table>
