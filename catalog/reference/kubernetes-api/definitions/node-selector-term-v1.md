---
type: api-reference
title: NodeSelectorTerm
description: A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/node-selector-term-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/node-selector-term-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/node-selector-term-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/node-selector-term-v1/
    title: NodeSelectorTerm | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/node-selector-term-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: NodeSelectorTerm
---

# NodeSelectorTerm

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## NodeSelectorTerm {#NodeSelectorTerm}

A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>matchExpressions</code><br/><em><a href="/reference/kubernetes-api/resource/resource-slice-v1.md#NodeSelectorRequirement">NodeSelectorRequirement array</a></em></td>
      <td>A list of node selector requirements by node's labels.</td>
    </tr>
    <tr>
      <td><code>matchFields</code><br/><em><a href="/reference/kubernetes-api/resource/resource-slice-v1.md#NodeSelectorRequirement">NodeSelectorRequirement array</a></em></td>
      <td>A list of node selector requirements by node's fields.</td>
    </tr>
  </tbody>
</table>
