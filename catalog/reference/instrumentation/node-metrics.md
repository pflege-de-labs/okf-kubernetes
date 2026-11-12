---
type: reference
title: Node metrics data
description: Mechanisms for accessing metrics at node, volume, pod and container level, as seen by the kubelet.
resource: https://kubernetes.io/docs/reference/instrumentation/node-metrics/
tags: [kubernetes, reference, instrumentation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/instrumentation/node-metrics.md
    title: 'Kubernetes website source: reference/instrumentation/node-metrics.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/instrumentation/node-metrics/
    title: Node metrics data | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/instrumentation/node-metrics.md
---

# Node metrics data

The [kubelet](/reference/command-line-tools-reference/kubelet.md)
gathers metric statistics at the node, volume, pod and container level,
and emits this information in the
[Summary API](/reference/kubelet-api/stats.v1alpha1.md).

You can send a proxied request to the stats summary API via the
Kubernetes API server.

Here is an example of a Summary API request for a node named `minikube`:

```shell
kubectl get --raw "/api/v1/nodes/minikube/proxy/stats/summary"
```

Here is the same API call using `curl`:

```shell
# You need to run "kubectl proxy" first
# Change 8080 to the port that "kubectl proxy" assigns
curl http://localhost:8080/api/v1/nodes/minikube/proxy/stats/summary
```

**Note:**
Beginning with `metrics-server` 0.6.x, `metrics-server` queries the `/metrics/resource`
kubelet endpoint, and not `/stats/summary`.

## Summary metrics API source {#summary-api-source}

By default, Kubernetes fetches node summary metrics data using an embedded
[cAdvisor](https://github.com/google/cadvisor) that runs within the kubelet. If you 
enable the `PodAndContainerStatsFromCRI` [feature gate](/reference/command-line-tools-reference/feature-gates.md) 
in your cluster, and you use a container runtime that supports statistics access via
[Container Runtime Interface](/reference/glossary/cri.md) (CRI), then
the kubelet [fetches Pod- and container-level metric data using CRI](/reference/instrumentation/cri-pod-container-metrics.md), and not via cAdvisor.

## Pressure Stall Information (PSI) {#psi}

**Feature state:** Kubernetes v1.36 [stable] (feature gate `KubeletPSI`, enabled by default)

As a stable feature, Kubernetes lets you configure kubelet to collect Linux kernel
[Pressure Stall Information](https://docs.kernel.org/accounting/psi.html)
(PSI) for CPU, memory, and I/O usage. The information is collected at node, pod and container level.
See [Summary API](/reference/kubelet-api/stats.v1alpha1.md) for detailed schema.
Starting with Kubernetes v.1.36, the `KubeletPSI` [feature gate](/reference/command-line-tools-reference/feature-gates.md) is locked to true and cannot be disabled. The information is also exposed in
[Prometheus metrics](/concepts/cluster-administration/system-metrics.md#psi-metrics).

You can learn how to interpret the PSI metrics in [Understand PSI Metrics](/reference/instrumentation/understand-psi-metrics.md).

### Requirements

Pressure Stall Information requires:

- [Linux kernel versions 4.20 or later](/reference/node/kernel-version-requirements.md#requirements-psi).
- [cgroup v2](/concepts/architecture/cgroups.md)

## What's next

The task pages for [Troubleshooting Clusters](/tasks/debug/debug-cluster.md) discuss
how to use a metrics pipeline that rely on these data.
