---
type: glossary-term
title: RBAC (Role-Based Access Control)
description: Manages authorization decisions, allowing admins to dynamically configure access policies through the Kubernetes API.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-rbac
tags: [kubernetes, reference, glossary, security, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/rbac.md
    title: 'Kubernetes website source: reference/glossary/rbac.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-rbac
    title: RBAC (Role-Based Access Control) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
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
