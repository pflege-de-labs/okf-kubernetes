---
type: glossary-term
title: Immutable Infrastructure
description: Immutable Infrastructure refers to computer infrastructure (virtual machines, containers, network appliances) that cannot be changed once deployed
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-immutable-infrastructure
tags: [kubernetes, reference, glossary, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/immutable-infrastructure.md
    title: 'Kubernetes website source: reference/glossary/immutable-infrastructure.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-immutable-infrastructure
    title: Immutable Infrastructure | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/immutable-infrastructure.md
---

# Immutable Infrastructure

Immutable Infrastructure refers to computer infrastructure (virtual machines, containers, network appliances) that cannot be changed once deployed.

Immutability can be enforced by an automated process that overwrites unauthorized changes or through a system that won’t allow changes in the first place.
[Containers](/reference/glossary/container.md) are a good example of immutable infrastructure because persistent changes to containers
can only be made by creating a new version of the container or recreating the existing container from its image.

By preventing or identifying unauthorized changes, immutable infrastructures make it easier to identify and mitigate security risks. 
Operating such a system becomes a lot more straightforward because administrators can make assumptions about it.
After all, they know no one made mistakes or changes they forgot to communicate.
Immutable infrastructure goes hand-in-hand with infrastructure as code where all automation needed
to create infrastructure is stored in version control (such as Git).
This combination of immutability and version control means that there is a durable audit log of every authorized change to a system.
