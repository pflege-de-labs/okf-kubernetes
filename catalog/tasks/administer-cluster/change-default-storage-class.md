---
type: task
title: Change the default StorageClass
description: This page shows how to change the default Storage Class that is used to provision volumes for PersistentVolumeClaims that have no special requirements.
resource: https://kubernetes.io/docs/tasks/administer-cluster/change-default-storage-class/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/change-default-storage-class.md
    title: 'Kubernetes website source: tasks/administer-cluster/change-default-storage-class.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/change-default-storage-class/
    title: Change the default StorageClass | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/change-default-storage-class.md
---

# Change the default StorageClass

This page shows how to change the default Storage Class that is used to
provision volumes for PersistentVolumeClaims that have no special requirements.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Why change the default storage class?

Depending on the installation method, your Kubernetes cluster may be deployed with
an existing StorageClass that is marked as default. This default StorageClass
is then used to dynamically provision storage for PersistentVolumeClaims
that do not require any specific storage class. See
[PersistentVolumeClaim documentation](/concepts/storage/persistent-volumes.md#persistentvolumeclaims)
for details.

The pre-installed default StorageClass may not fit well with your expected workload;
for example, it might provision storage that is too expensive. If this is the case,
you can either change the default StorageClass or disable it completely to avoid
dynamic provisioning of storage.

Deleting the default StorageClass may not work, as it may be re-created
automatically by the addon manager running in your cluster. Please consult the docs for your installation
for details about addon manager and how to disable individual addons.

## Changing the default StorageClass

1. List the StorageClasses in your cluster:

   ```bash
   kubectl get storageclass
   ```

   The output is similar to this:

   ```bash
   NAME                 PROVISIONER               AGE
   standard (default)   kubernetes.io/gce-pd      1d
   gold                 kubernetes.io/gce-pd      1d
   ```

   The default StorageClass is marked by `(default)`.

1. Mark the default StorageClass as non-default:

   The default StorageClass has an annotation
   `storageclass.kubernetes.io/is-default-class` set to `true`. Any other value
   or absence of the annotation is interpreted as `false`.

   To mark a StorageClass as non-default, you need to change its value to `false`:

   ```bash
   kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'
   ```

   where `standard` is the name of your chosen StorageClass.

1. Mark a StorageClass as default:

   Similar to the previous step, you need to add/set the annotation
   `storageclass.kubernetes.io/is-default-class=true`.

   ```bash
   kubectl patch storageclass gold -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
   ```

   Please note you can have multiple `StorageClass` marked as default. If more 
   than one `StorageClass` is marked as default, a `PersistentVolumeClaim` without 
   an explicitly defined `storageClassName` will be created using the most recently 
   created default `StorageClass`.
   When a `PersistentVolumeClaim` is created with a specified `volumeName`, it remains 
   in a pending state if the static volume's `storageClassName` does not match the 
   `StorageClass` on the `PersistentVolumeClaim`.

1. Verify that your chosen StorageClass is default:

   ```bash
   kubectl get storageclass
   ```

   The output is similar to this:

   ```bash
   NAME             PROVISIONER               AGE
   standard         kubernetes.io/gce-pd      1d
   gold (default)   kubernetes.io/gce-pd      1d
   ```

## What's next

* Learn more about [PersistentVolumes](/concepts/storage/persistent-volumes.md).
