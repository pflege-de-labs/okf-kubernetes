---
type: concept
title: Compatibility Version For Kubernetes Control Plane Components
description: Since release v1.32, we introduced configurable version compatibility and emulation options to Kubernetes control plane components to make upgrades safer by providing more control and increasing the granularity of steps available to cluster administrators.
resource: https://kubernetes.io/docs/concepts/cluster-administration/compatibility-version/
tags: [kubernetes, concepts, cluster-administration]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/cluster-administration/compatibility-version.md
    title: 'Kubernetes website source: concepts/cluster-administration/compatibility-version.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/cluster-administration/compatibility-version/
    title: Compatibility Version For Kubernetes Control Plane Components | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/cluster-administration/compatibility-version.md
---

# Compatibility Version For Kubernetes Control Plane Components

Since release v1.32, we introduced configurable version compatibility and emulation options to Kubernetes control plane components to make upgrades safer by providing more control and increasing the granularity of steps available to cluster administrators.

## Emulated Version

The emulation option is set by the `--emulated-version` flag of control plane components. It allows the component to emulate the behavior (APIs, features, ...) of an earlier version of Kubernetes.

When used, the capabilities available will match the emulated version: 
* Any capabilities present in the binary version that were introduced after the emulation version will be unavailable. 
* Any capabilities removed after the emulation version will be available. 

This enables a binary from a particular Kubernetes release to emulate the behavior of a previous version with sufficient fidelity that interoperability with other system components can be defined in terms of the emulated version.

The `--emulated-version` must be <= `binaryVersion`. See the help message of the `--emulated-version` flag for supported range of emulated versions.
