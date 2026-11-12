---
type: glossary-term
title: CronJob
description: A repeating task (a Job) that runs on a regular schedule.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cronjob
tags: [kubernetes, reference, glossary, core-object, workload]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/cronjob.md
    title: 'Kubernetes website source: reference/glossary/cronjob.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cronjob
    title: CronJob | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/cronjob.md
---

# CronJob

Manages a [Job](/concepts/workloads/controllers/job.md) that runs on a periodic schedule.

Similar to a line in a *crontab* file, a CronJob object specifies a schedule using the [cron](https://en.wikipedia.org/wiki/Cron) format.
