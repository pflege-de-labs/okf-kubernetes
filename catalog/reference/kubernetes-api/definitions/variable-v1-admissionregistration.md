---
type: api-reference
title: Variable
description: Variable is the definition of a variable that is used for composition. A variable is defined as a named expression.
resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/variable-v1-admissionregistration/
tags: [kubernetes, reference, kubernetes-api]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/kubernetes-api/definitions/variable-v1-admissionregistration.md
    title: 'Kubernetes website source: reference/kubernetes-api/definitions/variable-v1-admissionregistration.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/kubernetes-api/definitions/variable-v1-admissionregistration/
    title: Variable | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/kubernetes-api/definitions/variable-v1-admissionregistration.md
api_metadata:
  apiVersion: admissionregistration.k8s.io/v1
  import: k8s.io/api/admissionregistration/v1
  kind: Variable
---

# Variable

`apiVersion: admissionregistration.k8s.io/v1`

`import "k8s.io/api/admissionregistration/v1"`

## Variable {#Variable}

Variable is the definition of a variable that is used for composition. A variable is defined as a named expression.

<hr>

<table>
  <thead><tr><th>Field</th><th>Description</th></tr></thead>
  <tbody>
    <tr>
      <td><code>expression</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>expression is the expression that will be evaluated as the value of the variable. The CEL expression has access to the same identifiers as the CEL expressions in Validation.</td>
    </tr>
    <tr>
      <td><code>name</code>&nbsp;<strong>*</strong><br/><em>string</em></td>
      <td>name is the name of the variable. The name must be a valid CEL identifier and unique among all variables. The variable can be accessed in other expressions through `variables` For example, if name is "foo", the variable will be available as `variables.foo`</td>
    </tr>
  </tbody>
</table>
