---
type: glossary-term
title: Finalizer
description: A namespaced key that tells Kubernetes to wait until specific conditions are met before it fully deletes an object marked for deletion.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-finalizer
tags: [kubernetes, reference, glossary, fundamental, operation]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/glossary/finalizer.md
    title: 'Kubernetes website source: reference/glossary/finalizer.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-finalizer
    title: Finalizer | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/glossary/finalizer.md
---

# Finalizer

Finalizers are namespaced keys that tell Kubernetes to wait until specific
conditions are met before it fully deletes [resources](/reference/glossary/api-resource.md)
that are marked for deletion.
Finalizers alert [controllers](/reference/glossary/controller.md)
to clean up resources the deleted object owned.

When you tell Kubernetes to delete an object that has finalizers specified for
it, the Kubernetes API marks the object for deletion by populating `.metadata.deletionTimestamp`,
and returns a `202` status code (HTTP "Accepted"). The target object remains in a terminating state while the
control plane, or other components, take the actions defined by the finalizers.
After these actions are complete, the controller removes the relevant finalizers
from the target object. When the `metadata.finalizers` field is empty,
Kubernetes considers the deletion complete and deletes the object.

You can use finalizers to control [garbage collection](/reference/glossary/garbage-collection.md)
of resources. For example, you can define a finalizer to clean up related
[API resources](/reference/glossary/api-resource.md) or infrastructure before the controller
deletes the object being finalized.
