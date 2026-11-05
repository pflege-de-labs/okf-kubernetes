---
type: concept
title: Kubernetes API Aggregation Layer
description: The aggregation layer allows Kubernetes to be extended with additional APIs, beyond what is offered by the core Kubernetes APIs.
resource: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/
tags: [kubernetes, concepts, extend-kubernetes]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation.md
    title: 'Kubernetes website source: concepts/extend-kubernetes/api-extension/apiserver-aggregation.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/
    title: Kubernetes API Aggregation Layer | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation.md
---

# Kubernetes API Aggregation Layer

The aggregation layer allows Kubernetes to be extended with additional APIs, beyond what is
offered by the core Kubernetes APIs.
The additional APIs can either be ready-made solutions such as a
[metrics server](https://github.com/kubernetes-sigs/metrics-server), or APIs that you develop yourself.

The aggregation layer is different from
[Custom Resource Definitions](/concepts/extend-kubernetes/api-extension/custom-resources.md),
which are a way to make the [kube-apiserver](/reference/glossary/kube-apiserver.md)
recognise new kinds of object.

## Aggregation layer

The aggregation layer runs in-process with the kube-apiserver. Until an extension resource is
registered, the aggregation layer will do nothing. To register an API, you add an _APIService_
object, which "claims" the URL path in the Kubernetes API. At that point, the aggregation layer
will proxy anything sent to that API path (e.g. `/apis/myextension.mycompany.io/v1/…`) to the
registered APIService.

The most common way to implement the APIService is to run an *extension API server* in Pod(s) that
run in your cluster. If you're using the extension API server to manage resources in your cluster,
the extension API server (also written as "extension-apiserver") is typically paired with one or
more [controllers](/reference/glossary/controller.md). The apiserver-builder
library provides a skeleton for both extension API servers and the associated controller(s).

### Response latency

Extension API servers should have low latency networking to and from the kube-apiserver.
Discovery requests are required to round-trip from the kube-apiserver in five seconds or less.

If your extension API server cannot achieve that latency requirement, consider making changes that
let you meet it.

## What's next

* To get the aggregator working in your environment, [configure the aggregation layer](/tasks/extend-kubernetes/configure-aggregation-layer.md).
* Then, [setup an extension api-server](/tasks/extend-kubernetes/setup-extension-api-server.md) to work with the aggregation layer.
* Read about [APIService](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/api-service-v1/) in the API reference
* Learn about [Declarative Validation Concepts](/reference/using-api/declarative-validation.md),
  an internal mechanism for defining validation rules that in the future will help support validation for extension API server development.

Alternatively: learn how to extend the Kubernetes API using
[Custom Resource Definitions](/tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md).
