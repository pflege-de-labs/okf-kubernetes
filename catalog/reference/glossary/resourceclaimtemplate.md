---
type: glossary-term
title: ResourceClaimTemplate
description: Defines a template for Kubernetes to create ResourceClaims. Used to provide per-Pod or per-PodGroup access to separate, similar resources.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resourceclaimtemplate
tags: [kubernetes, reference, glossary, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/resourceclaimtemplate.md
    title: 'Kubernetes website source: reference/glossary/resourceclaimtemplate.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-resourceclaimtemplate
    title: ResourceClaimTemplate | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/resourceclaimtemplate.md
---

# ResourceClaimTemplate

Defines a template that Kubernetes uses to create
[ResourceClaims](/reference/glossary/resourceclaim.md).
ResourceClaimTemplates are used in
[dynamic resource allocation (DRA)](/concepts/resource-management/dynamic-resource-allocation.md)
to provide _per-Pod or per-[PodGroup](/reference/glossary/podgroup.md) access to separate, similar resources_.

When a ResourceClaimTemplate is referenced in a workload specification,
Kubernetes automatically creates ResourceClaim objects based on the template.
Each ResourceClaim is bound to a specific Pod or PodGroup. When the Pod
terminates or the PodGroup is deleted, Kubernetes deletes the corresponding
ResourceClaim. PodGroup ResourceClaimTemplates require the
[`DRAWorkloadResourceClaims`](/reference/command-line-tools-reference/feature-gates.md#DRAWorkloadResourceClaims)
feature to be enabled.
