---
type: glossary-term
title: StatefulSet
description: A StatefulSet manages deployment and scaling of a set of Pods, with durable storage and persistent identifiers for each Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-statefulset
tags: [kubernetes, reference, glossary, fundamental, core-object, workload, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/statefulset.md
    title: 'Kubernetes website source: reference/glossary/statefulset.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-statefulset
    title: StatefulSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/statefulset.md
---

# StatefulSet

Manages the deployment and scaling of a set of [Pods](/reference/glossary/pod.md), *and provides guarantees about the ordering and uniqueness* of these Pods.

 

Like a [Deployment](/reference/glossary/deployment.md), a StatefulSet manages Pods that are based on an identical container spec. Unlike a Deployment, a StatefulSet maintains a sticky identity for each of its Pods. These pods are created from the same spec, but are not interchangeable&#58; each has a persistent identifier that it maintains across any rescheduling.

If you want to use storage volumes to provide persistence for your workload, you can use a StatefulSet as part of the solution. Although individual Pods in a StatefulSet are susceptible to failure, the persistent Pod identifiers make it easier to match existing volumes to the new Pods that replace any that have failed.
