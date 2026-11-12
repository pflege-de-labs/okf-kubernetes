---
type: concept
title: Pod Security Policies
description: PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25.
resource: https://kubernetes.io/docs/concepts/security/pod-security-policy/
tags: [kubernetes, concepts, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/security/pod-security-policy.md
    title: 'Kubernetes website source: concepts/security/pod-security-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/security/pod-security-policy/
    title: Pod Security Policies | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/security/pod-security-policy.md
---

# Pod Security Policies

**Removed feature**
PodSecurityPolicy was [deprecated](https://kubernetes.io/blog/2021/04/08/kubernetes-1-21-release-announcement/#podsecuritypolicy-deprecation)
in Kubernetes v1.21, and removed from Kubernetes in v1.25.

Instead of using PodSecurityPolicy, you can enforce similar restrictions on Pods using
either or both:

- [Pod Security Admission](/concepts/security/pod-security-admission.md)
- a 3rd party admission plugin, that you deploy and configure yourself

For a migration guide, see [Migrate from PodSecurityPolicy to the Built-In PodSecurity Admission Controller](/tasks/configure-pod-container/migrate-from-psp.md).
For more information on the removal of this API,
see [PodSecurityPolicy Deprecation: Past, Present, and Future](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/).

If you are not running Kubernetes v1.37, check the documentation for
your version of Kubernetes.
