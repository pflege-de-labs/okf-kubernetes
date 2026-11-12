---
type: concept
title: The kubectl command-line tool
description: kubectl is the primary command-line tool for communicating with a Kubernetes cluster. This page provides an overview of kubectl and its role in the Kubernetes ecosystem.
resource: https://kubernetes.io/docs/concepts/overview/kubectl/
tags: [kubernetes, concepts, overview]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/kubectl.md
    title: 'Kubernetes website source: concepts/overview/kubectl.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/overview/kubectl/
    title: The kubectl command-line tool | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/overview/kubectl.md
---

# The kubectl command-line tool

Kubernetes provides a command line tool for communicating with a Kubernetes cluster's
[control plane](/reference/glossary/control-plane.md),
using the Kubernetes API.

The `kubectl` tool communicates with your cluster through the [Kubernetes API](/concepts/overview/kubernetes-api.md).
For configuration, `kubectl` looks for a file named `config` in the `$HOME/.kube` directory.
You can specify other [kubeconfig](/concepts/configuration/organize-cluster-access-kubeconfig.md)
files by setting the `KUBECONFIG` environment variable or by setting the
[`--kubeconfig`](/concepts/configuration/organize-cluster-access-kubeconfig.md) flag.

## Role of kubectl

The `kubectl` tool is the primary interface for creating, inspecting, updating, and deleting Kubernetes objects.
It complements the [Kubernetes Components](/concepts/overview/components.md) that run inside your cluster
and the [Kubernetes API](/concepts/overview/kubernetes-api.md) that those components implement.
Whether you run `kubectl` from your laptop or from a Pod inside the cluster, it sends requests to the API server.
Other clients, such as [client libraries](/reference/using-api/client-libraries.md) and web dashboards
like [Headlamp](https://headlamp.dev/), also communicate through the same API.

## How kubectl works

The `kubectl` tool connects to the API server and authenticates using the cluster, user, and context defined in your
[kubeconfig](/concepts/configuration/organize-cluster-access-kubeconfig.md) file.
When you run `kubectl` from outside a cluster, it uses the kubeconfig file to find the API server address and credentials.
When `kubectl` runs inside a Pod (for example, in a CI/CD pipeline), it can use in-cluster authentication
based on the ServiceAccount token mounted in the Pod.

When you run a command, `kubectl` translates your intent into one or more HTTP requests to the
[Kubernetes API](/concepts/overview/kubernetes-api.md). The API server validates each request,
applies it to the cluster state stored in [etcd](/reference/glossary/etcd.md), and
returns the result. This means every `kubectl` action, whether creating a Deployment or reading logs,
follows the same API-driven path.

Because your kubeconfig can define multiple clusters, users, and contexts, you can use `kubectl` to
switch between clusters without reconfiguring your environment. Run `kubectl config use-context` to
change the active context.

## What you can do with kubectl

The `kubectl` tool supports many operations, which fall into these broad categories:

* **Manage resources** – Create, update, and delete objects such as Pods, Deployments, and Services.
  Use `kubectl apply` for declarative management from configuration files.
* **Inspect cluster state** – List and describe objects, view events, and check resource usage.
* **Debug** – View logs from containers, execute commands inside a running container, or port-forward to a Pod.
* **Cluster operations** – Drain nodes for maintenance, cordon nodes to prevent new workloads, and manage cluster configuration.
* **Script and automate** – Format output as JSON, YAML, or custom columns using [JSONPath](/reference/kubectl/jsonpath.md) for use in scripts and pipelines.

For syntax, command reference, and examples, see the [kubectl reference documentation](/reference/kubectl.md).

## Declarative vs imperative

For production workloads, prefer [declarative object management](/concepts/overview/working-with-objects/object-management.md)
using `kubectl apply` with version-controlled configuration files.
Declarative management helps you track changes, collaborate, and integrate with GitOps workflows.
Imperative commands (such as `kubectl create` or `kubectl run`) are useful for development and experimentation,
but are harder to reproduce and audit.

## Extending kubectl with plugins

You can extend `kubectl` with [plugins](/tasks/extend-kubectl/kubectl-plugins.md) that add new
sub-commands. Plugins are standalone binaries that follow the `kubectl-<plugin-name>` naming convention.
The Kubernetes community maintains many plugins, and you can manage them with the
[Krew](https://krew.sigs.k8s.io/) plugin manager.

## Version compatibility

The `kubectl` tool supports a version skew of plus-or-minus one minor version relative to the cluster's
control plane. For example, `kubectl` v1.32 works with control planes at v1.31, v1.32, and v1.33.
Using a compatible version avoids unexpected behavior. See the
[version skew policy](/releases/version-skew-policy.md) for details.

## What's next

* Read the [kubectl reference](/reference/kubectl.md) for syntax and command details.
* [Install kubectl](/tasks/tools.md#kubectl) on your machine.
* Learn about [The Kubernetes API](/concepts/overview/kubernetes-api.md) that `kubectl` uses.
* Review [Kubernetes Components](/concepts/overview/components.md) that make up a cluster.
* Explore [Object Management](/concepts/overview/working-with-objects/object-management.md) and declarative configuration.
* Check the [version skew policy](/releases/version-skew-policy.md) for supported version combinations.
