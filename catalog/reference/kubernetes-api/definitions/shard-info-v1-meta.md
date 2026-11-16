---
type: api-reference
title: ShardInfo
description: ShardInfo describes the shard selector that was applied to produce a list response. Its presence on a list response indicates the list is a filtered subset.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/shard-info-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/kubernetes-api/definitions/shard-info-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/shard-info-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/shard-info-v1-meta/
    title: ShardInfo | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/kubernetes-api/definitions/shard-info-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: ShardInfo
---

# ShardInfo

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## ShardInfo {#ShardInfo}

ShardInfo describes the shard selector that was applied to produce a list response. Its presence on a list response indicates the list is a filtered subset.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>selector</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>selector is the shard selector string from the request, echoed back so clients can verify which shard they received and merge responses from multiple shards.</td>
    </tr>
  </tbody>
</table>
