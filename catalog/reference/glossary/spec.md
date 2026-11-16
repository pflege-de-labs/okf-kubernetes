---
type: glossary-term
title: Spec
description: Field in Kubernetes manifests that defines the desired state or configuration.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-spec
tags: [kubernetes, reference, glossary, fundamental, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/spec.md
    title: 'Kubernetes website source: reference/glossary/spec.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-spec
    title: Spec | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/spec.md
---

# Spec

Defines how each object, like Pods or Services, should be configured and its desired state.

Almost every Kubernetes object includes two nested object fields that govern the object's configuration: the object spec and the object status. For objects that have a spec, you have to set this when you create the object, providing a description of the characteristics you want the [resource](/reference/glossary/api-resource.md) to have: its desired state.

It varies for different objects like Pods, StatefulSets, and Services, detailing settings such as containers, volumes, replicas, ports,
and other specifications unique to each object type. This field encapsulates what state Kubernetes should maintain for the defined
object.
