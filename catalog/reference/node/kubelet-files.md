---
type: reference
title: Local Files And Paths Used By The Kubelet
description: The kubelet is mostly a stateless process running on a Kubernetes node.
resource: https://kubernetes.io/docs/reference/node/kubelet-files/
tags: [kubernetes, reference, node]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/node/kubelet-files.md
    title: 'Kubernetes website source: reference/node/kubelet-files.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/node/kubelet-files/
    title: Local Files And Paths Used By The Kubelet | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/node/kubelet-files.md
---

# Local Files And Paths Used By The Kubelet

The [kubelet](/reference/glossary/kubelet.md) is mostly a stateless
process running on a Kubernetes [node](/reference/glossary/node.md).
This document outlines files that kubelet reads and writes.

**Note:**

This document is for informational purposes and not describing any guaranteed behaviors or APIs.
It lists resources used by the kubelet, which is an implementation detail and subject to change at any release.

The kubelet typically uses the [control plane](/reference/glossary/control-plane.md) as
the source of truth on what needs to run on the Node, and the
[container runtime](/reference/glossary/container-runtime.md) to retrieve
the current state of containers. So long as you provide a _kubeconfig_ (API client configuration)
to the kubelet, the kubelet does connect to your control plane; otherwise the node operates in
_standalone mode_.

On Linux nodes, the kubelet also relies on reading cgroups and various system files to collect metrics.

On Windows nodes, the kubelet collects metrics via a different mechanism that does not rely on
paths.

There are also a few other files that are used by the kubelet as well,
as kubelet communicates using local Unix-domain sockets. Some are sockets that the
kubelet listens on, and for other sockets the kubelet discovers them and then connects
as a client.

**Note:**

