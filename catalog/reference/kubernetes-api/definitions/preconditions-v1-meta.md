---
type: api-reference
title: Preconditions
description: Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/preconditions-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/preconditions-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/preconditions-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/preconditions-v1-meta/
    title: Preconditions | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/preconditions-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: Preconditions
---

# Preconditions

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## Preconditions {#Preconditions}

Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>resourceVersion</code><br/><em>string</em></td>
      <td>Specifies the target ResourceVersion</td>
    </tr>
    <tr>
      <td><code>uid</code><br/><em>string</em></td>
      <td>Specifies the target UID.</td>
    </tr>
  </tbody>
</table>
