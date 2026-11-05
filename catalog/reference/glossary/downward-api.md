---
type: glossary-term
title: Downward API
description: A mechanism to expose Pod and container field values to code running in a container.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-downward-api
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/downward-api.md
    title: 'Kubernetes website source: reference/glossary/downward-api.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-downward-api
    title: Downward API | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/downward-api.md
---

# Downward API

Kubernetes' mechanism to expose Pod and container field values to code running in a container.

It is sometimes useful for a container to have information about itself, without
needing to make changes to the container code that directly couple it to Kubernetes.

The Kubernetes downward API allows containers to consume information about themselves
or their context in a Kubernetes cluster. Applications in containers can have
access to that information, without the application needing to act as a client of
the Kubernetes API.

There are two ways to expose Pod and container fields to a running container:

- using [environment variables](/tasks/inject-data-application/environment-variable-expose-pod-information.md)
- using [a `downwardAPI` volume](/tasks/inject-data-application/downward-api-volume-expose-pod-information.md)

Together, these two ways of exposing Pod and container fields are called the _downward API_.
