---
type: concept
title: Security For Linux Nodes
description: This page describes security considerations and best practices specific to the Linux operating system.
resource: https://kubernetes.io/docs/concepts/security/linux-security/
tags: [kubernetes, concepts, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/security/linux-security.md
    title: 'Kubernetes website source: concepts/security/linux-security.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/security/linux-security/
    title: Security For Linux Nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/security/linux-security.md
---

# Security For Linux Nodes

This page describes security considerations and best practices specific to the Linux operating system.

## Protection for Secret data on nodes

On Linux nodes, memory-backed volumes (such as [`secret`](/concepts/configuration/secret.md)
volume mounts, or [`emptyDir`](/concepts/storage/volumes.md#emptydir) with `medium: Memory`)
are implemented with a `tmpfs` filesystem.

If you have swap configured and use an older Linux kernel (or a current kernel and an unsupported configuration of Kubernetes),
**memory** backed volumes can have data written to persistent storage.

The Linux kernel officially supports the `noswap` option from version 6.3,
therefore it is recommended the used kernel version is 6.3 or later,
or supports the `noswap` option via a backport, if swap is enabled on the node.

Read [swap memory management](/concepts/cluster-administration/swap-memory-management.md#memory-backed-volumes)
for more info.
