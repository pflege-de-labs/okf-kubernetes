# Glossary

* [API Group](api-group.md) - A set of related paths in the Kubernetes API.
* [API resource](api-resource.md) - A Kubernetes entity, representing an endpoint on the Kubernetes API server.
* [API server](kube-apiserver.md) - Control plane component that serves the Kubernetes API.
* [API-initiated eviction](api-eviction.md) - API-initiated eviction is the process by which you use the Eviction API to create an Eviction object that triggers graceful pod termination.
* [Add-ons](addons.md) - Resources that extend the functionality of Kubernetes.
* [Admission Controller](admission-controller.md) - A piece of code that intercepts requests to the Kubernetes API server prior to persistence of the object.
* [Affinity](affinity.md) - Rules used by the scheduler to determine where to place pods
* [Aggregation Layer](aggregation-layer.md) - The aggregation layer lets you install additional Kubernetes-style APIs in your cluster.
* [Annotation](annotation.md) - A key-value pair that is used to attach arbitrary non-identifying metadata to objects.
* [App Container](app-container.md) - A container used to run part of a workload. Compare with init container.
* [Application Architect](application-architect.md) - A person responsible for the high-level design of an application.
* [Application Developer](application-developer.md) - A person who writes an application that runs in a Kubernetes cluster.
* [Applications](applications.md) - The layer where various containerized applications run.
* [Approver](approver.md) - A person who can review and approve Kubernetes code contributions.
* [CIDR](cidr.md) - CIDR is a notation for describing blocks of IP addresses and is used heavily in various networking configurations.
* [CLA (Contributor License Agreement)](cla.md) - Terms under which a contributor grants a license to an open source project for their contributions.
* [CRI-O](cri-o.md) - A lightweight container runtime specifically for Kubernetes
* [Certificate](certificate.md) - A cryptographically secure file used to validate access to the Kubernetes cluster.
* [Cloud Controller Manager](cloud-controller-manager.md) - Control plane component that integrates Kubernetes with third-party cloud providers.
* [Cloud Native Computing Foundation (CNCF)](cncf.md) - Cloud Native Computing Foundation
* [Cloud Provider](cloud-provider.md) - An organization that offers a cloud computing platform.
* [Cluster](cluster.md) - A set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.
* [Cluster Architect](cluster-architect.md) - A person who designs infrastructure that involves one or more Kubernetes clusters.
* [Cluster Infrastructure](cluster-infrastructure.md) - The infrastructure layer provides and maintains VMs, networking, security groups and others.
* [Cluster Operations](cluster-operations.md) - The work involved in managing a Kubernetes cluster.
* [Cluster Operator](cluster-operator.md) - A person who configures, controls, and monitors clusters.
* [Code Contributor](code-contributor.md) - A person who develops and contributes code to the Kubernetes open source codebase.
* [Common Expression Language](cel.md) - An expression language that's designed to be safe for executing user code.
* [Condition](condition.md) - A condition represents the current state of a Kubernetes resource, providing information about whether certain aspects of the resource are true.
* [ConfigMap](configmap.md) - An API object used to store non-confidential data in key-value pairs. Can be consumed as environment variables, command-line arguments, or configuration files in a volume.
* [Container](container.md) - A lightweight and portable executable image that contains software and all of its dependencies.
* [Container Device Interface (CDI)](cdi.md) - A CNCF specification for describing device configuration that container runtimes apply when creating containers.
* [Container Environment Variables](container-env-variables.md) - Container environment variables are name=value pairs that provide useful information into containers running in a Pod.
* [Container Lifecycle Hooks](container-lifecycle-hooks.md) - The lifecycle hooks expose events in the container management lifecycle and let the user run code when the events occur.
* [Container Runtime](container-runtime.md) - The container runtime is the software that is responsible for running containers.
* [Container Runtime Interface (CRI)](cri.md) - Protocol for communication between the kubelet and the local container runtime.
* [Container Storage Interface (CSI)](csi.md) - The Container Storage Interface (CSI) defines a standard interface to expose storage systems to containers.
* [Container network interface (CNI)](cni.md) - Container network interface (CNI) plugins are a type of Network plugin that adheres to the appc/CNI specification.
* [Contributor](contributor.md) - Someone who donates code, documentation, or their time to help the Kubernetes project or community.
* [Control Plane](control-plane.md) - The container orchestration layer that exposes the API and interfaces to define, deploy, and manage the lifecycle of containers.
* [Controller](controller.md) - A control loop that watches the shared state of the cluster through the apiserver and makes changes attempting to move the current state towards the desired state.
* [CronJob](cronjob.md) - A repeating task (a Job) that runs on a regular schedule.
* [CustomResourceDefinition](customresourcedefinition.md) - Custom code that defines a resource to add to your Kubernetes API server without building a complete custom server.
* [DaemonSet](daemonset.md) - Ensures a copy of a Pod is running across a set of nodes in a cluster.
* [Data Plane](data-plane.md) - The layer that provides capacity such as CPU, memory, network, and storage so that the containers can run and connect to a network.
* [Deployment](deployment.md) - Manages a replicated application on your cluster.
* [Developer (disambiguation)](developer.md) - May refer to&#58; Application Developer, Code Contributor, or Platform Developer.
* [Device](device.md) - Any resource that's directly or indirectly attached your cluster's nodes, like GPUs or circuit boards.
* [Device Plugin](device-plugin.md) - Software extensions to let Pods access devices that need vendor-specific initialization or setup
* [DeviceClass](deviceclass.md) - A category of devices in the cluster. Users can claim specific devices in a DeviceClass.
* [Disruption](disruption.md) - An event that leads to Pod(s) going out of service
* [Docker](docker.md) - Docker is a software technology providing operating-system-level virtualization also known as containers.
* [Dockershim](dockershim.md) - A component of Kubernetes v1.23 and earlier, which allows Kubernetes system components to communicate with Docker Engine.
* [Downstream (disambiguation)](downstream.md) - May refer to: code in the Kubernetes ecosystem that depends upon the core Kubernetes codebase or a forked repo.
* [Downward API](downward-api.md) - A mechanism to expose Pod and container field values to code running in a container.
* [Drain](drain.md) - Safely evicts Pods from a Node to prepare for maintenance or removal.
* [Duration](duration.md) - A string value representing an amount of time.
* [Dynamic Resource Allocation](dra.md) - A Kubernetes feature for requesting and sharing resources, like hardware accelerators, among Pods.
* [Dynamic Volume Provisioning](dynamic-volume-provisioning.md) - Allows users to request automatic creation of storage Volumes.
* [EndpointSlice](endpoint-slice.md) - EndpointSlices track the IP addresses of Pods for Services.
* [Endpoints](endpoint.md) - (Deprecated) API representing endpoints of a Service
* [Ephemeral Container](ephemeral-container.md) - A type of container that you can temporarily run inside a Pod.
* [Event](event.md) - Kubernetes objects that describe some state change in the cluster.
* [Eviction](eviction.md) - Process of terminating one or more Pods on Nodes
* [Extensions](extensions.md) - Extensions are software components that extend and deeply integrate with Kubernetes to support new types of hardware.
* [Feature gate](feature-gates.md) - A way to control whether or not a particular Kubernetes feature is enabled.
* [Finalizer](finalizer.md) - A namespaced key that tells Kubernetes to wait until specific conditions are met before it fully deletes an object marked for deletion.
* [FlexVolume](flexvolume.md) - FlexVolume is a deprecated interface for creating out-of-tree volume plugins. The Container Storage Interface is a newer interface that addresses several problems with FlexVolume.
* [Garbage Collection](garbage-collection.md) - A collective term for the various mechanisms Kubernetes uses to clean up cluster resources.
* [Gateway API](gateway.md) - An API for modeling service networking in Kubernetes.
* [Group Version Resource](group-version-resource.md) - The API group, API version and name of a Kubernetes API.
* [Helm Chart](helm-chart.md) - A package of pre-configured Kubernetes configurations that can be managed with the Helm tool.
* [Horizontal Pod Autoscaler](horizontal-pod-autoscaler.md) - Object that automatically scales the number of pod replicas based on targeted resource utilization or custom metric targets.
* [HostAliases](host-aliases.md) - A HostAliases is a mapping between the IP address and hostname to be injected into a Pod's hosts file.
* [Image](image.md) - Stored instance of a container that holds a set of software needed to run an application.
* [Immutable Infrastructure](immutable-infrastructure.md) - Immutable Infrastructure refers to computer infrastructure (virtual machines, containers, network appliances) that cannot be changed once deployed
* [Ingress](ingress.md) - An API object that manages external access to the services in a cluster, typically HTTP.
* [Init Container](init-container.md) - One or more initialization containers that must run to completion before any app containers run.
* [Istio](istio.md) - An open platform (not Kubernetes-specific) that provides a uniform way to integrate microservices, manage traffic flow, enforce policies, and aggregate telemetry data.
* [JSON Web Token (JWT)](jwt.md) - A means of representing claims to be transferred between two parties.
* [Job](job.md) - A finite or batch task that runs to completion.
* [Kubeadm](kubeadm.md) - A tool for quickly installing Kubernetes and setting up a secure cluster.
* [Kubectl](kubectl.md) - A command line tool for communicating with a Kubernetes cluster.
* [Kubelet](kubelet.md) - An agent that runs on each node in the cluster. It makes sure that containers are running in a pod.
* [Kubernetes API](kubernetes-api.md) - The HTTP API that lets users, cluster components, and external components communicate with one another and query or manipulate the state of Kubernetes objects.
* [Label](label.md) - Tags objects with identifying attributes that are meaningful and relevant to users.
* [LimitRange](limitrange.md) - Provides constraints to limit resource consumption per Containers or Pods in a namespace.
* [Logging](logging.md) - Logs are the list of events that are logged by cluster or application.
* [Managed Service](managed-service.md) - A software offering maintained by a third-party provider.
* [Manifest](manifest.md) - A serialized specification of one or more Kubernetes API objects.
* [Master](master.md) - Legacy term, used as synonym for nodes running the control plane.
* [Member](member.md) - A continuously active contributor in the K8s community.
* [Minikube](minikube.md) - A tool for running Kubernetes locally.
* [Mirror Pod](mirror-pod.md) - An object in the API server that tracks a static pod on a kubelet.
* [Mixed Version Proxy (MVP)](mixed-version-proxy.md) - Feature that lets a kube-apiserver proxy a resource request to a different peer API server.
* [Name](name.md) - A client-provided string that refers to an object in a resource URL, such as `/api/v1/pods/some-name`.
* [Namespace](namespace.md) - An abstraction used by Kubernetes to support isolation of groups of resources within a single cluster.
* [Network Policy](network-policy.md) - A specification of how groups of Pods are allowed to communicate with each other and with other network endpoints.
* [Node](node.md) - A node is a worker machine in Kubernetes.
* [Node-pressure eviction](node-pressure-eviction.md) - Node-pressure eviction is the process by which the kubelet proactively fails pods to reclaim resources on nodes.
* [Object](object.md) - An entity in the Kubernetes system, representing part of the state of your cluster.
* [Operator pattern](operator-pattern.md) - A specialized controller used to manage a custom resource
* [Persistent Volume](persistent-volume.md) - API object that represents a piece of storage in the cluster.
* [Persistent Volume Claim](persistent-volume-claim.md) - Claims storage resources defined in a PersistentVolume so that it can be mounted as a volume in a container.
* [Platform Developer](platform-developer.md) - A person who customizes the Kubernetes platform to fit the needs of their project.
* [Pod](pod.md) - A Pod represents a set of running containers in your cluster.
* [Pod Disruption](pod-disruption.md) - The process by which Pods on Nodes are terminated either voluntarily or involuntarily.
* [Pod Disruption Budget](pod-disruption-budget.md) - An object that limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions.
* [Pod Lifecycle](pod-lifecycle.md) - The sequence of states through which a Pod passes during its lifetime.
* [Pod Priority](pod-priority.md) - Pod Priority indicates the importance of a Pod relative to other Pods.
* [Pod Security Policy](pod-security-policy.md) - Removed API that enforced Pod security restrictions.
* [PodGroup](podgroup.md) - A PodGroup represents a set of Pods with common scheduling policy and constraints.
* [PodTemplate](pod-template.md) - A template for creating Pods.
* [Preemption](preemption.md) - Preemption logic in Kubernetes helps a pending Pod to find a suitable Node by evicting low priority Pods existing on that Node.
* [PriorityClass](priority-class.md) - A mapping from a class name to the scheduling priority that a Pod should have.
* [Probe](probe.md) - A check performed periodically by the kubelet on a container in a Pod.
* [Proxy](proxy.md) - An application acting as an intermediary between clients and servers
* [QoS Class](qos-class.md) - QoS Class (Quality of Service Class) provides a way for Kubernetes to classify pods within the cluster into several classes and make decisions about scheduling and eviction.
* [Quantity](quantity.md) - A whole-number representation of small or large numbers using SI suffixes.
* [RBAC (Role-Based Access Control)](rbac.md) - Manages authorization decisions, allowing admins to dynamically configure access policies through the Kubernetes API.
* [Replica](replica.md) - Replicas are copies of pods, ensuring availability, scalability, and fault tolerance by maintaining identical instances.
* [ReplicaSet](replica-set.md) - ReplicaSet ensures that a specified number of Pod replicas are running at one time
* [ReplicationController](replication-controller.md) - A (deprecated) API object that manages a replicated application.
* [Resource (infrastructure)](infrastructure-resource.md) - A defined amount of infrastructure available for consumption (CPU, memory, etc).
* [ResourceClaim](resourceclaim.md) - Describes the resources that a workload needs, such as devices. ResourceClaims can request devices from DeviceClasses.
* [ResourceClaimTemplate](resourceclaimtemplate.md) - Defines a template for Kubernetes to create ResourceClaims. Used to provide per-Pod or per-PodGroup access to separate, similar resources.
* [ResourceQuota](resource-quota.md) - Provides constraints that limit aggregate resource consumption per namespace.
* [ResourceSlice](resourceslice.md) - Represents one or more infrastructure resources, like devices, in a pool of similar resources.
* [Reviewer](reviewer.md) - A person who reviews code for quality and correctness on some part of the project.
* [SIG (special interest group)](sig.md) - Community members who collectively manage an ongoing piece or aspect of the larger Kubernetes open source project.
* [Secret](secret.md) - Stores sensitive information, such as passwords, OAuth tokens, and ssh keys.
* [Security Context](security-context.md) - The securityContext field defines privilege and access control settings for a Pod or container.
* [Selector](selector.md) - Allows users to filter a list of resources based on labels.
* [Service](service.md) - A way to expose an application running on a set of Pods as a network service.
* [Service Catalog](service-catalog.md) - A former extension API that enabled applications running in Kubernetes clusters to easily use external managed software offerings, such as a datastore service offered by a cloud provider.
* [ServiceAccount](service-account.md) - Provides an identity for processes that run in a Pod.
* [Shuffle-sharding](shuffle-sharding.md) - A technique for assigning requests to queues that provides better isolation than hashing modulo the number of queues.
* [Sidecar Container](sidecar-container.md) - An auxiliary container that stays running throughout the lifecycle of a Pod.
* [Spec](spec.md) - Field in Kubernetes manifests that defines the desired state or configuration.
* [StatefulSet](statefulset.md) - A StatefulSet manages deployment and scaling of a set of Pods, with durable storage and persistent identifiers for each Pod.
* [Static Pod](static-pod.md) - A pod managed directly by the kubelet daemon on a specific node.
* [Storage Class](storage-class.md) - A StorageClass provides a way for administrators to describe different available storage types.
* [Subject Alternative Name](san.md) - An X.509 certificate extension to identify what hostname or IP address the certificate applies to.
* [Taint](taint.md) - A core object consisting of three required properties: key, value, and effect. Taints prevent the scheduling of pods on nodes or node groups.
* [Toleration](toleration.md) - A core object consisting of three required properties: key, value, and effect. Tolerations enable the scheduling of pods on nodes or node groups that have a matching taint.
* [UID](uid.md) - A Kubernetes systems-generated string to uniquely identify objects.
* [Upstream (disambiguation)](upstream.md) - May refer to: core Kubernetes or the source repo from which a repo was forked.
* [Volume](volume.md) - A directory containing data, accessible to the containers in a pod.
* [Volume Plugin](volume-plugin.md) - A Volume Plugin enables integration of storage within a Pod.
* [WG (working group)](wg.md) - Facilitates the discussion and/or implementation of a short-lived, narrow, or decoupled project for a committee, SIG, or cross-SIG effort.
* [Watch](watch.md) - A verb that is used to track changes to an object in Kubernetes as a stream.
* [Workload](workload.md) - A workload is an application running on Kubernetes.
* [cAdvisor](cadvisor.md) - Tool that provides understanding of the resource usage and performance characteristics for containers
* [cgroup (control group)](cgroup.md) - A group of Linux processes with optional resource isolation, accounting and limits.
* [containerd](containerd.md) - A container runtime with an emphasis on simplicity, robustness and portability
* [etcd](etcd.md) - Consistent and highly-available key value store used as backing store of Kubernetes for all cluster data.
* [kOps (Kubernetes Operations)](kops.md) - kOps will not only help you create, destroy, upgrade and maintain production-grade, highly available, Kubernetes cluster, but it will also provision the necessary cloud infrastructure.
* [kube-controller-manager](kube-controller-manager.md) - Control Plane component that runs controller processes.
* [kube-proxy](kube-proxy.md) - `kube-proxy` is a network proxy that runs on each node in the cluster.
* [kube-scheduler](kube-scheduler.md) - Control plane component that watches for newly created pods with no assigned node, and selects a node for them to run on.
* [sysctl](sysctl.md) - An interface for getting and setting Unix kernel parameters
* [user namespace](userns.md) - A Linux kernel feature to emulate superuser privilege for unprivileged users.
