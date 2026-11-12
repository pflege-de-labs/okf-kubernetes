---
type: task
title: Use Cascading Deletion in a Cluster
description: This page shows you how to specify the type of cascading deletion to use in your cluster during garbage collection.
resource: https://kubernetes.io/docs/tasks/administer-cluster/use-cascading-deletion/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/use-cascading-deletion.md
    title: 'Kubernetes website source: tasks/administer-cluster/use-cascading-deletion.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/use-cascading-deletion/
    title: Use Cascading Deletion in a Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/use-cascading-deletion.md
---

# Use Cascading Deletion in a Cluster

This page shows you how to specify the type of
[cascading deletion](/concepts/architecture/garbage-collection.md#cascading-deletion)
to use in your cluster during [garbage collection](/reference/glossary/garbage-collection.md).

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

You also need to [create a sample Deployment](/tasks/run-application/run-stateless-application-deployment.md#creating-and-exploring-an-nginx-deployment) 
to experiment with the different types of cascading deletion. You will need to
recreate the Deployment for each type.

## Check owner references on your pods

Check that the `ownerReferences` field is present on your pods:

```shell 
kubectl get pods -l app=nginx --output=yaml
```

The output has an `ownerReferences` field similar to this:

```yaml
apiVersion: v1
    ...
    ownerReferences:
    - apiVersion: apps/v1
      blockOwnerDeletion: true
      controller: true
      kind: ReplicaSet
      name: nginx-deployment-6b474476c4
      uid: 4fdcd81c-bd5d-41f7-97af-3a3b759af9a7
    ...
```

## Use foreground cascading deletion {#use-foreground-cascading-deletion}

By default, Kubernetes uses [background cascading deletion](/concepts/architecture/garbage-collection.md#background-deletion)
to delete dependents of an object. You can switch to foreground cascading deletion
using either `kubectl` or the Kubernetes API, depending on the Kubernetes
version your cluster runs. To check the version, enter `kubectl version`.

You can delete objects using foreground cascading deletion using `kubectl` or the
Kubernetes API.

**Using kubectl**

Run the following command:

```shell
kubectl delete deployment nginx-deployment --cascade=foreground
```

**Using the Kubernetes API**

1. Start a local proxy session:

   ```shell
   kubectl proxy --port=8080
   ```

1. Use `curl` to trigger deletion:

   ```shell
   curl -X DELETE localhost:8080/apis/apps/v1/namespaces/default/deployments/nginx-deployment \
       -d '{"kind":"DeleteOptions","apiVersion":"v1","propagationPolicy":"Foreground"}' \
       -H "Content-Type: application/json"
   ```

   The output contains a `foregroundDeletion` [finalizer](/reference/glossary/finalizer.md)
   like this:

   ```
   "kind": "Deployment",
   "apiVersion": "apps/v1",
   "metadata": {
       "name": "nginx-deployment",
       "namespace": "default",
       "uid": "d1ce1b02-cae8-4288-8a53-30e84d8fa505",
       "resourceVersion": "1363097",
       "creationTimestamp": "2021-07-08T20:24:37Z",
       "deletionTimestamp": "2021-07-08T20:27:39Z",
       "finalizers": [
         "foregroundDeletion"
       ]
       ...
   ```

## Use background cascading deletion {#use-background-cascading-deletion}

1. [Create a sample Deployment](/tasks/run-application/run-stateless-application-deployment.md#creating-and-exploring-an-nginx-deployment).
1. Use either `kubectl` or the Kubernetes API to delete the Deployment,
   depending on the Kubernetes version your cluster runs. To check the version, enter `kubectl version`.

You can delete objects using background cascading deletion using `kubectl`
or the Kubernetes API.

Kubernetes uses background cascading deletion by default, and does so
even if you run the following commands without the `--cascade` flag or the
`propagationPolicy` argument.

**Using kubectl**

Run the following command:

```shell
kubectl delete deployment nginx-deployment --cascade=background
```

**Using the Kubernetes API**

1. Start a local proxy session:

   ```shell
   kubectl proxy --port=8080
   ```

1. Use `curl` to trigger deletion:

   ```shell
   curl -X DELETE localhost:8080/apis/apps/v1/namespaces/default/deployments/nginx-deployment \
       -d '{"kind":"DeleteOptions","apiVersion":"v1","propagationPolicy":"Background"}' \
       -H "Content-Type: application/json"
   ```

   The output is similar to this:

   ```
   "kind": "Status",
   "apiVersion": "v1",
   ...
   "status": "Success",
   "details": {
       "name": "nginx-deployment",
       "group": "apps",
       "kind": "deployments",
       "uid": "cc9eefb9-2d49-4445-b1c1-d261c9396456"
   }
   ```

## Delete owner objects and orphan dependents {#set-orphan-deletion-policy}

By default, when you tell Kubernetes to delete an object, the
[controller](/reference/glossary/controller.md) also deletes
dependent objects. You can make Kubernetes *orphan* these dependents using
`kubectl` or the Kubernetes API, depending on the Kubernetes version your
cluster runs. To check the version, enter `kubectl version`.

**Using kubectl**

Run the following command:

```shell
kubectl delete deployment nginx-deployment --cascade=orphan
```

**Using the Kubernetes API**

1. Start a local proxy session:

   ```shell
   kubectl proxy --port=8080
   ```

1. Use `curl` to trigger deletion:

   ```shell
   curl -X DELETE localhost:8080/apis/apps/v1/namespaces/default/deployments/nginx-deployment \
       -d '{"kind":"DeleteOptions","apiVersion":"v1","propagationPolicy":"Orphan"}' \
       -H "Content-Type: application/json"
   ```

   The output contains `orphan` in the `finalizers` field, similar to this:

   ```
   "kind": "Deployment",
   "apiVersion": "apps/v1",
   "namespace": "default",
   "uid": "6f577034-42a0-479d-be21-78018c466f1f",
   "creationTimestamp": "2021-07-09T16:46:37Z",
   "deletionTimestamp": "2021-07-09T16:47:08Z",
   "deletionGracePeriodSeconds": 0,
   "finalizers": [
     "orphan"
   ],
   ...
   ```

You can check that the Pods managed by the Deployment are still running:

```shell
kubectl get pods -l app=nginx
```

## What's next

* Learn about [owners and dependents](/concepts/overview/working-with-objects/owners-dependents.md) in Kubernetes.
* Learn about Kubernetes [finalizers](/concepts/overview/working-with-objects/finalizers.md).
* Learn about [garbage collection](/concepts/architecture/garbage-collection.md).
