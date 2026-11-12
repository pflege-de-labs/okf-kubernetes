---
type: concept
title: Dynamic Volume Provisioning
description: Dynamic volume provisioning allows storage volumes to be created on-demand.
resource: https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/
tags: [kubernetes, concepts, storage]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/storage/dynamic-provisioning.md
    title: 'Kubernetes website source: concepts/storage/dynamic-provisioning.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/
    title: Dynamic Volume Provisioning | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/concepts/storage/dynamic-provisioning.md
---

# Dynamic Volume Provisioning

Dynamic volume provisioning allows storage volumes to be created on-demand.
Without dynamic provisioning, cluster administrators have to manually make
calls to their cloud or storage provider to create new storage volumes, and
then create [`PersistentVolume` objects](/concepts/storage/persistent-volumes.md)
to represent them in Kubernetes. The dynamic provisioning feature eliminates
the need for cluster administrators to pre-provision storage. Instead, it
automatically provisions storage when users create
[`PersistentVolumeClaim` objects](/concepts/storage/persistent-volumes.md).

## Background

The implementation of dynamic volume provisioning is based on the API object `StorageClass`
from the API group `storage.k8s.io`. A cluster administrator can define as many
`StorageClass` objects as needed, each specifying a *volume plugin* (aka
*provisioner*) that provisions a volume and the set of parameters to pass to
that provisioner when provisioning.
A cluster administrator can define and expose multiple flavors of storage (from
the same or different storage systems) within a cluster, each with a custom set
of parameters. This design also ensures that end users don't have to worry
about the complexity and nuances of how storage is provisioned, but still
have the ability to select from multiple storage options.

For more details, see the [Storage Classes](/concepts/storage/storage-classes.md) concept.

## Enabling Dynamic Provisioning

To enable dynamic provisioning, a cluster administrator needs to pre-create
one or more StorageClass objects for users.
StorageClass objects define which provisioner should be used and what parameters
should be passed to that provisioner when dynamic provisioning is invoked.
The name of a StorageClass object must be a valid
[DNS subdomain name](/concepts/overview/working-with-objects/names.md#dns-subdomain-names).

The following manifest creates a storage class "slow" which provisions standard
disk-like persistent disks.

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: slow
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-standard
```

The following manifest creates a storage class "fast" which provisions
SSD-like persistent disks.

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
```

## Using Dynamic Provisioning

Users request dynamically provisioned storage by including a storage class in
their `PersistentVolumeClaim`. Before Kubernetes v1.6, this was done via the
`volume.beta.kubernetes.io/storage-class` annotation. However, this annotation
is deprecated since v1.9. Users now can and should instead use the
`storageClassName` field of the `PersistentVolumeClaim` object. The value of
this field must match the name of a `StorageClass` configured by the
administrator (see [Enabling Dynamic Provisioning](#enabling-dynamic-provisioning)).

To select the "fast" storage class, for example, a user would create the
following PersistentVolumeClaim:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: claim1
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: fast
  resources:
    requests:
      storage: 30Gi
```

This claim results in an SSD-like Persistent Disk being automatically
provisioned. When the claim is deleted, the volume is destroyed.

## Defaulting Behavior

Dynamic provisioning can be enabled on a cluster such that all claims are
dynamically provisioned if no storage class is specified. A cluster administrator
can enable this behavior by:

- Marking one `StorageClass` object as *default*.
- Making sure that the [`DefaultStorageClass` admission controller](/reference/access-authn-authz/admission-controllers.md#defaultstorageclass)
  is enabled on the API server.

An administrator can mark a specific `StorageClass` as default by adding the
[`storageclass.kubernetes.io/is-default-class` annotation](/reference/labels-annotations-taints.md#storageclass-kubernetes-io-is-default-class) to it.
When a default `StorageClass` exists in a cluster and a user creates a
`PersistentVolumeClaim` with `storageClassName` unspecified, the
`DefaultStorageClass` admission controller automatically adds the
`storageClassName` field pointing to the default storage class.

Note that if you set the `storageclass.kubernetes.io/is-default-class`
annotation to true on more than one StorageClass in your cluster, and you then
create a `PersistentVolumeClaim` with no `storageClassName` set, Kubernetes
uses the most recently created default StorageClass.

## Topology Awareness

In [Multi-Zone](https://kubernetes.io/docs/setup/best-practices/multiple-zones/) clusters, Pods can be spread across
Zones in a Region. Single-Zone storage backends should be provisioned in the Zones where
Pods are scheduled. This can be accomplished by setting the
[Volume Binding Mode](/concepts/storage/storage-classes.md#volume-binding-mode).
