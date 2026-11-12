---
type: glossary-term
title: Pod Security Policy
description: Removed API that enforced Pod security restrictions.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-security-policy
tags: [kubernetes, reference, glossary, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/pod-security-policy.md
    title: 'Kubernetes website source: reference/glossary/pod-security-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-pod-security-policy
    title: Pod Security Policy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/pod-security-policy.md
---

# Pod Security Policy

A former Kubernetes API that enforced security restrictions during [Pod](/reference/glossary/pod.md) creation and updates.

 

PodSecurityPolicy was deprecated as of Kubernetes v1.21, and removed in v1.25.
As an alternative, use [Pod Security Admission](/concepts/security/pod-security-admission.md) or a 3rd party admission plugin.
