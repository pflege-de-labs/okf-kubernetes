---
type: constraint
title: Object size limits
description: Documented size limits - ConfigMap and Secret 1 MiB, all annotations 256 KiB, StorageClass and VolumeAttributesClass parameters 512 entries / 256 KiB, termination messages 4096 bytes per container and 12 KiB per Pod.
tags: [kubernetes, requirements, limits, size, configmap, secret]
generated: { by: "claude-code/claude-opus-5-5", at: "2026-11-05T00:00:00Z" }
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
sources:
  - id: configmap
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/configuration/configmap.md#L35-L40
    title: ConfigMaps - size limit
  - id: secret
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/configuration/secret.md#L465-L471
    title: Secrets - size limit
  - id: annotations
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/overview/working-with-objects/annotations.md#L81
    title: Annotations - total size
  - id: storageclass
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/storage/storage-classes.md#L230-L232
    title: Storage Classes - parameter limits
  - id: vac
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/concepts/storage/volume-attributes-classes.md#L127-L128
    title: Volume Attributes Classes - parameter limits
  - id: termination
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/tasks/debug/debug-application/determine-reason-pod-failure.md#L94-L98
    title: Determine the Reason for Pod Failure - termination message length
---

# Object size limits

| What | Limit | Notes |
|---|---|---|
| ConfigMap data | 1 MiB[^configmap] | For larger settings, mount a volume or use a separate database or file service |
| Secret | 1 MiB per Secret[^secret] | Many small Secrets can also exhaust memory; limit their count with a ResourceQuota |
| Annotations | 256 KiB total per object, keys and values combined[^annotations] | See [labels and annotations](labels-and-annotations.md) |
| StorageClass `parameters` | at most 512 entries, 256 KiB total including keys and values[^storageclass] | |
| VolumeAttributesClass `parameters` | at most 512 entries, 256 KiB total including keys and values[^vac] | |
| Container termination message | 4096 bytes per message (kubelet truncates)[^termination] | |
| Termination messages per Pod | 12 KiB total, split equally across all `initContainers` and `containers`[^termination] | 12 containers leave 1024 bytes each |

Full text: [ConfigMaps](/concepts/configuration/configmap.md),
[Secrets](/concepts/configuration/secret.md),
[Storage Classes](/concepts/storage/storage-classes.md),
[Volume Attributes Classes](/concepts/storage/volume-attributes-classes.md),
[Determine the Reason for Pod Failure](/tasks/debug/debug-application/determine-reason-pod-failure.md).

[^configmap]: ConfigMaps - size limit
[^secret]: Secrets - size limit
[^annotations]: Annotations - total size
[^storageclass]: Storage Classes - parameter limits
[^vac]: Volume Attributes Classes - parameter limits
[^termination]: Determine the Reason for Pod Failure - termination message length
