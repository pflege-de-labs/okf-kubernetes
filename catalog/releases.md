---
type: section
title: Releases
description: The Kubernetes project maintains release branches for the most recent three minor releases (1.37, 1.36, 1.35).
resource: https://kubernetes.io/releases/
tags: [kubernetes, releases]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/releases/_index.md
    title: 'Kubernetes website source: releases/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/releases/
    title: Releases | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/releases/_index.md
---

# Releases

The Kubernetes project maintains release branches for the most recent three minor releases
(1.37, 1.36, 1.35).
Kubernetes 1.19 and newer receive
[approximately 1 year of patch support](/releases/patch-releases.md#support-period).
Kubernetes 1.18 and older received approximately 9 months of patch support.

Kubernetes versions are expressed as **x.y.z**,
where **x** is the major version, **y** is the minor version, and **z** is the patch version,
following [Semantic Versioning](https://semver.org/) terminology.

More information in the [version skew policy](/releases/version-skew-policy.md) document.

## Release History

See the [Release Lifecycle](/releases/lifecycle.md) for release, maintenance-mode and end-of-life dates.

## End-of-Life Releases

Older Kubernetes releases that are no longer maintained are listed below.

<details>
  <summary>End-of-life releases</summary>
  **Note:**
  These releases are no longer supported and do not receive security updates or bug fixes.
  If you are running one of these releases, the Kubernetes project strongly recommends upgrading to a [supported version](#release-history).
  
  
  See the [Release Lifecycle](/releases/lifecycle.md) for release, maintenance-mode and end-of-life dates.
</details>

## Upcoming Release

Check out the [schedule](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.38)
for the upcoming **1.38** Kubernetes release!

**Note:**
This schedule link may be temporarily unavailable during early release planning phases.  
Check the [SIG Release repository](https://github.com/kubernetes/sig-release/tree/master/releases) for the latest updates.

## Helpful Resources

Refer to the [Kubernetes Release Team](https://github.com/kubernetes/sig-release/tree/master/release-team) resources 
for key information on roles and the release process.
