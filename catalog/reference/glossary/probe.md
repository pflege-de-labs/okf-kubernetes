---
type: glossary-term
title: Probe
description: A check performed periodically by the kubelet on a container in a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-probe
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/probe.md
    title: 'Kubernetes website source: reference/glossary/probe.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-probe
    title: Probe | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/probe.md
---

# Probe

A check that the [kubelet](/reference/glossary/kubelet.md) periodically performs against a container that is 
running in a pod, that will define container's state and health and informing container's lifecycle.

 
To learn more, read [container probes](/concepts/workloads/pods/pod-lifecycle.md#container-probes).
