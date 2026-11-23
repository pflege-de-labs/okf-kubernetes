---
type: concept
title: Managing Workloads
description: Workload management in Kubernetes covers how you organize, deploy, update, scale, and maintain your applications.
resource: https://kubernetes.io/docs/concepts/workloads/management/
tags: [kubernetes, concepts, workloads]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/concepts/workloads/management.md
    title: 'Kubernetes website source: concepts/workloads/management.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/workloads/management/
    title: Managing Workloads | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/concepts/workloads/management.md
---

# Managing Workloads

Workload management in Kubernetes covers how you organize, deploy, update,
scale, and maintain your applications. After you deploy an application and
expose it through a Service, you need strategies for keeping it running
reliably: updating to new versions without disrupting users, scaling to
meet changing demand, and recovering from failures. This page explains
the key concepts and strategies behind these activities, and links to
task pages for step-by-step procedures.

## Organizing resource configurations

Most applications require multiple Kubernetes resources working together.
A typical web application might need a Deployment, a Service, a ConfigMap,
a PersistentVolumeClaim, and a Gateway (using the Gateway API add-on). Organizing these resources effectively makes
your applications easier to deploy, update, and troubleshoot.

Kubernetes supports several approaches to resource organization:

- **Single-file grouping**: combine related resources in one (YAML) _manifest_,
  separated by `---`. You can define resources in any order.
- **Directory-based organization**: group all manifests for an application
  in the same directory, organized by microservice or application tier.
- **URL-based configuration**: deploy directly from manifests hosted in
  source control systems.

For details on `kubectl` commands that support these approaches, see
[Command line tool (kubectl)](/reference/kubectl.md).

### External tools

Several tools extend Kubernetes object management beyond what `kubectl`
provides natively.

#### Helm {#external-tool-helm}

> Items on this page refer to third party products or projects that are not part of Kubernetes itself.

