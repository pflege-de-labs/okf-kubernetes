---
type: concept
title: Configure DNS for a Cluster
description: Kubernetes offers a DNS cluster addon, which most of the supported environments enable by default.
resource: https://kubernetes.io/docs/tasks/access-application-cluster/configure-dns-cluster/
tags: [kubernetes, tasks, access-application-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/tasks/access-application-cluster/configure-dns-cluster.md
    title: 'Kubernetes website source: tasks/access-application-cluster/configure-dns-cluster.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/access-application-cluster/configure-dns-cluster/
    title: Configure DNS for a Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/tasks/access-application-cluster/configure-dns-cluster.md
---

# Configure DNS for a Cluster

Kubernetes offers a DNS cluster addon, which most of the supported environments enable by default. In Kubernetes version 1.11 and later, CoreDNS is recommended and is installed by default with kubeadm.

For more information on how to configure CoreDNS for a Kubernetes cluster, see the [Customizing DNS Service](/tasks/administer-cluster/dns-custom-nameservers.md).
