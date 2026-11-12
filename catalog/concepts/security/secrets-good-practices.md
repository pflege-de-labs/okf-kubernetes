---
type: concept
title: Good practices for Kubernetes Secrets
description: Principles and practices for good Secret management for cluster administrators and application developers.
resource: https://kubernetes.io/docs/concepts/security/secrets-good-practices/
tags: [kubernetes, concepts, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/security/secrets-good-practices.md
    title: 'Kubernetes website source: concepts/security/secrets-good-practices.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/security/secrets-good-practices/
    title: Good practices for Kubernetes Secrets | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/security/secrets-good-practices.md
---

# Good practices for Kubernetes Secrets

In Kubernetes, a Secret is an object that stores sensitive information, such as passwords, OAuth tokens, and SSH keys.

Secrets give you more control over how sensitive information is used and reduces
the risk of accidental exposure. Secret values are encoded as base64 strings and
are stored unencrypted by default, but can be configured to be
[encrypted at rest](/tasks/administer-cluster/encrypt-data.md#ensure-all-secrets-are-encrypted).

A [Pod](/reference/glossary/pod.md) can reference the Secret in
a variety of ways, such as in a volume mount or as an environment variable.
Secrets are designed for confidential data and
[ConfigMaps](/tasks/configure-pod-container/configure-pod-configmap.md) are
designed for non-confidential data.

The following good practices are intended for both cluster administrators and
application developers. Use these guidelines to improve the security of your
sensitive information in Secret objects, as well as to more effectively manage
your Secrets.

## Cluster administrators

This section provides good practices that cluster administrators can use to
improve the security of confidential information in the cluster.

### Configure encryption at rest

By default, Secret objects are stored unencrypted in [etcd](/reference/glossary/etcd.md). You should configure encryption of your Secret
data in `etcd`. For instructions, refer to
[Encrypt Secret Data at Rest](/tasks/administer-cluster/encrypt-data.md).

### Configure least-privilege access to Secrets {#least-privilege-secrets}

When planning your access control mechanism, such as Kubernetes
[Role-based Access Control](/reference/glossary/rbac.md) [(RBAC)](/reference/access-authn-authz/rbac.md),
consider the following guidelines for access to `Secret` objects. You should
also follow the other guidelines in
[RBAC good practices](/concepts/security/rbac-good-practices.md).

- **Components**: Restrict `watch` or `list` access to only the most
  privileged, system-level components. Only grant `get` access for Secrets if
  the component's normal behavior requires it.
- **Humans**: Restrict `get`, `watch`, or `list` access to Secrets. Only allow
  cluster administrators to access `etcd`. This includes read-only access. For
  more complex access control, such as restricting access to Secrets with
  specific annotations, consider using third-party authorization mechanisms.

**Caution:**
Granting `list` access to Secrets implicitly lets the subject fetch the
contents of the Secrets.

A user who can create a Pod that uses a Secret can also see the value of that
Secret. Even if cluster policies do not allow a user to read the Secret
directly, the same user could have access to run a Pod that then exposes the
Secret. You can detect or limit the impact caused by Secret data being exposed,
either intentionally or unintentionally, by a user with this access. Some
recommendations include:

*  Use short-lived Secrets
*  Implement audit rules that alert on specific events, such as concurrent
   reading of multiple Secrets by a single user

#### Restrict Access for Secrets
Use separate namespaces to isolate access to mounted secrets.

### Improve etcd management policies

Consider wiping or shredding the durable storage used by `etcd` once it is
no longer in use.

If there are multiple `etcd` instances, configure encrypted SSL/TLS
communication between the instances to protect the Secret data in transit.

### Configure access to external Secrets

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

You can use third-party Secrets store providers to keep your confidential data
outside your cluster and then configure Pods to access that information.
The [Kubernetes Secrets Store CSI Driver](https://secrets-store-csi-driver.sigs.k8s.io/)
is a DaemonSet that lets the kubelet retrieve Secrets from external stores, and
mount the Secrets as a volume into specific Pods that you authorize to access
the data.

For a list of supported providers, refer to
[Providers for the Secret Store CSI Driver](https://secrets-store-csi-driver.sigs.k8s.io/concepts.html#provider-for-the-secrets-store-csi-driver).

## Good practices for using swap memory

For best practices for setting swap memory for Linux nodes, please refer to
[swap memory management](/concepts/cluster-administration/swap-memory-management.md#good-practice-for-using-swap-in-a-kubernetes-cluster).

## Developers

This section provides good practices for developers to use to improve the
security of confidential data when building and deploying Kubernetes resources.

### Restrict Secret access to specific containers

If you are defining multiple containers in a Pod, and only one of those
containers needs access to a Secret, define the volume mount or environment
variable configuration so that the other containers do not have access to that
Secret.

### Protect Secret data after reading

Applications still need to protect the value of confidential information after
reading it from an environment variable or volume. For example, your
application must avoid logging the secret data in the clear or transmitting it
to an untrusted party.

### Avoid sharing Secret manifests

If you configure a Secret through a
[manifest](/reference/glossary/manifest.md), with the secret
data encoded as base64, sharing this file or checking it in to a source
repository means the secret is available to everyone who can read the manifest.

**Caution:**
Base64 encoding is _not_ an encryption method, it provides no additional
confidentiality over plain text.
