---
type: section
title: API Access Control
description: For an introduction to how Kubernetes implements and controls API access, read Controlling Access to the Kubernetes API.
resource: https://kubernetes.io/docs/reference/access-authn-authz/
tags: [kubernetes, reference]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/access-authn-authz/_index.md
    title: 'Kubernetes website source: reference/access-authn-authz/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/access-authn-authz/
    title: API Access Control | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/access-authn-authz/_index.md
---

# API Access Control

For an introduction to how Kubernetes implements and controls API access,
read [Controlling Access to the Kubernetes API](/concepts/security/controlling-access.md).

Reference documentation:

- [Authenticating](/reference/access-authn-authz/authentication.md)
   - [Authenticating with Bootstrap Tokens](/reference/access-authn-authz/bootstrap-tokens.md)
- [Admission Controllers](/reference/access-authn-authz/admission-controllers.md)
   - [Dynamic Admission Control](/reference/access-authn-authz/extensible-admission-controllers.md)
   - [Manifest-Based Admission Control](/reference/access-authn-authz/manifest-admission-control.md)
- [Authorization](/reference/access-authn-authz/authorization.md)
   - [Role Based Access Control](/reference/access-authn-authz/rbac.md)
   - [Attribute Based Access Control](/reference/access-authn-authz/abac.md)
   - [Node Authorization](/reference/access-authn-authz/node.md)
   - [Webhook Authorization](/reference/access-authn-authz/webhook.md)
- [Certificate Signing Requests](/reference/access-authn-authz/certificate-signing-requests.md)
   - including [CSR approval](/reference/access-authn-authz/certificate-signing-requests.md#approval-rejection)
     and [certificate signing](/reference/access-authn-authz/certificate-signing-requests.md#signing)
- Service accounts
  - [Developer guide](/tasks/configure-pod-container/configure-service-account.md)
  - [Administration](/reference/access-authn-authz/service-accounts-admin.md)
- [Kubelet Authentication & Authorization](/reference/access-authn-authz/kubelet-authn-authz.md)
  - including kubelet [TLS bootstrapping](/reference/access-authn-authz/kubelet-tls-bootstrapping.md)
