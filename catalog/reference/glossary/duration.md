---
type: glossary-term
title: Duration
description: A string value representing an amount of time.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-duration
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/duration.md
    title: 'Kubernetes website source: reference/glossary/duration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-duration
    title: Duration | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/duration.md
---

# Duration

A string value representing an amount of time.

The format of a (Kubernetes) duration is based on the
[`time.Duration`](https://pkg.go.dev/time#Duration) type from the Go programming language.

In Kubernetes APIs that use durations, the value is expressed as series of a non-negative
integers combined with a time unit suffix. You can have more than one time quantity and
the duration is the sum of those time quantities.
The valid time units are "ns", "µs" (or "us"), "ms", "s", "m", and "h".

For example: `5s` represents a duration of five seconds, and `1m30s` represents a duration
of one minute and thirty seconds.
