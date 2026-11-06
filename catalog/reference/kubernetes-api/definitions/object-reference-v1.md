---
type: api-reference
title: ObjectReference
description: ObjectReference contains enough information to let you inspect or modify the referred object.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/object-reference-v1/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/object-reference-v1.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/object-reference-v1.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/object-reference-v1/
    title: ObjectReference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/object-reference-v1.md
api_metadata:
  apiVersion: v1
  import: k8s.io/api/core/v1
  kind: ObjectReference
---

# ObjectReference

`apiVersion: v1`

`import "k8s.io/api/core/v1"`

## ObjectReference {#ObjectReference}

ObjectReference contains enough information to let you inspect or modify the referred object.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>apiVersion</code><br/><em>string</em></td>
      <td>API version of the referent.</td>
    </tr>
    <tr>
      <td><code>fieldPath</code><br/><em>string</em></td>
      <td>If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.</td>
    </tr>
    <tr>
      <td><code>kind</code><br/><em>string</em></td>
      <td>Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
    </tr>
    <tr>
      <td><code>name</code><br/><em>string</em></td>
      <td>Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
    </tr>
    <tr>
      <td><code>namespace</code><br/><em>string</em></td>
      <td>Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</td>
    </tr>
    <tr>
      <td><code>resourceVersion</code><br/><em>string</em></td>
      <td>Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency</td>
    </tr>
    <tr>
      <td><code>uid</code><br/><em>string</em></td>
      <td>UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids</td>
    </tr>
  </tbody>
</table>
