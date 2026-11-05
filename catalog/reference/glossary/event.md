---
type: glossary-term
title: Event
description: Kubernetes objects that describe some state change in the cluster.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-event
tags: [kubernetes, reference, glossary, core-object, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/event.md
    title: 'Kubernetes website source: reference/glossary/event.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-event
    title: Event | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/event.md
---

# Event

A Kubernetes [object](/reference/glossary/object.md) that describes state changes
or notable occurrences in the cluster.

Events have a limited retention time and triggers and messages may evolve with time.
Event consumers should not rely on the timing of an event with a given reason reflecting a consistent underlying trigger,
or the continued existence of events with that reason.

Events should be treated as informative, best-effort, supplemental data.

In Kubernetes, [auditing](/tasks/debug/debug-cluster/audit.md) generates a different kind of
Event record (API group `audit.k8s.io`).
