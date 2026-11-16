---
type: glossary-term
title: kOps (Kubernetes Operations)
description: kOps will not only help you create, destroy, upgrade and maintain production-grade, highly available, Kubernetes cluster, but it will also provision the necessary cloud infrastructure.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kops
tags: [kubernetes, reference, glossary, tool, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/kops.md
    title: 'Kubernetes website source: reference/glossary/kops.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-kops
    title: kOps (Kubernetes Operations) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/kops.md
---

# kOps (Kubernetes Operations)

`kOps` will not only help you create, destroy, upgrade and maintain production-grade, highly available, Kubernetes cluster, but it will also provision the necessary cloud infrastructure.

 

**Note:**
AWS (Amazon Web Services) is currently officially supported, with DigitalOcean, GCE and OpenStack in beta support, and Azure in alpha.

`kOps` is an automated provisioning system:
  * Fully automated installation
  * Uses DNS to identify clusters
  * Self-healing: everything runs in Auto-Scaling Groups
  * Multiple OS support (Amazon Linux, Debian, Flatcar, RHEL, Rocky and Ubuntu)
  * High-Availability support
  * Can directly provision, or generate terraform manifests
