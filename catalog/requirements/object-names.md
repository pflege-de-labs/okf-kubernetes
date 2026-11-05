---
type: constraint
title: Object name formats and length limits
description: The four Kubernetes name formats (253-char DNS subdomain, 63-char RFC 1123 and RFC 1035 labels, path segment) and which resource kinds require which.
resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/
tags: [kubernetes, requirements, names, validation, limits, dns]
generated: { by: "claude-code/claude-opus-5-5", at: "2026-11-05T00:00:00Z" }
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
sources:
  - id: names-unique
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L23-L35
    title: Object Names and IDs - uniqueness
  - id: names-generate
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L41-L45
    title: Object Names and IDs - generateName
  - id: names-subdomain
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L49-L58
    title: Object Names and IDs - DNS subdomain names
  - id: names-rfc1123
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L60-L75
    title: Object Names and IDs - RFC 1123 label names
  - id: names-rfc1035
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L76-L92
    title: Object Names and IDs - RFC 1035 label names
  - id: names-path
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L94-L98
    title: Object Names and IDs - path segment names
  - id: names-extra
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L116-L118
    title: Object Names and IDs - additional per-type restrictions
  - id: names-uid
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md#L120-L125
    title: Object Names and IDs - UIDs
  - id: pod
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/pods/_index.md#L119-L123
    title: Pods - Pod name
  - id: deployment
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/deployment.md#L1117-L1122
    title: Deployments - Deployment name
  - id: replicaset
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/replicaset.md#L236-L241
    title: ReplicaSet - ReplicaSet name
  - id: replicationcontroller
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/replicationcontroller.md#L125-L130
    title: ReplicationController - name
  - id: job
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/job.md#L194-L199
    title: Jobs - Job name
  - id: cronjob
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/cron-jobs.md#L31-L40
    title: CronJob - CronJob name and 52-character limit
  - id: statefulset
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/statefulset.md#L85-L86
    title: StatefulSets - StatefulSet name
  - id: daemonset
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/controllers/daemonset.md#L59-L60
    title: DaemonSet - DaemonSet name
  - id: service
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/services-networking/service.md#L125-L126
    title: Service - Service name
  - id: namespace
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/namespaces.md#L115-L116
    title: Namespaces - namespace name
  - id: configmap
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/configuration/configmap.md#L52-L53
    title: ConfigMaps - ConfigMap name
  - id: secret
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/configuration/secret.md#L450-L451
    title: Secrets - Secret name
  - id: ingress
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/services-networking/ingress.md#L97-L98
    title: Ingress - Ingress name
  - id: endpointslice
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/services-networking/endpoint-slices.md#L34-L35
    title: EndpointSlices - EndpointSlice name
  - id: networkpolicy
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/administer-cluster/declare-network-policy.md#L90-L91
    title: Declare Network Policy - NetworkPolicy name
  - id: pv
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/storage/persistent-volumes.md#L554-L555
    title: Persistent Volumes - PersistentVolume name
  - id: pvc
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/storage/persistent-volumes.md#L795-L796
    title: Persistent Volumes - PersistentVolumeClaim name
  - id: storageclass
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/storage/dynamic-provisioning.md#L46-L47
    title: Dynamic Volume Provisioning - StorageClass name
  - id: node
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/architecture/nodes.md#L72-L73
    title: Nodes - Node name
  - id: serviceaccount
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/configure-pod-container/configure-service-account.md#L131-L132
    title: Configure Service Accounts - ServiceAccount name
  - id: hpa
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/workloads/autoscaling/horizontal-pod-autoscale.md#L266-L267
    title: Horizontal Pod Autoscaling - HorizontalPodAutoscaler name
  - id: limitrange
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/policy/limit-range.md#L41-L42
    title: Limit Ranges - LimitRange name
  - id: resourcequota
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/policy/resource-quotas.md#L83-L84
    title: Resource Quotas - ResourceQuota name
  - id: priorityclass
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/scheduling-eviction/pod-priority-preemption.md#L52-L53
    title: Pod Priority and Preemption - PriorityClass name
  - id: runtimeclass
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/containers/runtime-class.md#L50-L75
    title: Runtime Class - handler and RuntimeClass name
  - id: crd
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/extend-kubernetes/api-extension/custom-resources.md#L164-L166
    title: Custom Resources - CRD and custom object names
  - id: apiservice
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/tasks/extend-kubernetes/configure-aggregation-layer.md#L339-L340
    title: Configure the Aggregation Layer - APIService name
  - id: container-name
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/core/pod-v1.md#L757-L758
    title: Pod API reference - container name
---

# Object name formats and length limits

A name is unique per API group, resource type, namespace (for namespaced
resources) and name; the API version is not part of the identity, so two
objects of the same resource in one namespace cannot share a name via
different API versions.[^names-unique] Every object also has a cluster-wide
unique UID, which is a UUID (ISO/IEC 9834-8, ITU-T X.667).[^names-uid]

Full text: [Object Names and IDs](/concepts/overview/working-with-objects/names.md).

# Name formats

| Format | Max length | Allowed characters | Start | End |
|---|---|---|---|---|
| DNS subdomain (RFC 1123)[^names-subdomain] | 253 | lowercase alphanumeric, `-`, `.` | alphanumeric | alphanumeric |
| RFC 1123 label ("DNS label")[^names-rfc1123] | 63 | lowercase alphanumeric, `-` | alphabetic | alphanumeric |
| RFC 1035 label[^names-rfc1035] | 63 | lowercase alphanumeric, `-` | alphabetic | alphanumeric |
| Path segment[^names-path] | not stated | anything except `/` and `%`; not `.` or `..` | - | - |

