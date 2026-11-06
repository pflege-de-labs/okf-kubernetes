---
type: task
title: Upgrading Linux nodes
description: This page explains how to upgrade a Linux Worker Nodes created with kubeadm.
resource: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes.md
    title: 'Kubernetes website source: tasks/administer-cluster/kubeadm/upgrading-linux-nodes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes/
    title: Upgrading Linux nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes.md
---

# Upgrading Linux nodes

This page explains how to upgrade a Linux Worker Nodes created with kubeadm.

## Before you begin

You need to have shell access to all the nodes, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial 
on a cluster with at least two nodes that are not acting as control plane hosts. To check the version, enter `kubectl version`.
* Familiarize yourself with [the process for upgrading the rest of your kubeadm
cluster](/tasks/administer-cluster/kubeadm/kubeadm-upgrade.md). You will want to
upgrade the control plane nodes before upgrading your Linux Worker nodes.

## Changing the package repository

If you're using the community-owned package repositories (`pkgs.k8s.io`), you need to 
enable the package repository for the desired Kubernetes minor release. This is explained in
[Changing the Kubernetes package repository](/tasks/administer-cluster/kubeadm/change-package-repository.md)
document.

## Upgrading worker nodes

### Upgrade kubeadm and kubectl

Upgrade kubeadm and kubectl:

**Note:**
kubectl must be within one minor version of kube-apiserver as per the
[version skew policy](https://kubernetes.io/releases/version-skew-policy/#kubectl).
If you are upgrading the control plane across multiple minor versions before
upgrading worker nodes, the node-local kubectl may fall outside the allowed
skew. Upgrading kubectl alongside kubeadm here ensures it stays within
the supported version range throughout the node upgrade process.

**Ubuntu, Debian or HypriotOS:**

```shell
# replace x in 1.37.x-* with the latest patch version
sudo apt-mark unhold kubeadm kubectl && \
sudo apt-get update && sudo apt-get install -y kubeadm='1.37.x-*' kubectl='1.37.x-*' && \
sudo apt-mark hold kubeadm kubectl
```

**CentOS, RHEL or Fedora:**

For systems with DNF4 (Fedora < 41, RHEL/Centos < 10):
```shell
# replace x in 1.37.x-* with the latest patch version
sudo dnf install -y kubeadm-'1.37.x-*' kubectl-'1.37.x-*' --disableexcludes=kubernetes
```
For systems with DNF5:
```shell
# replace x in 1.37.x-* with the latest patch version
sudo dnf install -y kubeadm-'1.37.x-*' kubectl-'1.37.x-*' --setopt=disable_excludes=kubernetes
```

### Call "kubeadm upgrade"

For worker nodes this upgrades the local kubelet configuration:

```shell
sudo kubeadm upgrade node
```

### Drain the node

Prepare the node for maintenance by marking it unschedulable and evicting the workloads:

```shell
# execute this command on a control plane node
# replace <node-to-drain> with the name of your node you are draining
kubectl drain <node-to-drain> --ignore-daemonsets
```

### Upgrade kubelet

1. Upgrade the kubelet:

   
   **Ubuntu, Debian or HypriotOS:**

   ```shell
   # replace x in 1.37.x-* with the latest patch version
   sudo apt-mark unhold kubelet && \
   sudo apt-get update && sudo apt-get install -y kubelet='1.37.x-*' && \
   sudo apt-mark hold kubelet
   ```
   
   **CentOS, RHEL or Fedora:**

   For systems with DNF4 (Fedora < 41, RHEL/CentOS < 10):
   ```shell
   # replace x in 1.37.x-* with the latest patch version
   sudo dnf install -y kubelet-'1.37.x-*' --disableexcludes=kubernetes
   ```
   For systems with DNF5:
   ```shell
   # replace x in 1.37.x-* with the latest patch version
   sudo dnf install -y kubelet-'1.37.x-*' --setopt=disable_excludes=kubernetes
   ```
   
   

1. Restart the kubelet:

   ```shell
   sudo systemctl daemon-reload
   sudo systemctl restart kubelet
   ```

### Uncordon the node

Bring the node back online by marking it schedulable:

```shell
# execute this command on a control plane node
# replace <node-to-uncordon> with the name of your node
kubectl uncordon <node-to-uncordon>
```

## What's next

* See how to [Upgrade Windows nodes](/tasks/administer-cluster/kubeadm/upgrading-windows-nodes.md).
