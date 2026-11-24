---
type: glossary-term
title: etcd
description: Consistent and highly-available key value store used as backing store of Kubernetes for all cluster data.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-etcd
tags: [kubernetes, reference, glossary, architecture, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/etcd.md
    title: 'Kubernetes website source: reference/glossary/etcd.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-etcd
    title: etcd | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/etcd.md
---

# etcd

Consistent and highly-available key value store used as Kubernetes' backing store for all cluster data.

If your Kubernetes cluster uses etcd as its backing store, make sure you have a
[back up](/tasks/administer-cluster/configure-upgrade-etcd.md#backing-up-an-etcd-cluster) plan
for the data.

You can find in-depth information about etcd in the official [documentation](https://etcd.io/docs/).
