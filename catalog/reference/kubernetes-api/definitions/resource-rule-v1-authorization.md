---
type: api-reference
title: ResourceRule
description: ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn&#39;t significant, may contain duplicates, and possibly be incomplete.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/resource-rule-v1-authorization/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/resource-rule-v1-authorization.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/resource-rule-v1-authorization.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/resource-rule-v1-authorization/
    title: ResourceRule | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/resource-rule-v1-authorization.md
api_metadata:
  apiVersion: authorization.k8s.io/v1
  import: k8s.io/api/authorization/v1
  kind: ResourceRule
---

# ResourceRule

`apiVersion: authorization.k8s.io/v1`

`import "k8s.io/api/authorization/v1"`

## ResourceRule {#ResourceRule}

ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn&#39;t significant, may contain duplicates, and possibly be incomplete.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiGroups</code><br/><em>string array</em></td>
      <td>apiGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>resourceNames</code><br/><em>string array</em></td>
      <td>resourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  "*" means all.</td>
    </tr>
    <tr>
      <td><code>resources</code><br/><em>string array</em></td>
      <td>resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.  "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.</td>
    </tr>
    <tr>
      <td><code>verbs</code>&nbsp;<strong>*</strong><br/><em>string array</em></td>
      <td>verbs is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  "*" means all.</td>
    </tr>
  </tbody>
</table>
