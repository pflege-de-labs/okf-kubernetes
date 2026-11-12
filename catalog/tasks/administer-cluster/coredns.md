---
type: task
title: Using CoreDNS for Service Discovery
description: This page describes the CoreDNS upgrade process and how to install CoreDNS.
resource: https://kubernetes.io/docs/tasks/administer-cluster/coredns/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/coredns.md
    title: 'Kubernetes website source: tasks/administer-cluster/coredns.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/coredns/
    title: Using CoreDNS for Service Discovery | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/coredns.md
min_kubernetes_server_version: v1.9
---

# Using CoreDNS for Service Discovery

This page describes the CoreDNS upgrade process and how to install CoreDNS.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) Your Kubernetes server must be at or later than version v1.9. To check the version, enter `kubectl version`.

## About CoreDNS

[CoreDNS](https://coredns.io) is a flexible, extensible DNS server
that is the default implementation of Kubernetes cluster DNS.
Like Kubernetes, the CoreDNS project is hosted by the
[CNCF](/reference/glossary/cncf.md).

## Installing CoreDNS

For manual deployment, see the documentation at the
[CoreDNS website](https://coredns.io/manual/installation/).

## Upgrading CoreDNS

You can check the version of CoreDNS that kubeadm installs for each version of
Kubernetes in the page
[CoreDNS version in Kubernetes](https://github.com/coredns/deployment/blob/master/kubernetes/CoreDNS-k8s_version.md).

CoreDNS can be upgraded manually in case you want to only upgrade CoreDNS
or use your own custom image.
There is a helpful [guideline and walkthrough](https://github.com/coredns/deployment/blob/master/kubernetes/Upgrading_CoreDNS.md)
available to ensure a smooth upgrade.
Make sure the existing CoreDNS configuration ("Corefile") is retained when
upgrading your cluster.

If you are upgrading your cluster using the `kubeadm` tool, `kubeadm`
can take care of retaining the existing CoreDNS configuration automatically.

## Tuning CoreDNS

When resource utilisation is a concern, it may be useful to tune the
configuration of CoreDNS. For more details, check out the
[documentation on scaling CoreDNS](https://github.com/coredns/deployment/blob/master/kubernetes/Scaling_CoreDNS.md).

## What's next

You can configure [CoreDNS](https://coredns.io) to support many use cases beyond
basic service name resolution by modifying the CoreDNS configuration ("Corefile").
For more information, see the [documentation](https://coredns.io/plugins/kubernetes/)
for the `kubernetes` CoreDNS plugin, or read the 
[Custom DNS Entries for Kubernetes](https://coredns.io/2017/05/08/custom-dns-entries-for-kubernetes/).
in the CoreDNS blog.
