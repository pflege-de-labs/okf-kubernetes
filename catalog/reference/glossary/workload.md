---
type: glossary-term
title: Workload
description: A workload is an application running on Kubernetes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-workload
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/workload.md
    title: 'Kubernetes website source: reference/glossary/workload.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-workload
    title: Workload | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/workload.md
---

# Workload

A workload is an application running on Kubernetes.

 

Various core objects that represent different types or parts of a workload
include the DaemonSet, Deployment, Job, ReplicaSet, and StatefulSet objects.

For example, a workload that has a web server and a database might run the
database in one [StatefulSet](/reference/glossary/statefulset.md) and the web server
in a [Deployment](/reference/glossary/deployment.md).
