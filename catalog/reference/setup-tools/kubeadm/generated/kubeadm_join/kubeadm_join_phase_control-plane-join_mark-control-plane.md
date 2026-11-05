---
type: reference
title: kubeadm_join_phase_control-plane-join_mark-control-plane
description: Mark a node as a control-plane
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join_mark-control-plane/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join_mark-control-plane.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join_mark-control-plane.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join_mark-control-plane/
    title: kubeadm_join_phase_control-plane-join_mark-control-plane | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_join/kubeadm_join_phase_control-plane-join_mark-control-plane.md
---

# kubeadm_join_phase_control-plane-join_mark-control-plane

### Synopsis

Mark a node as a control-plane

```
kubeadm join phase control-plane-join mark-control-plane [flags]
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--config string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a kubeadm configuration file.</p></td>
</tr>

<tr>
<td colspan="2">--control-plane</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Create a new control plane instance on this node</p></td>
</tr>

<tr>
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Don't apply any changes; just output what would be done.</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for mark-control-plane</p></td>
</tr>

<tr>
<td colspan="2">--node-name string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Specify the node name.</p></td>
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
