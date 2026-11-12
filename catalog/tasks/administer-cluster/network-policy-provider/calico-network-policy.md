---
type: task
title: Use Calico for NetworkPolicy
description: This page shows a couple of quick ways to create a Calico cluster on Kubernetes.
resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/calico-network-policy/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/network-policy-provider/calico-network-policy.md
    title: 'Kubernetes website source: tasks/administer-cluster/network-policy-provider/calico-network-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/calico-network-policy/
    title: Use Calico for NetworkPolicy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/network-policy-provider/calico-network-policy.md
---

# Use Calico for NetworkPolicy

This page shows a couple of quick ways to create a Calico cluster on Kubernetes.

## Before you begin

Decide whether you want to deploy a [cloud](#creating-a-calico-cluster-with-google-kubernetes-engine-gke) or [local](#creating-a-local-calico-cluster-with-kubeadm) cluster.

## Creating a Calico cluster with Google Kubernetes Engine (GKE)

**Prerequisite**: [gcloud](https://cloud.google.com/sdk/docs/quickstarts).

1.  To launch a GKE cluster with Calico, include the `--enable-network-policy` flag.

    **Syntax**
    ```shell
    gcloud container clusters create [CLUSTER_NAME] --enable-network-policy
    ```

    **Example**
    ```shell
    gcloud container clusters create my-calico-cluster --enable-network-policy
    ```

1.  To verify the deployment, use the following command.

    ```shell
    kubectl get pods --namespace=kube-system
    ```

    The Calico pods begin with `calico`. Check to make sure each one has a status of `Running`.

## Creating a local Calico cluster with kubeadm

To get a local single-host Calico cluster in fifteen minutes using kubeadm, refer to the 
[Calico Quickstart](https://projectcalico.docs.tigera.io/getting-started/kubernetes/).

## What's next

Once your cluster is running, you can follow the [Declare Network Policy](/tasks/administer-cluster/declare-network-policy.md) to try out Kubernetes NetworkPolicy.
