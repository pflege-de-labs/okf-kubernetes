---
type: reference
title: kubeadm_init_phase_certs_sa
description: Generate a private key for signing service account tokens along with its public key
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_certs_sa/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_certs_sa.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_certs_sa.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_certs_sa/
    title: kubeadm_init_phase_certs_sa | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_init/kubeadm_init_phase_certs_sa.md
---

# kubeadm_init_phase_certs_sa

Generate a private key for signing service account tokens along with its public key

### Synopsis

Generate the private key for signing service account tokens along with its public key, and save them into sa.key and sa.pub files.

If both files already exist, kubeadm skips the generation step and existing files will be used.

```
kubeadm init phase certs sa [flags]
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

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
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Don't apply any changes; just output what would be done.</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for sa</p></td>
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
