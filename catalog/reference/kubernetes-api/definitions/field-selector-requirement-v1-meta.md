---
type: api-reference
title: FieldSelectorRequirement
description: FieldSelectorRequirement is a selector that contains values, a key, and an operator that relates the key and values.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/field-selector-requirement-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/field-selector-requirement-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/field-selector-requirement-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/field-selector-requirement-v1-meta/
    title: FieldSelectorRequirement | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/field-selector-requirement-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: FieldSelectorRequirement
---

# FieldSelectorRequirement

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## FieldSelectorRequirement {#FieldSelectorRequirement}

FieldSelectorRequirement is a selector that contains values, a key, and an operator that relates the key and values.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>key</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>key is the field selector key that the requirement applies to.</td>
    </tr>
    <tr>
      <td><code>operator</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. The list of operators may grow in the future.</td>
    </tr>
    <tr>
      <td><code>values</code><br/><em>string array</em></td>
      <td>values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.</td>
    </tr>
  </tbody>
</table>
