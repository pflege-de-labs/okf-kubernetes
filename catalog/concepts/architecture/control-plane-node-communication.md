---
type: concept
title: Communication between Nodes and the Control Plane
description: This document catalogs the communication paths between the API server and the Kubernetes cluster.
resource: https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/
tags: [kubernetes, concepts, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/architecture/control-plane-node-communication.md
    title: 'Kubernetes website source: concepts/architecture/control-plane-node-communication.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/
    title: Communication between Nodes and the Control Plane | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/architecture/control-plane-node-communication.md
---

# Communication between Nodes and the Control Plane

This document catalogs the communication paths between the [API server](/reference/glossary/kube-apiserver.md)
and the Kubernetes [cluster](/reference/glossary/cluster.md).
The intent is to allow users to customize their installation to harden the network configuration
such that the cluster can be run on an untrusted network (or on fully public IPs on a cloud
provider).

## Node to Control Plane

Kubernetes has a "hub-and-spoke" API pattern. All API usage from nodes (or the pods they run)
terminates at the API server. None of the other control plane components are designed to expose
remote services. The API server is configured to listen for remote connections on a secure HTTPS
port (typically 443) with one or more forms of client
[authentication](/reference/access-authn-authz/authentication.md) enabled.
One or more forms of [authorization](/reference/access-authn-authz/authorization.md) should be
enabled, especially if [anonymous requests](/reference/access-authn-authz/authentication.md#anonymous-requests)
or [service account tokens](/reference/access-authn-authz/authentication.md#service-account-tokens)
are allowed.

Nodes should be provisioned with the public root [certificate](/reference/glossary/certificate.md) for the cluster such that they can
connect securely to the API server along with valid client credentials. A good approach is that the
client credentials provided to the kubelet are in the form of a client certificate. See
[kubelet TLS bootstrapping](/reference/access-authn-authz/kubelet-tls-bootstrapping.md)
for automated provisioning of kubelet client certificates.

[Pods](/reference/glossary/pod.md) that wish to connect to the API server can do so securely by leveraging a service account so
that Kubernetes will automatically inject the public root certificate and a valid bearer token
into the pod when it is instantiated.
The `kubernetes` service (in `default` namespace) is configured with a virtual IP address that is
redirected (via `[kube-proxy](/reference/glossary/kube-proxy.md)`) to the HTTPS endpoint on the API server.

The control plane components also communicate with the API server over the secure port.

As a result, the default operating mode for connections from the nodes and pod running on the
nodes to the control plane is secured by default and can run over untrusted and/or public
networks.

## Control plane to node

There are two primary communication paths from the control plane (the API server) to the nodes.
The first is from the API server to the [kubelet](/reference/glossary/kubelet.md) process which runs on each node in the cluster.
The second is from the API server to any node, pod, or service through the API server's _proxy_
functionality.

### API server to kubelet

The connections from the API server to the kubelet are used for:

* Fetching logs for pods.
* Attaching (usually through `kubectl`) to running pods.
* Providing the kubelet's port-forwarding functionality.

These connections terminate at the kubelet's HTTPS endpoint. By default, the API server does not
verify the kubelet's serving certificate, which makes the connection subject to man-in-the-middle
attacks and **unsafe** to run over untrusted and/or public networks.

To verify this connection, use the `--kubelet-certificate-authority` flag to provide the API
server with a root certificate bundle to use to verify the kubelet's serving certificate.

If that is not possible, use [SSH tunneling](#ssh-tunnels) between the API server and kubelet if
required to avoid connecting over an
untrusted or public network.

Finally, [Kubelet authentication and/or authorization](/reference/access-authn-authz/kubelet-authn-authz.md)
should be enabled to secure the kubelet API.

### API server to nodes, pods, and services

The connections from the API server to a node, pod, or service default to plain HTTP connections
and are therefore neither authenticated nor encrypted. They can be run over a secure HTTPS
connection by prefixing `https:` to the node, pod, or service name in the API URL, but they will
not validate the certificate provided by the HTTPS endpoint nor provide client credentials. So
while the connection will be encrypted, it will not provide any guarantees of integrity. These
connections **are not currently safe** to run over untrusted or public networks.

### SSH tunnels

Kubernetes supports [SSH tunnels](https://www.ssh.com/academy/ssh/tunneling) to protect the control plane to nodes communication paths. In this
configuration, the API server initiates an SSH tunnel to each node in the cluster (connecting to
the SSH server listening on port 22) and passes all traffic destined for a kubelet, node, pod, or
service through the tunnel.
This tunnel ensures that the traffic is not exposed outside of the network in which the nodes are
running.

**Note:**
SSH tunnels are currently deprecated, so you shouldn't opt to use them unless you know what you
are doing. The [Konnectivity service](#konnectivity-service) is a replacement for this
communication channel.

### Konnectivity service

**Feature state:** Kubernetes v1.18 [beta]

As a replacement to the SSH tunnels, the Konnectivity service provides TCP level proxy for the
control plane to cluster communication. The Konnectivity service consists of two parts: the
Konnectivity server in the control plane network and the Konnectivity agents in the nodes network.
The Konnectivity agents initiate connections to the Konnectivity server and maintain the network
connections.
After enabling the Konnectivity service, all control plane to nodes traffic goes through these
connections.

Follow the [Konnectivity service task](/tasks/extend-kubernetes/setup-konnectivity.md) to set
up the Konnectivity service in your cluster.

## What's next

* Read about the [Kubernetes control plane components](/concepts/architecture.md#control-plane-components)
* Learn more about [Hubs and Spoke model](https://book.kubebuilder.io/multiversion-tutorial/conversion-concepts.html#hubs-spokes-and-other-wheel-metaphors)
* Learn how to [Secure a Cluster](/tasks/administer-cluster/securing-a-cluster.md) 
* Learn more about the [Kubernetes API](/concepts/overview/kubernetes-api.md)
* [Set up Konnectivity service](/tasks/extend-kubernetes/setup-konnectivity.md)
* [Use Port Forwarding to Access Applications in a Cluster](/tasks/access-application-cluster/port-forward-access-application-cluster.md)
* Learn how to [Fetch logs for Pods](/tasks/debug/debug-application/debug-running-pod.md#examine-pod-logs), [use kubectl port-forward](/tasks/access-application-cluster/port-forward-access-application-cluster.md#forward-a-local-port-to-a-port-on-the-pod)
