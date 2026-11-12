---
type: task
title: Perform a Rollback on a DaemonSet
description: This page shows how to perform a rollback on a DaemonSet.
resource: https://kubernetes.io/docs/tasks/manage-daemon/rollback-daemon-set/
tags: [kubernetes, tasks, manage-daemon]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/manage-daemon/rollback-daemon-set.md
    title: 'Kubernetes website source: tasks/manage-daemon/rollback-daemon-set.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/manage-daemon/rollback-daemon-set/
    title: Perform a Rollback on a DaemonSet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/manage-daemon/rollback-daemon-set.md
min_kubernetes_server_version: "1.7"
---

# Perform a Rollback on a DaemonSet

This page shows how to perform a rollback on a [DaemonSet](/reference/glossary/daemonset.md).

## Before you begin

You need to have a Kubernetes cluster, and the kubectl command-line tool must
be configured to communicate with your cluster. It is recommended to run this tutorial on a cluster with at least two nodes that are not acting as control plane hosts. If you do not already have a
cluster, you can create one by using
[minikube](https://minikube.sigs.k8s.io/docs/tutorials/multi_node/)
or you can use one of these Kubernetes playgrounds:

* [iximiuz Labs](https://labs.iximiuz.com/playgrounds?category=kubernetes&filter=all)
* [Killercoda](https://killercoda.com/playgrounds/scenario/kubernetes)
* [KodeKloud](https://kodekloud.com/public-playgrounds) Your Kubernetes server must be at or later than version 1.7. To check the version, enter `kubectl version`.

You should already know how to [perform a rolling update on a
 DaemonSet](/tasks/manage-daemon/update-daemon-set.md).

## Performing a rollback on a DaemonSet

### Step 1: Find the DaemonSet revision you want to roll back to

You can skip this step if you only want to roll back to the last revision.

List all revisions of a DaemonSet:

```shell
kubectl rollout history daemonset <daemonset-name>
```

This returns a list of DaemonSet revisions:

```
daemonsets "<daemonset-name>"
REVISION        CHANGE-CAUSE
1               ...
2               ...
...
```

* Change cause is copied from DaemonSet annotation `kubernetes.io/change-cause`
  to its revisions upon creation. You may specify `--record=true` in `kubectl`
  to record the command executed in the change cause annotation.

To see the details of a specific revision:

```shell
kubectl rollout history daemonset <daemonset-name> --revision=1
```

This returns the details of that revision:

```
daemonsets "<daemonset-name>" with revision #1
Pod Template:
Labels:       foo=bar
Containers:
app:
 Image:        ...
 Port:         ...
 Environment:  ...
 Mounts:       ...
Volumes:      ...
```

### Step 2: Roll back to a specific revision

```shell
# Specify the revision number you get from Step 1 in --to-revision
kubectl rollout undo daemonset <daemonset-name> --to-revision=<revision>
```

If it succeeds, the command returns:

```
daemonset "<daemonset-name>" rolled back
```

**Note:**
If `--to-revision` flag is not specified, kubectl picks the most recent revision.

### Step 3: Watch the progress of the DaemonSet rollback

`kubectl rollout undo daemonset` tells the server to start rolling back the
DaemonSet. The real rollback is done asynchronously inside the cluster
[control plane](/reference/glossary/control-plane.md).

To watch the progress of the rollback:

```shell
kubectl rollout status ds/<daemonset-name>
```

When the rollback is complete, the output is similar to:

```
daemonset "<daemonset-name>" successfully rolled out
```

## Understanding DaemonSet revisions

In the previous `kubectl rollout history` step, you got a list of DaemonSet
revisions. Each revision is stored in a resource named ControllerRevision.

To see what is stored in each revision, find the DaemonSet revision raw
resources:

```shell
kubectl get controllerrevision -l <daemonset-selector-key>=<daemonset-selector-value>
```

This returns a list of ControllerRevisions:

```
NAME                               CONTROLLER                     REVISION   AGE
<daemonset-name>-<revision-hash>   DaemonSet/<daemonset-name>     1          1h
<daemonset-name>-<revision-hash>   DaemonSet/<daemonset-name>     2          1h
```

Each ControllerRevision stores the annotations and template of a DaemonSet
revision.

`kubectl rollout undo` takes a specific ControllerRevision and replaces
DaemonSet template with the template stored in the ControllerRevision.
`kubectl rollout undo` is equivalent to updating DaemonSet template to a
previous revision through other commands, such as `kubectl edit` or `kubectl
apply`.

**Note:**
DaemonSet revisions only roll forward. That is to say, after a
rollback completes, the revision number (`.revision` field) of the
ControllerRevision being rolled back to will advance. For example, if you
have revision 1 and 2 in the system, and roll back from revision 2 to revision
1, the ControllerRevision with `.revision: 1` will become `.revision: 3`.

## Troubleshooting

* See [troubleshooting DaemonSet rolling
  update](/tasks/manage-daemon/update-daemon-set.md#troubleshooting).
