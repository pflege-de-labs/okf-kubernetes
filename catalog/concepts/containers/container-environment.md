---
type: concept
title: Container Environment
description: This page describes the resources available to Containers in the Container environment.
resource: https://kubernetes.io/docs/concepts/containers/container-environment/
tags: [kubernetes, concepts, containers]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/containers/container-environment.md
    title: 'Kubernetes website source: concepts/containers/container-environment.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/containers/container-environment/
    title: Container Environment | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/containers/container-environment.md
---

# Container Environment

This page describes the resources available to Containers in the Container environment. 

## Container environment

The Kubernetes Container environment provides several important resources to Containers:

* A filesystem, which is a combination of an [image](/concepts/containers/images.md) and one or more [volumes](/concepts/storage/volumes.md).
* Information about the Container itself.
* Information about other objects in the cluster.

### Container information

The *hostname* of a Container is the name of the Pod in which the Container is running.
It is available through the `hostname` command or the
[`gethostname`](https://man7.org/linux/man-pages/man2/gethostname.2.html)
function call in libc.

The Pod name and namespace are available as environment variables through the
[downward API](/tasks/inject-data-application/downward-api-volume-expose-pod-information.md).

User-defined environment variables from the Pod definition are also available to the Container,
as are any environment variables specified statically in the container image.

### Cluster information

A list of all services that were running when a Container was created is available to that Container as environment variables.
This list is limited to services within the same namespace as the new Container's Pod and Kubernetes control plane services.

For a service named *foo* that exposes a set of Pods, each running a container named *bar*,
the following variables are defined:

```shell
FOO_SERVICE_HOST=<the host the service is running on>
FOO_SERVICE_PORT=<the port the service is running on>
```

Services have dedicated IP addresses and are available to the Container via DNS,
if [DNS addon](https://releases.k8s.io/v1.37.1/cluster/addons/dns/) is enabled. 

## What's next

* Learn more about [Container lifecycle hooks](/concepts/containers/container-lifecycle-hooks.md).
* Get hands-on experience
  [attaching handlers to Container lifecycle events](/tasks/configure-pod-container/attach-handler-lifecycle-event.md).
