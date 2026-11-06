# Scheduling, Preemption and Eviction

* [Kubernetes Scheduler](kube-scheduler.md) - In Kubernetes, _scheduling_ refers to making sure that Pods are matched to Nodes so that Kubelet can run them.
* [Running multiple schedulers](running-multiple-schedulers.md) - Kubernetes allows you to run multiple schedulers within a single cluster.
* [Assigning Pods to Nodes](assign-pod-node.md) - You can constrain a Pod so that it is _restricted_ to run on particular node(s), or to _prefer_ to run on particular nodes.
* [Pod Overhead](pod-overhead.md) - When you run a Pod on a Node, the Pod itself takes an amount of system resources.
* [Pod Scheduling Readiness](pod-scheduling-readiness.md) - Pods were considered ready for scheduling once created.
* [Pod Topology Spread Constraints](topology-spread-constraints.md) - You can use _topology spread constraints_ to control how Pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains.
* [Taints and Tolerations](taint-and-toleration.md) - is a property of Pods that *attracts* them to a set of nodes (either as a preference or a hard requirement).
* [Scheduling Framework](scheduling-framework.md) - The _scheduling framework_ is a pluggable architecture for the Kubernetes scheduler.
* [Gang Scheduling](gang-scheduling.md) - Gang scheduling ensures that a group of Pods are scheduled on an "all-or-nothing" basis.
* [Scheduler Performance Tuning](scheduler-perf-tuning.md) - is the Kubernetes default scheduler.
* [PodGroup Scheduling](podgroup-scheduling.md) - The standard Kubernetes scheduler evaluates Pods sequentially.
* [Resource Bin Packing](resource-bin-packing.md) - This article applies to resource bin packing in context of scheduling of a single pod.
* [Workload-Aware Preemption](workload-aware-preemption.md) - In v1.36, the workload-aware preemption logic was gated by `WorkloadAwarePreemption` feature gate.
* [Topology-Aware Workload Scheduling](topology-aware-scheduling.md) - that allows finding the optimal placement for the considered PodGroup, guaranteeing that all pods will be collocated within the same topology domain.
* [Pod Priority and Preemption](pod-priority-preemption.md) - importance of a Pod relative to other Pods.
* [Node-pressure Eviction](node-pressure-eviction.md) - Node-pressure eviction is the process by which the kubelet proactively terminates pods to reclaim resource on nodes.</br>
* [API-initiated Eviction](api-eviction.md) - API-initiated eviction is the process by which you use the Eviction API to create an `Eviction` object that triggers graceful pod termination.
* [Node Declared Features](node-declared-features.md) - Kubernetes nodes use _declared features_ to report the availability of specific features that are new or feature-gated.
