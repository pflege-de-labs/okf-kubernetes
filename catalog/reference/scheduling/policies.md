---
type: concept
title: Scheduling Policies
description: In Kubernetes versions before v1.23, a scheduling policy can be used to specify the *predicates* and *priorities* process.
resource: https://kubernetes.io/docs/reference/scheduling/policies/
tags: [kubernetes, reference, scheduling]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/scheduling/policies.md
    title: 'Kubernetes website source: reference/scheduling/policies.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/scheduling/policies/
    title: Scheduling Policies | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/scheduling/policies.md
---

# Scheduling Policies

In Kubernetes versions before v1.23, a scheduling policy can be used to specify the *predicates* and *priorities* process. For example, you can set a scheduling policy by
running `kube-scheduler --policy-config-file <filename>` or `kube-scheduler --policy-configmap <ConfigMap>`.

This scheduling policy is not supported since Kubernetes v1.23. Associated flags `policy-config-file`, `policy-configmap`, `policy-configmap-namespace` and `use-legacy-policy-config` are also not supported. Instead, use the [Scheduler Configuration](/reference/scheduling/config.md) to achieve similar behavior.

## What's next

* Learn about [scheduling](/concepts/scheduling-eviction/kube-scheduler.md)
* Learn about [kube-scheduler Configuration](/reference/scheduling/config.md)
* Read the [kube-scheduler configuration reference (v1)](/reference/config-api/kube-scheduler-config.v1.md)
