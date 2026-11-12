---
type: concept
title: Service Internal Traffic Policy
description: If two Pods in your cluster want to communicate, and both Pods are actually running on the same node, use _Service Internal Traffic Policy_ to keep network traffic within that node. Avoiding a round trip via the cluster network can help with reliability, performance (network latency and throughput), or cost.
resource: https://kubernetes.io/docs/concepts/services-networking/service-traffic-policy/
tags: [kubernetes, concepts, services-networking]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/services-networking/service-traffic-policy.md
    title: 'Kubernetes website source: concepts/services-networking/service-traffic-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/services-networking/service-traffic-policy/
    title: Service Internal Traffic Policy | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/services-networking/service-traffic-policy.md
---

# Service Internal Traffic Policy

**Feature state:** Kubernetes v1.26 [stable]

_Service Internal Traffic Policy_ enables internal traffic restrictions to only route
internal traffic to endpoints within the node the traffic originated from. The
"internal" traffic here refers to traffic originated from Pods in the current
cluster. This can help to reduce costs and improve performance.

## Using Service Internal Traffic Policy

You can enable the internal-only traffic policy for a
[Service](/reference/glossary/service.md), by setting its
`.spec.internalTrafficPolicy` to `Local`. This tells kube-proxy to only use node local
endpoints for cluster internal traffic.

**Note:**
For pods on nodes with no endpoints for a given Service, the Service
behaves as if it has zero endpoints (for Pods on this node) even if the service
does have endpoints on other nodes.

The following example shows what a Service looks like when you set
`.spec.internalTrafficPolicy` to `Local`:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app.kubernetes.io/name: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
  internalTrafficPolicy: Local
```

## How it works

The kube-proxy filters the endpoints it routes to based on the
`spec.internalTrafficPolicy` setting. When it's set to `Local`, only node local
endpoints are considered. When it's `Cluster` (the default), or is not set,
Kubernetes considers all endpoints.

## What's next

* Read about [Topology Aware Routing](/concepts/services-networking/topology-aware-routing.md)
* Read about [Service External Traffic Policy](/tasks/access-application-cluster/create-external-load-balancer.md#preserving-the-client-source-ip)
* Follow the [Connecting Applications with Services](https://kubernetes.io/docs/tutorials/services/connect-applications-service/) tutorial
