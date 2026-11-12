---
type: task
title: Configure a Pod Quota for a Namespace
description: Restrict how many Pods you can create within a namespace.
resource: https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/quota-pod-namespace/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/manage-resources/quota-pod-namespace.md
    title: 'Kubernetes website source: tasks/administer-cluster/manage-resources/quota-pod-namespace.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/quota-pod-namespace/
    title: Configure a Pod Quota for a Namespace | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/manage-resources/quota-pod-namespace.md
---

# Configure a Pod Quota for a Namespace

This page shows how to set a quota for the total number of Pods that can run
in a [Namespace](/reference/glossary/namespace.md). You specify quotas in a
[ResourceQuota](/reference/kubernetes-api/core/resource-quota-v1.md)
object.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

You must have access to create namespaces in your cluster.

## Create a namespace

Create a namespace so that the resources you create in this exercise are
isolated from the rest of your cluster.

```shell
kubectl create namespace quota-pod-example
```

## Create a ResourceQuota

Here is an example manifest for a ResourceQuota:

[`admin/resource/quota-pod.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/admin/resource/quota-pod.yaml)

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: pod-demo
spec:
  hard:
    pods: "2"
```

Create the ResourceQuota:

```shell
kubectl apply -f https://k8s.io/examples/admin/resource/quota-pod.yaml --namespace=quota-pod-example
```

View detailed information about the ResourceQuota:

```shell
kubectl get resourcequota pod-demo --namespace=quota-pod-example --output=yaml
```

The output shows that the namespace has a quota of two Pods, and that currently there are
no Pods; that is, none of the quota is used.

```yaml
spec:
  hard:
    pods: "2"
status:
  hard:
    pods: "2"
  used:
    pods: "0"
```

Here is an example manifest for a [Deployment](/reference/glossary/deployment.md):

[`admin/resource/quota-pod-deployment.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/admin/resource/quota-pod-deployment.yaml)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pod-quota-demo
spec:
  selector:
    matchLabels:
      purpose: quota-demo
  replicas: 3
  template:
    metadata:
      labels:
        purpose: quota-demo
    spec:
      containers:
      - name: pod-quota-demo
        image: nginx
```

In that manifest, `replicas: 3` tells Kubernetes to attempt to create three new Pods, all
running the same application.

Create the Deployment:

```shell
kubectl apply -f https://k8s.io/examples/admin/resource/quota-pod-deployment.yaml --namespace=quota-pod-example
```

View detailed information about the Deployment:

```shell
kubectl get deployment pod-quota-demo --namespace=quota-pod-example --output=yaml
```

The output shows that even though the Deployment specifies three replicas, only two
Pods were created because of the quota you defined earlier:

```yaml
spec:
  ...
  replicas: 3
...
status:
  availableReplicas: 2
...
lastUpdateTime: 2021-04-02T20:57:05Z
    message: 'unable to create pods: pods "pod-quota-demo-1650323038-" is forbidden:
      exceeded quota: pod-demo, requested: pods=1, used: pods=2, limited: pods=2'
```

### Choice of resource

In this task you have defined a ResourceQuota that limited the total number of Pods, but
you could also limit the total number of other kinds of object. For example, you
might decide to limit how many [CronJobs](/reference/glossary/cronjob.md)
that can live in a single namespace.

## Clean up

Delete your namespace:

```shell
kubectl delete namespace quota-pod-example
```

## What's next

### For cluster administrators

* [Configure Default Memory Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/memory-default-namespace.md)

* [Configure Default CPU Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/cpu-default-namespace.md)

* [Configure Minimum and Maximum Memory Constraints for a Namespace](/tasks/administer-cluster/manage-resources/memory-constraint-namespace.md)

* [Configure Minimum and Maximum CPU Constraints for a Namespace](/tasks/administer-cluster/manage-resources/cpu-constraint-namespace.md)

* [Configure Memory and CPU Quotas for a Namespace](/tasks/administer-cluster/manage-resources/quota-memory-cpu-namespace.md)

* [Configure Quotas for API Objects](/tasks/administer-cluster/quota-api-object.md)

### For app developers

* [Assign Memory Resources to Containers and Pods](/tasks/configure-pod-container/assign-memory-resource.md)

* [Assign CPU Resources to Containers and Pods](/tasks/configure-pod-container/assign-cpu-resource.md)

* [Assign Pod-level CPU and memory resources](/tasks/configure-pod-container/assign-pod-level-resources.md)

* [Configure Quality of Service for Pods](/tasks/configure-pod-container/quality-service-pod.md)
