---
type: glossary-term
title: Controller
description: A control loop that watches the shared state of the cluster through the apiserver and makes changes attempting to move the current state towards the desired state.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-controller
tags: [kubernetes, reference, glossary, architecture, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/controller.md
    title: 'Kubernetes website source: reference/glossary/controller.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-controller
    title: Controller | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/controller.md
---

# Controller

In Kubernetes, controllers are control loops that watch the state of your
[cluster](/reference/glossary/cluster.md), then make or request
changes where needed.
Each controller tries to move the current cluster state closer to the desired
state.

Controllers watch the shared state of your cluster through the
[apiserver](/reference/glossary/kube-apiserver.md) (part of the
[Control Plane](/reference/glossary/control-plane.md)).

Some controllers also run inside the control plane, providing control loops that
are core to Kubernetes' operations. For example: the deployment controller, the
daemonset controller, the namespace controller, and the persistent volume
controller (and others) all run within the
[kube-controller-manager](/reference/glossary/kube-controller-manager.md).
