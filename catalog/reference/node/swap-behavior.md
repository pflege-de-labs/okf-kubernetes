---
type: reference
title: Linux Node Swap Behaviors
description: 'To allow Kubernetes workloads to use swap, on a Linux node, you must disable the kubelet''s default behavior of failing when swap is detected, and specify memory-swap behavior as `LimitedSwap`:'
resource: https://kubernetes.io/docs/reference/node/swap-behavior/
tags: [kubernetes, reference, node]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/node/swap-behavior.md
    title: 'Kubernetes website source: reference/node/swap-behavior.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/node/swap-behavior/
    title: Linux Node Swap Behaviors | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/node/swap-behavior.md
---

# Linux Node Swap Behaviors

To allow Kubernetes workloads to use swap, on a Linux node,
you must disable the kubelet's default behavior of failing when swap is detected,
and specify memory-swap behavior as `LimitedSwap`:

The available choices for swap behavior are:

`NoSwap`
: (default) Workloads running as Pods on this node do not and cannot use swap. However, processes
  outside of Kubernetes' scope, such as system daemons (including the kubelet itself!) **can** utilize swap.
  This behavior is beneficial for protecting the node from system-level memory spikes,
  but it does not safeguard the workloads themselves from such spikes.

`LimitedSwap`
: Kubernetes workloads can utilize swap memory. The amount of swap available to a Pod is determined automatically.

To learn more, read [swap memory management](/concepts/cluster-administration/swap-memory-management.md).
