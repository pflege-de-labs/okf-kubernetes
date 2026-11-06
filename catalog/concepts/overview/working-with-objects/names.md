---
type: concept
title: Object Names and IDs
description: Each object in your cluster has a _Name_ that is unique for that type of resource.
resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/
tags: [kubernetes, concepts, overview]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/names.md
    title: 'Kubernetes website source: concepts/overview/working-with-objects/names.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/
    title: Object Names and IDs | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/concepts/overview/working-with-objects/names.md
---

# Object Names and IDs

Each [object](/reference/glossary/object.md) in your cluster has a [_Name_](#names) that is unique for that type of resource.
Every Kubernetes object also has a [_UID_](#uids) that is unique across your whole cluster.

For example, you can only have one Pod named `myapp-1234` within the same [namespace](/concepts/overview/working-with-objects/namespaces.md), but you can have one Pod and one Deployment that are each named `myapp-1234`.

For non-unique user-provided attributes, Kubernetes provides [labels](/concepts/overview/working-with-objects/labels.md) and [annotations](/concepts/overview/working-with-objects/annotations.md).

## Names

A client-provided string that refers to an object in a [resource](/reference/glossary/api-resource.md)
URL, such as `/api/v1/pods/some-name`.

 

Only one object of a given kind can have a given name at a time. However, if you delete the object, you can make a new object with the same name.

Names must be unique across all [API versions](/concepts/overview/kubernetes-api.md#api-groups-and-versioning) of the same resource. 

Kubernetes uniquely identifies objects using a combination of four attributes:
* **API group** (e.g., `apps`)
* **Resource type** (e.g., `deployments`)
* **Namespace** (for namespaced resources)
* **Name**

While you can access a resource through different API versions (such as `v1` or `v1beta1`), the version is simply a different representation of the same underlying object. Because the version is not part of the unique identification, you cannot create two objects with the same name and resource type in the same namespace by using different API versions.

**Note:**
In cases when objects represent a physical entity, like a Node representing a physical host, when the host is re-created under the same name without deleting and re-creating the Node, Kubernetes treats the new host as the old one, which may lead to inconsistencies.

The server may generate a name when `generateName` is provided instead of `name` in a resource create request.
When `generateName` is used, the provided value is used as a name prefix, which server appends a generated suffix
to. Even though the name is generated, it may conflict with existing names resulting in an HTTP 409 response. This
became far less likely to happen in Kubernetes v1.31 and later, since the server will make up to 8 attempts to generate a
unique name before returning an HTTP 409 response.

Below are four types of commonly used name constraints for resources.

### DNS Subdomain Names

Most resource types require a name that can be used as a DNS subdomain name
as defined in [RFC 1123](https://tools.ietf.org/html/rfc1123).
This means the name must:

- contain no more than 253 characters
- contain only lowercase alphanumeric characters, '-' or '.'
- start with an alphanumeric character
- end with an alphanumeric character

### RFC 1123 Label Names {#dns-label-names}

Some resource types require their names to follow the DNS
label standard as defined in [RFC 1123](https://tools.ietf.org/html/rfc1123).
This means the name must:

- contain at most 63 characters
- contain only lowercase alphanumeric characters or '-'
- start with an alphabetic character
- end with an alphanumeric character

**Note:**
When the `RelaxedServiceNameValidation` feature gate is enabled,
Service object names are allowed to start with a digit.

### RFC 1035 Label Names

Some resource types require their names to follow the DNS
label standard as defined in [RFC 1035](https://tools.ietf.org/html/rfc1035).
This means the name must:

- contain at most 63 characters
- contain only lowercase alphanumeric characters or '-'
- start with an alphabetic character
- end with an alphanumeric character

**Note:**
While RFC 1123 technically allows labels to start with digits, the current
Kubernetes implementation requires both RFC 1035 and RFC 1123 labels to start
with an alphabetic character. The exception is when the `RelaxedServiceNameValidation`
feature gate is enabled for Service objects, which allows Service names to start with digits.

### Path Segment Names

Some resource types require their names to be able to be safely encoded as a
path segment. In other words, the name may not be "." or ".." and the name may
not contain "/" or "%".

Here's an example manifest for a Pod named `nginx-demo`.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-demo
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
```

**Note:**
Some resource types have additional restrictions on their names.

## UIDs

A Kubernetes systems-generated string to uniquely identify objects.

 

Every object created over the whole lifetime of a Kubernetes cluster has a distinct UID. It is intended to distinguish between historical occurrences of similar entities.

Kubernetes UIDs are universally unique identifiers (also known as UUIDs).
UUIDs are standardized as ISO/IEC 9834-8 and as ITU-T X.667.

## What's next

* Read about [labels](/concepts/overview/working-with-objects/labels.md) and [annotations](/concepts/overview/working-with-objects/annotations.md) in Kubernetes.
* See the [Identifiers and Names in Kubernetes](https://git.k8s.io/design-proposals-archive/architecture/identifiers.md) design document.
