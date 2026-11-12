---
type: api-reference
title: SelfSubjectReview
description: SelfSubjectReview contains the user information that the kube-apiserver has about the user making this request. When using impersonation, users will receive the user info of the user being impersonated. If impersonation or request header authentication is used, any extra keys will have their case ignored and returned as lowercase.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/self-subject-review-v1-authentication/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/self-subject-review-v1-authentication.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/self-subject-review-v1-authentication.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/self-subject-review-v1-authentication/
    title: SelfSubjectReview | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/self-subject-review-v1-authentication.md
api_metadata:
  apiVersion: authentication.k8s.io/v1
  import: k8s.io/api/authentication/v1
  kind: SelfSubjectReview
---

# SelfSubjectReview

`apiVersion: authentication.k8s.io/v1`

`import "k8s.io/api/authentication/v1"`

## SelfSubjectReview {#SelfSubjectReview}

SelfSubjectReview contains the user information that the kube-apiserver has about the user making this request. When using impersonation, users will receive the user info of the user being impersonated.  If impersonation or request header authentication is used, any extra keys will have their case ignored and returned as lowercase.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>metadata</code><br/><em><a href="/reference/kubernetes-api/definitions/object-meta-v1-meta.md#ObjectMeta">ObjectMeta</a></em></td>
      <td>metadata is standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata</td>
    </tr>
    <tr>
      <td><code>status</code><br/><em><a href="/reference/kubernetes-api/definitions.md#SelfSubjectReviewStatus">SelfSubjectReviewStatus</a></em></td>
      <td>status is filled in by the server with the user attributes.</td>
    </tr>
  </tbody>
</table>

## SelfSubjectReviewStatus {#SelfSubjectReviewStatus}

SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>userInfo</code><br/><em><a href="/reference/kubernetes-api/definitions/user-info-v1-authentication.md#UserInfo">UserInfo</a></em></td>
      <td>userInfo is a set of attributes belonging to the user making this request.</td>
    </tr>
  </tbody>
</table>
