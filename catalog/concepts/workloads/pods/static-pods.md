---
type: concept
title: Static Pods
description: _Static Pods_ are managed directly by the kubelet daemon on a specific node, without the API server observing them.
resource: https://kubernetes.io/docs/concepts/workloads/pods/static-pods/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/workloads/pods/static-pods.md
    title: 'Kubernetes website source: concepts/workloads/pods/static-pods.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/pods/static-pods/
    title: Static Pods | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/workloads/pods/static-pods.md
---

# Static Pods

_Static Pods_ are managed directly by the kubelet daemon on a specific node,
without the [API server](/reference/glossary/kube-apiserver.md)
observing them.
Unlike Pods that are managed by the control plane (for example, a
[Deployment](/reference/glossary/deployment.md)),
the kubelet watches each static Pod and restarts it if it fails.

Static Pods are always bound to one [kubelet](/reference/glossary/kubelet.md) on a specific node.

The main use for static Pods is to run a self-hosted control plane: in other words,
using the kubelet to supervise the individual
[control plane components](/concepts/overview/components.md#control-plane-components).
For example, [kubeadm](/reference/setup-tools/kubeadm.md) uses static Pods to run
`kube-apiserver`, `kube-controller-manager`, `kube-scheduler`, and `etcd` on control plane nodes.

**Note:**
If your cluster runs control plane components as Pods, they are likely
static Pods. You can recognize their mirror Pods in the `kube-system` namespace
by the `kubernetes.io/config.mirror` annotation.

## Mirror Pods {#mirror-pods}

The kubelet automatically tries to create a
[mirror Pod](/reference/glossary/mirror-pod.md)
on the Kubernetes API server for each static Pod.
This means that the Pods running on a node are visible on the API server,
but cannot be controlled from there.
The Pod names will be suffixed with the node hostname with a leading hyphen.

The kubelet propagates [labels](/reference/glossary/label.md)
from the static Pod to the mirror Pod. You can use those labels as normal via
[selectors](/reference/glossary/selector.md).

If you try to use `kubectl` to delete the mirror Pod from the API server,
the kubelet _does not_ remove the static Pod. The kubelet will recreate
the mirror Pod.

## Limitations {#limitations}

The spec of a static Pod cannot refer to other API objects,
such as [ServiceAccount](/reference/glossary/service-account.md),
[ConfigMap](/reference/glossary/configmap.md), or
[Secret](/reference/glossary/secret.md).

Static Pods do not support [ephemeral containers](/concepts/workloads/pods/ephemeral-containers.md).

## Static Pods vs DaemonSets {#static-pods-vs-daemonsets}

If you are running clustered Kubernetes and are using static Pods to run a Pod
on every node, you should probably be using a
[DaemonSet](/reference/glossary/daemonset.md) instead.

Static Pods are not managed by the control plane, so they cannot be rolled out,
rolled back, or scaled using standard Kubernetes mechanisms. DaemonSets provide
these capabilities and are the recommended approach for running node-level workloads.

Static Pods are started by the kubelet before the API server is available, which
makes them suitable for bootstrapping control plane components. DaemonSets require
a running control plane.

## What's next

- Learn how to [create static Pods](/tasks/configure-pod-container/static-pod.md).
- Learn about [Kubernetes components](/concepts/overview/components.md) and how the control plane uses static Pods.
- Learn about [DaemonSets](/concepts/workloads/controllers/daemonset.md) as an alternative to static Pods.