- Although RFC 1123 allows a leading digit, the current Kubernetes
  implementation requires both RFC 1123 and RFC 1035 labels to start with a
  letter. Exception: with the `RelaxedServiceNameValidation` feature gate,
  Service names may start with a digit.[^names-rfc1123][^names-rfc1035]
- Some resource types have additional restrictions on their names.[^names-extra]

# generateName

With `generateName` instead of `name`, the server appends a generated suffix
to the given prefix. A collision returns HTTP 409; since Kubernetes v1.31 the
server makes up to 8 attempts before returning 409.[^names-generate]

# Required format per resource kind

"Subdomain, label recommended" means the name must be a valid DNS subdomain,
but because the name is the basis for Pod hostnames the docs recommend the
stricter DNS label (63 chars) rules.

| Kind | Required name format | Notes |
|---|---|---|
| Pod | DNS subdomain, label recommended[^pod] | |
| Deployment | DNS subdomain, label recommended[^deployment] | |
| ReplicaSet | DNS subdomain, label recommended[^replicaset] | |
| ReplicationController | DNS subdomain, label recommended[^replicationcontroller] | |
| Job | DNS subdomain, label recommended[^job] | Job name is limited to 63 characters (see CronJob)[^cronjob] |
| CronJob | DNS subdomain, label recommended; **max 52 characters**[^cronjob] | The controller appends 11 characters to form the Job name |
| StatefulSet | DNS label[^statefulset] | |
| DaemonSet | DNS subdomain[^daemonset] | |
| Service | RFC 1123 label[^service] | Leading digit only with `RelaxedServiceNameValidation`[^names-rfc1123] |
| Namespace | RFC 1123 DNS label[^namespace] | See [namespace rules](namespaces.md) |
| ConfigMap | DNS subdomain[^configmap] | |
| Secret | DNS subdomain[^secret] | |
| Ingress | DNS subdomain[^ingress] | |
| EndpointSlice | DNS subdomain[^endpointslice] | |
| NetworkPolicy | DNS subdomain[^networkpolicy] | |
| PersistentVolume | DNS subdomain[^pv] | |
| PersistentVolumeClaim | DNS subdomain[^pvc] | |
| StorageClass | DNS subdomain[^storageclass] | |
| Node | DNS subdomain[^node] | |
| ServiceAccount | DNS subdomain[^serviceaccount] | |
| HorizontalPodAutoscaler | DNS subdomain[^hpa] | |
| LimitRange | DNS subdomain[^limitrange] | |
| ResourceQuota | DNS subdomain[^resourcequota] | |
| PriorityClass | DNS subdomain[^priorityclass] | |
| RuntimeClass | DNS subdomain; its `handler` must be a DNS label[^runtimeclass] | |
| CustomResourceDefinition | DNS subdomain, derived from resource name and group[^crd] | Objects of a CRD-defined kind must also be DNS subdomains[^crd] |
| APIService | Path segment[^apiservice] | |
| Container (in a Pod spec) | DNS label, unique within the Pod, immutable[^container-name] | |

Related: [labels and annotations](labels-and-annotations.md),
[size limits](size-limits.md), [ports and ranges](ports-and-ranges.md).

[^names-unique]: Object Names and IDs - uniqueness
[^names-generate]: Object Names and IDs - generateName
[^names-subdomain]: Object Names and IDs - DNS subdomain names
[^names-rfc1123]: Object Names and IDs - RFC 1123 label names
[^names-rfc1035]: Object Names and IDs - RFC 1035 label names
[^names-path]: Object Names and IDs - path segment names
[^names-extra]: Object Names and IDs - additional per-type restrictions
[^names-uid]: Object Names and IDs - UIDs
[^pod]: Pods - Pod name
[^deployment]: Deployments - Deployment name
[^replicaset]: ReplicaSet - ReplicaSet name
[^replicationcontroller]: ReplicationController - name
[^job]: Jobs - Job name
[^cronjob]: CronJob - CronJob name and 52-character limit
[^statefulset]: StatefulSets - StatefulSet name
[^daemonset]: DaemonSet - DaemonSet name
[^service]: Service - Service name
[^namespace]: Namespaces - namespace name
[^configmap]: ConfigMaps - ConfigMap name
[^secret]: Secrets - Secret name
[^ingress]: Ingress - Ingress name
[^endpointslice]: EndpointSlices - EndpointSlice name
[^networkpolicy]: Declare Network Policy - NetworkPolicy name
[^pv]: Persistent Volumes - PersistentVolume name
[^pvc]: Persistent Volumes - PersistentVolumeClaim name
[^storageclass]: Dynamic Volume Provisioning - StorageClass name
[^node]: Nodes - Node name
[^serviceaccount]: Configure Service Accounts - ServiceAccount name
[^hpa]: Horizontal Pod Autoscaling - HorizontalPodAutoscaler name
[^limitrange]: Limit Ranges - LimitRange name
[^resourcequota]: Resource Quotas - ResourceQuota name
[^priorityclass]: Pod Priority and Preemption - PriorityClass name
[^runtimeclass]: Runtime Class - handler and RuntimeClass name
[^crd]: Custom Resources - CRD and custom object names
[^apiservice]: Configure the Aggregation Layer - APIService name
[^container-name]: Pod API reference - container name