[Helm](https://helm.sh/) is a tool for managing packages of pre-configured
Kubernetes resources. These packages are known as _Helm charts_.

#### Kustomize {#external-tool-kustomize}

[Kustomize](https://kustomize.io/) traverses a Kubernetes manifest to add,
remove, or update configuration options. It is available both as a standalone
binary and as a
[native feature](/tasks/manage-kubernetes-objects/kustomization.md) of
`kubectl`.

## Update strategies

When you release a new version of your application, how Kubernetes transitions
from the old version to the new one determines whether users experience
downtime and how much risk the update carries. Kubernetes supports several
update strategies, each suited to different scenarios.

### Rolling updates

A rolling update gradually replaces old Pods with new ones. At any point
during the update, a mix of old and new Pods serves traffic, so the
application remains available throughout. This is the default strategy for
Deployments.

Two parameters control the pace of a rolling update:

- **`maxUnavailable`**: the maximum number of Pods that can be unavailable
  during the update (default: 25%).
- **`maxSurge`**: the maximum number of extra Pods that can exist above the
  desired count during the update (default: 25%).

Rolling updates also support pausing and resuming, which allows you to
inspect a partial update or batch multiple changes into a single rollout.
If a new version introduces problems, you can roll back to a previous
revision.

For step-by-step instructions, see
[Update a Deployment Without Downtime](/tasks/run-application/update-deployment-rolling.md).
For configuration details, see the
[Deployment](/concepts/workloads/controllers/deployment.md#strategy)
concept page.

### Recreate

The Recreate strategy terminates all existing Pods before creating new ones.
This causes a brief period of downtime but ensures that only one version of
the application runs at a time.

Use the Recreate strategy when:

- Your application cannot tolerate running two versions simultaneously
  (for example, due to database schema incompatibilities).
- You need a clean switchover with no overlap between old and new Pods.

For details, see the
[Deployment strategy](/concepts/workloads/controllers/deployment.md#strategy)
documentation.

You can also apply a Recreate strategy for a StatefulSet. Because these workloads
(or workload components) very often do manage some durable state, you need to
add your own mechanism to ensure that the application remains durable even if
you replace all of the Pods.

### Canary deployments

A canary deployment runs a new version of your application alongside the
existing stable version, routing only a small fraction of traffic to the new
version. This allows you to monitor the new version under real production
conditions before committing to a full rollout.

The canary pattern in Kubernetes uses labels to differentiate between stable
and canary Pods. Both Deployments share a common label (for example,
`app: my-app`) that allows a single Service to route traffic to both sets
of Pods. A `track` label (`stable` or `canary`) distinguishes the two
releases. By adjusting the replica counts, you control the traffic ratio
between stable and canary Pods.

For a hands-on walkthrough, see
[Deploy a Release Using a Canary Deployment](https://kubernetes.io/docs/tutorials/stateless-application/canary-deployment/).

If you want to use canarying with a StatefulSet, you will typically need to add your
own mechanism to do that. One detail that's relevant is that a StatefulSet doesn't
have to have an ordinal (Pod number) that starts at zero.
For example, you have a main StatefulSet with starting ordinal zero and 7 replicas (pod number 0 to pod number 6), and a canary StatefulSet with starting ordinal 7 (pod number 7). By managing how much traffic or work the canary Pod picks up,
you can gradually increase your confidence that the new code or configuration is
correct.

Fully developing this pattern for a stateful app is outside the scope of core Kubernetes.

### Comparing update strategies

| Strategy | Downtime | Risk level | Best for |
|----------|----------|------------|----------|
| Rolling update | None | Low — gradual rollout | Most workloads |
| Recreate | Yes | Low — clean switchover | Incompatible versions |
| Canary | None | Lowest — small traffic percentage | High-risk changes requiring production validation |

## Scaling workloads

Scaling adjusts the number of Pod replicas running your application to match
demand. Kubernetes supports two approaches:

- **Manual scaling**: you set the replica count directly. This works well for
  predictable, scheduled, or one-off capacity changes where you know the
  required number of replicas in advance.
- **Automatic scaling**: a HorizontalPodAutoscaler (HPA) monitors metrics
  such as CPU utilization, memory usage, or custom metrics, and adjusts the
  replica count automatically. This suits workloads with variable or
  unpredictable demand.

**Caution:**
If a HorizontalPodAutoscaler manages a Deployment or StatefulSet, do not set
the `replicas` field
manually. The HPA continuously reconciles the replica count and overrides
any manual changes.

For step-by-step instructions on manual scaling, see
[Scale a Deployment Manually](/tasks/run-application/scale-deployment.md).
For automatic scaling, see the
[HorizontalPodAutoscaler walkthrough](/tasks/run-application/horizontal-pod-autoscale-walkthrough.md).

## Updating resources in place

After deploying resources, you often need to make changes without replacing
the entire object. Kubernetes provides several approaches for in-place
updates, each suited to different workflows:

- **Declarative management** (`kubectl apply`): compares your configuration
  file against the live state and applies only the differences. This approach
  works well when you maintain configuration as code in source control. For
  details, see
  [server-side apply](/reference/using-api/server-side-apply.md).
- **Interactive editing** (`kubectl edit`): opens the live resource in a text
  editor, allowing you to make changes directly. Useful for one-off
  adjustments or exploration.
- **Patch operations** (`kubectl patch`): updates specific fields using JSON
  patch, JSON merge patch, or strategic merge patch. Best for scripted or
  automated changes. For details, see
  [Update API Objects in Place Using kubectl patch](/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch.md).

## Disruptive updates

Some resource fields are immutable after creation (for example, a Pod's
container name or a Service's `clusterIP`). When you need to change an
immutable field, you must delete and recreate the resource. The
`kubectl replace --force` command performs this operation in a single step.

**Caution:**
Disruptive updates cause downtime because Kubernetes deletes the existing
resource before creating the replacement. Use this approach only when
in-place updates are not possible.

## Cluster administrator considerations

As a cluster administrator, you can set guardrails that influence how
application teams manage their workloads:

- [**ResourceQuotas**](/concepts/policy/resource-quotas.md): limit the
  total compute resources a namespace can consume, preventing any single team
  from overusing cluster capacity.
- [**PodDisruptionBudgets**](/concepts/workloads/pods/disruptions.md):
  ensure a minimum number of Pods remain available during voluntary
  disruptions such as node maintenance or Deployment rollouts. See also
  [Specifying a PodDisruptionBudget](/tasks/run-application/configure-pdb.md).
- [**Topology spread constraints**](/concepts/scheduling-eviction/topology-spread-constraints.md):
  encourage or require Pods to spread across failure domains (zones, nodes)
  for higher availability.
- [**Taints and tolerations**](/concepts/scheduling-eviction/taint-and-toleration.md):
  control which Pods can schedule onto specific nodes, useful for dedicating
  hardware or isolating workloads.

## What's next

- [Scale a Deployment Manually](/tasks/run-application/scale-deployment.md)
- [Update a Deployment Without Downtime](/tasks/run-application/update-deployment-rolling.md)
- [Deploy a Release Using a Canary Deployment](https://kubernetes.io/docs/tutorials/stateless-application/canary-deployment/)
- [Horizontal Pod Autoscaling](/tasks/run-application/horizontal-pod-autoscale-walkthrough.md)
- [Deployments](/concepts/workloads/controllers/deployment.md)
- [Application introspection and debugging](/tasks/debug/debug-application/debug-running-pod.md)
