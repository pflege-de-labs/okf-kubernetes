---
type: task
title: Assign Pods to Nodes
description: This page shows how to assign a Kubernetes Pod to a particular node in a Kubernetes cluster.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/assign-pods-nodes.md
    title: 'Kubernetes website source: tasks/configure-pod-container/assign-pods-nodes.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/
    title: Assign Pods to Nodes | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/configure-pod-container/assign-pods-nodes.md
---

# Assign Pods to Nodes

This page shows how to assign a Kubernetes Pod to a particular node in a
Kubernetes cluster.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Add a label to a node

1. List the [nodes](/reference/glossary/node.md) in your cluster, along with their labels:

    ```shell
    kubectl get nodes --show-labels
    ```

    The output is similar to this:

    ```shell
    NAME      STATUS    ROLES    AGE     VERSION        LABELS
    worker0   Ready     <none>   1d      v1.13.0        ...,kubernetes.io/hostname=worker0
    worker1   Ready     <none>   1d      v1.13.0        ...,kubernetes.io/hostname=worker1
    worker2   Ready     <none>   1d      v1.13.0        ...,kubernetes.io/hostname=worker2
    ```
1. Choose one of your nodes, and add a label to it:

    ```shell
    kubectl label nodes <your-node-name> disktype=ssd
    ```

    where `<your-node-name>` is the name of your chosen node.

1. Verify that your chosen node has a `disktype=ssd` label:

    ```shell
    kubectl get nodes --show-labels
    ```

    The output is similar to this:

    ```shell
    NAME      STATUS    ROLES    AGE     VERSION        LABELS
    worker0   Ready     <none>   1d      v1.13.0        ...,disktype=ssd,kubernetes.io/hostname=worker0
    worker1   Ready     <none>   1d      v1.13.0        ...,kubernetes.io/hostname=worker1
    worker2   Ready     <none>   1d      v1.13.0        ...,kubernetes.io/hostname=worker2
    ```

    In the preceding output, you can see that the `worker0` node has a
    `disktype=ssd` label.

## Create a pod that gets scheduled to your chosen node

This pod configuration file describes a pod that has a node selector,
`disktype: ssd`. This means that the pod will get scheduled on a node that has
a `disktype=ssd` label.

[`pods/pod-nginx.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/pod-nginx.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    disktype: ssd
```

1. Use the configuration file to create a pod that will get scheduled on your
   chosen node:
    
    ```shell
    kubectl apply -f https://k8s.io/examples/pods/pod-nginx.yaml
    ```

1. Verify that the pod is running on your chosen node:

    ```shell
    kubectl get pods --output=wide
    ```

    The output is similar to this:
    
    ```shell
    NAME     READY     STATUS    RESTARTS   AGE    IP           NODE
    nginx    1/1       Running   0          13s    10.200.0.4   worker0
    ```
## Create a pod that gets scheduled to specific node

You can also schedule a pod to one specific node via setting `nodeName`.

[`pods/pod-nginx-specific-node.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/pods/pod-nginx-specific-node.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  nodeName: foo-node # schedule pod to specific node
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
```

Use the configuration file to create a pod that will get scheduled on `foo-node` only.

## What's next

* Learn more about [labels and selectors](/concepts/overview/working-with-objects/labels.md).
* Learn more about [nodes](/concepts/architecture/nodes.md).
