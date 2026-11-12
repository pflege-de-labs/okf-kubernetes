---
type: task
title: Running Pods on Only Some Nodes
description: This page demonstrates how can you run Pods on only some Nodes as part of a DaemonSet
resource: https://kubernetes.io/docs/tasks/manage-daemon/pods-some-nodes/
tags: [kubernetes, tasks, manage-daemon]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/manage-daemon/pods-some-nodes.md
    title: 'Kubernetes website source: tasks/manage-daemon/pods-some-nodes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/manage-daemon/pods-some-nodes/
    title: Running Pods on Only Some Nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/manage-daemon/pods-some-nodes.md
---

# Running Pods on Only Some Nodes

This page demonstrates how can you run [Pods](/reference/glossary/pod.md) on only some [Nodes](/reference/glossary/node.md) as part of a [DaemonSet](/reference/glossary/daemonset.md)

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

## Running Pods on only some Nodes

Imagine that you want to run a [DaemonSet](/reference/glossary/daemonset.md), but you only need to run those daemon pods
on nodes that have local solid state (SSD) storage. For example, the Pod might provide cache service to the
node, and the cache is only useful when low-latency local storage is available.

### Step 1: Add labels to your nodes

Add the label `ssd=true` to the nodes which have SSDs.

```shell
kubectl label nodes example-node-1 example-node-2 ssd=true
```

### Step 2: Create the manifest

Let's create a [DaemonSet](/reference/glossary/daemonset.md) which will provision the daemon pods on the SSD labeled [nodes](/reference/glossary/node.md) only.

Next, use a `nodeSelector` to ensure that the DaemonSet only runs Pods on nodes
with the `ssd` label set to `"true"`.

[`controllers/daemonset-label-selector.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/controllers/daemonset-label-selector.yaml)

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: ssd-driver
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: ssd-driver-pod
  template:
    metadata:
      labels:
        app: ssd-driver-pod
    spec:
      nodeSelector:
        ssd: "true"
      containers:
        - name: example-container
          image: example-image
```

### Step 3: Create the DaemonSet

Create the DaemonSet from the manifest by using `kubectl create` or `kubectl apply`

Let's label another node as `ssd=true`.

```shell
kubectl label nodes example-node-3 ssd=true
```

Labelling the node automatically triggers the control plane (specifically, the DaemonSet controller)
to run a new daemon pod on that node.

```shell
kubectl get pods -o wide
```
The output is similar to:

```
NAME                              READY     STATUS    RESTARTS   AGE    IP      NODE
<daemonset-name><some-hash-01>    1/1       Running   0          13s    .....   example-node-1
<daemonset-name><some-hash-02>    1/1       Running   0          13s    .....   example-node-2
<daemonset-name><some-hash-03>    1/1       Running   0          5s     .....   example-node-3
```
