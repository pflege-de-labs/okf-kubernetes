---
type: glossary-term
title: Mirror Pod
description: An object in the API server that tracks a static pod on a kubelet.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-mirror-pod
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/mirror-pod.md
    title: 'Kubernetes website source: reference/glossary/mirror-pod.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-mirror-pod
    title: Mirror Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/mirror-pod.md
---

# Mirror Pod

A [pod](/reference/glossary/pod.md) object that a [kubelet](/reference/glossary/kubelet.md) uses
 to represent a [static pod](/reference/glossary/static-pod.md)

 

When the kubelet finds a static pod in its configuration, it automatically tries to
create a Pod object on the Kubernetes API server for it. This means that the pod
will be visible on the API server, but cannot be controlled from there.

(For example, removing a mirror pod will not stop the kubelet daemon from running it).
