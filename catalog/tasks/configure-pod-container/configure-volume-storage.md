---
type: task
title: Configure a Pod to Use a Volume for Storage
description: This page shows how to configure a Pod to use a Volume for storage.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/configure-volume-storage.md
    title: 'Kubernetes website source: tasks/configure-pod-container/configure-volume-storage.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/
    title: Configure a Pod to Use a Volume for Storage | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configure-pod-container/configure-volume-storage.md
---

# Configure a Pod to Use a Volume for Storage

This page shows how to configure a Pod to use a Volume for storage.

A Container's file system lives only as long as the Container does. So when a
Container terminates and restarts, filesystem changes are lost. For more
consistent storage that is independent of the Container, you can use a
[Volume](/concepts/storage/volumes.md). This is especially important for stateful
applications, such as key-value stores (such as Redis) and databases.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Configure a volume for a Pod

In this exercise, you create a Pod that runs one Container. This Pod has a
Volume of type
[emptyDir](/concepts/storage/volumes.md#emptydir)
that lasts for the life of the Pod, even if the Container terminates and
restarts. Here is the configuration file for the Pod:

[`pods/storage/redis.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/storage/redis.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: redis-storage
      mountPath: /data/redis
  volumes:
  - name: redis-storage
    emptyDir: {}
```

1. Create the Pod:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/storage/redis.yaml
   ```

1. Verify that the Pod's Container is running, and then watch for changes to
   the Pod:

   ```shell
   kubectl get pod redis --watch
   ```

   The output looks like this:

   ```console
   NAME      READY     STATUS    RESTARTS   AGE
   redis     1/1       Running   0          13s
   ```

1. In another terminal, get a shell to the running Container:

   ```shell
   kubectl exec -it redis -- /bin/bash
   ```

1. In your shell, go to `/data/redis`, and then create a file:

   ```shell
   root@redis:/data# cd /data/redis/
   root@redis:/data/redis# echo Hello > test-file
   ```

1. In your shell, list the running processes:

   ```shell
   root@redis:/data/redis# apt-get update
   root@redis:/data/redis# apt-get install procps
   root@redis:/data/redis# ps aux
   ```

   The output is similar to this:

   ```console
   USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
   redis        1  0.1  0.1  33308  3828 ?        Ssl  00:46   0:00 redis-server *:6379
   root        12  0.0  0.0  20228  3020 ?        Ss   00:47   0:00 /bin/bash
   root        15  0.0  0.0  17500  2072 ?        R+   00:48   0:00 ps aux
   ```

1. In your shell, kill the Redis process:

   ```shell
   root@redis:/data/redis# kill <pid>
   ```

   where `<pid>` is the Redis process ID (PID).

1. In your original terminal, watch for changes to the Redis Pod. Eventually,
   you will see something like this:

   ```console
   NAME      READY     STATUS     RESTARTS   AGE
   redis     1/1       Running    0          13s
   redis     0/1       Completed  0         6m
   redis     1/1       Running    1         6m
   ```

At this point, the Container has terminated and restarted. This is because the
Redis Pod has a
[restartPolicy](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#podspec-v1-core)
of `Always`.

1. Get a shell into the restarted Container:

   ```shell
   kubectl exec -it redis -- /bin/bash
   ```

1. In your shell, go to `/data/redis`, and verify that `test-file` is still there.

   ```shell
   root@redis:/data/redis# cd /data/redis/
   root@redis:/data/redis# ls
   test-file
   ```

1. Delete the Pod that you created for this exercise:

   ```shell
   kubectl delete pod redis
   ```

## What's next

- See [Volume](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#volume-v1-core).

- See [Pod](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#pod-v1-core).

- In addition to the local disk storage provided by `emptyDir`, Kubernetes
  supports many different network-attached storage solutions, including PD on
  GCE and EBS on EC2, which are preferred for critical data and will handle
  details such as mounting and unmounting the devices on the nodes. See
  [Volumes](/concepts/storage/volumes.md) for more details.
