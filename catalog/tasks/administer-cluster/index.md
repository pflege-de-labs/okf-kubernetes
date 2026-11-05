# Administer a Cluster

* [Administration with kubeadm](kubeadm.md) - If you don't yet have a cluster, visit bootstrapping clusters with `kubeadm`. ([contents](kubeadm/index.md))
* [Overprovision Node Capacity For A Cluster](node-overprovisioning.md) - This page guides you through configuring Node overprovisioning in your Kubernetes cluster.
* [Migrating from dockershim](migrating-from-dockershim.md) - This section presents information you need to know when migrating from dockershim to other container runtimes. ([contents](migrating-from-dockershim/index.md))
* [Generate Certificates Manually](certificates.md) - When using client certificate authentication, you can generate certificates manually through `easyrsa`, `openssl` or `cfssl`.
* [Manage Memory, CPU, and API Resources](manage-resources.md) - See the contents of this section. ([contents](manage-resources/index.md))
* [Install a Network Policy Provider](network-policy-provider.md) - See the contents of this section. ([contents](network-policy-provider/index.md))
* [Access Clusters Using the Kubernetes API](access-cluster-api.md) - This page shows how to access clusters using the Kubernetes API.
* [Enable Or Disable Feature Gates](configure-feature-gates.md) - This page shows how to enable or disable feature gates to control specific Kubernetes features in your cluster.
* [Advertise Extended Resources for a Node](extended-resource-node.md) - This page shows how to specify extended resources for a Node.
* [Autoscale the DNS Service in a Cluster](dns-horizontal-autoscaling.md) - This page shows how to enable and configure autoscaling of the DNS service in your Kubernetes cluster.
* [Change the Access Mode of a PersistentVolume to ReadWriteOncePod](change-pv-access-mode-readwriteoncepod.md) - This page shows how to change the access mode on an existing PersistentVolume to use `ReadWriteOncePod`.
* [Change the default StorageClass](change-default-storage-class.md) - This page shows how to change the default Storage Class that is used to provision volumes for PersistentVolumeClaims that have no special requirements.
* [Switching from Polling to CRI Event-based Updates to Container Status](switch-to-evented-pleg.md) - This page shows how to migrate nodes to use event based updates for container status.
* [Change the Reclaim Policy of a PersistentVolume](change-pv-reclaim-policy.md) - This page shows how to change the reclaim policy of a Kubernetes PersistentVolume.
* [Cloud Controller Manager Administration](running-cloud-controller.md) - Since cloud providers develop and release at a different pace compared to the Kubernetes project, abstracting the provider-specific code to the `cloud-controller-manager` binary allows cloud vendors to evolve independently from the core Kubernetes code.
* [Configure a kubelet image credential provider](kubelet-credential-provider.md) - Starting from Kubernetes v1.20, the kubelet can dynamically retrieve credentials for a container image registry using exec plugins.
* [Configure Quotas for API Objects](quota-api-object.md) - This page shows how to configure quotas for API objects, including PersistentVolumeClaims and Services.
* [Control CPU Management Policies on the Node](cpu-management-policies.md) - Kubernetes keeps many aspects of how pods execute on nodes abstracted from the user.
* [Control Memory Management Policies on a Node](memory-manager.md) - The Kubernetes *Memory Manager* enables the feature of guaranteed memory (and hugepages) allocation for pods in the `Guaranteed` QoS class.
* [Control Topology Management Policies on a node](topology-manager.md) - An increasing number of systems leverage a combination of CPUs and hardware accelerators to support latency-critical execution and high-throughput parallel computation.
* [Customizing DNS Service](dns-custom-nameservers.md) - This page explains how to configure your DNS Pod(s) and customize the DNS resolution process in your cluster.
* [Debugging DNS Resolution](dns-debugging-resolution.md) - This page provides hints on diagnosing DNS problems.
* [Declare Network Policy](declare-network-policy.md) - This document helps you get started using the Kubernetes NetworkPolicy API to declare network policies that govern how pods communicate with each other.
* [Developing Cloud Controller Manager](developing-cloud-controller-manager.md) - The cloud-controller-manager is a Kubernetes control plane component that embeds cloud-specific control logic.
* [Enable Or Disable A Kubernetes API](enable-disable-api.md) - This page shows how to enable or disable an API version from your cluster's control plane.
* [Encrypting Confidential Data at Rest](encrypt-data.md) - All of the APIs in Kubernetes that let you write persistent API resource data support at-rest encryption.
* [Decrypt Confidential Data that is Already Encrypted at Rest](decrypt-data.md) - All of the APIs in Kubernetes that let you write persistent API resource data support at-rest encryption.
* [Guaranteed Scheduling For Critical Add-On Pods](guaranteed-scheduling-critical-addon-pods.md) - Kubernetes core components such as the API server, scheduler, and controller-manager run on a control plane node.
* [IP Masquerade Agent User Guide](ip-masq-agent.md) - This page shows how to configure and enable the `ip-masq-agent`.
* [Limit Storage Consumption](limit-storage-consumption.md) - This example demonstrates how to limit the amount of storage consumed in a namespace.
* [Migrate Replicated Control Plane To Use Cloud Controller Manager](controller-manager-leader-migration.md) - The cloud-controller-manager is a Kubernetes control plane component that embeds cloud-specific control logic.
* [Operating etcd clusters for Kubernetes](configure-upgrade-etcd.md) - etcd is a consistent and highly-available key value store used as Kubernetes' backing store for all cluster data.
* [Reserve Compute Resources for System Daemons](reserve-compute-resources.md) - Kubernetes nodes can be scheduled to `Capacity`.
* [Running Kubernetes Node Components as a Non-root User](kubelet-in-userns.md) - This document describes how to run Kubernetes Node components such as kubelet, CRI, OCI, and CNI without root privileges, by using a user namespace.
* [Safely Drain a Node](safely-drain-node.md) - This page shows how to safely drain a node, optionally respecting the PodDisruptionBudget you have defined.
* [Securing a Cluster](securing-a-cluster.md) - This document covers topics related to protecting a cluster from accidental or malicious access and provides recommendations on overall security.
* [Harden Dynamic Resource Allocation in Your Cluster](hardening-dra.md) - This page shows cluster administrators how to harden authorization for Dynamic Resource Allocation (DRA), with a focus on least-privilege access for `ResourceClaim` status updates.
* [Set Kubelet Parameters Via A Configuration File](kubelet-config-file.md) - Some steps in this page use the `jq` tool.
* [Share a Cluster with Namespaces](namespaces.md) - This page shows how to view, work in, and delete namespaces.
* [Upgrade A Cluster](cluster-upgrade.md) - This page provides an overview of the steps you should follow to upgrade a Kubernetes cluster.
* [Use Cascading Deletion in a Cluster](use-cascading-deletion.md) - This page shows you how to specify the type of cascading deletion to use in your cluster during garbage collection.
* [Using a KMS provider for data encryption](kms-provider.md) - This page shows how to configure a Key Management Service (KMS) provider and plugin to enable secret data encryption.
* [Using CoreDNS for Service Discovery](coredns.md) - This page describes the CoreDNS upgrade process and how to install CoreDNS.
* [Using NodeLocal DNSCache in Kubernetes Clusters](nodelocaldns.md) - This page provides an overview of NodeLocal DNSCache feature in Kubernetes.
* [Using sysctls in a Kubernetes Cluster](sysctl-cluster.md) - This document describes how to configure and use kernel parameters within a Kubernetes cluster using the sysctl interface.
* [Verify Signed Kubernetes Artifacts](verify-signed-artifacts.md) - You will need to have the following tools installed:
