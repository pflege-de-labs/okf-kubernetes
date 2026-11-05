---
type: api-reference
title: Eviction
description: Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod. A request to cause such an eviction is created by POSTing to .../pods/&lt;pod name&gt;/evictions.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/eviction-v1-policy/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/eviction-v1-policy.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/eviction-v1-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/eviction-v1-policy/
    title: Eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/eviction-v1-policy.md
api_metadata:
  apiVersion: policy/v1
  import: k8s.io/api/policy/v1
  kind: Eviction
---

# Eviction

`apiVersion: policy/v1`

`import "k8s.io/api/policy/v1"`

## Eviction {#Eviction}

Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/&lt;pod name&gt;/evictions.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</td>
    </tr>
    <tr>
      <td><code>deleteOptions</code><br/><em><a href="/reference/kubernetes-api/definitions/delete-options-v1-meta.md#DeleteOptions">DeleteOptions</a></em></td>
      <td>deleteOptions may be provided</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>metadata</code><br/><em><a href="/reference/kubernetes-api/definitions/object-meta-v1-meta.md#ObjectMeta">ObjectMeta</a></em></td>
      <td>metadata describes the pod that is being evicted.</td>
    </tr>
  </tbody>
</table>
