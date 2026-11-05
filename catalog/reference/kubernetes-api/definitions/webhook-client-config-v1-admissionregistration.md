---
type: api-reference
title: WebhookClientConfig
description: WebhookClientConfig contains the information to make a TLS connection with the webhook
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/webhook-client-config-v1-admissionregistration/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/webhook-client-config-v1-admissionregistration.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/webhook-client-config-v1-admissionregistration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/webhook-client-config-v1-admissionregistration/
    title: WebhookClientConfig | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/webhook-client-config-v1-admissionregistration.md
api_metadata:
  apiVersion: admissionregistration.k8s.io/v1
  import: k8s.io/api/admissionregistration/v1
  kind: WebhookClientConfig
---

# WebhookClientConfig

`apiVersion: admissionregistration.k8s.io/v1`

`import "k8s.io/api/admissionregistration/v1"`

## WebhookClientConfig {#WebhookClientConfig}

WebhookClientConfig contains the information to make a TLS connection with the webhook

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>caBundle</code><br/><em>string</em></td>
      <td>caBundle is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.</td>
    </tr>
    <tr>
      <td><code>service</code><br/><em><a href="/reference/kubernetes-api/apiregistration/api-service-v1.md#ServiceReference">ServiceReference</a></em></td>
      <td>service is a reference to the service for this webhook. Either `service` or `url` must be specified.  If the webhook is running within the cluster, then you should use `service`.</td>
    </tr>
    <tr>
      <td><code>url</code><br/><em>string</em></td>
      <td>url gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.  The `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.  Please note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be "https"; the URL must begin with "https://".  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. "user:password@" is not allowed. Fragments ("#...") and query parameters ("?...") are not allowed, either.</td>
    </tr>
  </tbody>
</table>
