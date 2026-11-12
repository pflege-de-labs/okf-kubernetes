---
type: glossary-term
title: RBAC (Role-Based Access Control)
description: Manages authorization decisions, allowing admins to dynamically configure access policies through the Kubernetes API.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-rbac
tags: [kubernetes, reference, glossary, security, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/rbac.md
    title: 'Kubernetes website source: reference/glossary/rbac.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-rbac
    title: RBAC (Role-Based Access Control) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/rbac.md
---

# RBAC (Role-Based Access Control)

Manages authorization decisions, allowing admins to dynamically configure access policies through the [Kubernetes API](/reference/glossary/kubernetes-api.md).

 

RBAC utilizes four kinds of Kubernetes objects:

Role
: Defines permission rules in a specific namespace.

ClusterRole
: Defines permission rules cluster-wide.

RoleBinding
: Grants the permissions defined in a role to a set of users in a specific namespace.

ClusterRoleBinding
: Grants the permissions defined in a role to a set of users cluster-wide.

For more information, see [RBAC](/reference/access-authn-authz/rbac.md).
