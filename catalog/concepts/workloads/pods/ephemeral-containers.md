---
type: concept
title: Ephemeral Containers
description: 'This page provides an overview of ephemeral containers: a special type of container that runs temporarily in an existing Pod to accomplish user-initiated actions such as troubleshooting.'
resource: https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/workloads/pods/ephemeral-containers.md
    title: 'Kubernetes website source: concepts/workloads/pods/ephemeral-containers.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/
    title: Ephemeral Containers | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/workloads/pods/ephemeral-containers.md
---

# Ephemeral Containers

**Feature state:** Kubernetes v1.25 [stable]

This page provides an overview of ephemeral containers: a special type of container
that runs temporarily in an existing [Pod](/reference/glossary/pod.md) to
accomplish user-initiated actions such as troubleshooting. You use ephemeral
containers to inspect services rather than to build applications.

## Understanding ephemeral containers

[Pods](/reference/glossary/pod.md) are the fundamental building
block of Kubernetes applications. Since Pods are intended to be disposable and
replaceable, you cannot add a container to a Pod once it has been created.
Instead, you usually delete and replace Pods in a controlled fashion using
[deployments](/reference/glossary/deployment.md).

Sometimes it's necessary to inspect the state of an existing Pod, however, for
example to troubleshoot a hard-to-reproduce bug. In these cases you can run
an ephemeral container in an existing Pod to inspect its state and run
arbitrary commands.

### What is an ephemeral container?

Ephemeral containers differ from other containers in that they lack guarantees
for resources or execution, and they will never be automatically restarted, so
they are not appropriate for building applications.  Ephemeral containers are
described using the same `ContainerSpec` as regular containers, but many fields
are incompatible and disallowed for ephemeral containers.

- Ephemeral containers may not have ports, so fields such as `ports`,
  `livenessProbe`, `readinessProbe` are disallowed.
- Pod resource allocations are immutable, so setting `resources` is disallowed.
- For a complete list of allowed fields, see the [EphemeralContainer reference
  documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#ephemeralcontainer-v1-core).

Ephemeral containers are created using a special `ephemeralcontainers` handler
in the API rather than by adding them directly to `pod.spec`, so it's not
possible to add an ephemeral container using `kubectl edit`.

Like regular containers, you may not change or remove an ephemeral container
after you have added it to a Pod.

**Note:**
Ephemeral containers are not supported by [static pods](/tasks/configure-pod-container/static-pod.md).

## Uses for ephemeral containers

Ephemeral containers are useful for interactive troubleshooting when `kubectl
exec` is insufficient because a container has crashed or a container image
doesn't include debugging utilities.

In particular, [distroless images](https://github.com/GoogleContainerTools/distroless)
enable you to deploy minimal container images that reduce attack surface
and exposure to bugs and vulnerabilities. Since distroless images do not include a
shell or any debugging utilities, it's difficult to troubleshoot distroless
images using `kubectl exec` alone.

When using ephemeral containers, it's helpful to enable [process namespace
sharing](/tasks/configure-pod-container/share-process-namespace.md) so
you can view processes in other containers.

## Access control {#access-control}

Adding an ephemeral container to a running Pod uses the Pod's
`ephemeralcontainers` subresource rather than being writable via `pod.spec`.
This is a deliberate design choice. Some admission plugins and security policies don't recognize ephemeral containers, and treating
ephemeral container creation as a distinct API surface lets cluster
administrators make an explicit decision about who may add containers to
already-running Pods.

For the same reason, the `pods/ephemeralcontainers` subresource is **not**
included in the default `admin` and `edit`
[user-facing ClusterRoles](/reference/access-authn-authz/rbac.md#user-facing-roles).
To allow users to add ephemeral containers (for example, by running
[`kubectl debug`](/reference/kubectl/generated/kubectl_debug.md)),
cluster administrators must grant the `patch` (and typically `update`) verb
on `pods/ephemeralcontainers` explicitly. See
[Debug Running Pods: Permissions for ephemeral containers](/tasks/debug/debug-application/debug-running-pod.md#ephemeral-container-permissions)
for recommended approaches.

## What's next

* Learn how to [debug pods using ephemeral containers](/tasks/debug/debug-application/debug-running-pod.md#ephemeral-container).
