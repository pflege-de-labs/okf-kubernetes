# Configure Pods and Containers

* [Assign Memory Resources to Containers and Pods](assign-memory-resource.md) - This page shows how to assign a memory *request* and a memory *limit* to a Container.
* [Assign CPU Resources to Containers and Pods](assign-cpu-resource.md) - This page shows how to assign a CPU *request* and a CPU *limit* to a container.
* [Assign Devices to Pods and Containers](assign-resources.md) - Assign infrastructure resources to your Kubernetes workloads. ([contents](assign-resources/index.md))
* [Assign Pod-level CPU and memory resources](assign-pod-level-resources.md) - This page shows how to specify CPU and memory resources for a Pod at pod-level in addition to container-level resource specifications.
* [Configure GMSA for Windows Pods and containers](configure-gmsa.md) - This page shows how to configure Group Managed Service Accounts (GMSA) for Pods and containers that will run on Windows nodes.
* [Resize CPU and Memory Resources assigned to Containers](resize-container-resources.md) - This page explains how to change the CPU and memory resource requests and limits assigned to a container *without recreating the Pod*.
* [Resize CPU and Memory Resources assigned to Pods](resize-pod-resources.md) - This page explains how to change the CPU and memory resources set at the Pod level without recreating the Pod.
* [Configure RunAsUserName for Windows pods and containers](configure-runasusername.md) - This page shows how to use the `runAsUserName` setting for Pods and containers that will run on Windows nodes.
* [Create a Windows HostProcess Pod](create-hostprocess-pod.md) - Windows HostProcess containers enable you to run containerized workloads on a Windows host.
* [Configure Quality of Service for Pods](quality-service-pod.md) - This page shows how to configure Pods so that they will be assigned particular Quality of Service (QoS) classes.
* [Assign Extended Resources to a Container](extended-resource.md) - This page shows how to assign extended resources to a Container.
* [Configure a Pod to Use a Volume for Storage](configure-volume-storage.md) - This page shows how to configure a Pod to use a Volume for storage.
* [Configure a Pod to Use a Projected Volume for Storage](configure-projected-volume-storage.md) - This page shows how to use a `projected` Volume to mount several existing volume sources into the same directory.
* [Configure a Security Context for a Pod or Container](security-context.md) - A security context defines privilege and access control settings for a Pod or Container.
* [Configure Service Accounts for Pods](configure-service-account.md) - Kubernetes offers two distinct ways for clients that run within your cluster, or that otherwise have a relationship to your cluster's control plane to authenticate to the API server.
* [Pull an Image from a Private Registry](pull-image-private-registry.md) - This page shows how to create a Pod that uses a Secret to pull an image from a private container image registry or repository.
* [Configure Liveness, Readiness and Startup Probes](configure-liveness-readiness-startup-probes.md) - This page shows how to configure liveness, readiness and startup probes for containers.
* [Assign Pods to Nodes](assign-pods-nodes.md) - This page shows how to assign a Kubernetes Pod to a particular node in a Kubernetes cluster.
* [Assign Pods to Nodes using Node Affinity](assign-pods-nodes-using-node-affinity.md) - This page shows how to assign a Kubernetes Pod to a particular node using Node Affinity in a Kubernetes cluster.
* [Configure Pod Initialization](configure-pod-initialization.md) - This page shows how to use an Init Container to initialize a Pod before an application Container runs.
* [Attach Handlers to Container Lifecycle Events](attach-handler-lifecycle-event.md) - This page shows how to attach handlers to Container lifecycle events.
* [Configure a Pod to Use a ConfigMap](configure-pod-configmap.md) - Many applications rely on configuration which is used during either application initialization or runtime.
* [Share Process Namespace between Containers in a Pod](share-process-namespace.md) - This page shows how to configure process namespace sharing for a pod.
* [Use a User Namespace With a Pod](user-namespaces.md) - This page shows how to configure a user namespace for pods.
* [Use an Image Volume With a Pod](image-volumes.md) - This page shows how to configure a pod using image volumes.
* [Set Permissions on an emptyDir Volume](configure-emptydir-volume-mode.md) - This page shows how to set Unix permission bits on an `emptyDir` volume directory using the `mode` field.
* [Set Bind Mount Options on Volume Mounts](configure-bind-mount-options.md) - This page shows how to apply security-related bind mount options (`noexec`, `nodev`, `nosuid`) to volume mounts in a Pod.
* [Create static Pods](static-pod.md) - This page shows you how to create _static Pods_ on a node.
* [Translate a Docker Compose File to Kubernetes Resources](translate-compose-kubernetes.md) - What's Kompose?
* [Enforce Pod Security Standards by Configuring the Built-in Admission Controller](enforce-standards-admission-controller.md) - Kubernetes provides a built-in admission controller to enforce the Pod Security Standards.
* [Enforce Pod Security Standards with Namespace Labels](enforce-standards-namespace-labels.md) - Namespaces can be labeled to enforce the Pod Security Standards.
* [Migrate from PodSecurityPolicy to the Built-In PodSecurity Admission Controller](migrate-from-psp.md) - This page describes the process of migrating from PodSecurityPolicies to the built-in PodSecurity admission controller.
