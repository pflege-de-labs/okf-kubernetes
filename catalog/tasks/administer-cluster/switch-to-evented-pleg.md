---
type: task
title: Switching from Polling to CRI Event-based Updates to Container Status
description: This page shows how to migrate nodes to use event based updates for container status.
resource: https://kubernetes.io/docs/tasks/administer-cluster/switch-to-evented-pleg/
tags: [kubernetes, tasks, administer-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/administer-cluster/switch-to-evented-pleg.md
    title: 'Kubernetes website source: tasks/administer-cluster/switch-to-evented-pleg.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/administer-cluster/switch-to-evented-pleg/
    title: Switching from Polling to CRI Event-based Updates to Container Status | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/tasks/administer-cluster/switch-to-evented-pleg.md
min_kubernetes_server_version: "1.26"
---

# Switching from Polling to CRI Event-based Updates to Container Status

**Feature state:** Kubernetes v1.26 [alpha] (feature gate `EventedPLEG`, disabled by default)

This page shows how to migrate nodes to use event based updates for container status. The event-based
implementation reduces node resource consumption by the kubelet, compared to the legacy approach
that relies on polling.
You may know this feature as _evented Pod lifecycle event generator (PLEG)_. That's the name used
internally within the Kubernetes project for a key implementation detail.

The polling based approach is referred to as _generic PLEG_.

## Before you begin

* You need to run a version of Kubernetes that provides this feature.
  Kubernetes v1.27 includes beta support for event-based container
  status updates. The feature is beta but is _disabled_ by default
  because it requires support from the container runtime.
* Your Kubernetes server must be at or later than version 1.26. To check the version, enter `kubectl version`.
  If you are running a different version of Kubernetes, check the documentation for that release.
* The container runtime in use must support container lifecycle events.
  The kubelet automatically switches back to the legacy generic PLEG
  mechanism if the container runtime does not announce support for
  container lifecycle events, even if you have this feature gate enabled.

## Why switch to Evented PLEG?

* The _Generic PLEG_ incurs non-negligible overhead due to frequent polling of container statuses.
* This overhead is exacerbated by Kubelet's parallelized polling of container states, thus limiting
  its scalability and causing poor performance and reliability problems.
* The goal of _Evented PLEG_ is to reduce unnecessary work during inactivity
  by replacing periodic polling.

## Switching to Evented PLEG

1. Start the Kubelet with the [feature gate](/reference/command-line-tools-reference/feature-gates.md)
   `EventedPLEG` enabled. You can manage the kubelet feature gates editing the kubelet
   [config file](/tasks/administer-cluster/kubelet-config-file.md) and restarting the kubelet service.
   You need to do this on each node where you are using this feature.

2. Make sure the node is [drained](/tasks/administer-cluster/safely-drain-node.md) before proceeding. 

3. Start the container runtime with the container event generation enabled. 

   
   **Containerd:**

   Version 1.7+
   
   **CRI-O:**

   Version 1.26+

   Check if the CRI-O is already configured to emit CRI events by verifying the configuration,

   ```shell
   crio config | grep enable_pod_events
   ```

   If it is enabled, the output should be similar to the following:

   ```none
   enable_pod_events = true
   ```

   To enable it, start the CRI-O daemon with the flag `--enable-pod-events=true` or
   use a dropin config with the following lines:

   ```toml
   [crio.runtime]
   enable_pod_events = true
   ```
   
   

   Your Kubernetes server must be at or later than version 1.26. To check the version, enter `kubectl version`.

4. Verify that the kubelet is using event-based container stage change monitoring.
   To check, look for the term `EventedPLEG` in the kubelet logs.

   The output should be similar to this:

   ```console
   I0314 11:10:13.909915 1105457 feature_gate.go:249] feature gates: &{map[EventedPLEG:true]}
   ```

   If you have set `--v` to 4 and above, you might see more entries that indicate
   that the kubelet is using event-based container state monitoring.

   ```console
   I0314 11:12:42.009542 1110177 evented.go:238] "Evented PLEG: Generated pod status from the received event" podUID=3b2c6172-b112-447a-ba96-94e7022912dc
   I0314 11:12:44.623326 1110177 evented.go:238] "Evented PLEG: Generated pod status from the received event" podUID=b3fba5ea-a8c5-4b76-8f43-481e17e8ec40
   I0314 11:12:44.714564 1110177 evented.go:238] "Evented PLEG: Generated pod status from the received event" podUID=b3fba5ea-a8c5-4b76-8f43-481e17e8ec40
   ```

## What's next

* Learn more about the design in the Kubernetes Enhancement Proposal (KEP):
  [Kubelet Evented PLEG for Better Performance](https://github.com/kubernetes/enhancements/blob/5b258a990adabc2ffdc9d84581ea6ed696f7ce6c/keps/sig-node/3386-kubelet-evented-pleg/README.md).
