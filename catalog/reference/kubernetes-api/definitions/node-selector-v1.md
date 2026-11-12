---
type: api-reference
title: NodeSelector
description: A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/node-selector-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/node-selector-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/node-selector-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/node-selector-v1/
    title: NodeSelector | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/node-selector-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: NodeSelector
---

# NodeSelector

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## NodeSelector {#NodeSelector}

A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>nodeSelectorTerms</code>&nbsp;<strong>*</strong><br/><em><a href="/reference/kubernetes-api/definitions/node-selector-term-v1.md#NodeSelectorTerm">NodeSelectorTerm array</a></em></td>
      <td>Required. A list of node selector terms. The terms are ORed.</td>
    </tr>
  </tbody>
</table>
