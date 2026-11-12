---
type: task
title: Use a User Namespace With a Pod
description: This page shows how to configure a user namespace for pods.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/user-namespaces/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/configure-pod-container/user-namespaces.md
    title: 'Kubernetes website source: tasks/configure-pod-container/user-namespaces.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/user-namespaces/
    title: Use a User Namespace With a Pod | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/configure-pod-container/user-namespaces.md
min_kubernetes_server_version: v1.25
---

# Use a User Namespace With a Pod

**Feature state:** Kubernetes v1.36 [stable] (feature gate `UserNamespacesSupport`, enabled by default)

This page shows how to configure a user namespace for pods. This allows you to
isolate the user running inside the container from the one in the host.

A process running as root in a container can run as a different (non-root) user
in the host; in other words, the process has full privileges for operations
inside the user namespace, but is unprivileged for operations outside the
namespace.

You can use this feature to reduce the damage a compromised container can do to
the host or other pods in the same node. There are [several security
vulnerabilities][KEP-vulns] rated either **HIGH** or **CRITICAL** that were not
exploitable when user namespaces is active. It is expected user namespace will
mitigate some future vulnerabilities too.

Without using a user namespace a container running as root, in the case of a
container breakout, has root privileges on the node. And if some capability were
granted to the container, the capabilities are valid on the host too. None of
this is true when user namespaces are used.

[KEP-vulns]: https://github.com/kubernetes/enhancements/tree/217d790720c5aef09b8bd4d6ca96284a0affe6c2/keps/sig-node/127-user-namespaces#motivation

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) Your Kubernetes server must be at or later than version v1.25. To check the version, enter `kubectl version`.

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

* The node OS needs to be Linux
* You need to exec commands in the host
* You need to be able to exec into pods

The cluster that you're using **must** include at least one node that meets the
[requirements](/concepts/workloads/pods/user-namespaces.md#before-you-begin)
for using user namespaces with Pods.

If you have a mixture of nodes and only some of the nodes provide user namespace support for
Pods, you also need to ensure that the user namespace Pods are
[scheduled](/concepts/scheduling-eviction/assign-pod-node.md) to suitable nodes.

## Run a Pod that uses a user namespace {#create-pod}

A user namespace for a pod is enabled setting the `hostUsers` field of `.spec`
to `false`. For example:

[`pods/user-namespaces-stateless.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/user-namespaces-stateless.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: userns
spec:
  hostUsers: false
  containers:
  - name: shell
    command: ["sleep", "infinity"]
    image: debian
```

1. Create the pod on your cluster:

   ```shell
   kubectl apply -f https://k8s.io/examples/pods/user-namespaces-stateless.yaml
   ```

1. Exec into the pod and run `readlink /proc/self/ns/user`:

   ```shell
   kubectl exec -ti userns -- bash
   ```

Run this command:

```shell
readlink /proc/self/ns/user
```

The output is similar to:

```shell
user:[4026531837]
```

Also run:

```shell
cat /proc/self/uid_map
```

The output is similar to:
```shell
0  833617920      65536
```

Then, open a shell in the host and run the same commands.

The `readlink` command shows the user namespace the process is running in. It
should be different when it is run on the host and inside the container.

The last number of the `uid_map` file inside the container must be 65536, on the
host it must be a bigger number.

If you are running the kubelet inside a user namespace, you need to compare the
output from running the command in the pod to the output of running in the host:

```shell
readlink /proc/$pid/ns/user
```

replacing `$pid` with the kubelet PID.
