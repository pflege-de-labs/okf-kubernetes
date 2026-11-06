---
type: concept
title: Field Selectors
description: _Field selectors_ let you select Kubernetes objects based on the value of one or more resource fields.
resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/
tags: [kubernetes, concepts, overview]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/field-selectors.md
    title: 'Kubernetes website source: concepts/overview/working-with-objects/field-selectors.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/
    title: Field Selectors | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/overview/working-with-objects/field-selectors.md
---

# Field Selectors

_Field selectors_ let you select Kubernetes [objects](/reference/glossary/object.md) based on the
value of one or more resource fields. Here are some examples of field selector queries:

* `metadata.name=my-service`
* `metadata.namespace!=default`
* `status.phase=Pending`

This `kubectl` command selects all Pods for which the value of the [`status.phase`](/concepts/workloads/pods/pod-lifecycle.md#pod-phase) field is `Running`:

```shell
kubectl get pods --field-selector status.phase=Running
```

**Note:**
Field selectors are essentially resource *filters*. By default, no selectors/filters are applied, meaning that all resources of the specified type are selected. This makes the `kubectl` queries `kubectl get pods` and `kubectl get pods --field-selector ""` equivalent.

## Supported fields

Supported field selectors vary by Kubernetes resource type. All resource types support the `metadata.name` and `metadata.namespace` fields. Using unsupported field selectors produces an error. For example:

```shell
kubectl get ingress --field-selector foo.bar=baz
```
```
Error from server (BadRequest): Unable to find "ingresses" that match label selector "", field selector "foo.bar=baz": "foo.bar" is not a known field selector: only "metadata.name", "metadata.namespace"
```

### List of supported fields

| Kind                      | Fields                                                                                                                                                                                                                                                          |
| ------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Pod                       | `spec.nodeName`<br>`spec.restartPolicy`<br>`spec.schedulerName`<br>`spec.serviceAccountName`<br>`spec.hostNetwork`<br>`status.phase`<br>`status.podIP`<br>`status.podIPs`<br>`status.nominatedNodeName`                                                                            |
| Event                     | `involvedObject.kind`<br>`involvedObject.namespace`<br>`involvedObject.name`<br>`involvedObject.uid`<br>`involvedObject.apiVersion`<br>`involvedObject.resourceVersion`<br>`involvedObject.fieldPath`<br>`reason`<br>`reportingComponent`<br>`source`<br>`type` |
| Secret                    | `type`                                                                                                                                                                                                                                                          |
| Service                   | `spec.clusterIP`<br>`spec.type`                                                                                                                                                                                                                                 |
| Namespace                 | `status.phase`                                                                                                                                                                                                                                                  |
| ReplicaSet                | `status.replicas`                                                                                                                                                                                                                                               |
| ReplicationController     | `status.replicas`                                                                                                                                                                                                                                               |
| Job                       | `status.successful`                                                                                                                                                                                                                                             |
| Node                      | `spec.unschedulable`                                                                                                                                                                                                                                            |
| CertificateSigningRequest | `spec.signerName`                                                                                                                                                                                                                                               |

### Custom resources fields

All custom resource types support the `metadata.name` and `metadata.namespace` fields.

Additionally, the `spec.versions[*].selectableFields` field of a [CustomResourceDefinition](/reference/glossary/customresourcedefinition.md)
declares which other fields in a custom resource may be used in field selectors. See [selectable fields for custom resources](/tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#crd-selectable-fields)
for more information about how to use field selectors with CustomResourceDefinitions.

## Supported operators

You can use the `=`, `==`, and `!=` operators with field selectors (`=` and `==` mean the same thing). This `kubectl` command, for example, selects all Kubernetes Services that aren't in the `default` namespace:

```shell
kubectl get services  --all-namespaces --field-selector metadata.namespace!=default
```
**Note:**
[Set-based operators](/concepts/overview/working-with-objects/labels.md#set-based-requirement)
(`in`, `notin`, `exists`) are not supported for field selectors. 

## Chained selectors

As with [label](/concepts/overview/working-with-objects/labels.md) and other selectors, field selectors can be chained together as a comma-separated list. This `kubectl` command selects all Pods for which the `status.phase` does not equal `Running` and the `spec.restartPolicy` field equals `Always`:

```shell
kubectl get pods --field-selector=status.phase!=Running,spec.restartPolicy=Always
```

## Multiple resource types

You can use field selectors across multiple resource types. This `kubectl` command selects all Statefulsets and Services that are not in the `default` namespace:

```shell
kubectl get statefulsets,services --all-namespaces --field-selector metadata.namespace!=default
```
