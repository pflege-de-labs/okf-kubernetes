---
type: api-reference
title: NonResourceAttributes
description: NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/non-resource-attributes-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/non-resource-attributes-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/non-resource-attributes-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/non-resource-attributes-v1-authorization/
    title: NonResourceAttributes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/non-resource-attributes-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: NonResourceAttributes
---

# NonResourceAttributes

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## NonResourceAttributes {#NonResourceAttributes}

NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>path</code><br/><em>string</em></td>
      <td>path is the URL path of the request</td>
    </tr>
    <tr>
      <td><code>verb</code><br/><em>string</em></td>
      <td>verb is the standard HTTP verb</td>
    </tr>
  </tbody>
</table>
