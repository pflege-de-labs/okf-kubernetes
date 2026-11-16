---
type: glossary-term
title: ConfigMap
description: An API object used to store non-confidential data in key-value pairs. Can be consumed as environment variables, command-line arguments, or configuration files in a volume.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-configmap
tags: [kubernetes, reference, glossary, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/configmap.md
    title: 'Kubernetes website source: reference/glossary/configmap.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-configmap
    title: ConfigMap | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/configmap.md
---

# ConfigMap

An API object used to store non-confidential data in key-value pairs.
[Pods](/reference/glossary/pod.md) can consume ConfigMaps as
environment variables, command-line arguments, or as configuration files in a
[volume](/reference/glossary/volume.md).

 

A ConfigMap allows you to decouple environment-specific configuration from your [container images](/reference/glossary/image.md), so that your applications are easily portable.
