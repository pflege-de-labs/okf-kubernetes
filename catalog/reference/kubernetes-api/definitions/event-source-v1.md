---
type: api-reference
title: EventSource
description: EventSource contains information for an event.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/event-source-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/event-source-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/event-source-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/event-source-v1/
    title: EventSource | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/event-source-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: EventSource
---

# EventSource

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## EventSource {#EventSource}

EventSource contains information for an event.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>component</code><br/><em>string</em></td>
      <td>Component from which the event is generated.</td>
    </tr>
    <tr>
      <td><code>host</code><br/><em>string</em></td>
      <td>Node name on which the event is generated.</td>
    </tr>
  </tbody>
</table>
