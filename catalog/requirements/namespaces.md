---
type: constraint
title: Namespace naming and reserved namespaces
description: Namespace names must be RFC 1123 DNS labels; default, kube-node-lease, kube-public and kube-system exist initially, and the kube- prefix is reserved for system namespaces.
resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
tags: [kubernetes, requirements, namespaces, names, reserved]
generated: { by: "claude-code/claude-opus-5-5", at: "2026-11-05T00:00:00Z" }
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
sources:
  - id: ns-scope
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L27-L29
    title: Namespaces - scope of names
  - id: ns-default-note
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L38-L40
    title: Namespaces - avoid the default namespace in production
  - id: ns-initial
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L42-L56
    title: Namespaces - initial namespaces
  - id: ns-kube-prefix
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L63-L65
    title: Namespaces - kube- prefix reserved
  - id: ns-name
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L113-L116
    title: Namespaces - namespace names must be RFC 1123 DNS labels
---

# Namespace naming and reserved namespaces

Full text: [Namespaces](/concepts/overview/working-with-objects/namespaces.md),
[Share a Cluster with Namespaces](/tasks/administer-cluster/namespaces.md).

# Rules

- Namespace names must be valid RFC 1123 DNS labels (max 63 characters, see
  [object names](object-names.md)).[^ns-name]
- Avoid the `kube-` prefix; it is reserved for Kubernetes system namespaces.[^ns-kube-prefix]
- Names of resources are unique within a namespace, not across namespaces.
  Namespaces cannot be nested, and each resource is in at most one
  namespace.[^ns-scope]
- For a production cluster, consider not using the `default` namespace.[^ns-default-note]

# Initial namespaces

| Namespace | Purpose |
|---|---|
| `default` | Lets you use a new cluster without first creating a namespace |
| `kube-node-lease` | Holds the per-node Lease objects used for kubelet heartbeats |
| `kube-public` | Readable by all clients, including unauthenticated ones; public by convention only |
| `kube-system` | Objects created by the Kubernetes system |

Source: [^ns-initial]

[^ns-scope]: Namespaces - scope of names
[^ns-default-note]: Namespaces - avoid the default namespace in production
[^ns-initial]: Namespaces - initial namespaces
[^ns-kube-prefix]: Namespaces - kube- prefix reserved
[^ns-name]: Namespaces - namespace names must be RFC 1123 DNS labels
