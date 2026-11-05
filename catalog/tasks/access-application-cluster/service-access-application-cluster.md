---
type: tutorial
title: Use a Service to Access an Application in a Cluster
description: This page shows how to create a Kubernetes Service object that external clients can use to access an application running in a cluster.
resource: https://kubernetes.io/docs/tasks/access-application-cluster/service-access-application-cluster/
tags: [kubernetes, tasks, access-application-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/access-application-cluster/service-access-application-cluster.md
    title: 'Kubernetes website source: tasks/access-application-cluster/service-access-application-cluster.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/access-application-cluster/service-access-application-cluster/
    title: Use a Service to Access an Application in a Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/access-application-cluster/service-access-application-cluster.md
---

# Use a Service to Access an Application in a Cluster

This page shows how to create a Kubernetes Service object that external
clients can use to access an application running in a cluster. The Service
provides load balancing for an application that has two running instances.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

## Objectives

- Run two instances of a Hello World application.
- Create a Service object that exposes a node port.
- Use the Service object to access the running application.

## Creating a service for an application running in two pods

Here is the configuration file for the application Deployment:

[`service/access/hello-application.yaml`](https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/examples/service/access/hello-application.yaml)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  selector:
    matchLabels:
      run: load-balancer-example
  replicas: 2
  template:
    metadata:
      labels:
        run: load-balancer-example
    spec:
      containers:
        - name: hello-world
          image: us-docker.pkg.dev/google-samples/containers/gke/hello-app:2.0
          ports:
            - containerPort: 8080
              protocol: TCP
```

1. Run a Hello World application in your cluster:
   Create the application Deployment using the file above:

   ```shell
   kubectl apply -f https://k8s.io/examples/service/access/hello-application.yaml
   ```

   The preceding command creates a
   [Deployment](/reference/glossary/deployment.md)
   and an associated
   [ReplicaSet](/reference/glossary/replica-set.md).
   The ReplicaSet has two
   [Pods](/reference/glossary/pod.md)
   each of which runs the Hello World application.

1. Display information about the Deployment:

   ```shell
   kubectl get deployments hello-world
   kubectl describe deployments hello-world
   ```

1. Display information about your ReplicaSet objects:

   ```shell
   kubectl get replicasets
   kubectl describe replicasets
   ```

1. Create a Service object that exposes the deployment:

   ```shell
   kubectl expose deployment hello-world --type=NodePort --name=example-service
   ```

1. Display information about the Service:

   ```shell
   kubectl describe services example-service
   ```

   The output is similar to this:

   ```none
   Name:                   example-service
   Namespace:              default
   Labels:                 run=load-balancer-example
   Annotations:            <none>
   Selector:               run=load-balancer-example
   Type:                   NodePort
   IP:                     10.32.0.16
   Port:                   <unset> 8080/TCP
   TargetPort:             8080/TCP
   NodePort:               <unset> 31496/TCP
   Endpoints:              10.200.1.4:8080,10.200.2.5:8080
   Session Affinity:       None
   Events:                 <none>
   ```

   Make a note of the NodePort value for the Service. For example,
   in the preceding output, the NodePort value is 31496.

1. List the pods that are running the Hello World application:

   ```shell
   kubectl get pods --selector="run=load-balancer-example" --output=wide
   ```

   The output is similar to this:

   ```none
   NAME                           READY   STATUS    ...  IP           NODE
   hello-world-2895499144-bsbk5   1/1     Running   ...  10.200.1.4   worker1
   hello-world-2895499144-m1pwt   1/1     Running   ...  10.200.2.5   worker2
   ```

1. Get the public IP address of one of your nodes that is running
   a Hello World pod. How you get this address depends on how you set
   up your cluster. For example, if you are using Minikube, you can
   see the node address by running `kubectl cluster-info`. If you are
   using Google Compute Engine instances, you can use the
   `gcloud compute instances list` command to see the public addresses of your
   nodes.

1. On your chosen node, create a firewall rule that allows TCP traffic
   on your node port. For example, if your Service has a NodePort value of
   31568, create a firewall rule that allows TCP traffic on port 31568. Different
   cloud providers offer different ways of configuring firewall rules.

1. Use the node address and node port to access the Hello World application:

   ```shell
   curl http://<public-node-ip>:<node-port>
   ```

   where `<public-node-ip>` is the public IP address of your node,
   and `<node-port>` is the NodePort value for your service. The
   response to a successful request is a hello message:

   ```none
   Hello, world!
   Version: 2.0.0
   Hostname: hello-world-cdd4458f4-m47c8
   ```

## Using a service configuration file

As an alternative to using `kubectl expose`, you can use a
[service configuration file](/concepts/services-networking/service.md)
to create a Service.

## Cleaning up

To delete the Service, enter this command:

    kubectl delete services example-service

To delete the Deployment, the ReplicaSet, and the Pods that are running
the Hello World application, enter this command:

    kubectl delete deployment hello-world

## What's next

Follow the
[Connecting Applications with Services](https://kubernetes.io/docs/tutorials/services/connect-applications-service/)
tutorial.
