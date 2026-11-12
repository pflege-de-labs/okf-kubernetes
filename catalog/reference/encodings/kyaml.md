---
type: concept
title: KYAML Reference
description: This reference describes KYAML syntax.
resource: https://kubernetes.io/docs/reference/encodings/kyaml/
tags: [kubernetes, reference, encodings]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/encodings/kyaml.md
    title: 'Kubernetes website source: reference/encodings/kyaml.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/encodings/kyaml/
    title: KYAML Reference | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/encodings/kyaml.md
---

# KYAML Reference

**KYAML** is a safer and less ambiguous subset of YAML. Designed specifically for Kubernetes, KYAML addresses common YAML pitfalls such as whitespace sensitivity and implicit type coercion while maintaining full compatibility with existing YAML parsers and tooling. 

This reference describes KYAML syntax.

## Getting started with KYAML

YAML’s reliance on indentation and implicit type coercion often leads to configuration errors, especially in CI/CD pipelines and templating systems like Helm. KYAML eliminates these issues by enforcing explicit syntax and structure, making configurations more reliable and easier to debug.

### Basic Structure

KYAML uses *flow style* syntax with `{}` for objects and `[]` for arrays. All string values must be **double-quoted**.

```yaml
---
{
  apiVersion: "v1",
  kind: "Pod",
  metadata: {
    name: "my-pod",
    labels: {
      app: "demo"
    },
  },
  spec: {
    containers: [{
      name: "nginx",
      image: "nginx:1.20"
    }]
  }
}
```
