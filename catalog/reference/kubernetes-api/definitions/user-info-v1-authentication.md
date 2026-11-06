---
type: api-reference
title: UserInfo
description: UserInfo holds the information about the user needed to implement the user.Info interface.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/user-info-v1-authentication/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/user-info-v1-authentication.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/user-info-v1-authentication.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/user-info-v1-authentication/
    title: UserInfo | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/user-info-v1-authentication.md
api_metadata:
  apiVersion: authentication.k8s.io/v1
  import: k8s.io/api/authentication/v1
  kind: UserInfo
---

# UserInfo

`apiVersion: authentication.k8s.io/v1`

`import "k8s.io/api/authentication/v1"`

## UserInfo {#UserInfo}

UserInfo holds the information about the user needed to implement the user.Info interface.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>extra</code><br/><em>object</em></td>
      <td>extra is any additional information provided by the authenticator.</td>
    </tr>
    <tr>
      <td><code>groups</code><br/><em>string array</em></td>
      <td>groups is the names of groups this user is a part of.</td>
    </tr>
    <tr>
      <td><code>uid</code><br/><em>string</em></td>
      <td>uid is a unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.</td>
    </tr>
    <tr>
      <td><code>username</code><br/><em>string</em></td>
      <td>username is the name that uniquely identifies this user among all active users.</td>
    </tr>
  </tbody>
</table>
