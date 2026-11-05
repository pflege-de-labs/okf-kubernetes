---
type: task
title: Configure Quality of Service for Pods
description: This page shows how to configure Pods so that they will be assigned particular Quality of Service (QoS) classes.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/quality-service-pod.md
    title: 'Kubernetes website source: tasks/configure-pod-container/quality-service-pod.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/
    title: Configure Quality of Service for Pods | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configure-pod-container/quality-service-pod.md
---

# Configure Quality of Service for Pods

This page shows how to configure Pods so that they will be assigned particular
[Quality of Service (QoS) classes](/reference/glossary/qos-class.md).
Kubernetes uses QoS classes to make decisions about evicting Pods when Node resources are exceeded.

When Kubernetes creates a Pod it assigns one of these QoS classes to the Pod:

* [Guaranteed](/concepts/workloads/pods/pod-qos.md#guaranteed)
* [Burstable](/concepts/workloads/pods/pod-qos.md#burstable)
* [BestEffort](/concepts/workloads/pods/pod-qos.md#besteffort)

**Note:**
Kubernetes assigns the QoS class when the Pod is created, and it remains unchanged
for the lifetime of the Pod. If you attempt to
[resize the Pod's resources](/tasks/configure-pod-container/resize-container-resources.md)
to values that would result in a different QoS class, control plane rejects your request with an error message.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

You also need to be able to create and delete namespaces.

## Create a namespace

Create a namespace so that the resources you create in this exercise are
isolated from the rest of your cluster.

```shell
kubectl create namespace qos-example
```

## Create a Pod that gets assigned a QoS class of Guaranteed

For a Pod to be given a QoS class of `Guaranteed`:

* Every Container in the Pod must have a memory limit and a memory request.
* For every Container in the Pod, the memory limit must equal the memory request.
* Every Container in the Pod must have a CPU limit and a CPU request.
* For every Container in the Pod, the CPU limit must equal the CPU request.

These restrictions apply to init containers and app containers equally.
[Ephemeral containers](/concepts/workloads/pods/ephemeral-containers.md)
cannot define resources so these restrictions do not apply.

Here is a manifest for a Pod that has one Container. The Container has a memory limit and a
memory request, both equal to 200 MiB. The Container has a CPU limit and a CPU request, both equal to 700 milliCPU:

[`pods/qos/qos-pod.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/qos/qos-pod.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: qos-demo
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-ctr
    image: nginx
    resources:
      limits:
        memory: "200Mi"
        cpu: "700m"
      requests:
        memory: "200Mi"
        cpu: "700m"
```

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/qos/qos-pod.yaml --namespace=qos-example
```

View detailed information about the Pod:

```shell
kubectl get pod qos-demo --namespace=qos-example --output=yaml
```

The output shows that Kubernetes gave the Pod a QoS class of `Guaranteed`. The output also
verifies that the Pod Container has a memory request that matches its memory limit, and it has
a CPU request that matches its CPU limit.

```yaml
spec:
  containers:
    ...
    resources:
      limits:
        cpu: 700m
        memory: 200Mi
      requests:
        cpu: 700m
        memory: 200Mi
    ...
status:
  qosClass: Guaranteed
```

**Note:**
If a Container specifies its own memory limit, but does not specify a memory request, Kubernetes
automatically assigns a memory request that matches the limit. Similarly, if a Container specifies its own
CPU limit, but does not specify a CPU request, Kubernetes automatically assigns a CPU request that matches
the limit.

#### Clean up {#clean-up-guaranteed}

Delete your Pod:

```shell
kubectl delete pod qos-demo --namespace=qos-example
```

## Create a Pod that gets assigned a QoS class of Burstable

A Pod is given a QoS class of `Burstable` if:

* The Pod does not meet the criteria for QoS class `Guaranteed`.
* At least one Container in the Pod has a memory or CPU request or limit.

Here is a manifest for a Pod that has one Container. The Container has a memory limit of 200 MiB
and a memory request of 100 MiB.

[`pods/qos/qos-pod-2.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/qos/qos-pod-2.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-2
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-2-ctr
    image: nginx
    resources:
      limits:
        memory: "200Mi"
      requests:
        memory: "100Mi"
```

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/qos/qos-pod-2.yaml --namespace=qos-example
```

View detailed information about the Pod:

```shell
kubectl get pod qos-demo-2 --namespace=qos-example --output=yaml
```

The output shows that Kubernetes gave the Pod a QoS class of `Burstable`:

```yaml
spec:
  containers:
  - image: nginx
    imagePullPolicy: Always
    name: qos-demo-2-ctr
    resources:
      limits:
        memory: 200Mi
      requests:
        memory: 100Mi
  ...
status:
  qosClass: Burstable
```

#### Clean up {#clean-up-burstable}

Delete your Pod:

```shell
kubectl delete pod qos-demo-2 --namespace=qos-example
```

## Create a Pod that gets assigned a QoS class of BestEffort

For a Pod to be given a QoS class of `BestEffort`, the Containers in the Pod must not
have any memory or CPU limits or requests.

Here is a manifest for a Pod that has one Container. The Container has no memory or CPU
limits or requests:

[`pods/qos/qos-pod-3.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/qos/qos-pod-3.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-3
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-3-ctr
    image: nginx
```

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/qos/qos-pod-3.yaml --namespace=qos-example
```

View detailed information about the Pod:

```shell
kubectl get pod qos-demo-3 --namespace=qos-example --output=yaml
```

The output shows that Kubernetes gave the Pod a QoS class of `BestEffort`:

```yaml
spec:
  containers:
    ...
    resources: {}
  ...
status:
  qosClass: BestEffort
```

#### Clean up {#clean-up-besteffort}

Delete your Pod:

```shell
kubectl delete pod qos-demo-3 --namespace=qos-example
```

## Create a Pod that has two Containers

Here is a manifest for a Pod that has two Containers. One container specifies a memory
request of 200 MiB. The other Container does not specify any requests or limits.

[`pods/qos/qos-pod-4.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/qos/qos-pod-4.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-4
  namespace: qos-example
spec:
  containers:

  - name: qos-demo-4-ctr-1
    image: nginx
    resources:
      requests:
        memory: "200Mi"

  - name: qos-demo-4-ctr-2
    image: redis
```

Notice that this Pod meets the criteria for QoS class `Burstable`. That is, it does not meet the
criteria for QoS class `Guaranteed`, and one of its Containers has a memory request.

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/qos/qos-pod-4.yaml --namespace=qos-example
```

View detailed information about the Pod:

```shell
kubectl get pod qos-demo-4 --namespace=qos-example --output=yaml
```

The output shows that Kubernetes gave the Pod a QoS class of `Burstable`:

```yaml
spec:
  containers:
    ...
    name: qos-demo-4-ctr-1
    resources:
      requests:
        memory: 200Mi
    ...
    name: qos-demo-4-ctr-2
    resources: {}
    ...
status:
  qosClass: Burstable
```

## Retrieve the QoS class for a Pod

Rather than see all the fields, you can view just the field you need:

```bash
kubectl --namespace=qos-example get pod qos-demo-4 -o jsonpath='{ .status.qosClass}{"\n"}'
```

```none
Burstable
```

## Clean up

Delete your namespace:

```shell
kubectl delete namespace qos-example
```

## What's next

### For app developers

* [Assign Memory Resources to Containers and Pods](/tasks/configure-pod-container/assign-memory-resource.md)

* [Assign CPU Resources to Containers and Pods](/tasks/configure-pod-container/assign-cpu-resource.md)

### For cluster administrators

* [Configure Default Memory Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/memory-default-namespace.md)

* [Configure Default CPU Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/cpu-default-namespace.md)

* [Configure Minimum and Maximum Memory Constraints for a Namespace](/tasks/administer-cluster/manage-resources/memory-constraint-namespace.md)

* [Configure Minimum and Maximum CPU Constraints for a Namespace](/tasks/administer-cluster/manage-resources/cpu-constraint-namespace.md)

* [Configure Memory and CPU Quotas for a Namespace](/tasks/administer-cluster/manage-resources/quota-memory-cpu-namespace.md)

* [Configure a Pod Quota for a Namespace](/tasks/administer-cluster/manage-resources/quota-pod-namespace.md)

* [Configure Quotas for API Objects](/tasks/administer-cluster/quota-api-object.md)

* [Control Topology Management policies on a node](/tasks/administer-cluster/topology-manager.md)
