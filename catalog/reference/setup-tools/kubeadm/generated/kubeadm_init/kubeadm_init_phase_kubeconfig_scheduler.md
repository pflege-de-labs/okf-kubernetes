---
type: reference
title: kubeadm_init_phase_kubeconfig_scheduler
description: Generate a kubeconfig file for the scheduler to use
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_kubeconfig_scheduler/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_kubeconfig_scheduler.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_kubeconfig_scheduler.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_kubeconfig_scheduler/
    title: kubeadm_init_phase_kubeconfig_scheduler | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_kubeconfig_scheduler.md
---

# kubeadm_init_phase_kubeconfig_scheduler

Generate a kubeconfig file for the scheduler to use

### Synopsis

Generate the kubeconfig file for the scheduler to use and save it to scheduler.conf file.

```
kubeadm init phase kubeconfig scheduler [flags]
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--apiserver-advertise-address string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.</p></td>
</tr>

<tr>
<td colspan="2">--apiserver-bind-port int32&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 6443</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Port for the API Server to bind to.</p></td>
</tr>

<tr>
<td colspan="2">--cert-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/pki"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path where to save and store the certificates.</p></td>
</tr>

<tr>
<td colspan="2">--config string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a kubeadm configuration file.</p></td>
</tr>

<tr>
<td colspan="2">--control-plane-endpoint string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Specify a stable IP address or DNS name for the control plane.</p></td>
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for scheduler</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path where to save the kubeconfig file.</p></td>
</tr>

<tr>
<td colspan="2">--kubernetes-version string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "stable-1"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Choose a specific Kubernetes version for the control plane.</p></td>
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
