---
type: api-reference
title: GroupResource
description: GroupResource specifies a Group and a Resource, but does not force a version. This is useful for identifying concepts during lookup stages without having partially valid types
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/group-resource-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/kubernetes-api/definitions/group-resource-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/group-resource-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/group-resource-v1-meta/
    title: GroupResource | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/kubernetes-api/definitions/group-resource-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: GroupResource
---

# GroupResource

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## GroupResource {#GroupResource}

GroupResource specifies a Group and a Resource, but does not force a version.  This is useful for identifying concepts during lookup stages without having partially valid types

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>group</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td></td>
    </tr>
    <tr>
      <td><code>resource</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td></td>
    </tr>
  </tbody>
</table>
