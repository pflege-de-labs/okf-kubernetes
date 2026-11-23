---
type: glossary-term
title: Network Policy
description: A specification of how groups of Pods are allowed to communicate with each other and with other network endpoints.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-network-policy
tags: [kubernetes, reference, glossary, networking, architecture, extension, core-object]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/reference/glossary/network-policy.md
    title: 'Kubernetes website source: reference/glossary/network-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-network-policy
    title: Network Policy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/reference/glossary/network-policy.md
---

# Network Policy

A specification of how groups of Pods are allowed to communicate with each other and with other network endpoints.

 

NetworkPolicies help you declaratively configure which Pods are allowed to connect to each other, which namespaces are allowed to communicate,
and more specifically which port numbers to enforce each policy on. NetworkPolicy objects use [labels](/reference/glossary/label.md)
to select Pods and define rules which specify what traffic is allowed to the selected Pods.

NetworkPolicies are implemented by a supported network plugin provided by a network provider.
Be aware that creating a NetworkPolicy object without a controller to implement it will have no effect.
