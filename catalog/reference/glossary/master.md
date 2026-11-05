---
type: glossary-term
title: Master
description: Legacy term, used as synonym for nodes running the control plane.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-master
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/master.md
    title: 'Kubernetes website source: reference/glossary/master.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-master
    title: Master | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/master.md
---

# Master

Legacy term, used as synonym for [nodes](/reference/glossary/node.md) hosting the [control plane](/reference/glossary/control-plane.md).

The term is still being used by some provisioning tools, such as [kubeadm](/reference/glossary/kubeadm.md), and managed services, to [label](/reference/glossary/label.md) [nodes](/reference/glossary/node.md) with `kubernetes.io/role` and control placement of [control plane](/reference/glossary/control-plane.md) [pods](/reference/glossary/pod.md).
