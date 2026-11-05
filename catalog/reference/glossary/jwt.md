---
type: glossary-term
title: JSON Web Token (JWT)
description: A means of representing claims to be transferred between two parties.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-jwt
tags: [kubernetes, reference, glossary, security, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/jwt.md
    title: 'Kubernetes website source: reference/glossary/jwt.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-jwt
    title: JSON Web Token (JWT) | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/jwt.md
---

# JSON Web Token (JWT)

A means of representing claims to be transferred between two parties.

JWTs can be digitally signed and encrypted. Kubernetes uses JWTs as
authentication tokens to verify the identity of entities that want to perform
actions in a cluster.
