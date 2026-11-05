---
type: section
title: Workloads
description: Understand Pods, the smallest deployable compute object in Kubernetes, and the higher-level abstractions that help you to run them.
resource: https://kubernetes.io/docs/concepts/workloads/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/_index.md
    title: 'Kubernetes website source: concepts/workloads/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/
    title: Workloads | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/workloads/_index.md
---

# Workloads

A workload is an application running on Kubernetes.
Whether your workload is a single component or several that work together, on Kubernetes you run
it inside a set of [_pods_](/concepts/workloads/pods.md).
In Kubernetes, a Pod represents a set of one or more running
[containers](/reference/glossary/container.md) on your cluster.

Kubernetes pods have a [defined lifecycle](/concepts/workloads/pods/pod-lifecycle.md).
For example, once a pod is running in your cluster then a critical fault on the
[node](/reference/glossary/node.md) where that pod is running means that
all the pods on that node fail. Kubernetes treats that level of failure as final: you
would need to create a new Pod to recover, even if the node later becomes healthy.

However, to make life considerably easier, you don't need to manage each Pod directly.
Instead, you can use _workload resources_ that manage a set of pods on your behalf.
These resources configure [controllers](/reference/glossary/controller.md)
that make sure the right number of the right kind of pod are running, to match the state
you specified.

Kubernetes provides several built-in workload resources:

* [Deployment](/concepts/workloads/controllers/deployment.md) and [ReplicaSet](/concepts/workloads/controllers/replicaset.md)
  (replacing the legacy resource
  [ReplicationController](/reference/glossary/replication-controller.md)).
  Deployment is a good fit for managing a stateless application workload on your cluster,
  where any Pod in the Deployment is interchangeable and can be replaced if needed.
* [StatefulSet](/concepts/workloads/controllers/statefulset.md) lets you
  run one or more related Pods that do track state somehow. For example, if your workload
  records data persistently, you can run a StatefulSet that matches each Pod with a
  [PersistentVolume](/concepts/storage/persistent-volumes.md). Your code, running in the
  Pods for that StatefulSet, can replicate data to other Pods in the same StatefulSet
  to improve overall resilience.
* [DaemonSet](/concepts/workloads/controllers/daemonset.md) defines Pods that provide
  facilities that are local to nodes.
  Every time you add a node to your cluster that matches the specification in a DaemonSet,
  the control plane schedules a Pod for that DaemonSet onto the new node.
  Each pod in a DaemonSet performs a job similar to a system daemon on a classic Unix / POSIX
  server. A DaemonSet might be fundamental to the operation of your cluster, such as
  a plugin to run [cluster networking](/concepts/cluster-administration/networking.md#how-to-implement-the-kubernetes-network-model),
  it might help you to manage the node,
  or it could provide optional behavior that enhances the container platform you are running.
* [Job](/concepts/workloads/controllers/job.md) and
  [CronJob](/concepts/workloads/controllers/cron-jobs.md) provide different ways to
  define tasks that run to completion and then stop.
  You can use a [Job](/concepts/workloads/controllers/job.md)  to
  define a task that runs to completion, just once. You can use a
  [CronJob](/concepts/workloads/controllers/cron-jobs.md) to run
  the same Job multiple times according a schedule.

In the wider Kubernetes ecosystem, you can find third-party workload resources that provide
additional behaviors. Using a
[custom resource definition](/concepts/extend-kubernetes/api-extension/custom-resources.md),
you can add in a third-party workload resource if you want a specific behavior that's not part
of Kubernetes' core. For example, if you wanted to run a group of Pods for your application but
stop work unless _all_ the Pods are available (perhaps for some high-throughput distributed task),
then you can implement or install an extension that does provide that feature.

## Workload placement

**Feature state:** Kubernetes v1.37 [beta] (feature gate `GenericWorkload`, disabled by default)

While standard workload resources (like Deployments and Jobs) manage the lifecycle of Pods,
you may have complex scheduling requirements where groups of Pods must be treated as a single unit.

The [Workload API](/concepts/workloads/workload-api.md) allows you to define `PodGroupTemplates` to group Pods and apply advanced scheduling policies to them, 
such as [gang scheduling](/concepts/scheduling-eviction/gang-scheduling.md).
Controllers create [PodGroup](/concepts/workloads/podgroup-api.md) objects from these templates at runtime, 
and `Pods` reference their `PodGroup` via the
`spec.schedulingGroup` field. This is particularly useful for batch processing and machine
learning workloads where "all-or-nothing" placement is required.

## What's next

As well as reading about each API kind for workload management, you can read how to
do specific tasks:

* [Run a stateless application using a Deployment](/tasks/run-application/run-stateless-application-deployment.md)
* Run a stateful application either as a [single instance](/tasks/run-application/run-single-instance-stateful-application.md)
  or as a [replicated set](/tasks/run-application/run-replicated-stateful-application.md)
* [Run automated tasks with a CronJob](/tasks/job/automated-tasks-with-cron-jobs.md)

To learn about Kubernetes' mechanisms for separating code from configuration,
visit [Configuration](/concepts/configuration.md).

There are two supporting concepts that provide backgrounds about how Kubernetes manages pods
for applications:
* [Garbage collection](/concepts/architecture/garbage-collection.md) tidies up objects
  from your cluster after their _owning resource_ has been removed.
* The [_time-to-live after finished_ controller](/concepts/workloads/controllers/ttlafterfinished.md)
  removes Jobs once a defined time has passed since they completed.

Once your application is running, you might want to make it available on the internet as
a [Service](/concepts/services-networking/service.md) or, for web application only,
using an [Ingress](/concepts/services-networking/ingress.md).
