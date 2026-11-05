---
type: concept
title: Scheduling, Preemption and Eviction
description: In Kubernetes, scheduling refers to making sure that Pods are matched to Nodes so that the kubelet can run them.
resource: https://kubernetes.io/docs/concepts/scheduling-eviction/
tags: [kubernetes, concepts]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/scheduling-eviction/_index.md
    title: 'Kubernetes website source: concepts/scheduling-eviction/_index.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/scheduling-eviction/
    title: Scheduling, Preemption and Eviction | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/scheduling-eviction/_index.md
---

# Scheduling, Preemption and Eviction

In Kubernetes, scheduling refers to making sure that [Pods](/reference/glossary/pod.md)
are matched to [Nodes](/reference/glossary/node.md) so that the
[kubelet](/reference/glossary/kubelet.md) can run them. Preemption
is the process of terminating Pods with lower [Priority](/reference/glossary/pod-priority.md)
so that Pods with higher Priority can schedule on Nodes. Eviction is the process
of terminating one or more Pods on Nodes.

## Scheduling

* [Kubernetes Scheduler](/concepts/scheduling-eviction/kube-scheduler.md)
* [Running Multiple Schedulers](/concepts/scheduling-eviction/running-multiple-schedulers.md)
* [Assigning Pods to Nodes](/concepts/scheduling-eviction/assign-pod-node.md)
* [Pod Overhead](/concepts/scheduling-eviction/pod-overhead.md)
* [Pod Topology Spread Constraints](/concepts/scheduling-eviction/topology-spread-constraints.md)
* [Taints and Tolerations](/concepts/scheduling-eviction/taint-and-toleration.md)
* [Scheduling Framework](/concepts/scheduling-eviction/scheduling-framework.md)
* [Scheduler Performance Tuning](/concepts/scheduling-eviction/scheduler-perf-tuning.md)
* [Resource Bin Packing for Extended Resources](/concepts/scheduling-eviction/resource-bin-packing.md)
* [Pod Scheduling Readiness](/concepts/scheduling-eviction/pod-scheduling-readiness.md)
* [PodGroup Scheduling](/concepts/scheduling-eviction/podgroup-scheduling.md)
* [Gang Scheduling](/concepts/scheduling-eviction/gang-scheduling.md)
* [Topology-aware Scheduling](/concepts/scheduling-eviction/topology-aware-scheduling.md)
* [Workload-aware preemption](/concepts/scheduling-eviction/workload-aware-preemption.md)
* [Descheduler](https://github.com/kubernetes-sigs/descheduler#descheduler-for-kubernetes)
* [Node Declared Features](/concepts/scheduling-eviction/node-declared-features.md)

## Pod Disruption

[Pod disruption](/concepts/workloads/pods/disruptions.md) is the process by which 
Pods on Nodes are terminated either voluntarily or involuntarily. 

 

Voluntary disruptions are started intentionally by application owners or cluster 
administrators. Involuntary disruptions are unintentional and can be triggered by 
unavoidable issues like Nodes running out of [resources](/reference/glossary/infrastructure-resource.md),
or by accidental deletions.

* [Pod Priority and Preemption](/concepts/scheduling-eviction/pod-priority-preemption.md)
* [Node-pressure Eviction](/concepts/scheduling-eviction/node-pressure-eviction.md)
* [API-initiated Eviction](/concepts/scheduling-eviction/api-eviction.md)
