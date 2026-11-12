---
type: constraint
title: Port numbers, port names and NodePort range
description: Container and Service port numbers are 1-65535, port names must be IANA_SVC_NAMEs unique within the Pod, NodePorts default to 30000-32767, plus the well-known control plane and node ports.
tags: [kubernetes, requirements, ports, networking, service, nodeport]
generated: { by: "claude-code/claude-opus-5-5", at: "2026-11-05T00:00:00Z" }
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
sources:
  - id: container-port-name
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/core/pod-v1.md#L875-L876
    title: Pod API reference - ContainerPort name
  - id: probe-port
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/core/pod-v1.md#L1589
    title: Pod API reference - port number or name range
  - id: target-port
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/kubernetes-api/core/service-v1.md#L324
    title: Service API reference - targetPort
  - id: nodeport-range
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/services-networking/service.md#L476-L477
    title: Service - NodePort range
  - id: nodeport-bands
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/services-networking/service.md#L530-L537
    title: Service - NodePort static and dynamic bands
  - id: ports
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/networking/ports-and-protocols.md#L14-L32
    title: Ports and Protocols
---

# Port numbers, port names and NodePort range

# Port numbers and names

- A port number referenced by a probe or a Service `targetPort` must be in the
  range 1 to 65535; a port name must be an `IANA_SVC_NAME`.[^probe-port][^target-port]
- A named container port must be an `IANA_SVC_NAME` and unique within the Pod;
  Services can refer to it by name.[^container-port-name]
- A string `targetPort` is looked up as a named port in the target Pod's
  containers; if unset, `targetPort` defaults to `port`.[^target-port]

# NodePort

- The control plane allocates NodePorts from the range set by the
  kube-apiserver flag `--service-node-port-range`, default **30000-32767**.[^nodeport-range]
- The range is split into a static band (30000-30085 by default) for manually
  chosen ports and a dynamic band (30086-32767) used first for automatic
  assignment, which falls back to the static band only when the dynamic band
  is exhausted.[^nodeport-bands]

# Well-known ports

| Component | Protocol | Port | Used by |
|---|---|---|---|
| Kubernetes API server | TCP | 6443 | All |
| etcd server client API | TCP | 2379-2380 | kube-apiserver, etcd |
| Kubelet API | TCP | 10250 | Self, control plane |
| kube-scheduler | TCP | 10259 | Self |
| kube-controller-manager | TCP | 10257 | Self |
| kube-proxy | TCP | 10256 | Self, load balancers |
| NodePort Services | TCP, UDP | 30000-32767 | All |

Source: [^ports]

Full text: [Service](/concepts/services-networking/service.md),
[Ports and Protocols](/reference/networking/ports-and-protocols.md),
[Pod API](/reference/kubernetes-api/core/pod-v1.md),
[Service API](/reference/kubernetes-api/core/service-v1.md).

[^container-port-name]: Pod API reference - ContainerPort name
[^probe-port]: Pod API reference - port number or name range
[^target-port]: Service API reference - targetPort
[^nodeport-range]: Service - NodePort range
[^nodeport-bands]: Service - NodePort static and dynamic bands
[^ports]: Ports and Protocols
