---
type: task
title: Change the Reclaim Policy of a PersistentVolume
description: This page shows how to change the reclaim policy of a Kubernetes PersistentVolume.
resource: https://kubernetes.io/docs/tasks/administer-cluster/change-pv-reclaim-policy/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/administer-cluster/change-pv-reclaim-policy.md
    title: 'Kubernetes website source: tasks/administer-cluster/change-pv-reclaim-policy.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/change-pv-reclaim-policy/
    title: Change the Reclaim Policy of a PersistentVolume | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/administer-cluster/change-pv-reclaim-policy.md
---

# Change the Reclaim Policy of a PersistentVolume

This page shows how to change the reclaim policy of a Kubernetes
PersistentVolume.

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) To check the version, enter `kubectl version`.

## Why change reclaim policy of a PersistentVolume

PersistentVolumes can have various reclaim policies, including "Retain",
"Recycle", and "Delete". For dynamically provisioned PersistentVolumes,
the default reclaim policy is "Delete". This means that a dynamically provisioned
volume is automatically deleted when a user deletes the corresponding
PersistentVolumeClaim. This automatic behavior might be inappropriate if the volume
contains precious data. In that case, it is more appropriate to use the "Retain"
policy. With the "Retain" policy, if a user deletes a PersistentVolumeClaim,
the corresponding PersistentVolume will not be deleted. Instead, it is moved to the
Released phase, where all of its data can be manually recovered.

## Changing the reclaim policy of a PersistentVolume

1. List the PersistentVolumes in your cluster:

   ```shell
   kubectl get pv
   ```

   The output is similar to this:

   ```none
   NAME                                       CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM             STORAGECLASS     REASON    AGE
   pvc-b6efd8da-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Delete          Bound     default/claim1    manual                     10s
   pvc-b95650f8-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Delete          Bound     default/claim2    manual                     6s
   pvc-bb3ca71d-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Delete          Bound     default/claim3    manual                     3s
   ```

   This list also includes the name of the claims that are bound to each volume
   for easier identification of dynamically provisioned volumes.

1. Choose one of your PersistentVolumes and change its reclaim policy:

   ```shell
   kubectl patch pv <your-pv-name> -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'
   ```

   where `<your-pv-name>` is the name of your chosen PersistentVolume.

   **Note:**
   On Windows, you must _double_ quote any JSONPath template that contains spaces (not single
   quote as shown above for bash). This in turn means that you must use a single quote or escaped
   double quote around any literals in the template. For example:

   ```cmd
   kubectl patch pv <your-pv-name> -p "{\"spec\":{\"persistentVolumeReclaimPolicy\":\"Retain\"}}"
   ```
   

1. Verify that your chosen PersistentVolume has the right policy:

   ```shell
   kubectl get pv
   ```

   The output is similar to this:

   ```none
   NAME                                       CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM             STORAGECLASS     REASON    AGE
   pvc-b6efd8da-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Delete          Bound     default/claim1    manual                     40s
   pvc-b95650f8-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Delete          Bound     default/claim2    manual                     36s
   pvc-bb3ca71d-b7b5-11e6-9d58-0ed433a7dd94   4Gi        RWO           Retain          Bound     default/claim3    manual                     33s
   ```

   In the preceding output, you can see that the volume bound to claim
   `default/claim3` has reclaim policy `Retain`. It will not be automatically
   deleted when a user deletes claim `default/claim3`.

## What's next

* Learn more about [PersistentVolumes](/concepts/storage/persistent-volumes.md).
* Learn more about [PersistentVolumeClaims](/concepts/storage/persistent-volumes.md#persistentvolumeclaims).

### References {#reference}

* [PersistentVolume](/reference/kubernetes-api/core/persistent-volume-v1.md)
  * Pay attention to the `.spec.persistentVolumeReclaimPolicy`
    [field](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-v1/#PersistentVolumeSpec)
    of PersistentVolume.
* [PersistentVolumeClaim](/reference/kubernetes-api/core/persistent-volume-claim-v1.md)
