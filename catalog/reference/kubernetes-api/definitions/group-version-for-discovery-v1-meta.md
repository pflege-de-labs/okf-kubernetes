---
type: api-reference
title: GroupVersionForDiscovery
description: GroupVersion contains the &#34;group/version&#34; and &#34;version&#34; string of a version. It is made a struct to keep extensibility.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/group-version-for-discovery-v1-meta/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/definitions/group-version-for-discovery-v1-meta.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/group-version-for-discovery-v1-meta.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/group-version-for-discovery-v1-meta/
    title: GroupVersionForDiscovery | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/kubernetes-api/definitions/group-version-for-discovery-v1-meta.md
api_metadata:
  apiVersion: meta/v1
  import: k8s.io/apimachinery/pkg/apis/meta/v1
  kind: GroupVersionForDiscovery
---

# GroupVersionForDiscovery

`apiVersion: meta/v1`

`import "k8s.io/apimachinery/pkg/apis/meta/v1"`

## GroupVersionForDiscovery {#GroupVersionForDiscovery}

GroupVersion contains the &#34;group/version&#34; and &#34;version&#34; string of a version. It is made a struct to keep extensibility.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>groupVersion</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>groupVersion specifies the API group and version in the form "group/version"</td>
    </tr>
    <tr>
      <td><code>version</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.</td>
    </tr>
  </tbody>
</table>
