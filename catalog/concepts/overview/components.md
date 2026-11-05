---
type: concept
title: Kubernetes Components
description: An overview of the key components that make up a Kubernetes cluster.
resource: https://kubernetes.io/docs/concepts/overview/components/
tags: [kubernetes, concepts, overview]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/components.md
    title: 'Kubernetes website source: concepts/overview/components.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/overview/components/
    title: Kubernetes Components | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/overview/components.md
---

# Kubernetes Components

This page provides a high-level overview of the essential components that make up a Kubernetes cluster.

![Components of Kubernetes](https://kubernetes.io/images/docs/components-of-kubernetes.svg)

*The components of a Kubernetes cluster*

## Core Components

A Kubernetes cluster consists of a control plane and one or more worker nodes.
Here's a brief overview of the main components:

### Control Plane Components

Manage the overall state of the cluster:

[kube-apiserver](/concepts/architecture.md#kube-apiserver)
: The core component server that exposes the Kubernetes HTTP API.

[etcd](/concepts/architecture.md#etcd)
: Consistent and highly-available key value store for all API server data.

[kube-scheduler](/concepts/architecture.md#kube-scheduler)
: Looks for Pods not yet bound to a node, and assigns each Pod to a suitable node.

[kube-controller-manager](/concepts/architecture.md#kube-controller-manager)
: Runs [controllers](/reference/glossary/controller.md) to implement Kubernetes API behavior.

[cloud-controller-manager](/concepts/architecture.md#cloud-controller-manager) (optional)
: Integrates with underlying cloud provider(s).

### Node Components

Run on every node, maintaining running pods and providing the Kubernetes runtime environment:

[kubelet](/concepts/architecture.md#kubelet)
: Ensures that Pods are running, including their containers.

[kube-proxy](/concepts/architecture.md#kube-proxy) (optional)
: Maintains network rules on nodes to implement [Services](/reference/glossary/service.md).

[Container runtime](/concepts/architecture.md#container-runtime)
: Software responsible for running containers. Read
  [Container Runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/) to learn more.

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

Your cluster may require additional software on each node; for example, you might also
run [systemd](https://systemd.io/) on a Linux node to supervise local components.

## Addons

Addons extend the functionality of Kubernetes. A few important examples include:

[DNS](/concepts/architecture.md#dns)
: For cluster-wide DNS resolution.

[Web UI](/concepts/architecture.md#web-ui-dashboard) (Dashboard)
: For cluster management via a web interface.

[Container Resource Monitoring](/concepts/architecture.md#container-resource-monitoring)
: For collecting and storing container metrics.

[Cluster-level Logging](/concepts/architecture.md#cluster-level-logging)
: For saving container logs to a central log store.

## Flexibility in Architecture

Kubernetes allows for flexibility in how these components are deployed and managed.
The architecture can be adapted to various needs, from small development environments
to large-scale production deployments.

For more detailed information about each component and various ways to configure your
cluster architecture, see the [Cluster Architecture](/concepts/architecture.md) page.
