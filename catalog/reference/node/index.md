# Node Reference Information

* [node-lifecycle-conditions](node-lifecycle-conditions.md) - Node lifecycle conditions report whether a Node is undergoing a lifecycle event such as drain, maintenance, or Graceful Node Shutdown.
* [Kubelet Checkpoint API](kubelet-checkpoint-api.md) - Checkpointing a container is the functionality to create a stateful copy of a running container.
* [Linux Kernel Version Requirements](kernel-version-requirements.md) - Many features rely on specific kernel functionalities and have minimum kernel version requirements.
* [Standard Device Attributes for Dynamic Resource Allocation](dra-standard-device-attributes.md) - Kubernetes defines a set of standard device attributes under the `resource.kubernetes.io` domain for use with Dynamic Resource Allocation (DRA).
* [Articles on dockershim Removal and on Using CRI-compatible Runtimes](topics-on-dockershim-and-cri-compatible-runtimes.md) - This is a list of articles and other pages that are either about the Kubernetes' deprecation and removal of _dockershim_, or about using CRI-compatible container runtimes, in connection with that removal.
* [Kubelet Pods API](kubelet-pods-api.md) - The Kubelet Pods API provides a way for Node-local components to query information about Pods running on the Node directly from the `kubelet`.
* [Pod-level resource managers reference](pod-level-resource-managers.md) - This document provides reference details for the pod-level resource managers implementation in the `kubelet`, including PodResources API reporting structures and state checkpoint format migrations during version upgrades and downgrades.
* [Node Labels Populated By The Kubelet](node-labels.md) - Kubernetes nodes come pre-populated with a standard set of labels.
* [Kubelet Sync Loop](kubelet-sync-loop.md) - The kubelet is the primary "node agent" that creates and watches Pods on each node.
* [Local Files And Paths Used By The Kubelet](kubelet-files.md) - The kubelet is mostly a stateless process running on a Kubernetes node.
* [Kubelet Configuration Directory Merging](kubelet-config-directory-merging.md) - When using the kubelet's `--config-dir` flag to specify a drop-in directory for configuration, there is some specific behavior on how different types are merged.
* [Kubelet Device Manager API Versions](device-plugin-api-versions.md) - This page provides details of version compatibility between the Kubernetes device plugin API, and different versions of Kubernetes itself.
* [Kubelet Systemd Watchdog](systemd-watchdog.md) - On Linux nodes, Kubernetes 1.37 supports integrating with systemd to allow the operating system supervisor to recover a failed kubelet.
* [Node Status](node-status.md) - The status of a node in Kubernetes is a critical aspect of managing a Kubernetes cluster.
* [Seccomp and Kubernetes](seccomp.md) - Seccomp stands for secure computing mode and has been a feature of the Linux kernel since version 2.6.12.
* [What Happens After A Node Restart](what-happens-on-restart.md) - System components on a node sometimes restart, either because of an upgrade, a crash, or an explicit operator action.
* [Linux Node Swap Behaviors](swap-behavior.md) - To allow Kubernetes workloads to use swap, on a Linux node, you must disable the kubelet's default behavior of failing when swap is detected, and specify memory-swap behavior as `LimitedSwap`:
