---
type: api-reference
title: Status
description: Status is a return value for calls that don&#39;t return other objects.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/status-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/status-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/status-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/status-v1-meta/
    title: Status | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/status-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: Status
---

# Status

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## Status {#Status}

Status is a return value for calls that don&#39;t return other objects.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</td>
    </tr>
    <tr>
      <td><code>code</code><br/><em>integer</em></td>
      <td>Suggested HTTP return code for this status, 0 if not set.</td>
    </tr>
    <tr>
      <td><code>details</code><br/><em><a href="/reference/kubernetes-api/definitions/status-details-v1-meta.md#StatusDetails">StatusDetails</a></em></td>
      <td>Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>message</code><br/><em>string</em></td>
      <td>A human-readable description of the status of this operation.</td>
    </tr>
    <tr>
      <td><code>metadata</code><br/><em><a href="/reference/kubernetes-api/definitions/list-meta-v1-meta.md#ListMeta">ListMeta</a></em></td>
      <td>Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>reason</code><br/><em>string</em></td>
      <td>A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.</td>
    </tr>
    <tr>
      <td><code>status</code><br/><em>string</em></td>
      <td>Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</td>
    </tr>
  </tbody>
</table>
