# Requirements and limits

Curated, line-cited constraints from the Kubernetes v1.37 documentation. Hand-maintained; not touched by the generator.

* [Object name formats and length limits](object-names.md) - The four name formats (253-char DNS subdomain, 63-char RFC 1123 and RFC 1035 labels, path segment) and which resource kinds require which.
* [Label and annotation syntax and limits](labels-and-annotations.md) - Key prefix/name rules, 63-char label values, reserved prefixes, 256 KiB annotation total.
* [Namespace naming and reserved namespaces](namespaces.md) - RFC 1123 label names, initial namespaces, reserved `kube-` prefix.
* [Object size limits](size-limits.md) - ConfigMap and Secret 1 MiB, annotations 256 KiB, class parameters, termination messages.
* [Port numbers, port names and NodePort range](ports-and-ranges.md) - Port number and name rules, NodePort range and bands, well-known component ports.
