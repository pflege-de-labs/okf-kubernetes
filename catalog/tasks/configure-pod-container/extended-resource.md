---
type: task
title: Assign Extended Resources to a Container
description: This page shows how to assign extended resources to a Container.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/extended-resource/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/extended-resource.md
    title: 'Kubernetes website source: tasks/configure-pod-container/extended-resource.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/extended-resource/
    title: Assign Extended Resources to a Container | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configure-pod-container/extended-resource.md
---

# Assign Extended Resources to a Container

**Feature state:** Kubernetes  [stable]

This page shows how to assign extended resources to a Container.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

Before you do this exercise, do the exercise in
[Advertise Extended Resources for a Node](/tasks/administer-cluster/extended-resource-node.md).
That will configure one of your Nodes to advertise a dongle resource.

## Assign an extended resource to a Pod

To request an extended resource, include the `resources.requests.<resource_name>` field
in the container manifest.
`*.kubernetes.io/`. Valid extended resource names have the form `example.com/foo` where
`example.com` is replaced with your organization's domain and `foo` is a
descriptive resource name.

Here is the configuration file for a Pod that has one Container:

[`pods/resource/extended-resource-pod.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/resource/extended-resource-pod.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: extended-resource-demo
spec:
  containers:
  - name: extended-resource-demo-ctr
    image: nginx
    resources:
      requests:
        example.com/dongle: 3
      limits:
        example.com/dongle: 3
```

In the configuration file, you can see that the Container requests 3 dongles.

Create a Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/resource/extended-resource-pod.yaml
```

Verify that the Pod is running:

```shell
kubectl get pod extended-resource-demo
```

Describe the Pod:

```shell
kubectl describe pod extended-resource-demo
```

The output shows dongle requests:

```yaml
Limits:
  example.com/dongle: 3
Requests:
  example.com/dongle: 3
```

## Attempt to create a second Pod

Here is the configuration file for a Pod that has one Container. The Container requests
two dongles.

[`pods/resource/extended-resource-pod-2.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/resource/extended-resource-pod-2.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: extended-resource-demo-2
spec:
  containers:
  - name: extended-resource-demo-2-ctr
    image: nginx
    resources:
      requests:
        example.com/dongle: 2
      limits:
        example.com/dongle: 2
```

Kubernetes will not be able to satisfy the request for two dongles, because the first Pod
used three of the four available dongles.

Attempt to create a Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/resource/extended-resource-pod-2.yaml
```

Describe the Pod

```shell
kubectl describe pod extended-resource-demo-2
```

The output shows that the Pod cannot be scheduled, because there is no Node that has
2 dongles available:

```
Conditions:
  Type    Status
  PodScheduled  False
...
Events:
  ...
  ... Warning   FailedScheduling  pod (extended-resource-demo-2) failed to fit in any node
fit failure summary on nodes : Insufficient example.com/dongle (1)
```

View the Pod status:

```shell
kubectl get pod extended-resource-demo-2
```

The output shows that the Pod was created, but not scheduled to run on a Node.
It has a status of Pending:

```yaml
NAME                       READY     STATUS    RESTARTS   AGE
extended-resource-demo-2   0/1       Pending   0          6m
```

## Clean up

Delete the Pods that you created for this exercise:

```shell
kubectl delete pod extended-resource-demo
kubectl delete pod extended-resource-demo-2
```

## Extended resources backed by DRA

**Feature state:** Kubernetes v1.37 [stable] (feature gate `DRAExtendedResource`, enabled by default)

The exercise above uses an extended resource that a Node advertises. Extended
resources can also be backed by
[Dynamic Resource Allocation](/reference/glossary/dra.md) (DRA):
a DeviceClass sets an `extendedResourceName`, and the scheduler satisfies matching
extended-resource requests using DRA devices instead of Node-advertised capacity.

A Pod requests the resource the same way in both cases — through
`resources.requests.<resource_name>` — so a workload does not need to know whether a
device plugin or DRA provides it. The same resource name can even be provided by a
device plugin on some Nodes and by DRA on others.

For the DeviceClass setup and examples, see
[Extended resource allocation by DRA](/concepts/resource-management/dynamic-resource-allocation/dra-api.md#extended-resource).

## What's next

### For application developers

* [Assign Memory Resources to Containers and Pods](/tasks/configure-pod-container/assign-memory-resource.md)
* [Assign CPU Resources to Containers and Pods](/tasks/configure-pod-container/assign-cpu-resource.md)

### For cluster administrators

* [Advertise Extended Resources for a Node](/tasks/administer-cluster/extended-resource-node.md)
