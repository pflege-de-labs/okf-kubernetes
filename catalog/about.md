---
type: reference
title: 'About this bundle: Kubernetes v1.37'
description: This bundle documents Kubernetes v1.37 (kubernetes/website@a81f5383d579); how to apply it to clusters running other versions.
tags: [kubernetes, about]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/hugo.toml
    title: 'Kubernetes website source: hugo.toml'
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: hugo.toml
---

# About this bundle: Kubernetes v1.37

This bundle documents **Kubernetes v1.37** (latest patch release 1.37.1 per the upstream schedule data).
It was generated from [kubernetes/website@a81f5383d579](https://github.com/kubernetes/website/tree/a81f5383d57959d479d946ae3ad2c620ceb11e6a), as of 2026-11-23.

Every page describes the behavior of v1.37. Before applying it to a cluster, compare v1.37 with
the cluster's version (`kubectl version`):

* Cluster older than v1.37: features may be missing, still alpha or beta, or disabled by default.
  Check the stage history in [Feature Gates](/reference/command-line-tools-reference/feature-gates.md)
  and suggest the upgrade that makes a needed feature available.
* Cluster newer than v1.37: features may have graduated, been removed, or changed. Check
  [removed feature gates](/reference/command-line-tools-reference/feature-gates-removed.md) and the
  [Deprecated API Migration Guide](/reference/using-api/deprecation-guide.md), and say that the bundle is older.
* Support status (supported, maintenance mode, end of life) and dates of every release:
  [Release Lifecycle](/releases/lifecycle.md).
* Supported version differences between components, and upgrade order:
  [Version Skew Policy](/releases/version-skew-policy.md).
* Curated naming rules and limits: [Requirements and limits](/requirements/index.md).
