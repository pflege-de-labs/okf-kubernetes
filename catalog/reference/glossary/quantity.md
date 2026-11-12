---
type: glossary-term
title: Quantity
description: A whole-number representation of small or large numbers using SI suffixes.
resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-quantity
tags: [kubernetes, reference, glossary, fundamental]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/glossary/quantity.md
    title: 'Kubernetes website source: reference/glossary/quantity.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/glossary/?all=true#term-quantity
    title: Quantity | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/glossary/quantity.md
---

# Quantity

A whole-number representation of small or large numbers using [SI](https://en.wikipedia.org/wiki/International_System_of_Units) suffixes.

Quantities are representations of small or large numbers using a compact,
whole-number notation with SI suffixes.  Fractional numbers are represented
using milli units, while large numbers can be represented using kilo,
mega, or giga units.

For instance, the number `1.5` is represented as `1500m`, while the number `1000`
can be represented as `1k`, and `1000000` as `1M`. You can also specify
[binary-notation](https://en.wikipedia.org/wiki/Binary_prefix) suffixes; the number 2048 can be written as `2Ki`.

The accepted decimal (power-of-10) units are `m` (milli), `k` (kilo,
intentionally lowercase), `M` (mega), `G` (giga), `T` (tera), `P` (peta),
`E` (exa).

The accepted binary (power-of-2) units are `Ki` (kibi), `Mi` (mebi), `Gi` (gibi),
`Ti` (tebi), `Pi` (pebi), `Ei` (exbi).