This page lists paths as Linux paths, which map to the Windows paths by adding a root disk
`C:\` in place of `/` (unless specified otherwise).
For example, `/var/lib/kubelet/device-plugins` maps to `C:\var\lib\kubelet\device-plugins`.

## Configuration

### Kubelet configuration files

The path to the kubelet configuration file can be configured
using the command line argument `--config`. The kubelet also supports
[drop-in configuration files](/tasks/administer-cluster/kubelet-config-file.md#kubelet-conf-d)
to enhance configuration.

### Certificates

Certificates and private keys are typically located at `/var/lib/kubelet/pki`,
but can be configured using the `--cert-dir` kubelet command line argument.
Names of certificate files are also configurable.

### Manifests

Manifests for static pods are typically located in `/etc/kubernetes/manifests`.
Location can be configured using the `staticPodPath` kubelet configuration option.

### Systemd unit settings

When kubelet is running as a systemd unit, some kubelet configuration may be declared
in systemd unit settings file. Typically it includes:

- command line arguments to [run kubelet](/reference/command-line-tools-reference/kubelet.md)
- environment variables, used by kubelet or [configuring golang runtime](https://pkg.go.dev/runtime#hdr-Environment_Variables)

## State

### Checkpoint files for resource managers {#resource-managers-state}

All resource managers keep the mapping of Pods to allocated resources in state files.
State files are located in the kubelet's base directory, also termed the _root directory_
(but not the same as `/`, the node root directory). You can configure the base directory
for the kubelet
using the kubelet command line argument `--root-dir`.

Names of files:

- `memory_manager_state` for the [Memory Manager](/tasks/administer-cluster/memory-manager.md)
- `cpu_manager_state` for the [CPU Manager](/tasks/administer-cluster/cpu-management-policies.md)
- `dra_manager_state` for [DRA](/concepts/resource-management/dynamic-resource-allocation.md)

### Checkpoint file for device manager {#device-manager-state}

Device manager creates checkpoints in the same directory with socket files: `/var/lib/kubelet/device-plugins/`.
This path is hardcoded and is not relative to the kubelet root directory.
The name of a checkpoint file is `kubelet_internal_checkpoint` for
[Device Manager](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#device-plugin-integration-with-the-topology-manager)

### Pod resource checkpoints

**Feature state:** Kubernetes v1.35 [stable] (feature gate `InPlacePodVerticalScaling`, enabled by default)

If a node has enabled the `InPlacePodVerticalScaling`[feature gate](/reference/command-line-tools-reference/feature-gates.md),
the kubelet stores a local record of _allocated_ and _actuated_ Pod resources.
See [Resize CPU and Memory Resources assigned to Containers](/tasks/configure-pod-container/resize-container-resources.md)
for more details on how these records are used.

Names of files:

- `allocated_pods_state` records the resources allocated to each pod running on the node
- `actuated_pods_state` records the resources that have been accepted by the runtime
  for each pod running on the node

The files are located within the kubelet base directory
(`/var/lib/kubelet` by default on Linux; configurable using `--root-dir`).

### Container runtime

Kubelet communicates with the container runtime using socket configured via the
configuration parameters:

- `containerRuntimeEndpoint` for runtime operations
- `imageServiceEndpoint` for image management operations

The actual values of those endpoints depend on the container runtime being used.

### Device plugins

The kubelet exposes a socket at the path `/var/lib/kubelet/device-plugins/kubelet.sock` for
various [Device Plugins to register](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#device-plugin-implementation).

When a device plugin registers itself, it provides its socket path for the kubelet to connect.

The device plugin socket must be in the directory `/var/lib/kubelet/device-plugins/`.
This path is hardcoded and is not relative to the kubelet base directory (root directory).
On Linux, this path is always `/var/lib/kubelet/device-plugins`.

### Pod resources API

[Pod Resources API](/concepts/extend-kubernetes/compute-storage-net/device-plugins.md#monitoring-device-plugin-resources)
will be exposed at the path `pod-resources` within the kubelet base directory (root directory).
On a typical Linux node, this means `/var/lib/kubelet/pod-resources`.

### DRA, CSI, and Device plugins

The kubelet looks for socket files created by device plugins managed via [DRA](/concepts/resource-management/dynamic-resource-allocation.md),
device manager, or storage plugins, and then attempts to connect
to these sockets. The directory that the kubelet looks in is `plugins_registry` within the kubelet base
directory, so on a typical Linux node this means `/var/lib/kubelet/plugins_registry`.

Note, for the device plugins there are two alternative registration mechanisms
Only one should be used for a given plugin.

The types of plugins that can place socket files into that directory are:

- CSI plugins
- DRA plugins
- Device Manager plugins

(typically `/var/lib/kubelet/plugins_registry`).

### Graceful node shutdown

**Feature state:** Kubernetes v1.21 [beta] (feature gate `GracefulNodeShutdown`, enabled by default)

[Graceful node shutdown](/concepts/cluster-administration/node-shutdown.md#graceful-node-shutdown)
stores state locally at `/var/lib/kubelet/graceful_node_shutdown_state`.

### Image Pull Records

**Feature state:** Kubernetes v1.35 [beta] (feature gate `KubeletEnsureSecretPulledImages`, enabled by default)

The kubelet stores records of attempted and successful image pulls, and uses it
to verify that the image was previously successfully pulled with the same credentials.

These records are cached as files in the `image_registry` directory within
the kubelet base directory. On a typical Linux node, this means `/var/lib/kubelet/image_manager`.
There are two subdirectories to `image_manager`:
* `pulling` - stores records about images the Kubelet is attempting to pull.
* `pulled` - stores records about images that were successfully pulled by the Kubelet, 
  along with metadata about the credentials used for the pulls.

See [Ensure Image Pull Credential Verification](/concepts/containers/images.md#ensureimagepullcredentialverification)
for details.

## Security profiles & configuration

### Seccomp

Seccomp profile files referenced from Pods should be placed in `/var/lib/kubelet/seccomp`.
See the [seccomp reference](/reference/node/seccomp.md) for details.

### AppArmor

The kubelet does not load or refer to AppArmor profiles by a Kubernetes-specific path.
AppArmor profiles are loaded via the node operating system rather than referenced by their path.

## Locking

**Feature state:** Kubernetes v1.2 [alpha]

A lock file for the kubelet; typically `/var/run/kubelet.lock`. The kubelet uses this to ensure
that two different kubelets don't try to run in conflict with each other.
You can configure the path to the lock file using the `--lock-file` kubelet command line argument.

If two kubelets on the same node use a different value for the lock file path, they will not be able to
detect a conflict when both are running.

## What's next

- Learn about the kubelet [command line arguments](/reference/command-line-tools-reference/kubelet.md).
- Review the [Kubelet Configuration (v1beta1) reference](/reference/config-api/kubelet-config.v1beta1.md)
