---
type: concept
title: Kubeadm
description: kubeadm performs the actions necessary to get a minimum viable cluster up and running.
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/_index.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/
    title: Kubeadm | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/_index.md
---

# Kubeadm

<img src="https://kubernetes.io/images/kubeadm-stacked-color.png" align="right" width="150px">Kubeadm is a tool built to provide `kubeadm init` and `kubeadm join` as best-practice "fast paths" for creating Kubernetes clusters.

kubeadm performs the actions necessary to get a minimum viable cluster up and running. By design, it cares only about bootstrapping, not about provisioning machines. Likewise, installing various nice-to-have addons, like the Kubernetes Dashboard, monitoring solutions, and cloud-specific addons, is not in scope.

Instead, we expect higher-level and more tailored tooling to be built on top of kubeadm, and ideally, using kubeadm as the basis of all deployments will make it easier to create conformant clusters.

## How to install

To install kubeadm, see the [installation guide](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm).

## What's next

* [kubeadm init](/reference/setup-tools/kubeadm/kubeadm-init.md) to bootstrap a Kubernetes control-plane node
* [kubeadm join](/reference/setup-tools/kubeadm/kubeadm-join.md) to bootstrap a Kubernetes worker node and join it to the cluster
* [kubeadm upgrade](/reference/setup-tools/kubeadm/kubeadm-upgrade.md) to upgrade a Kubernetes cluster to a newer version
* [kubeadm config](/reference/setup-tools/kubeadm/kubeadm-config.md) if you initialized your cluster using kubeadm v1.7.x or lower, to configure your cluster for `kubeadm upgrade`
* [kubeadm token](/reference/setup-tools/kubeadm/kubeadm-token.md) to manage tokens for `kubeadm join`
* [kubeadm reset](/reference/setup-tools/kubeadm/kubeadm-reset.md) to revert any changes made to this host by `kubeadm init` or `kubeadm join`
* [kubeadm certs](/reference/setup-tools/kubeadm/kubeadm-certs.md) to manage Kubernetes certificates
* [kubeadm kubeconfig](/reference/setup-tools/kubeadm/kubeadm-kubeconfig.md) to manage kubeconfig files
* [kubeadm version](/reference/setup-tools/kubeadm/kubeadm-version.md) to print the kubeadm version
* [kubeadm alpha](/reference/setup-tools/kubeadm/kubeadm-alpha.md) to preview a set of features made available for gathering feedback from the community
