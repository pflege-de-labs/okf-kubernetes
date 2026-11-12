---
type: task
title: Use Antrea for NetworkPolicy
description: This page shows how to install and use Antrea CNI plugin on Kubernetes.
resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/antrea-network-policy/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/network-policy-provider/antrea-network-policy.md
    title: 'Kubernetes website source: tasks/administer-cluster/network-policy-provider/antrea-network-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/antrea-network-policy/
    title: Use Antrea for NetworkPolicy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/network-policy-provider/antrea-network-policy.md
---

# Use Antrea for NetworkPolicy

This page shows how to install and use Antrea CNI plugin on Kubernetes.
For background on Project Antrea, read the [Introduction to Antrea](https://antrea.io/docs/).

## Before you begin

You need to have a Kubernetes cluster. Follow the
[kubeadm getting started guide](/reference/setup-tools/kubeadm.md) to bootstrap one.

## Deploying Antrea with kubeadm

Follow [Getting Started](https://github.com/vmware-tanzu/antrea/blob/main/docs/getting-started.md) guide to deploy Antrea for kubeadm.

## What's next

Once your cluster is running, you can follow the [Declare Network Policy](/tasks/administer-cluster/declare-network-policy.md) to try out Kubernetes NetworkPolicy.
