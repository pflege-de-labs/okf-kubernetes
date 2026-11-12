---
type: task
title: Upgrading Windows nodes
description: This page explains how to upgrade a Windows node created with kubeadm.
resource: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-windows-nodes/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/kubeadm/upgrading-windows-nodes.md
    title: 'Kubernetes website source: tasks/administer-cluster/kubeadm/upgrading-windows-nodes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-windows-nodes/
    title: Upgrading Windows nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/kubeadm/upgrading-windows-nodes.md
min_kubernetes_server_version: "1.17"
---

# Upgrading Windows nodes

**Feature state:** Kubernetes v1.18 [beta]

This page explains how to upgrade a Windows node created with kubeadm.

## Before you begin
 
You need to have shell access to all the nodes, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial 
on a cluster with at least two nodes that are not acting as control plane hosts. Your Kubernetes server must be at or later than version 1.17. To check the version, enter `kubectl version`.
* Familiarize yourself with [the process for upgrading the rest of your kubeadm
cluster](/tasks/administer-cluster/kubeadm/kubeadm-upgrade.md). You will want to
upgrade the control plane nodes before upgrading your Windows nodes.

## Upgrading worker nodes

### Upgrade kubeadm

1.  From the Windows node, upgrade kubeadm:

    ```powershell
    # replace 1.37.1 with your desired version
    curl.exe -Lo <path-to-kubeadm.exe>  "https://dl.k8s.io/v1.37.1/bin/windows/amd64/kubeadm.exe"
    ```

### Drain the node

1.  From a machine with access to the Kubernetes API,
    prepare the node for maintenance by marking it unschedulable and evicting the workloads:

    ```shell
    # replace <node-to-drain> with the name of your node you are draining
    kubectl drain <node-to-drain> --ignore-daemonsets
    ```

    You should see output similar to this:

    ```
    node/ip-172-31-85-18 cordoned
    node/ip-172-31-85-18 drained
    ```

### Upgrade the kubelet configuration

1.  From the Windows node, call the following command to sync new kubelet configuration:

    ```powershell
    kubeadm upgrade node
    ```

### Upgrade kubelet and kube-proxy

1.  From the Windows node, upgrade and restart the kubelet:

    ```powershell
    stop-service kubelet
    curl.exe -Lo <path-to-kubelet.exe> "https://dl.k8s.io/v1.37.1/bin/windows/amd64/kubelet.exe"
    restart-service kubelet
    ```

2. From the Windows node, upgrade and restart the kube-proxy.

    ```powershell
    stop-service kube-proxy
    curl.exe -Lo <path-to-kube-proxy.exe> "https://dl.k8s.io/v1.37.1/bin/windows/amd64/kube-proxy.exe"
    restart-service kube-proxy
    ```

**Note:**
If you are running kube-proxy in a HostProcess container within a Pod, and not as a Windows Service,
you can upgrade kube-proxy by applying a newer version of your kube-proxy manifests.

### Uncordon the node

1.  From a machine with access to the Kubernetes API,
bring the node back online by marking it schedulable:

    ```shell
    # replace <node-to-drain> with the name of your node
    kubectl uncordon <node-to-drain>
    ```
 ## What's next

* See how to [Upgrade Linux nodes](/tasks/administer-cluster/kubeadm/upgrading-linux-nodes.md).
