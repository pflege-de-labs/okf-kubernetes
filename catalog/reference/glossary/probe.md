---
type: glossary-term
title: Probe
description: A check performed periodically by the kubelet on a container in a Pod.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-probe
tags: [kubernetes, reference, glossary, tool]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/probe.md
    title: 'Kubernetes website source: reference/glossary/probe.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-probe
    title: Probe | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/probe.md
---

# Probe

A check that the [kubelet](/reference/glossary/kubelet.md) periodically performs against a container that is 
running in a pod, that will define container's state and health and informing container's lifecycle.

 
To learn more, read [container probes](/concepts/workloads/pods/pod-lifecycle.md#container-probes).
