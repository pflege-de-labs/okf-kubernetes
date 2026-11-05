---
type: api-reference
title: StatusCause
description: StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/status-cause-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/status-cause-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/status-cause-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/status-cause-v1-meta/
    title: StatusCause | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/status-cause-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: StatusCause
---

# StatusCause

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## StatusCause {#StatusCause}

StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>field</code><br/><em>string</em></td>
      <td>The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.  Examples:   "name" - the field "name" on the current resource   "items[0].name" - the field "name" on the first array entry in "items"</td>
    </tr>
    <tr>
      <td><code>message</code><br/><em>string</em></td>
      <td>A human-readable description of the cause of the error.  This field may be presented as-is to a reader.</td>
    </tr>
    <tr>
      <td><code>reason</code><br/><em>string</em></td>
      <td>A machine-readable description of the cause of the error. If this value is empty there is no information available.</td>
    </tr>
  </tbody>
</table>
