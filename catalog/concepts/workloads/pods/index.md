# Pods

* [Pod Lifecycle](pod-lifecycle.md) - This page describes the lifecycle of a Pod.
* [Pod Conditions](pod-condition.md) - In Kubernetes, many objects have _conditions_.
* [Init Containers](init-containers.md) - This page provides an overview of init containers: specialized containers that run before app containers in a Pod.
* [Sidecar Containers](sidecar-containers.md) - Sidecar containers are the secondary containers that run along with the main application container within the same Pod.
* [Ephemeral Containers](ephemeral-containers.md) - This page provides an overview of ephemeral containers: a special type of container that runs temporarily in an existing Pod to accomplish user-initiated actions such as troubleshooting.
* [Liveness, Readiness, and Startup Probes](probes.md) - Kubernetes lets you define _probes_ to continuously monitor the health of containers in a Pod.
* [Disruptions](disruptions.md) - This guide is for application owners who want to build highly available applications, and thus need to understand what types of disruptions can happen to Pods.
* [Pod Hostname](pod-hostname.md) - This page explains how to set a Pod's hostname, potential side effects after configuration, and the underlying mechanics.
* [Pod Quality of Service Classes](pod-qos.md) - This page introduces _Quality of Service (QoS) classes_ in Kubernetes, and explains how Kubernetes assigns a QoS class to each Pod as a consequence of the resource constraints that you specify for the containers in that Pod.
* [Scheduling Group](scheduling-group.md) - You can link a `Pod` to a PodGroup to indicate that the `Pod` belongs to a group of `Pods` scheduled together.
* [Static Pods](static-pods.md) - _Static Pods_ are managed directly by the kubelet daemon on a specific node, without the API server observing them.
* [User Namespaces](user-namespaces.md) - This page explains how user namespaces are used in Kubernetes pods.
* [Downward API](downward-api.md) - There are two ways to expose Pod and container fields to a running container: environment variables, and as files that are populated by a special volume type. Together, these two ways of exposing Pod and container fields are called the downward API.
* [Advanced Pod Configuration](advanced-pod-config.md) - This page covers advanced Pod configuration topics including PriorityClasses, RuntimeClasses, security context within Pods, and introduces aspects of scheduling.
