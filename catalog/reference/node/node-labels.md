---
type: reference
title: Node Labels Populated By The Kubelet
description: Kubernetes nodes come pre-populated with a standard set of labels.
resource: https://kubernetes.io/docs/reference/node/node-labels/
tags: [kubernetes, reference, node]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/node/node-labels.md
    title: 'Kubernetes website source: reference/node/node-labels.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/node/node-labels/
    title: Node Labels Populated By The Kubelet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/node/node-labels.md
---

# Node Labels Populated By The Kubelet

Kubernetes [nodes](/reference/glossary/node.md) come pre-populated
with a standard set of [labels](/reference/glossary/label.md).

You can also set your own labels on nodes, either through the kubelet configuration or
using the Kubernetes API.

## Preset labels

The preset labels that Kubernetes sets on nodes are:

* [`kubernetes.io/arch`](/reference/labels-annotations-taints.md#kubernetes-io-arch)
* [`kubernetes.io/hostname`](/reference/labels-annotations-taints.md#kubernetesiohostname)
* [`kubernetes.io/os`](/reference/labels-annotations-taints.md#kubernetes-io-os)
* [`node.kubernetes.io/instance-type`](/reference/labels-annotations-taints.md#nodekubernetesioinstance-type)
  (if known to the kubelet &ndash; Kubernetes may not have this information to set the label)
* [`topology.kubernetes.io/region`](/reference/labels-annotations-taints.md#topologykubernetesioregion)
  (if known to the kubelet &ndash; Kubernetes may not have this information to set the label)
* [`topology.kubernetes.io/zone`](/reference/labels-annotations-taints.md#topologykubernetesiozone)
  (if known to the kubelet &ndash; Kubernetes may not have this information to set the label)

**Note:**
The value of these labels is cloud provider specific and is not guaranteed to be reliable.
For example, the value of `kubernetes.io/hostname` may be the same as the node name in some environments
and a different value in other environments.

## What's next

- See [Well-Known Labels, Annotations and Taints](/reference/labels-annotations-taints.md) for a list of common labels.
- Learn how to [add a label to a node](/tasks/configure-pod-container/assign-pods-nodes.md#add-a-label-to-a-node).
