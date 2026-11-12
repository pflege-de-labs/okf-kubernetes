---
type: task
title: Use an HTTP Proxy to Access the Kubernetes API
description: This page shows how to use an HTTP proxy to access the Kubernetes API.
resource: https://kubernetes.io/docs/tasks/extend-kubernetes/http-proxy-access-api/
tags: [kubernetes, tasks, extend-kubernetes]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/extend-kubernetes/http-proxy-access-api.md
    title: 'Kubernetes website source: tasks/extend-kubernetes/http-proxy-access-api.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/extend-kubernetes/http-proxy-access-api/
    title: Use an HTTP Proxy to Access the Kubernetes API | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/extend-kubernetes/http-proxy-access-api.md
---

# Use an HTTP Proxy to Access the Kubernetes API

This page shows how to use an HTTP proxy to access the Kubernetes API.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

If you do not already have an application running in your cluster, start
a Hello world application by entering this command:

```shell
kubectl create deployment hello-app --image=gcr.io/google-samples/hello-app:2.0 --port=8080
```

## Using kubectl to start a proxy server

This command starts a proxy to the Kubernetes API server:

    kubectl proxy --port=8080

## Exploring the Kubernetes API

When the proxy server is running, you can explore the API using `curl`, `wget`,
or a browser.

Get the API versions:

    curl http://localhost:8080/api/

The output should look similar to this:

    {
      "kind": "APIVersions",
      "versions": [
        "v1"
      ],
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "10.0.2.15:8443"
        }
      ]
    }

Get a list of pods:

    curl http://localhost:8080/api/v1/namespaces/default/pods

The output should look similar to this:

    {
      "kind": "PodList",
      "apiVersion": "v1",
      "metadata": {
        "resourceVersion": "33074"
      },
      "items": [
        {
          "metadata": {
            "name": "kubernetes-bootcamp-2321272333-ix8pt",
            "generateName": "kubernetes-bootcamp-2321272333-",
            "namespace": "default",
            "uid": "ba21457c-6b1d-11e6-85f7-1ef9f1dab92b",
            "resourceVersion": "33003",
            "creationTimestamp": "2016-08-25T23:43:30Z",
            "labels": {
              "pod-template-hash": "2321272333",
              "run": "kubernetes-bootcamp"
            },
            ...
    }

## What's next

Learn more about [kubectl proxy](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#proxy).
