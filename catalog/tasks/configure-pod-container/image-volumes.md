---
type: task
title: Use an Image Volume With a Pod
description: This page shows how to configure a pod using image volumes.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/image-volumes/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/configure-pod-container/image-volumes.md
    title: 'Kubernetes website source: tasks/configure-pod-container/image-volumes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/image-volumes/
    title: Use an Image Volume With a Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/configure-pod-container/image-volumes.md
min_kubernetes_server_version: v1.31
---

# Use an Image Volume With a Pod

**Feature state:** Kubernetes v1.36 [stable] (feature gate `ImageVolume`, enabled by default)

This page shows how to configure a pod using image volumes. This allows you to
mount content from OCI registries inside containers.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) Your Kubernetes server must be at or later than version v1.31. To check the version, enter `kubectl version`.

- The container runtime needs to support the image volumes feature
- You need to exec commands in the host
- You need to be able to exec into pods

## Run a Pod that uses an image volume {#create-pod}

An image volume for a pod is enabled by setting the `volumes[*].image` field of `.spec`
to a valid reference and consuming it in the `volumeMounts` of the container. For example:

[`pods/image-volumes.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/image-volumes.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: image-volume
spec:
  containers:
  - name: shell
    command: ["sleep", "infinity"]
    image: debian
    volumeMounts:
    - name: volume
      mountPath: /volume
  volumes:
  - name: volume
    image:
      reference: quay.io/crio/artifact:v2
      pullPolicy: IfNotPresent
```

1. Create the pod on your cluster:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/image-volumes.yaml
   ```

1. Attach to the container:

   ```shell
   kubectl exec image-volume -it -- bash
   ```

1. Check the content of a file in the volume:

   ```shell
   cat /volume/dir/file
   ```

   The output is similar to:

   ```none
   1
   ```

   You can also check another file in a different path:

   ```shell
   cat /volume/file
   ```

   The output is similar to:

   ```none
   2
   ```

## Use `subPath` (or `subPathExpr`)

It is possible to utilize
[`subPath`](/concepts/storage/volumes.md#using-subpath) or
[`subPathExpr`](/concepts/storage/volumes.md#using-subpath-expanded-environment)
from Kubernetes v1.33 when using the image volume feature.

[`pods/image-volumes-subpath.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/image-volumes-subpath.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: image-volume
spec:
  containers:
  - name: shell
    command: ["sleep", "infinity"]
    image: debian
    volumeMounts:
    - name: volume
      mountPath: /volume
      subPath: dir
  volumes:
  - name: volume
    image:
      reference: quay.io/crio/artifact:v2
      pullPolicy: IfNotPresent
```

1. Create the pod on your cluster:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/image-volumes-subpath.yaml
   ```

1. Attach to the container:

   ```shell
   kubectl exec image-volume -it -- bash
   ```

1. Check the content of the file from the `dir` sub path in the volume:

   ```shell
   cat /volume/file
   ```

   The output is similar to:

   ```none
   1
   ```

## Further reading

- [`image` volumes](/concepts/storage/volumes.md#image)
