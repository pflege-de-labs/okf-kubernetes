---
type: task
title: Building a Basic DaemonSet
description: This page demonstrates how to build a basic DaemonSet that runs a Pod on every node in a Kubernetes cluster.
resource: https://kubernetes.io/docs/tasks/manage-daemon/create-daemon-set/
tags: [kubernetes, tasks, manage-daemon]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/manage-daemon/create-daemon-set.md
    title: 'Kubernetes website source: tasks/manage-daemon/create-daemon-set.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/manage-daemon/create-daemon-set/
    title: Building a Basic DaemonSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/manage-daemon/create-daemon-set.md
---

# Building a Basic DaemonSet

This page demonstrates how to build a basic [DaemonSet](/reference/glossary/daemonset.md)
that runs a Pod on every node in a Kubernetes cluster.
It covers a simple use case of mounting a file from the host, logging its contents using
an [init container](/concepts/workloads/pods/init-containers.md), and utilizing a pause container.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

A Kubernetes cluster with at least two nodes (one control plane node and one worker node)
to demonstrate the behavior of DaemonSets.

## Define the DaemonSet

In this task, a basic DaemonSet is created which ensures that the copy of a Pod is scheduled on every node.
The Pod will use an init container to read and log the contents of `/etc/machine-id` from the host,
while the main container will be a `pause` container, which keeps the Pod running.

[`application/basic-daemonset.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/application/basic-daemonset.yaml)

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: example-daemonset
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: example
  template:
    metadata:
      labels:
        app.kubernetes.io/name: example
    spec:
      containers:
      - name: pause
        image: registry.k8s.io/pause
      initContainers:
      - name: log-machine-id
        image: busybox:1.37
        command: ['sh', '-c', 'cat /etc/machine-id > /var/log/machine-id.log']
        volumeMounts:
        - name: machine-id
          mountPath: /etc/machine-id
          readOnly: true
        - name: log-dir
          mountPath: /var/log
      volumes:
      - name: machine-id
        hostPath:
          path: /etc/machine-id
          type: File
      - name: log-dir
        hostPath:
          path: /var/log
```

1. Create a DaemonSet based on the (YAML) manifest:

   ```shell
   kubectl apply -f https://k8s.io/examples/application/basic-daemonset.yaml
   ```

1. Once applied, you can verify that the DaemonSet is running a Pod on every node in the cluster:

   ```shell
   kubectl get pods -o wide
   ```

   The output will list one Pod per node, similar to:

   ```
   NAME                                READY   STATUS    RESTARTS   AGE    IP       NODE
   example-daemonset-xxxxx             1/1     Running   0          5m     x.x.x.x  node-1
   example-daemonset-yyyyy             1/1     Running   0          5m     x.x.x.x  node-2
   ```

1. You can inspect the contents of the logged `/etc/machine-id` file by checking
   the log directory mounted from the host:

   ```shell
   kubectl exec <pod-name> -- cat /var/log/machine-id.log
   ```

   Where `<pod-name>` is the name of one of your Pods.

## Cleaning up

To delete the DaemonSet, run this command:

```shell
kubectl delete --cascade=foreground --ignore-not-found --now daemonsets/example-daemonset
```

This simple DaemonSet example introduces key components like init containers and host path volumes,
which can be expanded upon for more advanced use cases. For more details refer to
[DaemonSet](/concepts/workloads/controllers/daemonset.md).

## What's next

* See [Performing a rolling update on a DaemonSet](/tasks/manage-daemon/update-daemon-set.md)
* See [Creating a DaemonSet to adopt existing DaemonSet pods](/concepts/workloads/controllers/daemonset.md)
