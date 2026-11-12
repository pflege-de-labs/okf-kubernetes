---
type: concept
title: Container Runtime Interface (CRI)
description: The CRI is a plugin interface which enables the kubelet to use a wide variety of container runtimes, without having a need to recompile the cluster components.
resource: https://kubernetes.io/docs/concepts/containers/cri/
tags: [kubernetes, concepts, containers]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/containers/cri.md
    title: 'Kubernetes website source: concepts/containers/cri.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/containers/cri/
    title: Container Runtime Interface (CRI) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/containers/cri.md
---

# Container Runtime Interface (CRI)

The CRI is a plugin interface which enables the kubelet to use a wide variety of
container runtimes, without having a need to recompile the cluster components.

You need a working
[container runtime](/reference/glossary/container-runtime.md) on
each Node in your cluster, so that the
[kubelet](/reference/glossary/kubelet.md) can launch
[Pods](/reference/glossary/pod.md) and their containers.

The Container Runtime Interface (CRI) is the main protocol for the communication between the [kubelet](/reference/glossary/kubelet.md) and Container Runtime.

The Kubernetes Container Runtime Interface (CRI) defines the main
[gRPC](https://grpc.io) protocol for the communication between the
[node components](/concepts/architecture.md#node-components)
[kubelet](/reference/glossary/kubelet.md) and
[container runtime](/reference/glossary/container-runtime.md).

## The API {#api}

**Feature state:** Kubernetes v1.23 [stable]

The kubelet acts as a client when connecting to the container runtime via gRPC.
The runtime and image service endpoints have to be available in the container
runtime, which can be configured separately within the kubelet by using the
`--container-runtime-endpoint`
[command line flag](/reference/command-line-tools-reference/kubelet.md).

For Kubernetes v1.26 and later, the kubelet requires that the container runtime
supports the `v1` CRI API. If a container runtime does not support the `v1` API,
the kubelet will not register the node.

## Upgrading

When upgrading the Kubernetes version on a node, the kubelet restarts. If the
container runtime does not support the `v1` CRI API, the kubelet will fail to
register and report an error. If a gRPC re-dial is required because the container
runtime has been upgraded, the runtime must support the `v1` CRI API for the
connection to succeed. This might require a restart of the kubelet after the
container runtime is correctly configured.

## List streaming {#list-streaming}

**Feature state:** Kubernetes v1.36 [alpha] (feature gate `CRIListStreaming`, disabled by default)

The standard CRI list RPCs (`ListContainers`, `ListPodSandbox`, `ListImages`) return
all results in a single unary response. On nodes with a large number of containers
(for example, more than roughly 10,000 including both running and stopped), these
responses can exceed gRPC's default 16 MiB message size limit, causing the kubelet
to fail when reconciling state with the container runtime.

With the `CRIListStreaming` feature gate enabled, the kubelet uses server-side
streaming RPCs (such as `StreamContainers`, `StreamPodSandboxes`,
`StreamImages`) that allow the container runtime to divide results across
multiple response messages, bypassing the per-message size limit. This is
particularly useful for:

- High container churn environments (CI/CD systems)
- Large-scale batch processing workloads

If the container runtime does not support streaming RPCs, the kubelet
automatically falls back to the standard unary RPCs for backward
compatibility.

## What's next

- Learn more about the CRI [protocol definition](https://github.com/kubernetes/cri-api/blob/v0.33.1/pkg/apis/runtime/v1/api.proto)
