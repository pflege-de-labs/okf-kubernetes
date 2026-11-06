---
type: concept
title: Containers
description: Technology for packaging an application along with its runtime dependencies.
resource: https://kubernetes.io/docs/concepts/containers/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/containers/_index.md
    title: 'Kubernetes website source: concepts/containers/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/containers/
    title: Containers | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/containers/_index.md
---

# Containers

This page will discuss containers and container images, as well as their use in operations and solution development.

The word _container_ is an overloaded term. Whenever you use the word, check whether your audience uses the same definition.

Each container that you run is repeatable; the standardization from having
dependencies included means that you get the same behavior wherever you
run it.

Containers decouple applications from the underlying host infrastructure.
This makes deployment easier in different cloud or OS environments.

Each [node](/reference/glossary/node.md) in a Kubernetes
cluster runs the containers that form the
[Pods](/concepts/workloads/pods.md) assigned to that node.
Containers in a Pod are co-located and co-scheduled to run on the same node.

## Container images
A [container image](/concepts/containers/images.md) is a ready-to-run
software package containing everything needed to run an application:
the code and any runtime it requires, application and system libraries,
and default values for any essential settings.

Containers are intended to be stateless and
[immutable](https://glossary.cncf.io/immutable-infrastructure/):
you should not change
the code of a container that is already running. If you have a containerized
application and want to make changes, the correct process is to build a new
image that includes the change, then recreate the container to start from the
updated image.

## Container runtimes

A fundamental component that empowers Kubernetes to run containers effectively.
 It is responsible for managing the execution and lifecycle of containers within the Kubernetes environment.

Kubernetes supports container runtimes such as
[containerd](/reference/glossary/containerd.md), [CRI-O](/reference/glossary/cri-o.md),
and any other implementation of the [Kubernetes CRI (Container Runtime
Interface)](https://github.com/kubernetes/community/blob/main/contributors/devel/sig-node/container-runtime-interface.md).

Usually, you can allow your cluster to pick the default container runtime
for a Pod. If you need to use more than one container runtime in your cluster,
you can specify the [RuntimeClass](/concepts/containers/runtime-class.md)
for a Pod to make sure that Kubernetes runs those containers using a
particular container runtime.

You can also use RuntimeClass to run different Pods with the same container
runtime but with different settings.
