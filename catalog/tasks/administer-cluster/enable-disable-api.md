---
type: task
title: Enable Or Disable A Kubernetes API
description: This page shows how to enable or disable an API version from your cluster's control plane.
resource: https://kubernetes.io/docs/tasks/administer-cluster/enable-disable-api/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/tasks/administer-cluster/enable-disable-api.md
    title: 'Kubernetes website source: tasks/administer-cluster/enable-disable-api.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/enable-disable-api/
    title: Enable Or Disable A Kubernetes API | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/tasks/administer-cluster/enable-disable-api.md
---

# Enable Or Disable A Kubernetes API

This page shows how to enable or disable an API version from your cluster's
[control plane](/reference/glossary/control-plane.md).

Specific API versions can be turned on or off by passing `--runtime-config=api/<version>` as a
command line argument to the API server. The values for this argument are a comma-separated
list of API versions. Later values override earlier values.

The `runtime-config` command line argument also supports 2 special keys:

- `api/all`, representing all known APIs
- `api/legacy`, representing only legacy APIs. Legacy APIs are any APIs that have been
   explicitly [deprecated](/reference/deprecation-policy.md).

For example, to turn off all API versions except v1, pass `--runtime-config=api/all=false,api/v1=true`
to the `kube-apiserver`.

## What's next

Read the [full documentation](/reference/command-line-tools-reference/kube-apiserver.md)
for the `kube-apiserver` component.
