# Scheduling

* [CompositePodGroup](composite-pod-group-v1alpha3.md) - CompositePodGroup represents a runtime instance of pod groups grouped together. CompositePodGroups are created by workload controllers (LWS, JobSet, etc...) from Workload.compositePodGroupTemplates. CompositePodGroup API enablement is toggled by the CompositePodGroup feature gate.
* [PodGroup](pod-group-v1beta1.md) - PodGroup represents a runtime instance of pods grouped together. PodGroups are created by workload controllers (Job, LWS, JobSet, etc...) from Workload.podGroupTemplates. PodGroup API enablement is toggled by the GenericWorkload feature gate.
* [PriorityClass](priority-class-v1.md) - PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
* [Workload](workload-v1beta1.md) - Workload allows for expressing scheduling constraints that should be used when managing the lifecycle of workloads from the scheduling perspective, including scheduling, preemption, eviction and other phases. Workload API enablement is toggled by the GenericWorkload feature gate.
