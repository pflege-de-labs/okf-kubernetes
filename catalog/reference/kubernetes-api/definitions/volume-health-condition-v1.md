---
type: api-reference
title: VolumeHealthCondition
description: VolumeHealthCondition represents an adverse health condition reported for a volume.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/volume-health-condition-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/volume-health-condition-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/volume-health-condition-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/volume-health-condition-v1/
    title: VolumeHealthCondition | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/volume-health-condition-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: VolumeHealthCondition
---

# VolumeHealthCondition

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## VolumeHealthCondition {#VolumeHealthCondition}

VolumeHealthCondition represents an adverse health condition reported for a volume.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>message</code><br/><em>string</em></td>
      <td>message is a human-readable description. Maximum permitted length of a message is 1024 bytes.</td>
    </tr>
    <tr>
      <td><code>reason</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>reason is a brief CamelCase machine-parseable reason. Together with status it forms the unique identity of a condition entry. Maximum permitted length of a reason is 256 bytes.</td>
    </tr>
    <tr>
      <td><code>status</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>status is the machine-parseable health category. Possible values: - "Inaccessible": the volume cannot be accessed. - "DataLoss": data loss has been detected on the volume. - "Degraded": the volume is functioning with reduced capability.<br/><br/>Possible enum values:<br/> - `"DataLoss"` indicates data loss has been detected on the volume.<br/> - `"Degraded"` indicates the volume is functioning but with reduced capability.<br/> - `"Inaccessible"` indicates the volume cannot be accessed.</td>
    </tr>
  </tbody>
</table>
