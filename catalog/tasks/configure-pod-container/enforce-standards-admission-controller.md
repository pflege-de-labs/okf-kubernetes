---
type: task
title: Enforce Pod Security Standards by Configuring the Built-in Admission Controller
description: Kubernetes provides a built-in admission controller to enforce the Pod Security Standards.
resource: https://kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-admission-controller/
tags: [kubernetes, tasks, configure-pod-container]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/configure-pod-container/enforce-standards-admission-controller.md
    title: 'Kubernetes website source: tasks/configure-pod-container/enforce-standards-admission-controller.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-admission-controller/
    title: Enforce Pod Security Standards by Configuring the Built-in Admission Controller | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/tasks/configure-pod-container/enforce-standards-admission-controller.md
---

# Enforce Pod Security Standards by Configuring the Built-in Admission Controller

Kubernetes provides a built-in [admission controller](/reference/access-authn-authz/admission-controllers.md#podsecurity)
to enforce the [Pod Security Standards](/concepts/security/pod-security-standards.md).
You can configure this admission controller to set cluster-wide defaults and [exemptions](/concepts/security/pod-security-admission.md#exemptions).

## Before you begin

Following an alpha release in Kubernetes v1.22,
Pod Security Admission became available by default in Kubernetes v1.23, as
a beta. From version 1.25 onwards, Pod Security Admission is generally
available.

To check the version, enter `kubectl version`.

If you are not running Kubernetes 1.37, you can switch
to viewing this page in the documentation for the Kubernetes version that you
are running.

## Configure the Admission Controller

**Note:**
`pod-security.admission.config.k8s.io/v1` configuration requires v1.25+.
For v1.23 and v1.24, use [v1beta1](https://v1-24.docs.kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-admission-controller/).
For v1.22, use [v1alpha1](https://v1-22.docs.kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-admission-controller/).

```yaml
apiVersion: apiserver.config.k8s.io/v1
kind: AdmissionConfiguration
plugins:
- name: PodSecurity
  configuration:
    apiVersion: pod-security.admission.config.k8s.io/v1 # see compatibility note
    kind: PodSecurityConfiguration
    # Defaults applied when a mode label is not set.
    #
    # Level label values must be one of:
    # - "privileged" (default)
    # - "baseline"
    # - "restricted"
    #
    # Version label values must be one of:
    # - "latest" (default) 
    # - specific version like "v1.37"
    defaults:
      enforce: "privileged"
      enforce-version: "latest"
      audit: "privileged"
      audit-version: "latest"
      warn: "privileged"
      warn-version: "latest"
    exemptions:
      # Array of authenticated usernames to exempt.
      usernames: []
      # Array of runtime class names to exempt.
      runtimeClasses: []
      # Array of namespaces to exempt.
      namespaces: []
```

**Note:**
The above manifest needs to be specified via the `--admission-control-config-file` to kube-apiserver.
