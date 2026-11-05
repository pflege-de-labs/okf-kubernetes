---
type: section
title: Windows in Kubernetes
description: Kubernetes supports nodes that run Microsoft Windows.
resource: https://kubernetes.io/docs/concepts/windows/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/windows/_index.md
    title: 'Kubernetes website source: concepts/windows/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/windows/
    title: Windows in Kubernetes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/windows/_index.md
---

# Windows in Kubernetes

Kubernetes supports worker [nodes](/reference/glossary/node.md)
running either Linux or Microsoft Windows.

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

The CNCF and its parent the Linux Foundation take a vendor-neutral approach
towards compatibility. It is possible to join your [Windows server](https://www.microsoft.com/en-us/windows-server)
as a worker node to a Kubernetes cluster.

You can [install and set up kubectl on Windows](/tasks/tools/install-kubectl-windows.md)
no matter what operating system you use within your cluster.

If you are using Windows nodes, you can read:

* [Networking On Windows](/concepts/services-networking/windows-networking.md)
* [Windows Storage In Kubernetes](/concepts/storage/windows-storage.md)
* [Resource Management for Windows Nodes](/concepts/configuration/windows-resource-management.md)
* [Configure RunAsUserName for Windows Pods and Containers](/tasks/configure-pod-container/configure-runasusername.md)
* [Create A Windows HostProcess Pod](/tasks/configure-pod-container/create-hostprocess-pod.md)
* [Configure Group Managed Service Accounts for Windows Pods and Containers](/tasks/configure-pod-container/configure-gmsa.md)
* [Security For Windows Nodes](/concepts/security/windows-security.md)
* [Windows Debugging Tips](/tasks/debug/debug-cluster/windows.md)
* [Guide for Scheduling Windows Containers in Kubernetes](/concepts/windows/user-guide.md)

or, for an overview, read:
