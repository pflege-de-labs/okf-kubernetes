---
type: api-reference
title: SubjectRulesReviewStatus
description: SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it&#39;s safe to assume the subject has that permission, even if that list is incomplete.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization/
    title: SubjectRulesReviewStatus | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/subject-rules-review-status-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: SubjectRulesReviewStatus
---

# SubjectRulesReviewStatus

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## SubjectRulesReviewStatus {#SubjectRulesReviewStatus}

SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is configured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it&#39;s safe to assume the subject has that permission, even if that list is incomplete.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>evaluationError</code><br/><em>string</em></td>
      <td>evaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.</td>
    </tr>
    <tr>
      <td><code>incomplete</code>&nbsp;<strong>*</strong><br/><em>boolean</em></td>
      <td>incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.</td>
    </tr>
    <tr>
      <td><code>nonResourceRules</code>&nbsp;<strong>*</strong><br/><em><a href="/reference/kubernetes-api/definitions/non-resource-rule-v1-authorization.md#NonResourceRule">NonResourceRule array</a></em></td>
      <td>nonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.</td>
    </tr>
    <tr>
      <td><code>resourceRules</code>&nbsp;<strong>*</strong><br/><em><a href="/reference/kubernetes-api/definitions/resource-rule-v1-authorization.md#ResourceRule">ResourceRule array</a></em></td>
      <td>resourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.</td>
    </tr>
  </tbody>
</table>
