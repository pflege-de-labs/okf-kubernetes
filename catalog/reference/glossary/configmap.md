---
type: glossary-term
title: ConfigMap
description: An API object used to store non-confidential data in key-value pairs. Can be consumed as environment variables, command-line arguments, or configuration files in a volume.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-configmap
tags: [kubernetes, reference, glossary, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/configmap.md
    title: 'Kubernetes website source: reference/glossary/configmap.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-configmap
    title: ConfigMap | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/configmap.md
---

# ConfigMap

An API object used to store non-confidential data in key-value pairs.
[Pods](/reference/glossary/pod.md) can consume ConfigMaps as
environment variables, command-line arguments, or as configuration files in a
[volume](/reference/glossary/volume.md).

 

A ConfigMap allows you to decouple environment-specific configuration from your [container images](/reference/glossary/image.md), so that your applications are easily portable.
