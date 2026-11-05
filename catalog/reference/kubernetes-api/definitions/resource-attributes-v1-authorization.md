---
type: api-reference
title: ResourceAttributes
description: ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/resource-attributes-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/resource-attributes-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/resource-attributes-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/resource-attributes-v1-authorization/
    title: ResourceAttributes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/resource-attributes-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: ResourceAttributes
---

# ResourceAttributes

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## ResourceAttributes {#ResourceAttributes}

ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>fieldSelector</code><br/><em><a href="/reference/kubernetes-api/definitions/field-selector-attributes-v1-authorization.md#FieldSelectorAttributes">FieldSelectorAttributes</a></em></td>
      <td>fieldSelector describes the limitation on access based on field.  It can only limit access, not broaden it.</td>
    </tr>
    <tr>
      <td><code>group</code><br/><em>string</em></td>
      <td>group is the API Group of the Resource.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>labelSelector</code><br/><em><a href="/reference/kubernetes-api/definitions/label-selector-attributes-v1-authorization.md#LabelSelectorAttributes">LabelSelectorAttributes</a></em></td>
      <td>labelSelector describes the limitation on access based on labels.  It can only limit access, not broaden it.</td>
    </tr>
    <tr>
      <td><code>name</code><br/><em>string</em></td>
      <td>name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.</td>
    </tr>
    <tr>
      <td><code>namespace</code><br/><em>string</em></td>
      <td>namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview</td>
    </tr>
    <tr>
      <td><code>resource</code><br/><em>string</em></td>
      <td>resource is one of the existing resource types.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>subresource</code><br/><em>string</em></td>
      <td>subresource is one of the existing resource types.  "" means none.</td>
    </tr>
    <tr>
      <td><code>verb</code><br/><em>string</em></td>
      <td>verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>version</code><br/><em>string</em></td>
      <td>version is the API Version of the Resource.  "*" means all.</td>
    </tr>
  </tbody>
</table>
