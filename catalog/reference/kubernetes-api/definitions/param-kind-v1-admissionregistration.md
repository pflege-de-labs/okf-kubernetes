---
type: api-reference
title: ParamKind
description: ParamKind is a tuple of Group Kind and Version.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/param-kind-v1-admissionregistration/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/kubernetes-api/definitions/param-kind-v1-admissionregistration.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/param-kind-v1-admissionregistration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/param-kind-v1-admissionregistration/
    title: ParamKind | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/kubernetes-api/definitions/param-kind-v1-admissionregistration.md
api_metadata:
  apiVersion: admissionregistration.k8s.io/v1
  import: k8s.io/api/admissionregistration/v1
  kind: ParamKind
---

# ParamKind

`apiVersion: admissionregistration.k8s.io/v1`

`import "k8s.io/api/admissionregistration/v1"`

## ParamKind {#ParamKind}

ParamKind is a tuple of Group Kind and Version.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>apiVersion is the API group version the resources belong to. In format of "group/version". Required.</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>kind is the API kind the resources belong to. Required.</td>
    </tr>
  </tbody>
</table>
