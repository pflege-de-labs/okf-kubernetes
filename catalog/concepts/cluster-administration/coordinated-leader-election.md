---
type: concept
title: Coordinated Leader Election
description: Kubernetes 1.37 includes a beta feature that allows control plane components to deterministically select a leader via _coordinated leader election_.
resource: https://kubernetes.io/docs/concepts/cluster-administration/coordinated-leader-election/
tags: [kubernetes, concepts, cluster-administration]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/cluster-administration/coordinated-leader-election.md
    title: 'Kubernetes website source: concepts/cluster-administration/coordinated-leader-election.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/cluster-administration/coordinated-leader-election/
    title: Coordinated Leader Election | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/cluster-administration/coordinated-leader-election.md
---

# Coordinated Leader Election

**Feature state:** Kubernetes v1.33 [beta] (feature gate `CoordinatedLeaderElection`, disabled by default)

Kubernetes 1.37 includes a beta feature that allows [control plane](/reference/glossary/control-plane.md) components to
deterministically select a leader via _coordinated leader election_.
This is useful to satisfy Kubernetes version skew constraints during cluster upgrades.
Currently, the only builtin selection strategy is `OldestEmulationVersion`,
preferring the leader with the lowest emulation version, followed by binary
version, followed by creation timestamp.

## Enabling coordinated leader election

Ensure that `CoordinatedLeaderElection` [feature
gate](/reference/command-line-tools-reference/feature-gates.md) is enabled
when you start the [API Server](/reference/glossary/kube-apiserver.md): and that the `coordination.k8s.io/v1beta1` API group is
enabled.

This can be done by setting flags `--feature-gates="CoordinatedLeaderElection=true"` and
`--runtime-config="coordination.k8s.io/v1beta1=true"`.

## Component configuration

Provided that you have enabled the `CoordinatedLeaderElection` feature gate _and_  
have the `coordination.k8s.io/v1beta1` API group enabled, compatible control plane  
components automatically use the LeaseCandidate and Lease APIs to elect a leader  
as needed.  

For Kubernetes 1.37, two control plane components  
(kube-controller-manager and kube-scheduler) automatically use coordinated  
leader election when the feature gate and API group are enabled.

## Leader selection for Kubernetes components

Kubernetes uses the [Lease API](/concepts/architecture/leases.md) to perform leader election among multiple instances of the same control-plane component in a high-availability cluster, such as `kube-controller-manager` or `kube-scheduler`.

A [Lease](/concepts/architecture/leases.md) acts as a lightweight distributed lock, stored by the [Kubernetes API server](/reference/command-line-tools-reference/kube-apiserver.md).
All running instances of a component watch or periodically read the relevant Lease object
to determine which instance is currently acting as the leader.

The [Lease API](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/lease-v1/) defines fields
such as:

`holderIdentity`
: the identity (for example: pod name or hostname-based string) of the current leader.

`acquireTime`
: timestamp when leadership was acquired.

`renewTime`
: timestamp of the most recent renewal by the leader.

`leaseDurationSeconds`
: the validity period of the lease (candidates should wait this long plus a small grace period before attempting to acquire an expired lease).

`leaseTransitions`
: counter of how many times leadership has changed hands.

These fields indicate which instance holds leadership and how long that leadership remains valid.

When the [Lease](/concepts/architecture/leases.md) does not exist or has expired (current time > `renewTime` + `leaseDurationSeconds`), candidate instances attempt to update the Lease with their identity. Kubernetes relies on _optimistic concurrency control_ via the object's `resourceVersion`: only one update succeeds due to version mismatch on concurrent attempts. The instance whose update is accepted becomes the _leader_.

Kubernetes uses the [LeaseCandidate](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/lease-candidate-v1beta1/) 
API to manage leader elections. Control plane components such as `kube-controller-manager` and `kube-scheduler` register their role as a candidate by creating LeaseCandidate objects, which track all instances competing for leadership and carry metadata including the candidate's identity, binary version, and emulation version.

During an election, candidates coordinate through a shared [Lease](/concepts/architecture/leases.md). 
The Kubernetes control plane guarantees that only one candidate successfully acquires the [Lease](/concepts/architecture/leases.md) and assumes the role of _leader_, while all others remain as followers. If the current _leader_ fails to renew the [Lease](/concepts/architecture/leases.md) within the selected timeout period, the remaining candidates compete to acquire leadership and elect a new _leader_.

Once elected, the leader periodically renews its Lease by updating the `renewTime` field
(for example, performing renewal every `leaseDurationSeconds` ÷ 2, in order to avoid conflicts when the [Lease](/concepts/architecture/leases.md) is about to expire).
As long as renewals occur before the lease expires, the current leader instance retains leadership.
If the leader crashes, becomes unreachable, or stops renewing the Lease, that Lease expires. Other healthy instances detect the expired Lease and attempt a new election.

This mechanism ensures that even though multiple replicas of a component may be running for stability and recovery, _only one instance actively performs control tasks at a time_, while the others remain on standby, watching the Lease and ready to take over quickly if needed.
