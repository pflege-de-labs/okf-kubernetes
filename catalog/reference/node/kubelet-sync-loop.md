---
type: reference
title: Kubelet Sync Loop
description: The kubelet is the primary "node agent" that creates and watches Pods on each node.
resource: https://kubernetes.io/docs/reference/node/kubelet-sync-loop/
tags: [kubernetes, reference, node]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/node/kubelet-sync-loop.md
    title: 'Kubernetes website source: reference/node/kubelet-sync-loop.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/node/kubelet-sync-loop/
    title: Kubelet Sync Loop | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/node/kubelet-sync-loop.md
---

# Kubelet Sync Loop

The [kubelet](/reference/command-line-tools-reference/kubelet.md) is the
primary "node agent" that creates and watches Pods on each node. The `kubelet`
runs a sync loop that periodically reconciles the desired state (a Pod spec)
with the actual state of the running containers.

1.  _Sync Loop_: The Sync Loop queues work (aggregated from many sources) for
    the Pods assigned to its node (where `nodeName` matches the node). Over the
    course of each loop, subprocesses called pod workers will attempt to
    reconcile the desired state of these Pods against the current state of the
    running containers.
2.  _Sync Pod_: The majority of the `kubelet` logic is stored in a suite of
    functions within the `podSyncer` interface, including the `SyncPod` function
    and its variants (like `SyncTerminatingPod` and `SyncTerminatedPod`). During
    each Sync Loop, a relevant `podSyncer` function will be executed for each Pod
    in an attempt to drive its state on the node toward the desired state.
3.  _[Container Runtime Interface](/reference/glossary/cri.md)
    (CRI)_: To actually run the containers, the `kubelet` uses the CRI to talk
    to a container runtime (like containerd or CRI-O). The `kubelet` acts as the
    client, instructing the runtime to create a "pod sandbox" and then
    create/start the individual containers defined in the Pod spec.
4.  _PLEG (Pod Lifecycle Event Generator)_: The `kubelet` needs to know when
    containers start, stop, or fail. It relies on a component called PLEG to
    periodically poll the runtime for the standard state of all containers. PLEG
    generates events that wake up the Sync Loop to update the Pod status.

Because of this polling mechanism, the status seen in the API (like `kubectl get
pod`) might have a slight delay compared to the instant reality on the node.
