---
type: glossary-term
title: Subject Alternative Name
description: An X.509 certificate extension to identify what hostname or IP address the certificate applies to.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-san
tags: [kubernetes, reference, glossary, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/glossary/san.md
    title: 'Kubernetes website source: reference/glossary/san.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-san
    title: Subject Alternative Name | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/glossary/san.md
---

# Subject Alternative Name

Subject Alternative Name is an [X.509 certificate](/reference/glossary/certificate.md)
extension that allows identities to be bound to the subject of the certificate.

The [standard](https://datatracker.ietf.org/doc/html/rfc4985) defines identities represented
as an email address, a DNS name, an IP address or a Uniform Resource Identifier (URI).
