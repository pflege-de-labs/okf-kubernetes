---
type: concept
title: Effective Pod requests and limits
description: This page explains the _effective request_ and _effective limit_ of a Pod, and how Kubernetes calculates them from the resource requests and limits that you specify for the containers in that Pod.
resource: https://kubernetes.io/docs/concepts/resource-management/effective-pod-requests-limits/
tags: [kubernetes, concepts, resource-management]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/concepts/resource-management/effective-pod-requests-limits.md
    title: 'Kubernetes website source: concepts/resource-management/effective-pod-requests-limits.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/resource-management/effective-pod-requests-limits/
    title: Effective Pod requests and limits | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/concepts/resource-management/effective-pod-requests-limits.md
---

# Effective Pod requests and limits

This page explains the _effective request_ and _effective limit_ of a
[Pod](/reference/glossary/pod.md), and how Kubernetes calculates them
from the [resource requests and limits](/concepts/configuration/manage-resources-containers.md)
that you specify for the containers in that Pod.
Kubernetes uses these values for the Pod as a whole, for example when it schedules the Pod,
enforces resource quota, or sizes the Pod-level cgroup on Linux.

The [kube-scheduler](/reference/glossary/kube-scheduler.md)
places the Pod only on a node that has enough resources for the effective request,
and the node keeps the effective request reserved for the Pod,
even after initialization finishes.
The effective values are not always a simple sum of the container values,
because init containers run one at a time,
while sidecar containers keep running during and after initialization.
Understanding this calculation helps you set requests on init and sidecar containers
without reserving more node resources than the Pod needs.

## Calculating effective requests and limits

For each resource, Kubernetes calculates a request/limit for two phases:

* During initialization: init containers run one at a time, alongside any
  sidecar containers that have already started. The request/limit for this phase
  is the *effective init request/limit*, which is the highest amount needed at any
  point during initialization:
  * For each init container, take its request/limit for the resource and add the
    requests/limits of all sidecar containers that are listed before it in the
    `initContainers` field.
  * The highest of these values is the effective init request/limit.

  If any resource has no resource limit specified this is considered as the highest limit.
* After initialization: all sidecar containers and app containers run at the
  same time. The request/limit for this phase is the sum of the requests/limits of
  all non-init containers (app and sidecar containers).

The Pod's *effective request/limit* for a resource is the higher of the value during
initialization (the effective init request/limit) and the value after initialization
(the sum for all non-init containers), plus the
[pod overhead](/concepts/scheduling-eviction/pod-overhead.md).
See [Examples of effective requests](#examples-of-effective-requests) to learn how
this calculation works for specific Pods.

The following rules also apply:

* Scheduling is done based on effective requests/limits, which means
  init containers can reserve resources for initialization that are not used
  during the life of the Pod.
* The QoS (quality of service) tier of the Pod's *effective QoS tier* is the
  QoS tier for all init, sidecar and app containers alike.

Quota and limits are applied based on the effective Pod request and
limit.

## Examples of effective requests

The following examples show how the effective CPU request of a Pod is calculated.
None of these Pods has any pod overhead.

In this Pod, the sidecar container is listed before the init container:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: sidecar-before-init
spec:
  initContainers:
  - name: sidecar
    image: busybox:1.38
    command: ["sleep", "infinity"]
    restartPolicy: Always # make this container a sidecar
    resources:
      requests:
        cpu: 100m
  - name: init
    image: busybox:1.38
    command: ["sleep", "5"]
    resources:
      requests:
        cpu: 200m
  containers:
  - name: app
    image: busybox:1.38
    command: ["sleep", "infinity"]
    resources:
      requests:
        cpu: 50m
```

* During initialization, the `sidecar` container is already running when the `init`
  container starts, so the effective init request is 100m + 200m = 300m.
* After initialization, the `sidecar` and `app` containers run at the same time,
  so the sum of their requests is 100m + 50m = 150m.
* The effective CPU request of the Pod is the higher of the two values
  (300m and 150m), which is 300m.

This Pod has the same containers, but the init container is listed before the
sidecar container:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: init-before-sidecar
spec:
  initContainers:
  - name: init
    image: busybox:1.38
    command: ["sleep", "5"]
    resources:
      requests:
        cpu: 200m
  - name: sidecar
    image: busybox:1.38
    command: ["sleep", "infinity"]
    restartPolicy: Always # make this container a sidecar
    resources:
      requests:
        cpu: 100m
  containers:
  - name: app
    image: busybox:1.38
    command: ["sleep", "infinity"]
    resources:
      requests:
        cpu: 50m
```

* During initialization, no sidecar container is running when the `init` container
  starts, so the effective init request is 200m.
* After initialization, the `sidecar` and `app` containers run at the same time,
  so the sum of their requests is 100m + 50m = 150m.
* The effective CPU request of the Pod is the higher of the two values
  (200m and 150m), which is 200m.

In this Pod, the sidecar container is again listed before the init container, but this
time the app container requests more CPU:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: large-app
spec:
  initContainers:
  - name: sidecar
    image: busybox:1.38
    command: ["sleep", "infinity"]
    restartPolicy: Always # make this container a sidecar
    resources:
      requests:
        cpu: 100m
  - name: init
    image: busybox:1.38
    command: ["sleep", "5"]
    resources:
      requests:
        cpu: 200m
  containers:
  - name: app
    image: busybox:1.38
    command: ["sleep", "infinity"]
    resources:
      requests:
        cpu: 400m
```

* During initialization, the effective init request is 100m + 200m = 300m.
* After initialization, the `sidecar` and `app` containers run at the same time,
  so the sum of their requests is 100m + 400m = 500m.
* The effective CPU request of the Pod is the higher of the two values
  (300m and 500m), which is 500m.

## Linux cgroups {#cgroups}

On Linux, resource allocations for Pod level control groups (cgroups) are based on the effective Pod
request and limit, the same as the scheduler.

## What's next

* Read about [resource management for Pods and containers](/concepts/configuration/manage-resources-containers.md).
* Learn about [init containers](/concepts/workloads/pods/init-containers.md).
* Learn about [sidecar containers](/concepts/workloads/pods/sidecar-containers.md).
* Learn about [resource quotas](/concepts/policy/resource-quotas.md).
* Learn about [Pod Quality of Service classes](/concepts/workloads/pods/pod-qos.md).
