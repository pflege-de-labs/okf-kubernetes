---
type: constraint
title: Label and annotation syntax and limits
description: Label and annotation keys are an optional 253-char DNS subdomain prefix plus a 63-char name; label values are at most 63 chars; all annotations on an object total at most 256 KiB.
resource: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
tags: [kubernetes, requirements, labels, annotations, validation, limits]
generated: { by: "claude-code/claude-opus-5-5", at: "2026-11-05T00:00:00Z" }
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
sources:
  - id: label-key
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/labels.md#L57-L75
    title: Labels and Selectors - label key syntax
  - id: label-value
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/labels.md#L77-L81
    title: Labels and Selectors - label value syntax
  - id: annotation-key
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/annotations.md#L72-L76
    title: Annotations - key syntax
  - id: annotation-value
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/concepts/overview/working-with-objects/annotations.md#L78-L81
    title: Annotations - value syntax and total size
---

# Label and annotation syntax and limits

Full text: [Labels and Selectors](/concepts/overview/working-with-objects/labels.md),
[Annotations](/concepts/overview/working-with-objects/annotations.md),
[Well-Known Labels, Annotations and Taints](/reference/labels-annotations-taints.md).

# Keys (labels and annotations)

Keys are `[prefix/]name`.[^label-key][^annotation-key]

| Part | Required | Max length | Rules |
|---|---|---|---|
| prefix | no | 253 | DNS subdomain (dot-separated DNS labels), followed by `/` |
| name | yes | 63 | `[a-z0-9A-Z]` at start and end; `-`, `_`, `.` and alphanumerics between |

- Without a prefix the key is presumed private to the user. Automated system
  components that add labels or annotations to end-user objects must use a
  prefix.[^label-key][^annotation-key]
- The `kubernetes.io/` and `k8s.io/` prefixes are reserved for Kubernetes core
  components.[^label-key][^annotation-key]
- A label key must be unique for a given object.[^label-key]

# Label values

- At most 63 characters; may be empty.[^label-value]
- Unless empty, must begin and end with `[a-z0-9A-Z]`.[^label-value]
- May contain `-`, `_`, `.` and alphanumerics between.[^label-value]

# Annotation values

- No character set restrictions: any string, including whitespace and
  structured data such as JSON or YAML. Base64-encode binary data.[^annotation-value]
- The total size of **all** annotations on one object (keys and values
  combined) must not exceed **256 KiB**.[^annotation-value]

Related: [object names](object-names.md), [size limits](size-limits.md).

[^label-key]: Labels and Selectors - label key syntax
[^label-value]: Labels and Selectors - label value syntax
[^annotation-key]: Annotations - key syntax
[^annotation-value]: Annotations - value syntax and total size
