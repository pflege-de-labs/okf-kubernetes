---
type: task
title: Communicate Between Containers in the Same Pod Using a Shared Volume
description: This page shows how to use a Volume to communicate between two Containers running in the same Pod.
resource: https://kubernetes.io/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume/
tags: [kubernetes, tasks, access-application-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume.md
    title: 'Kubernetes website source: tasks/access-application-cluster/communicate-containers-same-pod-shared-volume.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume/
    title: Communicate Between Containers in the Same Pod Using a Shared Volume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume.md
---

# Communicate Between Containers in the Same Pod Using a Shared Volume

This page shows how to use a Volume to communicate between two Containers running
in the same Pod. See also how to allow processes to communicate by
[sharing process namespace](/tasks/configure-pod-container/share-process-namespace.md)
between containers.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

##  Creating a Pod that runs two Containers

In this exercise, you create a Pod that runs two Containers. The two containers
share a Volume that they can use to communicate. Here is the configuration file
for the Pod:

[`pods/two-container-pod.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/two-container-pod.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: two-containers
spec:

  restartPolicy: Never

  volumes:
  - name: shared-data
    emptyDir: {}

  containers:

  - name: nginx-container
    image: nginx
    volumeMounts:
    - name: shared-data
      mountPath: /usr/share/nginx/html

  - name: debian-container
    image: debian
    volumeMounts:
    - name: shared-data
      mountPath: /pod-data
    command: ["/bin/sh"]
    args: ["-c", "echo Hello from the debian container > /pod-data/index.html"]
```

In the configuration file, you can see that the Pod has a Volume named
`shared-data`.

The first container listed in the configuration file runs an nginx server. The
mount path for the shared Volume is `/usr/share/nginx/html`.
The second container is based on the debian image, and has a mount path of
`/pod-data`. The second container runs the following command and then terminates.

    echo Hello from the debian container > /pod-data/index.html

Notice that the second container writes the `index.html` file in the root
directory of the nginx server.

Create the Pod and the two Containers:

    kubectl apply -f https://k8s.io/examples/pods/two-container-pod.yaml

View information about the Pod and the Containers:

    kubectl get pod two-containers --output=yaml

Here is a portion of the output:

    apiVersion: v1
    kind: Pod
    metadata:
      ...
      name: two-containers
      namespace: default
      ...
    spec:
      ...
      containerStatuses:

      - containerID: docker://c1d8abd1 ...
        image: debian
        ...
        lastState:
          terminated:
            ...
        name: debian-container
        ...

      - containerID: docker://96c1ff2c5bb ...
        image: nginx
        ...
        name: nginx-container
        ...
        state:
          running:
        ...

You can see that the debian Container has terminated, and the nginx Container
is still running.

Get a shell to nginx Container:

    kubectl exec -it two-containers -c nginx-container -- /bin/bash

In your shell, verify that nginx is running:

    root@two-containers:/# apt-get update
    root@two-containers:/# apt-get install curl procps
    root@two-containers:/# ps aux

The output is similar to this:

    USER       PID  ...  STAT START   TIME COMMAND
    root         1  ...  Ss   21:12   0:00 nginx: master process nginx -g daemon off;

Recall that the debian Container created the `index.html` file in the nginx root
directory. Use `curl` to send a GET request to the nginx server:

```
root@two-containers:/# curl localhost
```

The output shows that nginx serves a web page written by the debian container:

```
Hello from the debian container
```

## Discussion

The primary reason that Pods can have multiple containers is to support
helper applications that assist a primary application. Typical examples of
helper applications are data pullers, data pushers, and proxies.
Helper and primary applications often need to communicate with each other.
Typically this is done through a shared filesystem, as shown in this exercise,
or through the loopback network interface, localhost. An example of this pattern is a
web server along with a helper program that polls a Git repository for new updates.

The Volume in this exercise provides a way for Containers to communicate during
the life of the Pod. If the Pod is deleted and recreated, any data stored in
the shared Volume is lost.

## What's next

* Learn more about [patterns for composite containers](https://kubernetes.io/blog/2015/06/the-distributed-system-toolkit-patterns/).

* Learn about [composite containers for modular architecture](https://www.slideshare.net/Docker/slideshare-burns).

* See [Configuring a Pod to Use a Volume for Storage](/tasks/configure-pod-container/configure-volume-storage.md).

* See [Configure a Pod to share process namespace between containers in a Pod](/tasks/configure-pod-container/share-process-namespace.md)

* See [Volume](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#volume-v1-core).

* See [Pod](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#pod-v1-core).
