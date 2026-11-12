---
type: task
title: Configure Quotas for API Objects
description: This page shows how to configure quotas for API objects, including PersistentVolumeClaims and Services.
resource: https://kubernetes.io/docs/tasks/administer-cluster/quota-api-object/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/quota-api-object.md
    title: 'Kubernetes website source: tasks/administer-cluster/quota-api-object.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/quota-api-object/
    title: Configure Quotas for API Objects | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/quota-api-object.md
---

# Configure Quotas for API Objects

This page shows how to configure quotas for API objects, including
PersistentVolumeClaims and Services. A quota restricts the number of
objects, of a particular type, that can be created in a namespace.
You specify quotas in a
[ResourceQuota](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#resourcequota-v1-core)
object.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Create a namespace

Create a namespace so that the resources you create in this exercise are
isolated from the rest of your cluster.

```shell
kubectl create namespace quota-object-example
```

## Create a ResourceQuota

Here is the configuration file for a ResourceQuota object:

[`admin/resource/quota-objects.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/admin/resource/quota-objects.yaml)

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: object-quota-demo
spec:
  hard:
    persistentvolumeclaims: "1"
    services.loadbalancers: "2"
    services.nodeports: "0"
```

Create the ResourceQuota:

```shell
kubectl apply -f https://k8s.io/examples/admin/resource/quota-objects.yaml --namespace=quota-object-example
```

View detailed information about the ResourceQuota:

```shell
kubectl get resourcequota object-quota-demo --namespace=quota-object-example --output=yaml
```

The output shows that in the quota-object-example namespace, there can be at most
one PersistentVolumeClaim, at most two Services of type LoadBalancer, and no Services
of type NodePort.

```yaml
status:
  hard:
    persistentvolumeclaims: "1"
    services.loadbalancers: "2"
    services.nodeports: "0"
  used:
    persistentvolumeclaims: "0"
    services.loadbalancers: "0"
    services.nodeports: "0"
```

## Create a PersistentVolumeClaim

Here is the configuration file for a PersistentVolumeClaim object:

[`admin/resource/quota-objects-pvc.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/admin/resource/quota-objects-pvc.yaml)

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-quota-demo
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
```

Create the PersistentVolumeClaim:

```shell
kubectl apply -f https://k8s.io/examples/admin/resource/quota-objects-pvc.yaml --namespace=quota-object-example
```

Verify that the PersistentVolumeClaim was created:

```shell
kubectl get persistentvolumeclaims --namespace=quota-object-example
```

The output shows that the PersistentVolumeClaim exists and has status Pending:

```
NAME             STATUS
pvc-quota-demo   Pending
```

## Attempt to create a second PersistentVolumeClaim

Here is the configuration file for a second PersistentVolumeClaim:

[`admin/resource/quota-objects-pvc-2.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/admin/resource/quota-objects-pvc-2.yaml)

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-quota-demo-2
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 4Gi
```

Attempt to create the second PersistentVolumeClaim:

```shell
kubectl apply -f https://k8s.io/examples/admin/resource/quota-objects-pvc-2.yaml --namespace=quota-object-example
```

The output shows that the second PersistentVolumeClaim was not created,
because it would have exceeded the quota for the namespace.

```
persistentvolumeclaims "pvc-quota-demo-2" is forbidden:
exceeded quota: object-quota-demo, requested: persistentvolumeclaims=1,
used: persistentvolumeclaims=1, limited: persistentvolumeclaims=1
```

## Notes

These are the strings used to identify API resources that can be constrained
by quotas:

<table>
<tr><th>String</th><th>API Object</th></tr>
<tr><td>"pods"</td><td>Pod</td></tr>
<tr><td>"services"</td><td>Service</td></tr>
<tr><td>"replicationcontrollers"</td><td>ReplicationController</td></tr>
<tr><td>"resourcequotas"</td><td>ResourceQuota</td></tr>
<tr><td>"secrets"</td><td>Secret</td></tr>
<tr><td>"configmaps"</td><td>ConfigMap</td></tr>
<tr><td>"persistentvolumeclaims"</td><td>PersistentVolumeClaim</td></tr>
<tr><td>"services.nodeports"</td><td>Service of type NodePort</td></tr>
<tr><td>"services.loadbalancers"</td><td>Service of type LoadBalancer</td></tr>
</table>

## Clean up

Delete your namespace:

```shell
kubectl delete namespace quota-object-example
```

## What's next

### For cluster administrators

* [Configure Default Memory Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/memory-default-namespace.md)

* [Configure Default CPU Requests and Limits for a Namespace](/tasks/administer-cluster/manage-resources/cpu-default-namespace.md)

* [Configure Minimum and Maximum Memory Constraints for a Namespace](/tasks/administer-cluster/manage-resources/memory-constraint-namespace.md)

* [Configure Minimum and Maximum CPU Constraints for a Namespace](/tasks/administer-cluster/manage-resources/cpu-constraint-namespace.md)

* [Configure Memory and CPU Quotas for a Namespace](/tasks/administer-cluster/manage-resources/quota-memory-cpu-namespace.md)

* [Configure a Pod Quota for a Namespace](/tasks/administer-cluster/manage-resources/quota-pod-namespace.md)

### For app developers

* [Assign Memory Resources to Containers and Pods](/tasks/configure-pod-container/assign-memory-resource.md)

* [Assign CPU Resources to Containers and Pods](/tasks/configure-pod-container/assign-cpu-resource.md)

* [Assign Pod-level CPU and memory resources](/tasks/configure-pod-container/assign-pod-level-resources.md)

* [Configure Quality of Service for Pods](/tasks/configure-pod-container/quality-service-pod.md)
