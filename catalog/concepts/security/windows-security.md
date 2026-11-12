---
type: concept
title: Security For Windows Nodes
description: This page describes security considerations and best practices specific to the Windows operating system.
resource: https://kubernetes.io/docs/concepts/security/windows-security/
tags: [kubernetes, concepts, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/security/windows-security.md
    title: 'Kubernetes website source: concepts/security/windows-security.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/security/windows-security/
    title: Security For Windows Nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/security/windows-security.md
---

# Security For Windows Nodes

This page describes security considerations and best practices specific to the Windows operating system.

## Protection for Secret data on nodes

On Windows, data from Secrets are written out in clear text onto the node's local
storage (as compared to using tmpfs / in-memory filesystems on Linux). As a cluster
operator, you should take both of the following additional measures:

1. Use file ACLs to secure the Secrets' file location.
1. Apply volume-level encryption using
   [BitLocker](https://docs.microsoft.com/windows/security/information-protection/bitlocker/bitlocker-how-to-deploy-on-windows-server).

## Container users

[RunAsUsername](/tasks/configure-pod-container/configure-runasusername.md)
can be specified for Windows Pods or containers to execute the container
processes as specific user. This is roughly equivalent to
[RunAsUser](/concepts/security/pod-security-policy.md#users-and-groups).

Windows containers offer two default user accounts, ContainerUser and ContainerAdministrator.
The differences between these two user accounts are covered in
[When to use ContainerAdmin and ContainerUser user accounts](https://docs.microsoft.com/virtualization/windowscontainers/manage-containers/container-security#when-to-use-containeradmin-and-containeruser-user-accounts)
within Microsoft's _Secure Windows containers_ documentation.

Local users can be added to container images during the container build process.

**Note:**

* [Nano Server](https://hub.docker.com/_/microsoft-windows-nanoserver) based images run as
  `ContainerUser` by default
* [Server Core](https://hub.docker.com/_/microsoft-windows-servercore) based images run as
  `ContainerAdministrator` by default

Windows containers can also run as Active Directory identities by utilizing
[Group Managed Service Accounts](/tasks/configure-pod-container/configure-gmsa.md)

## Pod-level security isolation

Linux-specific pod security context mechanisms (such as SELinux, AppArmor, Seccomp, or custom
POSIX capabilities) are not supported on Windows nodes.

Privileged containers are [not supported](/concepts/windows/intro.md#compatibility-v1-pod-spec-containers-securitycontext)
on Windows.
Instead [HostProcess containers](/tasks/configure-pod-container/create-hostprocess-pod.md)
can be used on Windows to perform many of the tasks performed by privileged containers on Linux.
