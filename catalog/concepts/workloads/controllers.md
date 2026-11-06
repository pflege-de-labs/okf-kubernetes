---
type: section
title: Workload Management
description: Kubernetes provides several built-in APIs for declarative management of your workloads and the components of those workloads.
resource: https://kubernetes.io/docs/concepts/workloads/controllers/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/_index.md
    title: 'Kubernetes website source: concepts/workloads/controllers/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/controllers/
    title: Workload Management | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/workloads/controllers/_index.md
---

# Workload Management

Kubernetes provides several built-in APIs for declarative management of your
[workloads](/reference/glossary/workload.md)
and the components of those workloads.

Ultimately, your applications run as containers inside
[Pods](/reference/glossary/pod.md); however, managing individual
Pods would be a lot of effort. For example, if a Pod fails, you probably want to
run a new Pod to replace it. Kubernetes can do that for you.

You use the Kubernetes API to create a workload
[object](/reference/glossary/object.md) that represents a higher abstraction level
than a Pod, and then the Kubernetes
[control plane](/reference/glossary/control-plane.md) automatically manages
Pod objects on your behalf, based on the specification for the workload object you defined.

The built-in APIs for managing workloads are:

[Deployment](/concepts/workloads/controllers/deployment.md) (and, indirectly, [ReplicaSet](/concepts/workloads/controllers/replicaset.md)),
the most common way to run an application on your cluster.
Deployment is a good fit for managing a stateless application workload on your cluster, where
any Pod in the Deployment is interchangeable and can be replaced if needed.
(Deployments are a replacement for the legacy
[ReplicationController](/reference/glossary/replication-controller.md) API).

A [StatefulSet](/concepts/workloads/controllers/statefulset.md) lets you
manage one or more Pods – all running the same application code – where the Pods rely
on having a distinct identity. This is different from a Deployment where the Pods are
expected to be interchangeable.
The most common use for a StatefulSet is to be able to make a link between its Pods and
their persistent storage. For example, you can run a StatefulSet that associates each Pod
with a [PersistentVolume](/concepts/storage/persistent-volumes.md). If one of the Pods
in the StatefulSet fails, Kubernetes makes a replacement Pod that is connected to the
same PersistentVolume.

A [DaemonSet](/concepts/workloads/controllers/daemonset.md) defines Pods that provide
facilities that are local to a specific [node](/reference/glossary/node.md);
for example, a driver that lets containers on that node access a storage system. You use a DaemonSet
when the driver, or other node-level service, has to run on the node where it's useful.
Each Pod in a DaemonSet performs a role similar to a system daemon on a classic Unix / POSIX
server.
A DaemonSet might be fundamental to the operation of your cluster,
such as a plugin to let that node access
[cluster networking](/concepts/cluster-administration/networking.md#how-to-implement-the-kubernetes-network-model),
it might help you to manage the node,
or it could provide less essential facilities that enhance the container platform you are running.
You can run DaemonSets (and their pods) across every node in your cluster, or across just a subset (for example,
only install the GPU accelerator driver on nodes that have a GPU installed).

You can use a [Job](/concepts/workloads/controllers/job.md) and / or
a [CronJob](/concepts/workloads/controllers/cron-jobs.md) to
define tasks that run to completion and then stop. A Job represents a one-off task,
whereas each CronJob repeats according to a schedule.

Other topics in this section:
