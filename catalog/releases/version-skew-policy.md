---
type: releases
title: Version Skew Policy
description: The maximum version skew supported between various Kubernetes components.
resource: https://kubernetes.io/releases/version-skew-policy/
tags: [kubernetes, releases]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/releases/version-skew-policy.md
    title: 'Kubernetes website source: releases/version-skew-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/releases/version-skew-policy/
    title: Version Skew Policy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/releases/version-skew-policy.md
---

# Version Skew Policy

This document describes the maximum version skew supported between various Kubernetes components.
Specific cluster deployment tools may place additional restrictions on version skew.

## Supported versions

Kubernetes versions are expressed as **x.y.z**, where **x** is the major version,
**y** is the minor version, and **z** is the patch version, following
[Semantic Versioning](https://semver.org/) terminology. For more information, see
[Kubernetes Release Versioning](https://git.k8s.io/sig-release/release-engineering/versioning.md#kubernetes-release-versioning).

The Kubernetes project maintains release branches for the most recent three minor releases
(1.37, 1.36, 1.35).
Kubernetes 1.19 and newer receive [approximately 1 year of patch support](/releases/patch-releases.md#support-period).
Kubernetes 1.18 and older received approximately 9 months of patch support.

Applicable fixes, including security fixes, may be backported to those three release branches,
depending on severity and feasibility. Patch releases are cut from those branches at a
[regular cadence](/releases/patch-releases.md#cadence), plus additional urgent releases, when required.

The [Release Managers](https://kubernetes.io/releases/release-managers/) group owns this decision.

For more information, see the Kubernetes [patch releases](/releases/patch-releases.md) page.

## Supported version skew

### kube-apiserver

In [highly-available (HA) clusters](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/high-availability/),
the newest and oldest `kube-apiserver` instances must be within one minor version.

Example:

* newest `kube-apiserver` is at **1.37**
* other `kube-apiserver` instances are supported at **1.37** and **1.36**

### kubelet

* `kubelet` must not be newer than `kube-apiserver`.
* `kubelet` may be up to three minor versions older than `kube-apiserver` (`kubelet` < 1.25 may only be up to two minor versions older than `kube-apiserver`).

Example:

* `kube-apiserver` is at **1.37**
* `kubelet` is supported at **1.37**, **1.36**,
  **1.35**, and **1.34**

**Note:**
If version skew exists between `kube-apiserver` instances in an HA cluster, this narrows the allowed `kubelet` versions.

Example:

* `kube-apiserver` instances are at **1.37** and **1.36**
* `kubelet` is supported at **1.36**, **1.35**,
  and **1.34** (**1.37** is not supported because that
  would be newer than the `kube-apiserver` instance at version **1.36**)

### kube-proxy

* `kube-proxy` must not be newer than `kube-apiserver`.
* `kube-proxy` may be up to three minor versions older than `kube-apiserver`
  (`kube-proxy` < 1.25 may only be up to two minor versions older than `kube-apiserver`).
* `kube-proxy` may be up to three minor versions older or newer than the `kubelet` instance
  it runs alongside (`kube-proxy` < 1.25 may only be up to two minor versions older or newer
  than the `kubelet` instance it runs alongside).

Example:

* `kube-apiserver` is at **1.37**
* `kube-proxy` is supported at **1.37**, **1.36**,
  **1.35**, and **1.34**

**Note:**
If version skew exists between `kube-apiserver` instances in an HA cluster, this narrows the allowed `kube-proxy` versions.

Example:

* `kube-apiserver` instances are at **1.37** and **1.36**
* `kube-proxy` is supported at **1.36**, **1.35**,
  and **1.34** (**1.37** is not supported because that would
  be newer than the `kube-apiserver` instance at version **1.36**)

### kube-controller-manager, kube-scheduler, and cloud-controller-manager

`kube-controller-manager`, `kube-scheduler`, and `cloud-controller-manager` must not be newer than the
`kube-apiserver` instances they communicate with. They are expected to match the `kube-apiserver` minor version,
but may be up to one minor version older (to allow live upgrades).

Example:

* `kube-apiserver` is at **1.37**
* `kube-controller-manager`, `kube-scheduler`, and `cloud-controller-manager` are supported
  at **1.37** and **1.36**

**Note:**
If version skew exists between `kube-apiserver` instances in an HA cluster, and these components
can communicate with any `kube-apiserver` instance in the cluster (for example, via a load balancer),
this narrows the allowed versions of these components.

Example:

* `kube-apiserver` instances are at **1.37** and **1.36**
* `kube-controller-manager`, `kube-scheduler`, and `cloud-controller-manager` communicate with a load balancer
  that can route to any `kube-apiserver` instance
* `kube-controller-manager`, `kube-scheduler`, and `cloud-controller-manager` are supported at
  **1.36** (**1.37** is not supported
  because that would be newer than the `kube-apiserver` instance at version **1.36**)

### kubectl

`kubectl` is supported within one minor version (older or newer) of `kube-apiserver`.

Example:

* `kube-apiserver` is at **1.37**
* `kubectl` is supported at **1.38**, **1.37**,
  and **1.36**

**Note:**
If version skew exists between `kube-apiserver` instances in an HA cluster, this narrows the supported `kubectl` versions.

Example:

* `kube-apiserver` instances are at **1.37** and **1.36**
* `kubectl` is supported at **1.37** and **1.36**
  (other versions would be more than one minor version skewed from one of the `kube-apiserver` components)

## Supported component upgrade order

The supported version skew between components has implications on the order
in which components must be upgraded. This section describes the order in
which components must be upgraded to transition an existing cluster from version
**1.36** to version **1.37**.

Optionally, when preparing to upgrade, the Kubernetes project recommends that
you do the following to benefit from as many regression and bug fixes as
possible during your upgrade:

* Ensure that components are on the most recent patch version of your current
  minor version.
* Upgrade components to the most recent patch version of the target minor
  version.

For example, if you're running version 1.36,
ensure that you're on the most recent patch version. Then, upgrade to the most
recent patch version of 1.37.

### kube-apiserver

Pre-requisites:

* In a single-instance cluster, the existing `kube-apiserver` instance is **1.36**
* In an HA cluster, all `kube-apiserver` instances are at **1.36** or
  **1.37** (this ensures maximum skew of 1 minor version between the oldest and newest `kube-apiserver` instance)
* The `kube-controller-manager`, `kube-scheduler`, and `cloud-controller-manager` instances that
  communicate with this server are at version **1.36**
  (this ensures they are not newer than the existing API server version, and are within 1 minor version of the new API server version)
* `kubelet` instances on all nodes are at version **1.36** or **1.35**
  (this ensures they are not newer than the existing API server version, and are within 2 minor versions of the new API server version)
* Registered admission webhooks are able to handle the data the new `kube-apiserver` instance will send them:
  * `ValidatingWebhookConfiguration` and `MutatingWebhookConfiguration` objects are updated to include
    any new versions of REST resources added in **1.37**
    (or use the [`matchPolicy: Equivalent` option](/reference/access-authn-authz/extensible-admission-controllers.md#matching-requests-matchpolicy) available in v1.15+)
  * The webhooks are able to handle any new versions of REST resources that will be sent to them,
    and any new fields added to existing versions in **1.37**

Upgrade `kube-apiserver` to **1.37**

**Note:**
Project policies for [API deprecation](/reference/deprecation-policy.md) and
[API change guidelines](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-architecture/api_changes.md)
require `kube-apiserver` to not skip minor versions when upgrading, even in single-instance clusters.

### kube-controller-manager, kube-scheduler, and cloud-controller-manager

Pre-requisites:

* The `kube-apiserver` instances these components communicate with are at **1.37**
  (in HA clusters in which these control plane components can communicate with any `kube-apiserver`
  instance in the cluster, all `kube-apiserver` instances must be upgraded before upgrading these components)

Upgrade `kube-controller-manager`, `kube-scheduler`, and
`cloud-controller-manager` to **1.37**. There is no
required upgrade order between `kube-controller-manager`, `kube-scheduler`, and
`cloud-controller-manager`. You can upgrade these components in any order, or
even simultaneously.

### kubelet

Pre-requisites:

* The `kube-apiserver` instances the `kubelet` communicates with are at **1.37**

Optionally upgrade `kubelet` instances to **1.37** (or they can be left at
**1.36**, **1.35**, or **1.34**)

**Note:**
Before performing a minor version `kubelet` upgrade, [drain](/tasks/administer-cluster/safely-drain-node.md) pods from that node.
In-place minor version `kubelet` upgrades are not supported.

**Warning:**
Running a cluster with `kubelet` instances that are persistently three minor versions behind
`kube-apiserver` means they must be upgraded before the control plane can be upgraded.

### kube-proxy

Pre-requisites:

* The `kube-apiserver` instances `kube-proxy` communicates with are at **1.37**

Optionally upgrade `kube-proxy` instances to **1.37**
(or they can be left at **1.36**, **1.35**,
or **1.34**)

**Warning:**
Running a cluster with `kube-proxy` instances that are persistently three minor versions behind
`kube-apiserver` means they must be upgraded before the control plane can be upgraded.
