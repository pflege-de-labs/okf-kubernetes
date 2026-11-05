---
type: glossary-term
title: Container Device Interface (CDI)
description: A CNCF specification for describing device configuration that container runtimes apply when creating containers.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cdi
tags: [kubernetes, reference, glossary, extension]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cdi.md
    title: 'Kubernetes website source: reference/glossary/cdi.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cdi
    title: Container Device Interface (CDI) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/cdi.md
---

# Container Device Interface (CDI)

The Container Device Interface (CDI) is a specification for how to configure
devices inside containers. Kubernetes uses CDI together with device plugins and
with Dynamic Resource Allocation so that workloads receive device setup such as
bind mounts or environment variables from the runtime.

* [Device Plugins](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md)
* [Container Device Interface](https://github.com/cncf-tags/container-device-interface)
  specification repository
