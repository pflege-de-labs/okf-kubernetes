---
type: glossary-term
title: Admission Controller
description: A piece of code that intercepts requests to the Kubernetes API server prior to persistence of the object.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-admission-controller
tags: [kubernetes, reference, glossary, extension, security]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/admission-controller.md
    title: 'Kubernetes website source: reference/glossary/admission-controller.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-admission-controller
    title: Admission Controller | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/admission-controller.md
---

# Admission Controller

A piece of code that intercepts requests to the Kubernetes API server prior to persistence of the object.

Admission controllers are configurable for the Kubernetes API server and may be "validating", "mutating", or
both. Any admission controller may reject the request. Mutating controllers may modify the objects they admit;
validating controllers may not.

* [Admission controllers in the Kubernetes documentation](/reference/access-authn-authz/admission-controllers.md)
