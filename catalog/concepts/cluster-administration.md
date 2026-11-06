---
type: concept
title: Cluster Administration
description: Lower-level detail relevant to creating or administering a Kubernetes cluster.
resource: https://kubernetes.io/docs/concepts/cluster-administration/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/cluster-administration/_index.md
    title: 'Kubernetes website source: concepts/cluster-administration/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/cluster-administration/
    title: Cluster Administration | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/cluster-administration/_index.md
---

# Cluster Administration

The cluster administration overview is for anyone creating or administering a Kubernetes cluster.
It assumes some familiarity with core Kubernetes [concepts](/concepts.md).

## Planning a cluster

See the guides in [Setup](https://kubernetes.io/docs/setup/) for examples of how to plan, set up, and configure
Kubernetes clusters. The solutions listed in this article are called *distros*.

**Note:**
Not all distros are actively maintained. Choose distros which have been tested with a recent
version of Kubernetes.

Before choosing a guide, here are some considerations:

- Do you want to try out Kubernetes on your computer, or do you want to build a high-availability,
  multi-node cluster? Choose distros best suited for your needs.
- Will you be using **a hosted Kubernetes cluster**, such as
  [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/), or **hosting your own cluster**?
- Will your cluster be **on-premises**, or **in the cloud (IaaS)**? Kubernetes does not directly
  support hybrid clusters. Instead, you can set up multiple clusters.
- **If you are configuring Kubernetes on-premises**, consider which
  [networking model](/concepts/cluster-administration/networking.md) fits best.
- Will you be running Kubernetes on **"bare metal" hardware** or on **virtual machines (VMs)**?
- Do you **want to run a cluster**, or do you expect to do **active development of Kubernetes project code**?
  If the latter, choose an actively-developed distro. Some distros only use binary releases, but
  offer a greater variety of choices.
- Familiarize yourself with the [components](/concepts/overview/components.md) needed to run a cluster.

## Managing a cluster

* Learn how to [manage nodes](/concepts/architecture/nodes.md).
  * Read about [Node autoscaling](/concepts/cluster-administration/node-autoscaling.md).

* Learn how to set up and manage the [resource quota](/concepts/policy/resource-quotas.md) for shared clusters.

## Securing a cluster

* [Generate Certificates](/tasks/administer-cluster/certificates.md) describes the steps to
  generate certificates using different tool chains.

* [Kubernetes Container Environment](/concepts/containers/container-environment.md) describes
  the environment for Kubelet managed containers on a Kubernetes node.

* [Controlling Access to the Kubernetes API](/concepts/security/controlling-access.md) describes
  how Kubernetes implements access control for its own API.

* [Authenticating](/reference/access-authn-authz/authentication.md) explains authentication in
  Kubernetes, including the various authentication options.

* [Authorization](/reference/access-authn-authz/authorization.md) is separate from
  authentication, and controls how HTTP calls are handled.

* [Using Admission Controllers](/reference/access-authn-authz/admission-controllers.md)
  explains plug-ins which intercepts requests to the Kubernetes API server after authentication
  and authorization.

* [Admission Webhook Good Practices](/concepts/cluster-administration/admission-webhooks-good-practices.md)
  provides good practices and considerations when designing mutating admission
  webhooks and validating admission webhooks.

* [Using Sysctls in a Kubernetes Cluster](/tasks/administer-cluster/sysctl-cluster.md)
  describes to an administrator how to use the `sysctl` command-line tool to set kernel parameters
.

* [Auditing](/tasks/debug/debug-cluster/audit.md) describes how to interact with Kubernetes'
  audit logs.

### Securing the kubelet

* [Control Plane-Node communication](/concepts/architecture/control-plane-node-communication.md)
* [TLS bootstrapping](/reference/access-authn-authz/kubelet-tls-bootstrapping.md)
* [Kubelet authentication/authorization](/reference/access-authn-authz/kubelet-authn-authz.md)

## Optional Cluster Services

* [DNS Integration](/concepts/services-networking/dns-pod-service.md) describes how to resolve
  a DNS name directly to a Kubernetes service.

* [Logging and Monitoring Cluster Activity](/concepts/cluster-administration/logging.md)
  explains how logging in Kubernetes works and how to implement it.
