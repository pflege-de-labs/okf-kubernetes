---
type: glossary-term
title: CronJob
description: A repeating task (a Job) that runs on a regular schedule.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cronjob
tags: [kubernetes, reference, glossary, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/cronjob.md
    title: 'Kubernetes website source: reference/glossary/cronjob.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cronjob
    title: CronJob | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/cronjob.md
---

# CronJob

Manages a [Job](/concepts/workloads/controllers/job.md) that runs on a periodic schedule.

Similar to a line in a *crontab* file, a CronJob object specifies a schedule using the [cron](https://en.wikipedia.org/wiki/Cron) format.
