---
type: concept
title: Kubernetes Self-Healing
description: Kubernetes is designed with self-healing capabilities that help maintain the health and availability of workloads.
resource: https://kubernetes.io/docs/concepts/architecture/self-healing/
tags: [kubernetes, concepts, architecture]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/architecture/self-healing.md
    title: 'Kubernetes website source: concepts/architecture/self-healing.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/architecture/self-healing/
    title: Kubernetes Self-Healing | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/architecture/self-healing.md
---

# Kubernetes Self-Healing

Kubernetes is designed with self-healing capabilities that help maintain the health and availability of workloads. 
It automatically replaces failed containers, reschedules workloads when nodes become unavailable, and ensures that the desired state of the system is maintained.

## Self-Healing capabilities {#self-healing-capabilities} 

- **Container-level restarts:** If a container inside a Pod fails, Kubernetes restarts it based on the [`restartPolicy`](/concepts/workloads/pods/pod-lifecycle.md#restart-policy).

- **Replica replacement:** If a Pod in a [Deployment](/concepts/workloads/controllers/deployment.md) or [StatefulSet](/concepts/workloads/controllers/statefulset.md) fails, Kubernetes creates a replacement Pod to maintain the specified number of replicas.
  If a Pod that is part of a [DaemonSet](/concepts/workloads/controllers/daemonset.md) fails, the control plane
  creates a replacement Pod to run on the same node.
  
- **Persistent storage recovery:** If a node is running a Pod with a PersistentVolume (PV) attached, and the node fails, Kubernetes can reattach the volume to a new Pod on a different node.

- **Load balancing for Services:** If a Pod behind a [Service](/concepts/services-networking/service.md) fails, Kubernetes automatically removes it from the Service's endpoints to route traffic only to healthy Pods.

Here are some of the key components that provide Kubernetes self-healing:

- **[kubelet](/concepts/architecture.md#kubelet):** Ensures that containers are running, and restarts those that fail.

- **Deployment (via ReplicaSet), ReplicaSet, StatefulSet and DaemonSet controllers:** Maintain the desired number of Pod replicas.

- **PersistentVolume controller:** Manages volume attachment and detachment for stateful workloads.

## Considerations {#considerations} 

- **Storage Failures:** If a persistent volume becomes unavailable, recovery steps may be required.

- **Application Errors:** Kubernetes can restart containers, but underlying application issues must be addressed separately.

## What's next 

- Read more about [Pods](/concepts/workloads/pods.md)
- Learn about [Kubernetes Controllers](/concepts/architecture/controller.md)
- Explore [PersistentVolumes](/concepts/storage/persistent-volumes.md)
- Read about [node autoscaling](/concepts/cluster-administration/node-autoscaling.md). Node autoscaling
  also provides automatic healing if or when nodes fail in your cluster.
