---
type: api-reference
title: NonResourceRule
description: NonResourceRule holds information that describes a rule for the non-resource
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/non-resource-rule-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/non-resource-rule-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/non-resource-rule-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/non-resource-rule-v1-authorization/
    title: NonResourceRule | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/non-resource-rule-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: NonResourceRule
---

# NonResourceRule

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## NonResourceRule {#NonResourceRule}

NonResourceRule holds information that describes a rule for the non-resource

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>nonResourceURLs</code><br/><em>string array</em></td>
      <td>nonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>verbs</code>&nbsp;<strong>*</strong><br/><em>string array</em></td>
      <td>verbs is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.</td>
    </tr>
  </tbody>
</table>
