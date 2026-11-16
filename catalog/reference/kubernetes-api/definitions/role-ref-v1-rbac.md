---
type: api-reference
title: RoleRef
description: RoleRef contains information that points to the role being used
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/role-ref-v1-rbac/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/kubernetes-api/definitions/role-ref-v1-rbac.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/role-ref-v1-rbac.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/role-ref-v1-rbac/
    title: RoleRef | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/kubernetes-api/definitions/role-ref-v1-rbac.md
api_metadata:
  apiVersion: rbac.authorization.k8s.io/v1
  import: k8s.io/api/rbac/v1
  kind: RoleRef
---

# RoleRef

`apiVersion: rbac.authorization.k8s.io/v1`

`import "k8s.io/api/rbac/v1"`

## RoleRef {#RoleRef}

RoleRef contains information that points to the role being used

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiGroup</code><br/><em>string</em></td>
      <td>apiGroup is the group for the resource being referenced</td>
    </tr>
    <tr>
      <td><code>kind</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>kind is the type of resource being referenced</td>
    </tr>
    <tr>
      <td><code>name</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>name is the name of resource being referenced</td>
    </tr>
  </tbody>
</table>
