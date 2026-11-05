---
type: glossary-term
title: Job
description: A finite or batch task that runs to completion.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-job
tags: [kubernetes, reference, glossary, fundamental, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/job.md
    title: 'Kubernetes website source: reference/glossary/job.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-job
    title: Job | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/job.md
---

# Job

A finite or batch task that runs to completion.

 

Creates one or more [Pod](/reference/glossary/pod.md) objects and ensures that a specified number of them successfully terminate. As Pods successfully complete, the Job tracks the successful completions.
