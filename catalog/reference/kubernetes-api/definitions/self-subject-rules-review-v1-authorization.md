---
type: api-reference
title: SelfSubjectRulesReview
description: SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server&#39;s authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/self-subject-rules-review-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/self-subject-rules-review-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/self-subject-rules-review-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/self-subject-rules-review-v1-authorization/
    title: SelfSubjectRulesReview | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/self-subject-rules-review-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: SelfSubjectRulesReview
---

# SelfSubjectRulesReview

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## SelfSubjectRulesReview {#SelfSubjectRulesReview}

SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server&#39;s authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.

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
      <td>metadata is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata</td>
    </tr>
    <tr>
      <td><code>spec</code>&nbsp;<strong>*</strong><br/><em><a href="/reference/kubernetes-api/definitions.md#SelfSubjectRulesReviewSpec">SelfSubjectRulesReviewSpec</a></em></td>
      <td>spec holds information about the request being evaluated.</td>
    </tr>
    <tr>
      <td><code>status</code><br/><em><a href="/reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization.md#SubjectRulesReviewStatus">SubjectRulesReviewStatus</a></em></td>
      <td>status is filled in by the server and indicates the set of actions a user can perform.</td>
    </tr>
  </tbody>
</table>

## SelfSubjectRulesReviewSpec {#SelfSubjectRulesReviewSpec}

SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>namespace</code><br/><em>string</em></td>
      <td>namespace to evaluate rules for. Required.</td>
    </tr>
  </tbody>
</table>
