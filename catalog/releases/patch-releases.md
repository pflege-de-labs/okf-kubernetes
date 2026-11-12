---
type: releases
title: Patch Releases
description: Schedule and team contact information for Kubernetes patch releases.
resource: https://kubernetes.io/releases/patch-releases/
tags: [kubernetes, releases]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/releases/patch-releases.md
    title: 'Kubernetes website source: releases/patch-releases.md'
  - id: k8s-docs
    resource: https://kubernetes.io/releases/patch-releases/
    title: Patch Releases | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/releases/patch-releases.md
---

# Patch Releases

Schedule and team contact information for Kubernetes patch releases.

For general information about Kubernetes release cycle, see the
[release process description].

## Cadence

Our typical patch release cadence is monthly. It is
commonly a bit faster (1 to 2 weeks) for the earliest patch releases
after a 1.X minor release. Critical bug fixes may cause a more
immediate release outside of the normal cadence. We also aim to not make
releases during major holiday periods.

## Contact

See the [Release Managers page][release-managers] for full contact details on the Patch Release Team.

Please give us a business day to respond - we may be in a different timezone!

In between releases the team is looking at incoming cherry pick
requests on a weekly basis. The team will get in touch with
submitters via GitHub PR, SIG channels in Slack, and direct messages
in Slack and [email](mailto:release-managers-private@kubernetes.io)
if there are questions on the PR.

## Cherry picks

Please follow the [cherry pick process][cherry-picks].

Cherry picks must be merge-ready in GitHub with proper labels (e.g.,
`approved`, `lgtm`, `release-note`) and passing CI tests ahead of the
cherry pick deadline. This is typically two days before the target
release, but may be more. Earlier PR readiness is better, as we
need time to get CI signal after merging your cherry picks ahead
of the actual release.

Cherry pick PRs which miss merge criteria will be carried over and tracked
for the next patch release.

## Support Period

In accordance with the [yearly support KEP][yearly-support], the Kubernetes
Community will support active patch release series for a period of roughly
fourteen (14) months.

The first twelve months of this timeframe will be considered the standard
period.

Towards the end of the twelve month, the following will happen:

- [Release Managers][release-managers] will cut a release
- The patch release series will enter maintenance mode

During the two-month maintenance mode period, Release Managers may cut
additional maintenance releases to resolve:

- [Vulnerabilities](/reference/issues-security/official-cve-feed.md) that have an assigned
  CVE ID (under the advisement of the Security Response Committee)
- dependency issues (including base image updates)
- critical core component issues

At the end of the two-month maintenance mode period, the patch release series
will be considered EOL (end of life) and cherry picks to the associated branch
are to be closed soon afterwards.

Note that the 28th of the month was chosen for maintenance mode and EOL target
dates for simplicity (every month has it).

## Upcoming Monthly Releases

Timelines may vary with the severity of bug fixes, but for easier planning we
will target the following monthly release points. Unplanned, critical
releases may also occur in between these.

See the [Release Lifecycle](/releases/lifecycle.md) for release, maintenance-mode and end-of-life dates.

## Detailed Release History for Active Branches

See the [Release Lifecycle](/releases/lifecycle.md) for release, maintenance-mode and end-of-life dates.

## Non-Active Branch history

These releases are no longer supported.

See the [Release Lifecycle](/releases/lifecycle.md) for release, maintenance-mode and end-of-life dates.

[cherry-picks]: https://github.com/kubernetes/community/blob/main/contributors/devel/sig-release/cherry-picks.md
[release-managers]: https://kubernetes.io/releases/release-managers
[release process description]: /releases/release.md
[yearly-support]: https://git.k8s.io/enhancements/keps/sig-release/1498-kubernetes-yearly-support-period/README.md
