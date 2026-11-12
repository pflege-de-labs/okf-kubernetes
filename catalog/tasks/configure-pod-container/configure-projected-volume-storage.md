---
type: task
title: Configure a Pod to Use a Projected Volume for Storage
description: This page shows how to use a `projected` Volume to mount several existing volume sources into the same directory.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-projected-volume-storage/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/configure-pod-container/configure-projected-volume-storage.md
    title: 'Kubernetes website source: tasks/configure-pod-container/configure-projected-volume-storage.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-projected-volume-storage/
    title: Configure a Pod to Use a Projected Volume for Storage | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/configure-pod-container/configure-projected-volume-storage.md
---

# Configure a Pod to Use a Projected Volume for Storage

This page shows how to use a [`projected`](/concepts/storage/volumes.md#projected) Volume to mount
several existing volume sources into the same directory. Currently, `secret`, `configMap`, `downwardAPI`,
and `serviceAccountToken` volumes can be projected.

**Note:**
`serviceAccountToken` is not a volume type.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Configure a projected volume for a pod

In this exercise, you create username and password [Secrets](/reference/glossary/secret.md) from local files. You then create a Pod that runs one container, using a [`projected`](/concepts/storage/volumes.md#projected) Volume to mount the Secrets into the same shared directory.

Here is the configuration file for the Pod:

[`pods/storage/projected.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/storage/projected.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-projected-volume
spec:
  containers:
  - name: test-projected-volume
    image: busybox:1.28
    args:
    - sleep
    - "86400"
    volumeMounts:
    - name: all-in-one
      mountPath: "/projected-volume"
      readOnly: true
  volumes:
  - name: all-in-one
    projected:
      sources:
      - secret:
          name: user
      - secret:
          name: pass
```

1. Create the Secrets:

    ```shell
    # Create files containing the username and password:
    echo -n "admin" > ./username.txt
    echo -n "1f2d1e2e67df" > ./password.txt

    # Package these files into secrets:
    kubectl create secret generic user --from-file=./username.txt
    kubectl create secret generic pass --from-file=./password.txt
    ```
1. Create the Pod:

    ```shell
    kubectl apply -f https://k8s.io/examples/pods/storage/projected.yaml
    ```
1. Verify that the Pod's container is running, and then watch for changes to
the Pod:

    ```shell
    kubectl get --watch pod test-projected-volume
    ```
    The output looks like this:
    ```
    NAME                    READY     STATUS    RESTARTS   AGE
    test-projected-volume   1/1       Running   0          14s
    ```
1. In another terminal, get a shell to the running container:

    ```shell
    kubectl exec -it test-projected-volume -- /bin/sh
    ```
1. In your shell, verify that the `projected-volume` directory contains your projected sources:

    ```shell
    ls /projected-volume/
    ```

## Clean up

Delete the Pod and the Secrets:

```shell
kubectl delete pod test-projected-volume
kubectl delete secret user pass
```

## What's next

* Learn more about [`projected`](/concepts/storage/volumes.md#projected) volumes.
* Read the [all-in-one volume](https://git.k8s.io/design-proposals-archive/node/all-in-one-volume.md) design document.
