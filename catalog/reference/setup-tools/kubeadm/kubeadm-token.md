---
type: concept
title: kubeadm token
description: Bootstrap tokens are used for establishing bidirectional trust between a node joining the cluster and a control-plane node, as described in authenticating with bootstrap tokens.
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-token/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/kubeadm-token.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/kubeadm-token.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-token/
    title: kubeadm token | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/kubeadm-token.md
---

# kubeadm token

Bootstrap tokens are used for establishing bidirectional trust between a node joining
the cluster and a control-plane node, as described in [authenticating with bootstrap tokens](/reference/access-authn-authz/bootstrap-tokens.md).

`kubeadm init` creates an initial token with a 24-hour TTL. The following commands allow you to manage
such a token and also to create and manage new ones.

## kubeadm token create {#cmd-token-create}

Create bootstrap tokens on the server

### Synopsis

This command will create a bootstrap token for you.
You can specify the usages for this token, the "time to live" and an optional human friendly description.

The [token] is the actual token to write.
This should be a securely generated random token of the form "[a-z0-9]{6}.[a-z0-9]{16}".
If no [token] is given, kubeadm will generate a random token instead.

```
kubeadm token create [token]
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--certificate-key string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>When used together with '--print-join-command', print the full 'kubeadm join' flag needed to join the cluster as a control-plane. To create a new certificate key you must use 'kubeadm init phase upload-certs --upload-certs'.</p></td>
</tr>

<tr>
<td colspan="2">--config string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a kubeadm configuration file.</p></td>
</tr>

<tr>
<td colspan="2">--description string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>A human friendly description of how this token is used.</p></td>
</tr>

<tr>
<td colspan="2">--groups strings&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "system:bootstrappers:kubeadm:default-node-token"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Extra groups that this token will authenticate as when used for authentication. Must match &quot;\Asystem:bootstrappers:[a-z0-9:-]{0,255}[a-z0-9]\z&quot;</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for create</p></td>
</tr>

<tr>
<td colspan="2">--print-join-command</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Instead of printing only the token, print the full 'kubeadm join' flag needed to join the cluster using the token.</p></td>
</tr>

<tr>
<td colspan="2">--ttl duration&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 24h0m0s</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire</p></td>
</tr>

<tr>
<td colspan="2">--usages strings&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "signing,authentication"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Describes the ways in which this token can be used. You can pass --usages multiple times or provide a comma separated list of options. Valid options: [signing,authentication]</p></td>
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
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Whether to enable dry-run mode or not</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/admin.conf"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.</p></td>
</tr>

<tr>
<td colspan="2">--rootfs string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.</p></td>
</tr>

</tbody>
</table>

## kubeadm token delete {#cmd-token-delete}

Delete bootstrap tokens on the server

### Synopsis

This command will delete a list of bootstrap tokens for you.

The [token-value] is the full Token of the form "[a-z0-9]{6}.[a-z0-9]{16}" or the
Token ID of the form "[a-z0-9]{6}" to delete.

```
kubeadm token delete [token-value] ...
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for delete</p></td>
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
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Whether to enable dry-run mode or not</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/admin.conf"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.</p></td>
</tr>

<tr>
<td colspan="2">--rootfs string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.</p></td>
</tr>

</tbody>
</table>

## kubeadm token generate {#cmd-token-generate}

Generate and print a bootstrap token, but do not create it on the server

### Synopsis

This command will print out a randomly-generated bootstrap token that can be used with
the "init" and "join" commands.

You don't have to use this command in order to generate a token. You can do so
yourself as long as it is in the format "[a-z0-9]{6}.[a-z0-9]{16}". This
command is provided for convenience to generate tokens in the given format.

You can also use "kubeadm init" without specifying a token and it will
generate and print one for you.

```
kubeadm token generate [flags]
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for generate</p></td>
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
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Whether to enable dry-run mode or not</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/admin.conf"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.</p></td>
</tr>

<tr>
<td colspan="2">--rootfs string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.</p></td>
</tr>

</tbody>
</table>

## kubeadm token list {#cmd-token-list}

List bootstrap tokens on the server

### Synopsis

This command will list all bootstrap tokens for you.

```
kubeadm token list [flags]
```

### Options

   <table style="width: 100%; table-layout: fixed;">
<colgroup>
<col span="1" style="width: 10px;" />
<col span="1" />
</colgroup>
<tbody>

<tr>
<td colspan="2">--allow-missing-template-keys&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: true</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for list</p></td>
</tr>

<tr>
<td colspan="2">-o, --output string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "text"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Output format. One of: text|json|yaml|kyaml|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.</p></td>
</tr>

<tr>
<td colspan="2">--show-managed-fields</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If true, keep the managedFields when printing objects in JSON or YAML format.</p></td>
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
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Whether to enable dry-run mode or not</p></td>
</tr>

<tr>
<td colspan="2">--kubeconfig string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/admin.conf"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.</p></td>
</tr>

<tr>
<td colspan="2">--rootfs string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.</p></td>
</tr>

</tbody>
</table>

## What's next

* [kubeadm join](/reference/setup-tools/kubeadm/kubeadm-join.md) to bootstrap a Kubernetes worker node and join it to the cluster
