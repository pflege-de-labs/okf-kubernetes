---
type: api-reference
title: ServerAddressByClientCIDR
description: ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/server-address-by-client-cidr-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/server-address-by-client-cidr-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/server-address-by-client-cidr-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/server-address-by-client-cidr-v1-meta/
    title: ServerAddressByClientCIDR | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/server-address-by-client-cidr-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: ServerAddressByClientCIDR
---

# ServerAddressByClientCIDR

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## ServerAddressByClientCIDR {#ServerAddressByClientCIDR}

ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>clientCIDR</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>The CIDR with which clients can match their IP to figure out the server address that they should use.</td>
    </tr>
    <tr>
      <td><code>serverAddress</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.</td>
    </tr>
  </tbody>
</table>
