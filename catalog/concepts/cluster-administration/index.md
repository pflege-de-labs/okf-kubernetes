# Cluster Administration

* [Node Shutdowns](node-shutdown.md) - In a Kubernetes cluster, a node can be shut down in a planned graceful way or unexpectedly because of reasons such as a power outage or something else external.
* [Swap memory management](swap-memory-management.md) - Kubernetes can be configured to use swap memory on a node, allowing the kernel to free up physical memory by swapping out pages to backing storage.
* [Node Autoscaling](node-autoscaling.md) - Automatically provision and consolidate the Nodes in your cluster to adapt to demand and optimize cost.
* [Certificates](certificates.md) - To learn how to generate certificates for your cluster, see Certificates.
* [Cluster Networking](networking.md) - Networking is a central part of Kubernetes, but it can be challenging to understand exactly how it is expected to work.
* [Observability](observability.md) - Understand how to gain end-to-end visibility of a Kubernetes cluster through the collection of metrics, logs, and traces.
* [Admission Webhook Good Practices](admission-webhooks-good-practices.md) - Recommendations for designing and deploying admission webhooks in Kubernetes.
* [Good practices for Dynamic Resource Allocation as a Cluster Admin](dra.md) - This page describes good practices when configuring a Kubernetes cluster utilizing Dynamic Resource Allocation (DRA).
* [Logging Architecture](logging.md) - Application logs can help you understand what is happening inside your application.
* [Compatibility Version For Kubernetes Control Plane Components](compatibility-version.md) - Since release v1.32, we introduced configurable version compatibility and emulation options to Kubernetes control plane components to make upgrades safer by providing more control and increasing the granularity of steps available to cluster administrators.
* [Metrics For Kubernetes System Components](system-metrics.md) - System component metrics can give a better look into what is happening inside them.
* [Metrics for Kubernetes Object States](kube-state-metrics.md) - kube-state-metrics, an add-on agent to generate and expose cluster-level metrics.
* [System Logs](system-logs.md) - System component logs record events happening in cluster, which can be very useful for debugging.
* [Traces For Kubernetes System Components](system-traces.md) - System component traces record the latency of and relationships between operations in the cluster.
* [Proxies in Kubernetes](proxies.md) - This page explains proxies used with Kubernetes.
* [API Priority and Fairness](flow-control.md) - Controlling the behavior of the Kubernetes API server in an overload situation is a key task for cluster administrators.
* [Installing Addons](addons.md) - Add-ons extend the functionality of Kubernetes.
* [Coordinated Leader Election](coordinated-leader-election.md) - Kubernetes 1.37 includes a beta feature that allows control plane components to deterministically select a leader via _coordinated leader election_.
