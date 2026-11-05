---
type: task
title: Harden Dynamic Resource Allocation in Your Cluster
description: This page shows cluster administrators how to harden authorization for Dynamic Resource Allocation (DRA), with a focus on least-privilege access for `ResourceClaim` status updates.
resource: https://kubernetes.io/docs/tasks/administer-cluster/hardening-dra/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/administer-cluster/hardening-dra.md
    title: 'Kubernetes website source: tasks/administer-cluster/hardening-dra.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/hardening-dra/
    title: Harden Dynamic Resource Allocation in Your Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/administer-cluster/hardening-dra.md
---

# Harden Dynamic Resource Allocation in Your Cluster

This page shows cluster administrators how to harden authorization for
Dynamic Resource Allocation (DRA), with a focus on least-privilege access for
`ResourceClaim` status updates.

## Before you begin

- You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.
- Dynamic Resource Allocation is configured in your cluster.
- You can edit RBAC resources and restart or roll out DRA components.

## Identify DRA components that write status

Document which identities (usually ServiceAccounts) update ResourceClaim
status in your cluster. Typical writers are:

- kube-scheduler or a custom allocation controller
- node-local DRA drivers
- multi-node DRA status controllers

## Grant least-privilege permissions for synthetic subresources

Starting in Kubernetes v1.36, DRA status updates require synthetic subresource
permissions in addition to `resourceclaims/status`.

### Grant scheduler and allocation-controller permissions

Apply a role that allows binding-related updates:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: dra-binding-updater
rules:
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/status"]
    verbs: ["get", "patch", "update"]
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/binding"]
    verbs: ["patch", "update"]
```

### Grant node-local driver permissions

Use node-aware verbs for node-local drivers:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: dra-node-driver-status-updater
rules:
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/status"]
    verbs: ["get", "patch", "update"]
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/driver"]
    verbs: ["associated-node:patch", "associated-node:update"]
    resourceNames: ["dra.example.com"]
```

### Grant multi-node controller permissions only when needed

Use `arbitrary-node:*` only for components that must update from any node:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: dra-multinode-status-updater
rules:
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/status"]
    verbs: ["get", "patch", "update"]
  - apiGroups: ["resource.k8s.io"]
    resources: ["resourceclaims/driver"]
    verbs: ["arbitrary-node:patch", "arbitrary-node:update"]
    resourceNames: ["dra.example.com"]
```

## Bind roles to explicit identities

Create `ClusterRoleBinding` objects for each component identity, and avoid
sharing a broad role across unrelated DRA components.

Restrict `resourceclaims/driver` rules with `resourceNames` where possible so
an identity can only write status for the specific DRA driver it operates.

## Validate and monitor

1. Verify each identity has only the required verbs and subresources.
1. Confirm DRA status updates work after rollout.
1. Watch API server audit events for denied `resourceclaims/binding` and
   `resourceclaims/driver` requests.

## What's next

- [Hardening Guide - Dynamic Resource Allocation](/concepts/security/hardening-guide/dynamic-resource-allocation.md)
- [Securing a Cluster](/tasks/administer-cluster/securing-a-cluster.md)
- [Authorization](/reference/access-authn-authz/authorization.md)
