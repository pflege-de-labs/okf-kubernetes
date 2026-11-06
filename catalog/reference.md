---
type: concept
title: Reference
description: This section of the Kubernetes documentation contains references.
resource: https://kubernetes.io/docs/reference/
tags: [kubernetes, reference.md]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/_index.md
    title: 'Kubernetes website source: reference/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/
    title: Reference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/_index.md
---

# Reference

This section of the Kubernetes documentation contains references.

## API Reference

* [Glossary](/reference/glossary.md) -  a comprehensive, standardized list of Kubernetes terminology

* [Kubernetes API Reference](/reference/kubernetes-api.md)
* [One-page API Reference for Kubernetes v1.37](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/)
* [Using The Kubernetes API](/reference/using-api.md) - overview of the API for Kubernetes.
* [API access control](/reference/access-authn-authz.md) - details on how Kubernetes controls API access
* [Well-Known Labels, Annotations and Taints](/reference/labels-annotations-taints.md)

## Officially supported client libraries

To call the Kubernetes API from a programming language, you can use
[client libraries](/reference/using-api/client-libraries.md). Officially supported
client libraries:

- [Kubernetes Go client library](https://github.com/kubernetes/client-go/)
- [Kubernetes Python client library](https://github.com/kubernetes-client/python)
- [Kubernetes Java client library](https://github.com/kubernetes-client/java)
- [Kubernetes JavaScript client library](https://github.com/kubernetes-client/javascript)
- [Kubernetes C# client library](https://github.com/kubernetes-client/csharp)
- [Kubernetes Haskell client library](https://github.com/kubernetes-client/haskell)

## CLI

* [kubectl](/reference/kubectl.md) - Main CLI tool for running commands and managing Kubernetes clusters.
  * [JSONPath](/reference/kubectl/jsonpath.md) - Syntax guide for using [JSONPath expressions](https://goessner.net/articles/JsonPath/) with kubectl.
* [kubeadm](/reference/setup-tools/kubeadm.md) - CLI tool to easily provision a secure Kubernetes cluster.

## Components

* [kubelet](/reference/command-line-tools-reference/kubelet.md) - The
  primary agent that runs on each node. The kubelet takes a set of PodSpecs
  and ensures that the described containers are running and healthy.
* [kube-apiserver](/reference/command-line-tools-reference/kube-apiserver.md) -
  REST API that validates and configures data for API objects such as  pods,
  services, replication controllers.
* [kube-controller-manager](/reference/command-line-tools-reference/kube-controller-manager.md) -
  Daemon that embeds the core control loops shipped with Kubernetes.
* [kube-proxy](/reference/command-line-tools-reference/kube-proxy.md) - Can
  do simple TCP/UDP stream forwarding or round-robin TCP/UDP forwarding across
  a set of back-ends.
* [kube-scheduler](/reference/command-line-tools-reference/kube-scheduler.md) -
  Scheduler that manages availability, performance, and capacity.
  
  * [Scheduler Policies](/reference/scheduling/policies.md)
  * [Scheduler Profiles](/reference/scheduling/config.md#profiles)

* List of [ports and protocols](/reference/networking/ports-and-protocols.md) that
  should be open on control plane and worker nodes

## Config APIs

This section hosts the documentation for "unpublished" APIs which are used to
configure  kubernetes components or tools. Most of these APIs are not exposed
by the API server in a RESTful way though they are essential for a user or an
operator to use or manage a cluster.

* [kubeconfig (v1)](/reference/config-api/kubeconfig.v1.md)
* [kuberc (v1alpha1)](/reference/config-api/kuberc.v1alpha1.md) and
  [kuberc (v1beta1)](/reference/config-api/kuberc.v1beta1.md)
* [kube-apiserver admission (v1)](/reference/config-api/apiserver-admission.v1.md)
  - `admission.k8s.io/v1`
    - [AdmissionReview](/reference/config-api/apiserver-admission.v1.md#admission-k8s-io-v1-AdmissionReview)
* [kube-apiserver configuration (v1alpha1)](/reference/config-api/apiserver-config.v1alpha1.md) and
  - `apiserver.k8s.io/v1`
    - [AdmissionConfiguration](/reference/config-api/apiserver-config.v1alpha1.md#apiserver-k8s-io-v1alpha1-AdmissionConfiguration)
    - [AuthenticationConfiguration](/reference/config-api/apiserver-config.v1alpha1.md#apiserver-k8s-io-v1alpha1-AuthenticationConfiguration)
    - [AuthorizationConfiguration](/reference/config-api/apiserver-config.v1alpha1.md#apiserver-k8s-io-v1alpha1-AuthorizationConfiguration)
    - [EgressSelectorConfiguration](/reference/config-api/apiserver-config.v1alpha1.md#apiserver-k8s-io-v1alpha1-EgressSelectorConfiguration)
    - [TracingConfiguration](/reference/config-api/apiserver-config.v1alpha1.md#apiserver-k8s-io-v1alpha1-TracingConfiguration)
 
* [kube-apiserver configuration (v1beta1)](/reference/config-api/apiserver-config.v1beta1.md)
  - `apiserver.k8s.io/v1beta1`
    - [AuthenticationConfiguration](/reference/config-api/apiserver-config.v1beta1.md#apiserver-k8s-io-v1beta1-AuthenticationConfiguration)
    - [AuthorizationConfiguration](/reference/config-api/apiserver-config.v1beta1.md#apiserver-k8s-io-v1beta1-AuthorizationConfiguration)
    - [EgressSelectorConfiguration](/reference/config-api/apiserver-config.v1beta1.md#apiserver-k8s-io-v1beta1-EgressSelectorConfiguration)
    - [TracingConfiguration](/reference/config-api/apiserver-config.v1beta1.md#apiserver-k8s-io-v1beta1-TracingConfiguration)
 
* [kube-apiserver configuration (v1)](/reference/config-api/apiserver-config.v1.md)
  - `apiserver.config.k8s.io/v1`
    - [AdmissionConfiguration](/reference/config-api/apiserver-config.v1.md#apiserver-config-k8s-io-v1-AdmissionConfiguration)
    - [AuthenticationConfiguration](/reference/config-api/apiserver-config.v1.md#apiserver-config-k8s-io-v1-AuthenticationConfiguration)
    - [AuthorizationConfiguration](/reference/config-api/apiserver-config.v1.md#apiserver-config-k8s-io-v1-AuthorizationConfiguration)
    - [EncryptionConfiguration](/reference/config-api/apiserver-config.v1.md#apiserver-config-k8s-io-v1-EncryptionConfiguration)
    - [TracingConfiguration](/reference/config-api/apiserver-config.v1.md#apiserver-config-k8s-io-v1-TracingConfiguration)
* [kube-apiserver event rate limit (v1alpha1)](/reference/config-api/apiserver-eventratelimit.v1alpha1.md)
  - `eventratelimit.admission.k8s.io/v1alpha1`
    - [Configuration](/reference/config-api/apiserver-eventratelimit.v1alpha1.md#eventratelimit-admission-k8s-io-v1alpha1-Configuration)
* [kubelet configuration (v1alpha1)](/reference/config-api/kubelet-config.v1alpha1.md)
  - `kubelet.config.k8s.io/v1alpha1`
    - [CredentialProviderConfig](/reference/config-api/kubelet-config.v1alpha1.md#kubelet-config-k8s-io-v1alpha1-CredentialProviderConfig)
    - [ImagePullIntent](/reference/config-api/kubelet-config.v1alpha1.md#kubelet-config-k8s-io-v1alpha1-ImagePullIntent)
    - [ImagePulledRecord](/reference/config-api/kubelet-config.v1alpha1.md#kubelet-config-k8s-io-v1alpha1-ImagePulledRecord)
* [kubelet configuration (v1beta1)](/reference/config-api/kubelet-config.v1beta1.md)
  - `kubelet.config.k8s.io/v1beta1`
    - [CredentialProviderConfig](/reference/config-api/kubelet-config.v1beta1.md#kubelet-config-k8s-io-v1beta1-CredentialProviderConfig)
    - [KubeletConfiguration](/reference/config-api/kubelet-config.v1beta1.md#kubelet-config-k8s-io-v1beta1-KubeletConfiguration)
    - [SerializedNodeConfigSource](/reference/config-api/kubelet-config.v1beta1.md#kubelet-config-k8s-io-v1beta1-SerializedNodeConfigSource)
* [kubelet configuration (v1)](/reference/config-api/kubelet-config.v1.md)
  - `kubelet.config.k8s.io/v1`
    - [CredentialProviderConfig](/reference/config-api/kubelet-config.v1.md#kubelet-config-k8s-io-v1-CredentialProviderConfig)
* [kubelet credential providers (v1)](/reference/config-api/kubelet-credentialprovider.v1.md)
  - `credentialprovider.kubelet.k8s.io/v1`
    - [CredentialProviderRequest](/reference/config-api/kubelet-credentialprovider.v1.md#credentialprovider-kubelet-k8s-io-v1-CredentialProviderRequest)
    - [CredentialProviderResponse](/reference/config-api/kubelet-credentialprovider.v1.md#credentialprovider-kubelet-k8s-io-v1-CredentialProviderResponse)
* [kube-scheduler configuration (v1)](/reference/config-api/kube-scheduler-config.v1.md)
  - `kubescheduler.config.k8s.io/v1`
    - [DefaultPreemptionArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-DefaultPreemptionArgs)
    - [DynamicResourcesArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-DynamicResourcesArgs)
    - [InterPodAffinityArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-InterPodAffinityArgs)
    - [KubeSchedulerConfiguration](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-KubeSchedulerConfiguration)
    - [NodeAffinityArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-NodeAffinityArgs)
    - [NodeResourcesBalancedAllocationArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-NodeResourcesBalancedAllocationArgs)
    - [NodeResourcesFitArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-NodeResourcesFitArgs)
    - [PodTopologySpreadArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-PodTopologySpreadArgs)
    - [VolumeBindingArgs](/reference/config-api/kube-scheduler-config.v1.md#kubescheduler-config-k8s-io-v1-VolumeBindingArgs)
* [kube-controller-manager configuration (v1alpha1)](/reference/config-api/kube-controller-manager-config.v1alpha1.md)
  - `kubecontrollermanager.config.k8s.io/v1alpha1`
    - [LeaderMigrationConfiguration](/reference/config-api/kube-controller-manager-config.v1alpha1.md#controllermanager-config-k8s-io-v1alpha1-LeaderMigrationConfiguration)
    - [KubeControllerManagerConfiguration](/reference/config-api/kube-controller-manager-config.v1alpha1.md#kubecontrollermanager-config-k8s-io-v1alpha1-KubeControllerManagerConfiguration)
  - `cloudcontrollermanager.config.k8s.io/v1alpha1`
    - [CloudControllerManagerConfiguration](/reference/config-api/kube-controller-manager-config.v1alpha1.md#cloudcontrollermanager-config-k8s-io-v1alpha1-CloudControllerManagerConfiguration)
* [kube-proxy configuration (v1alpha1)](/reference/config-api/kube-proxy-config.v1alpha1.md)
  - `kubeproxy.config.k8s.io/v1alpha1`
    - [KubeProxyConfiguration](/reference/config-api/kube-proxy-config.v1alpha1.md#kubeproxy-config-k8s-io-v1alpha1-KubeProxyConfiguration)
* [`audit.k8s.io/v1` API](/reference/config-api/apiserver-audit.v1.md)
  - `audit.k8s.io/v1`
    - [Event](/reference/config-api/apiserver-audit.v1.md#audit-k8s-io-v1-Event)
    - [EventList](/reference/config-api/apiserver-audit.v1.md#audit-k8s-io-v1-EventList)
    - [Policy](/reference/config-api/apiserver-audit.v1.md#audit-k8s-io-v1-Policy)
    - [PolicyList](/reference/config-api/apiserver-audit.v1.md#audit-k8s-io-v1-PolicyList)
* [Client authentication API (v1beta1)](/reference/config-api/client-authentication.v1beta1.md)
  - `client.authentication.k8s.io/v1beta1`
     - [ExecCredential](/reference/config-api/client-authentication.v1beta1.md#client-authentication-k8s-io-v1beta1-ExecCredential)
* [Client authentication API (v1)](/reference/config-api/client-authentication.v1.md)
  - `client.authentication.k8s.io/v1`
     - [ExecCredential](/reference/config-api/client-authentication.v1.md#client-authentication-k8s-io-v1-ExecCredential)
* [WebhookAdmission configuration (v1)](/reference/config-api/apiserver-webhookadmission.v1.md)
  - `apiserver.config.k8s.io/v1`
    - [WebhookAdmission](/reference/config-api/apiserver-webhookadmission.v1.md#apiserver-config-k8s-io-v1-WebhookAdmission)
* [ImagePolicy API (v1alpha1)](/reference/config-api/imagepolicy.v1alpha1.md)
  - `imagepolicy.k8s.io/v1alpha1`
    - [ImageReview](/reference/config-api/imagepolicy.v1alpha1.md#imagepolicy-k8s-io-v1alpha1-ImageReview)

## Config API for kubeadm

* [v1beta3](/reference/config-api/kubeadm-config.v1beta3.md)
* [v1beta4](/reference/config-api/kubeadm-config.v1beta4.md)

## External APIs

These are the APIs defined by the Kubernetes project, but are not implemented
by the core project:

* [Metrics API (v1beta1)](/reference/external-api/metrics.v1beta1.md)
* [Custom Metrics API (v1beta2)](https://kubernetes.io/docs/reference/external-api/custom-metrics.v1beta2)
* [External Metrics API (v1beta1)](https://kubernetes.io/docs/reference/external-api/external-metrics.v1beta1)

## Design Docs

An archive of the design docs for Kubernetes functionality. Good starting points are
[Kubernetes Architecture](https://git.k8s.io/design-proposals-archive/architecture/architecture.md) and
[Kubernetes Design Overview](https://git.k8s.io/design-proposals-archive).

## Encodings

Tools such as [kubectl](/reference/glossary/kubectl.md)
can work with different formats / encodings. These include:

* [CBOR](https://cbor.io/), used on the network but **not** available as a kubectl output format
  * See [CBOR resource encoding](/reference/using-api/api-concepts.md#cbor-encoding)
* [JSON](https://www.json.org/), available as a `kubectl` output format and also used at the HTTP layer
* [KYAML](/reference/encodings/kyaml.md), a Kubernetes dialect of YAML
  * KYAML is essentially an _output format_; any place where you can provide KYAML to Kubernetes, you can also provide any other valid YAML input
* [YAML](https://yaml.org/), available as a `kubectl` output format and also used at the HTTP layer

Kubernetes also has a custom [protobuf encoding](/reference/using-api/api-concepts.md#protobuf-encoding) that is only used within HTTP messages.

The `kubectl` tool supports some other output formats, such as _custom columns_;
see [output formats](/reference/kubectl.md#output-options) in the kubectl reference.
