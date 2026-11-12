---
type: task
title: Use Kube-router for NetworkPolicy
description: This page shows how to use Kube-router for NetworkPolicy.
resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/kube-router-network-policy/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/network-policy-provider/kube-router-network-policy.md
    title: 'Kubernetes website source: tasks/administer-cluster/network-policy-provider/kube-router-network-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/kube-router-network-policy/
    title: Use Kube-router for NetworkPolicy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/network-policy-provider/kube-router-network-policy.md
---

# Use Kube-router for NetworkPolicy

This page shows how to use [Kube-router](https://github.com/cloudnativelabs/kube-router) for NetworkPolicy.

## Before you begin

You need to have a Kubernetes cluster running. If you do not already have a cluster, you can create one by using any of the cluster installers like Kops, Bootkube, Kubeadm etc.

## Installing Kube-router addon
The Kube-router Addon comes with a Network Policy Controller that watches Kubernetes API server for any NetworkPolicy and pods updated and configures iptables rules and ipsets to allow or block traffic as directed by the policies. Please follow the [trying Kube-router with cluster installers](https://www.kube-router.io/docs/user-guide/#try-kube-router-with-cluster-installers) guide to install Kube-router addon.

## What's next

Once you have installed the Kube-router addon, you can follow the [Declare Network Policy](/tasks/administer-cluster/declare-network-policy.md) to try out Kubernetes NetworkPolicy.
