---
type: reference
title: kubeadm_certs_certificate-key
description: Generate certificate keys
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_certs/kubeadm_certs_certificate-key/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-12T15:41:02Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2674d5095a50d781d2f884c83973921f48c5cf28/content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_certs/kubeadm_certs_certificate-key.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/generated/kubeadm_certs/kubeadm_certs_certificate-key.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/generated/kubeadm_certs/kubeadm_certs_certificate-key/
    title: kubeadm_certs_certificate-key | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2674d5095a50d781d2f884c83973921f48c5cf28
source_path: content/en/docs/reference/setup-tools/kubeadm/generated/kubeadm_certs/kubeadm_certs_certificate-key.md
---

# kubeadm_certs_certificate-key

Generate certificate keys

### Synopsis

This command will print out a secure randomly-generated certificate key that can be used with
the "init" command.

You can also use "kubeadm init --upload-certs" without specifying a certificate key and it will
generate and print one for you.

```
kubeadm certs certificate-key [flags]
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for certificate-key</p></td>
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
