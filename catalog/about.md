---
type: reference
title: 'About this bundle: Kubernetes v1.37'
description: This bundle documents Kubernetes v1.37 (kubernetes/website@2674d5095a50); how to apply it to clusters running other versions.
tags: [kubernetes, about]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/hugo.toml
    title: 'Kubernetes website source: hugo.toml'
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: hugo.toml
---

# About this bundle: Kubernetes v1.37

This bundle documents **Kubernetes v1.37** (latest patch release 1.37.1 per the upstream schedule data).
It was generated from [kubernetes/website@2674d5095a50](https://github.com/kubernetes/website/tree/2674d5095a50d781d2f884c83973921f48c5cf28), as of 2026-11-12.

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
