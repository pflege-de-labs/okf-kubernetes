---
type: api-reference
title: ServiceReference
description: ServiceReference holds a reference to Service.legacy.k8s.io
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/service-reference-v1-admissionregistration/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/service-reference-v1-admissionregistration.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/service-reference-v1-admissionregistration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/service-reference-v1-admissionregistration/
    title: ServiceReference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/service-reference-v1-admissionregistration.md
api_metadata:
  apiVersion: admissionregistration.k8s.io/v1
  import: k8s.io/api/admissionregistration/v1
  kind: ServiceReference
---

# ServiceReference

`apiVersion: admissionregistration.k8s.io/v1`

`import "k8s.io/api/admissionregistration/v1"`

## ServiceReference {#ServiceReference}

ServiceReference holds a reference to Service.legacy.k8s.io

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>name</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>name is the name of the service. Required</td>
    </tr>
    <tr>
      <td><code>namespace</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>namespace is the namespace of the service. Required</td>
    </tr>
    <tr>
      <td><code>path</code><br/><em>string</em></td>
      <td>path is an optional URL path which will be sent in any request to this service.</td>
    </tr>
    <tr>
      <td><code>port</code><br/><em>integer</em></td>
      <td>port is the port on the service that hosts the webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).</td>
    </tr>
  </tbody>
</table>
