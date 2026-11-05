---
type: glossary-term
title: Cloud Provider
description: An organization that offers a cloud computing platform.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cloud-provider
tags: [kubernetes, reference, glossary, community]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/cloud-provider.md
    title: 'Kubernetes website source: reference/glossary/cloud-provider.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-cloud-provider
    title: Cloud Provider | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/cloud-provider.md
---

# Cloud Provider

A business or other organization that offers a cloud computing platform.

Cloud providers, sometimes called Cloud Service Providers (CSPs), offer
cloud computing platforms or services.

Many cloud providers offer managed infrastructure (also called
Infrastructure as a Service or IaaS).
With managed infrastructure the cloud provider is responsible for
servers, storage, and networking while you manage layers on top of that
such as running a Kubernetes cluster.

You can also find Kubernetes as a managed service; sometimes called
Platform as a Service, or PaaS. With managed Kubernetes, your
cloud provider is responsible for the Kubernetes control plane as well
as the [nodes](/reference/glossary/node.md) and the
infrastructure they rely on: networking, storage, and possibly other
elements such as load balancers.
