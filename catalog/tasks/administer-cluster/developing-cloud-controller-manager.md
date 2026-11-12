---
type: concept
title: Developing Cloud Controller Manager
description: The cloud-controller-manager is a Kubernetes control plane component that embeds cloud-specific control logic.
resource: https://kubernetes.io/docs/tasks/administer-cluster/developing-cloud-controller-manager/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/developing-cloud-controller-manager.md
    title: 'Kubernetes website source: tasks/administer-cluster/developing-cloud-controller-manager.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/developing-cloud-controller-manager/
    title: Developing Cloud Controller Manager | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/developing-cloud-controller-manager.md
---

# Developing Cloud Controller Manager

**Feature state:** Kubernetes v1.11 [beta]

The cloud-controller-manager is a Kubernetes [control plane](/reference/glossary/control-plane.md) component
that embeds cloud-specific control logic. The cloud controller manager lets you link your
cluster into your cloud provider's API, and separates out the components that interact
with that cloud platform from components that only interact with your cluster.

By decoupling the interoperability logic between Kubernetes and the underlying cloud
infrastructure, the cloud-controller-manager component enables cloud providers to release
features at a different pace compared to the main Kubernetes project.

## Background

Since cloud providers develop and release at a different pace compared to the Kubernetes project, abstracting the provider-specific code to the `cloud-controller-manager` binary allows cloud vendors to evolve independently from the core Kubernetes code.

The Kubernetes project provides skeleton cloud-controller-manager code with Go interfaces to allow you (or your cloud provider) to plug in your own implementations. This means that a cloud provider can implement a cloud-controller-manager by importing packages from Kubernetes core; each cloudprovider will register their own code by calling `cloudprovider.RegisterCloudProvider` to update a global variable of available cloud providers.

## Developing

### Out of tree

To build an out-of-tree cloud-controller-manager for your cloud:

1. Create a go package with an implementation that satisfies [cloudprovider.Interface](https://github.com/kubernetes/cloud-provider/blob/master/cloud.go).
2. Use [`main.go` in cloud-controller-manager](https://github.com/kubernetes/kubernetes/blob/master/cmd/cloud-controller-manager/main.go) from Kubernetes core as a template for your `main.go`. As mentioned above, the only difference should be the cloud package that will be imported.
3. Import your cloud package in `main.go`, ensure your package has an `init` block to run [`cloudprovider.RegisterCloudProvider`](https://github.com/kubernetes/cloud-provider/blob/master/plugins.go).

Many cloud providers publish their controller manager code as open source. If you are creating
a new cloud-controller-manager from scratch, you could take an existing out-of-tree cloud
controller manager as your starting point.

### In tree

For in-tree cloud providers, you can run the in-tree cloud controller manager as a [DaemonSet](/reference/glossary/daemonset.md) in your cluster. See [Cloud Controller Manager Administration](/tasks/administer-cluster/running-cloud-controller.md) for more details.
