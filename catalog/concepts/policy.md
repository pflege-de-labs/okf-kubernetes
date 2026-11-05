---
type: section
title: Policies
description: Manage security and best-practices with policies.
resource: https://kubernetes.io/docs/concepts/policy/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/policy/_index.md
    title: 'Kubernetes website source: concepts/policy/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/policy/
    title: Policies | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/policy/_index.md
---

# Policies

Kubernetes policies are configurations that manage other configurations or runtime behaviors. Kubernetes offers various forms of policies, described below:

## Apply policies using API objects

 Some API objects act as policies. Here are some examples:
* [NetworkPolicies](/concepts/services-networking/network-policies.md) can be used to restrict ingress and egress traffic for a workload.
* [LimitRanges](/concepts/policy/limit-range.md) manage resource allocation constraints across different object kinds.
* [ResourceQuotas](/concepts/policy/resource-quotas.md) limit resource consumption for a [namespace](/reference/glossary/namespace.md).

## Apply policies using admission controllers

An [admission controller](/reference/glossary/admission-controller.md)
runs in the API server
and can validate or mutate API requests. Some admission controllers act to apply policies.
For example, the [AlwaysPullImages](/reference/access-authn-authz/admission-controllers.md#alwayspullimages) admission controller modifies a new Pod to set the image pull policy to `Always`.

Kubernetes has several built-in admission controllers that are configurable via the API server `--enable-admission-plugins` flag.

Details on admission controllers, with the complete list of available admission controllers, are documented in a dedicated section:

* [Admission Controllers](/reference/access-authn-authz/admission-controllers.md)

## Apply policies using ValidatingAdmissionPolicy

Validating admission policies allow configurable validation checks to be executed in the API server using the Common Expression Language (CEL). For example, a `ValidatingAdmissionPolicy` can be used to disallow use of the `latest` image tag.

A `ValidatingAdmissionPolicy` operates on an API request and can be used to block, audit, and warn users about non-compliant configurations.

Details on the `ValidatingAdmissionPolicy` API, with examples, are documented in a dedicated section:
* [Validating Admission Policy](/reference/access-authn-authz/validating-admission-policy.md)

## Apply policies using dynamic admission control

Dynamic admission controllers (or admission webhooks) run outside the API server as separate applications that register to receive webhooks requests to perform validation or mutation of API requests. 

Dynamic admission controllers can be used to apply policies on API requests and trigger other policy-based workflows. A dynamic admission controller can perform complex checks including those that require retrieval of other cluster resources and external data. For example, an image verification check can lookup data from OCI registries to validate the container image signatures and attestations.

Details on dynamic admission control are documented in a dedicated section:
* [Dynamic Admission Control](/reference/access-authn-authz/extensible-admission-controllers.md)

### Implementations {#implementations-admission-control}

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

Dynamic Admission Controllers that act as flexible policy engines are being developed in the Kubernetes ecosystem, such as:
- [Kubewarden](https://github.com/kubewarden)
- [Kyverno](https://kyverno.io)
- [OPA Gatekeeper](https://github.com/open-policy-agent/gatekeeper)
- [Polaris](https://polaris.docs.fairwinds.com/admission-controller/)

## Apply policies using Kubelet configurations

Kubernetes allows configuring the Kubelet on each worker node.  Some Kubelet configurations act as policies:
* [Process ID limits and reservations](/concepts/policy/pid-limiting.md) are used to limit and reserve allocatable PIDs.
* [Node Resource Managers](https://kubernetes.io/docs/concepts/policy/node-resource-managers/) can manage compute, memory, and device resources for latency-critical and high-throughput workloads.
