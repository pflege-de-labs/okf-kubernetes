---
type: glossary-term
title: Control Plane
description: The container orchestration layer that exposes the API and interfaces to define, deploy, and manage the lifecycle of containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-control-plane
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/control-plane.md
    title: 'Kubernetes website source: reference/glossary/control-plane.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-control-plane
    title: Control Plane | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/control-plane.md
---

# Control Plane

The container orchestration layer that exposes the API and interfaces to define, deploy, and manage the lifecycle of containers.

  
 
 This layer is composed by many different components, such as (but not restricted to):

 * [etcd](/reference/glossary/etcd.md)
 * [API Server](/reference/glossary/kube-apiserver.md)
 * [Scheduler](/reference/glossary/kube-scheduler.md)
 * [Controller Manager](/reference/glossary/kube-controller-manager.md)
 * [Cloud Controller Manager](/reference/glossary/cloud-controller-manager.md)

 These components can be run as traditional operating system services (daemons) or as containers. The hosts running these components were historically called [masters](/reference/glossary/master.md).
