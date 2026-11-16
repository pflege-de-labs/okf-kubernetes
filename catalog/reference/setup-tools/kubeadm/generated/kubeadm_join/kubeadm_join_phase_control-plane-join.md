---
type: reference
title: kubeadm_join_phase_control-plane-join
description: Join a machine as a control plane instance
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-16T13:37:38Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/48a1f06596294a8ea2304fcc33047d3b639f135d/content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join/
    title: kubeadm_join_phase_control-plane-join | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 48a1f06596294a8ea2304fcc33047d3b639f135d
source_path: content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join.md
---

# kubeadm_join_phase_control-plane-join

### Synopsis

Join a machine as a control plane instance

```
kubeadm join phase control-plane-join [flags]
```

### Examples

```
  # Joins a machine as a control plane instance
  kubeadm join phase control-plane-join all
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for control-plane-join</p></td>
</tr>

</tbody>
</table>

### Options inherited from parent commands

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--rootfs string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.</p></td>
</tr>

</tbody>
</table>
