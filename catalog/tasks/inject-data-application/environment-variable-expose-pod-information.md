---
type: task
title: Expose Pod Information to Containers Through Environment Variables
description: This page shows how a Pod can use environment variables to expose information about itself to containers running in the Pod, using the _downward API_.
resource: https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
tags: [kubernetes, tasks, inject-data-application]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/inject-data-application/environment-variable-expose-pod-information.md
    title: 'Kubernetes website source: tasks/inject-data-application/environment-variable-expose-pod-information.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
    title: Expose Pod Information to Containers Through Environment Variables | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/inject-data-application/environment-variable-expose-pod-information.md
---

# Expose Pod Information to Containers Through Environment Variables

This page shows how a Pod can use environment variables to expose information
about itself to containers running in the Pod, using the _downward API_.
You can use environment variables to expose Pod fields, container fields, or both.

In Kubernetes, there are two ways to expose Pod and container fields to a running container:

* _Environment variables_, as explained in this task
* [Volume files](/tasks/inject-data-application/downward-api-volume-expose-pod-information.md)

Together, these two ways of exposing Pod and container fields are called the
downward API.

As Services are the primary mode of communication between containerized applications managed by Kubernetes, 
it is helpful to be able to discover them at runtime. 

Read more about accessing Services [here](https://kubernetes.io/docs/tutorials/services/connect-applications-service/#accessing-the-service).

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds)

## Use Pod fields as values for environment variables

In this part of exercise, you create a Pod that has one container, and you
project Pod-level fields into the running container as environment variables.

[`pods/inject/dapi-envars-pod.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/inject/dapi-envars-pod.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: dapi-envars-fieldref
spec:
  containers:
    - name: test-container
      image: registry.k8s.io/busybox:1.27.2
      command: [ "sh", "-c"]
      args:
      - while true; do
          echo -en '\n';
          printenv MY_NODE_NAME MY_POD_NAME MY_POD_NAMESPACE;
          printenv MY_POD_IP MY_POD_SERVICE_ACCOUNT;
          sleep 10;
        done;
      env:
        - name: MY_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: MY_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: MY_POD_SERVICE_ACCOUNT
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
  restartPolicy: Never
```

In that manifest, you can see five environment variables. The `env`
field is an array of
environment variable definitions.
The first element in the array specifies that the `MY_NODE_NAME` environment
variable gets its value from the Pod's `spec.nodeName` field. Similarly, the
other environment variables get their names from Pod fields.

**Note:**
The fields in this example are Pod fields. They are not fields of the
container in the Pod.

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/inject/dapi-envars-pod.yaml
```

Verify that the container in the Pod is running:

```shell
# If the new Pod isn't yet healthy, rerun this command a few times.
kubectl get pods
```

View the container's logs:

```shell
kubectl logs dapi-envars-fieldref
```

The output shows the values of selected environment variables:

```
minikube
dapi-envars-fieldref
default
172.17.0.4
default
```

To see why these values are in the log, look at the `command` and `args` fields
in the configuration file. When the container starts, it writes the values of
five environment variables to stdout. It repeats this every ten seconds.

Next, get a shell into the container that is running in your Pod:

```shell
kubectl exec -it dapi-envars-fieldref -- sh
```

In your shell, view the environment variables:

```shell
# Run this in a shell inside the container
printenv
```

The output shows that certain environment variables have been assigned the
values of Pod fields:

```
MY_POD_SERVICE_ACCOUNT=default
...
MY_POD_NAMESPACE=default
MY_POD_IP=172.17.0.4
...
MY_NODE_NAME=minikube
...
MY_POD_NAME=dapi-envars-fieldref
```

## Use container fields as values for environment variables

In the preceding exercise, you used information from Pod-level fields as the values
for environment variables.
In this next exercise, you are going to pass fields that are part of the Pod
definition, but taken from the specific
[container](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#Container)
rather than from the Pod overall.

Here is a manifest for another Pod that again has just one container:

[`pods/inject/dapi-envars-container.yaml`](https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/examples/pods/inject/dapi-envars-container.yaml)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: dapi-envars-resourcefieldref
spec:
  containers:
    - name: test-container
      image: registry.k8s.io/busybox:1.27.2
      command: [ "sh", "-c"]
      args:
      - while true; do
          echo -en '\n';
          printenv MY_CPU_REQUEST MY_CPU_LIMIT;
          printenv MY_MEM_REQUEST MY_MEM_LIMIT;
          sleep 10;
        done;
      resources:
        requests:
          memory: "32Mi"
          cpu: "125m"
        limits:
          memory: "64Mi"
          cpu: "250m"
      env:
        - name: MY_CPU_REQUEST
          valueFrom:
            resourceFieldRef:
              containerName: test-container
              resource: requests.cpu
        - name: MY_CPU_LIMIT
          valueFrom:
            resourceFieldRef:
              containerName: test-container
              resource: limits.cpu
        - name: MY_MEM_REQUEST
          valueFrom:
            resourceFieldRef:
              containerName: test-container
              resource: requests.memory
        - name: MY_MEM_LIMIT
          valueFrom:
            resourceFieldRef:
              containerName: test-container
              resource: limits.memory
  restartPolicy: Never
```

In this manifest, you can see four environment variables. The `env`
field is an array of
environment variable definitions.
The first element in the array specifies that the `MY_CPU_REQUEST` environment
variable gets its value from the `requests.cpu` field of a container named
`test-container`. Similarly, the other environment variables get their values
from fields that are specific to this container.

Create the Pod:

```shell
kubectl apply -f https://k8s.io/examples/pods/inject/dapi-envars-container.yaml
```

Verify that the container in the Pod is running:

```shell
# If the new Pod isn't yet healthy, rerun this command a few times.
kubectl get pods
```

View the container's logs:

```shell
kubectl logs dapi-envars-resourcefieldref
```

The output shows the values of selected environment variables:

```
1
1
33554432
67108864
```

## What's next

* Read [Defining Environment Variables for a Container](/tasks/inject-data-application/define-environment-variable-container.md)
* Read the [`spec`](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#PodSpec)
  API definition for Pod. This includes the definition of Container (part of Pod).
* Read the list of [available fields](/concepts/workloads/pods/downward-api.md#available-fields) that you
  can expose using the downward API.

Read about Pods, containers and environment variables in the legacy API reference:

* [PodSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#podspec-v1-core)
* [Container](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#container-v1-core)
* [EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#envvar-v1-core)
* [EnvVarSource](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#envvarsource-v1-core)
* [ObjectFieldSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#objectfieldselector-v1-core)
* [ResourceFieldSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.37/#resourcefieldselector-v1-core)
