---
type: task
title: Set Bind Mount Options on Volume Mounts
description: This page shows how to apply security-related bind mount options (`noexec`, `nodev`, `nosuid`) to volume mounts in a Pod.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-bind-mount-options/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/configure-bind-mount-options.md
    title: 'Kubernetes website source: tasks/configure-pod-container/configure-bind-mount-options.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-bind-mount-options/
    title: Set Bind Mount Options on Volume Mounts | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configure-pod-container/configure-bind-mount-options.md
min_kubernetes_server_version: v1.37
---

# Set Bind Mount Options on Volume Mounts

**Feature state:** Kubernetes v1.37 [alpha] (feature gate `VolumeBindMountOptions`, disabled by default)

This page shows how to apply security-related bind mount options (`noexec`,
`nodev`, `nosuid`) to volume mounts in a Pod.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) Your Kubernetes server must be at or later than version v1.37. To check the version, enter `kubectl version`.

You need to have the `VolumeBindMountOptions`
[feature gate](/reference/command-line-tools-reference/feature-gates.md) enabled
on the API server **and** the kubelet. The container runtime must also support
the `mount_options` field in the CRI `Mount` message.

## Create a Pod with bind mount options {#create-pod}

The `.spec.containers[*].volumeMounts[*].bindMountOptions` field accepts a list of bind mount flags.
The allowed values are `noexec`, `nodev`, and `nosuid`.

For example, to mount an emptyDir volume at `/tmp` with `noexec` and `nosuid`
so that binaries cannot be executed and set-user-ID bits are ignored:

[`pods/bind-mount-options.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/bind-mount-options.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: bind-mount-options-demo
spec:
  containers:
  - image: registry.k8s.io/busybox
    name: test-container
    command: ["sleep", "3600"]
    volumeMounts:
    - mountPath: /tmp
      name: tmp-volume
      bindMountOptions:
      - noexec
      - nosuid
  volumes:
  - name: tmp-volume
    emptyDir: {}
```

1. Create the pod on your cluster:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/bind-mount-options.yaml
   ```

1. Verify the pod is running:

   ```shell
   kubectl get pod bind-mount-options-demo
   ```

1. Check the mount options on the volume:

   ```shell
   kubectl exec bind-mount-options-demo -- mount | grep /tmp
   ```

   The output should include `noexec` and `nosuid` in the mount options.

1. Verify that executing a binary on the mount fails:

   ```shell
   kubectl exec bind-mount-options-demo -- sh -c 'cp /bin/ls /tmp/ls && /tmp/ls'
   ```

   The output is similar to:

   ```none
   sh: /tmp/ls: Permission denied
   ```

1. Delete the Pod that you created for this exercise:

   ```shell
   kubectl delete pod bind-mount-options-demo
   ```

## What's next

- Learn more about [bind mount options](/concepts/storage/volumes.md#bind-mount-options) for volumes
