---
type: api-reference
title: LocalObjectReference
description: LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/local-object-reference-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/kubernetes-api/definitions/local-object-reference-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/local-object-reference-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/local-object-reference-v1/
    title: LocalObjectReference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/kubernetes-api/definitions/local-object-reference-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: LocalObjectReference
---

# LocalObjectReference

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## LocalObjectReference {#LocalObjectReference}

LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>name</code><br/><em>string</em></td>
      <td>Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
    </tr>
  </tbody>
</table>
