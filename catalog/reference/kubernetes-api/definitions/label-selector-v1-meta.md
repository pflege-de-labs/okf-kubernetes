---
type: api-reference
title: LabelSelector
description: A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/label-selector-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/label-selector-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/label-selector-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/label-selector-v1-meta/
    title: LabelSelector | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/label-selector-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: LabelSelector
---

# LabelSelector

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## LabelSelector {#LabelSelector}

A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>matchExpressions</code><br/><em><a href="/reference/kubernetes-api/definitions/label-selector-requirement-v1-meta.md#LabelSelectorRequirement">LabelSelectorRequirement array</a></em></td>
      <td>matchExpressions is a list of label selector requirements. The requirements are ANDed.</td>
    </tr>
    <tr>
      <td><code>matchLabels</code><br/><em>object</em></td>
      <td>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.</td>
    </tr>
  </tbody>
</table>
