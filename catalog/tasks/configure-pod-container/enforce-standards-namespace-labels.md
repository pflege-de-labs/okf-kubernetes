---
type: task
title: Enforce Pod Security Standards with Namespace Labels
description: Namespaces can be labeled to enforce the Pod Security Standards.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-namespace-labels/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/configure-pod-container/enforce-standards-namespace-labels.md
    title: 'Kubernetes website source: tasks/configure-pod-container/enforce-standards-namespace-labels.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-namespace-labels/
    title: Enforce Pod Security Standards with Namespace Labels | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/configure-pod-container/enforce-standards-namespace-labels.md
---

# Enforce Pod Security Standards with Namespace Labels

Namespaces can be labeled to enforce the [Pod Security Standards](/concepts/security/pod-security-standards.md). The three policies
[privileged](/concepts/security/pod-security-standards.md#privileged), [baseline](/concepts/security/pod-security-standards.md#baseline)
and [restricted](/concepts/security/pod-security-standards.md#restricted) broadly cover the security spectrum
and are implemented by the [Pod Security](/concepts/security/pod-security-admission.md) [admission controller](/reference/glossary/admission-controller.md).

## Before you begin

Pod Security Admission was available by default in Kubernetes v1.23, as
a beta. From version 1.25 onwards, Pod Security Admission is generally
available.

To check the version, enter `kubectl version`.

## Requiring the `baseline` Pod Security Standard with namespace labels

This manifest defines a Namespace `my-baseline-namespace` that:

- _Blocks_ any pods that don't satisfy the `baseline` policy requirements.
- Generates a user-facing warning and adds an audit annotation to any created pod that does not
  meet the `restricted` policy requirements.
- Pins the versions of the `baseline` and `restricted` policies to v1.37.

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: my-baseline-namespace
  labels:
    pod-security.kubernetes.io/enforce: baseline
    pod-security.kubernetes.io/enforce-version: v1.37

    # We are setting these to our _desired_ `enforce` level.
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/audit-version: v1.37
    pod-security.kubernetes.io/warn: restricted
    pod-security.kubernetes.io/warn-version: v1.37
```

## Add labels to existing namespaces with `kubectl label`

**Note:**
When an `enforce` policy (or version) label is added or changed, the admission plugin will test
each pod in the namespace against the new policy. Violations are returned to the user as warnings.

It is helpful to apply the `--dry-run` flag when initially evaluating security profile changes for
namespaces. The Pod Security Standard checks will still be run in _dry run_ mode, giving you
information about how the new policy would treat existing pods, without actually updating a policy.

```shell
kubectl label --dry-run=server --overwrite ns --all \
    pod-security.kubernetes.io/enforce=baseline
```

### Applying to all namespaces

If you're just getting started with the Pod Security Standards, a suitable first step would be to
configure all namespaces with audit annotations for a stricter level such as `baseline`:

```shell
kubectl label --overwrite ns --all \
  pod-security.kubernetes.io/audit=baseline \
  pod-security.kubernetes.io/warn=baseline
```

Note that this is not setting an enforce level, so that namespaces that haven't been explicitly
evaluated can be distinguished. You can list namespaces without an explicitly set enforce level
using this command:

```shell
kubectl get namespaces --selector='!pod-security.kubernetes.io/enforce'
```

### Applying to a single namespace

You can update a specific namespace as well. This command adds the `enforce=restricted`
policy to `my-existing-namespace`, pinning the restricted policy version to v1.37.

```shell
kubectl label --overwrite ns my-existing-namespace \
  pod-security.kubernetes.io/enforce=restricted \
  pod-security.kubernetes.io/enforce-version=v1.37
```
