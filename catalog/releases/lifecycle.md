---
type: reference
title: Release Lifecycle
description: Release, maintenance-mode and end-of-life dates of every Kubernetes minor release, and how to derive whether a release is supported, in maintenance mode (security fixes only), or end of life.
resource: https://kubernetes.io/releases/
tags: [kubernetes, releases]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/data/releases/schedule.yaml
    title: 'Kubernetes website source: data/releases/schedule.yaml'
  - id: k8s-docs
    resource: https://kubernetes.io/releases/
    title: Release Lifecycle | Kubernetes documentation
  - id: k8s-eol-data
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/data/releases/eol.yaml
    title: 'Kubernetes website source: data/releases/eol.yaml'
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: data/releases/schedule.yaml
---

# Release Lifecycle

The Kubernetes project maintains release branches for the three most recent minor releases.
A new minor release ships roughly every four months. Each minor release gets about 14 months of
patch support: 12 months of standard support, then 2 months of maintenance mode, then end of life
([support period](/releases/patch-releases.md#support-period)).

The tables are a snapshot of kubernetes/website@2fb3b0d636e8 as of 2026-11-05. They do not state a status.
Derive the status of a release from today's date:

| Status | Condition | Fixes shipped |
|---|---|---|
| supported (standard support) | today < maintenance mode start | regular monthly patch releases |
| maintenance mode (security and critical fixes only) | maintenance mode start <= today < end of life | CVE fixes, dependency and base-image updates, critical core component fixes |
| end of life (unsupported) | today >= end of life | none; upgrade to a supported release |

A release listed under maintained releases whose end-of-life date has passed is end of life;
the upstream data is updated after the fact.

## Maintained releases

| Release | Released | Latest patch | Next patch (target date) | Maintenance mode start | End of life |
|---|---|---|---|---|---|
| 1.37 | 2026-08-26 | 1.37.1 | 1.37.2 (2026-10-13) | 2027-08-28 | 2027-10-28 |
| 1.36 | 2026-04-22 | 1.36.5 | 1.36.6 (2026-10-13) | 2027-04-28 | 2027-06-28 |
| 1.35 | 2025-12-17 | 1.35.9 | 1.35.10 (2026-10-13) | 2026-12-28 | 2027-02-28 |
| 1.34 | 2025-08-27 | 1.34.12 | 1.34.13 (2026-10-13) | 2026-08-27 | 2026-10-27 |

## End-of-life releases

| Release | Final patch | End of life | Note |
|---|---|---|---|
| 1.33 | 1.33.13 | 2026-06-28 |  |
| 1.32 | 1.32.13 | 2026-02-28 |  |
| 1.31 | 1.31.14 | 2025-11-11 |  |
| 1.30 | 1.30.14 | 2025-07-15 |  |
| 1.29 | 1.29.14 | 2025-02-28 |  |
| 1.28 | 1.28.15 | 2024-10-22 |  |
| 1.27 | 1.27.16 | 2024-07-16 |  |
| 1.26 | 1.26.15 | 2024-02-28 | 1.26.15 was released in March 2024 (after the EOL date) to pick up a new version of Go to [address several Go CVEs](https://groups.google.com/g/golang-dev/c/o1I1Vv8Rfgs/m/Wr8tD1RlAgAJ) |
| 1.25 | 1.25.16 | 2023-10-28 | 1.25.16 was released in November 2023 (after the EOL date) to fix [CVE-2023-5528](https://groups.google.com/g/kubernetes-announce/c/c3py6Fw0DTI/m/cScFSdk1BwAJ) |
| 1.24 | 1.24.17 | 2023-07-28 | 1.24.17 was released in August 2023 (after the EOL date) to fix CVE-2023-3676 and CVE-2023-3955 |
| 1.23 | 1.23.17 | 2023-02-28 |  |
| 1.22 | 1.22.17 | 2022-12-08 | 1.22.17 was released in December 2022 (after the EOL date) to backport registry changes and fix two critical issues. |
| 1.21 | 1.21.14 | 2022-06-28 |  |
| 1.20 | 1.20.15 | 2022-02-28 |  |
| 1.19 | 1.19.16 | 2021-10-28 |  |
| 1.18 | 1.18.20 | 2021-06-18 | Created to solve regression introduced in 1.18.19 |
| 1.17 | 1.17.17 | 2021-01-13 |  |
| 1.16 | 1.16.15 | 2020-09-02 |  |
| 1.15 | 1.15.12 | 2020-05-06 |  |
| 1.14 | 1.14.10 | 2019-12-11 |  |
| 1.13 | 1.13.12 | 2019-10-15 |  |
| 1.12 | 1.12.10 | 2019-07-08 |  |
| 1.11 | 1.11.10 | 2019-05-01 |  |
| 1.10 | 1.10.13 | 2019-02-13 |  |
| 1.9 | 1.9.11 | 2018-09-29 |  |
| 1.8 | 1.8.15 | 2018-07-12 |  |
| 1.7 | 1.7.16 | 2018-04-04 |  |
| 1.6 | 1.6.13 | 2017-11-23 |  |
| 1.5 | 1.5.8 | 2017-10-01 |  |
| 1.4 | 1.4.12 | 2017-04-21 |  |
| 1.3 | 1.3.10 | 2016-11-01 |  |
| 1.2 | 1.2.7 | 2016-10-23 |  |

## Upcoming patch release dates

| Target date | Cherry-pick deadline |
|---|---|
| 2026-10-13 | 2026-10-09 |
| 2026-11-10 | 2026-11-06 |
| 2026-12-15 | 2026-12-11 |
| 2027-01-12 | 2027-01-08 |

## Next minor release

At the snapshot date, v1.37 was the newest release and v1.38 was in development; neither v1.38
nor later releases are part of this bundle. Minor releases ship about every four months, so if
today is more than four months after the newest release date above, assume a newer minor exists.
The v1.38 schedule is in
[kubernetes/sig-release](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.38).
Features planned for an unreleased version can still slip.

## Feature availability by release

The pages of this bundle describe Kubernetes v1.37. To check whether a feature is available
in another release, use the per-stage history (stage, default, since, until) in
[Feature Gates](/reference/command-line-tools-reference/feature-gates.md) and
[removed feature gates](/reference/command-line-tools-reference/feature-gates-removed.md),
API removals in the [Deprecated API Migration Guide](/reference/using-api/deprecation-guide.md),
and supported component version combinations in the [Version Skew Policy](/releases/version-skew-policy.md).
