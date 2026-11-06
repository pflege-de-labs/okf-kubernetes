---
type: concept
title: kubeadm join
description: This command initializes a new Kubernetes node and joins it to the existing cluster.
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/kubeadm-join.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/kubeadm-join.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/
    title: kubeadm join | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/kubeadm-join.md
---

# kubeadm join

This command initializes a new Kubernetes node and joins it to the existing cluster.

Run this on any machine you wish to join an existing cluster

### Synopsis

When joining a kubeadm initialized cluster, we need to establish
bidirectional trust. This is split into discovery (having the Node
trust the Kubernetes Control Plane) and TLS bootstrap (having the
Kubernetes Control Plane trust the Node).

There are 2 main schemes for discovery. The first is to use a shared
token along with the IP address of the API server. The second is to
provide a file - a subset of the standard kubeconfig file. The
discovery/kubeconfig file supports token, client-go authentication
plugins ("exec"), "tokenFile", and "authProvider". This file can be a
local file or downloaded via an HTTPS URL. The forms are
kubeadm join --discovery-token abcdef.1234567890abcdef 1.2.3.4:6443,
kubeadm join --discovery-file path/to/file.conf, or kubeadm join
--discovery-file https://url/file.conf. Only one form can be used. If
the discovery information is loaded from a URL, HTTPS must be used.
Also, in that case the host installed CA bundle is used to verify
the connection.

If you use a shared token for discovery, you should also pass the
--discovery-token-ca-cert-hash flag to validate the public key of the
root certificate authority (CA) presented by the Kubernetes Control Plane.
The value of this flag is specified as "&lt;hash-type&gt;:&lt;hex-encoded-value&gt;",
where the supported hash type is "sha256". The hash is calculated over
the bytes of the Subject Public Key Info (SPKI) object (as in RFC7469).
This value is available in the output of "kubeadm init" or can be
calculated using standard tools. The --discovery-token-ca-cert-hash flag
may be repeated multiple times to allow more than one public key.

If you cannot know the CA public key hash ahead of time, you can pass
the --discovery-token-unsafe-skip-ca-verification flag to disable this
verification. This weakens the kubeadm security model since other nodes
can potentially impersonate the Kubernetes Control Plane.

The TLS bootstrap mechanism is also driven via a shared token. This is
used to temporarily authenticate with the Kubernetes Control Plane to submit a
certificate signing request (CSR) for a locally created key pair. By
default, kubeadm will set up the Kubernetes Control Plane to automatically
approve these signing requests. This token is passed in with the
--tls-bootstrap-token abcdef.1234567890abcdef flag.

Often times the same token is used for both parts. In this case, the
--token flag can be used instead of specifying each token individually.

The "join [api-server-endpoint]" command executes the following phases:
```
preflight               Run join pre-flight checks
control-plane-prepare   Prepare the machine for serving a control plane
  /download-certs         Download certificates shared among control-plane nodes from the kubeadm-certs Secret
  /certs                  Generate the certificates for the new control plane components
  /kubeconfig             Generate the kubeconfig for the new control plane components
  /control-plane          Generate the manifests for the new control plane components
kubelet-start           Write kubelet settings, certificates and (re)start the kubelet
etcd-join               Join etcd for control plane nodes
kubelet-wait-bootstrap  Wait for the kubelet to bootstrap itself
control-plane-join      Join a machine as a control plane instance
  /mark-control-plane     Mark a node as a control-plane
wait-control-plane      Wait for the control plane to start
```

```
kubeadm join [api-server-endpoint] [flags]
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.</p></td>
</tr>

<tr>
<td colspan="2">--apiserver-bind-port int32&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 6443</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>If the node should host a new control plane instance, the port for the API Server to bind to.</p></td>
</tr>

<tr>
<td colspan="2">--certificate-key string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Use this key to decrypt the certificate secrets uploaded by init. The certificate key is a hex encoded string that is an AES key of size 32 bytes.</p></td>
</tr>

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
<td colspan="2">--cri-socket string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.</p></td>
</tr>

<tr>
<td colspan="2">--discovery-file string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>For file-based discovery, a file or URL from which to load cluster information.</p></td>
</tr>

<tr>
<td colspan="2">--discovery-token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>For token-based discovery, the token used to validate cluster information fetched from the API server.</p></td>
</tr>

<tr>
<td colspan="2">--discovery-token-ca-cert-hash strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>For token-based discovery, validate that the root CA public key matches this hash (format: &quot;&lt;type&gt;:&lt;value&gt;&quot;).</p></td>
</tr>

<tr>
<td colspan="2">--discovery-token-unsafe-skip-ca-verification</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.</p></td>
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
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for join</p></td>
</tr>

<tr>
<td colspan="2">--ignore-preflight-errors strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.</p></td>
</tr>

<tr>
<td colspan="2">--node-name string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Specify the node name.</p></td>
</tr>

<tr>
<td colspan="2">--patches string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to a directory that contains files named &quot;target[suffix][+patchtype].extension&quot;. For example, &quot;kube-apiserver0+merge.yaml&quot; or just &quot;etcd.json&quot;. &quot;target&quot; can be one of &quot;kube-apiserver&quot;, &quot;kube-controller-manager&quot;, &quot;kube-scheduler&quot;, &quot;etcd&quot;, &quot;kubeletconfiguration&quot;, &quot;corednsdeployment&quot;, &quot;kubeproxydaemonset&quot;. &quot;patchtype&quot; can be one of &quot;strategic&quot;, &quot;merge&quot; or &quot;json&quot; and they match the patch formats supported by kubectl. The default &quot;patchtype&quot; is &quot;strategic&quot;. &quot;extension&quot; must be either &quot;json&quot; or &quot;yaml&quot;. &quot;suffix&quot; is an optional string that can be used to determine which patches are applied first alpha-numerically.</p></td>
</tr>

<tr>
<td colspan="2">--skip-phases strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>List of phases to be skipped</p></td>
</tr>

<tr>
<td colspan="2">--tls-bootstrap-token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.</p></td>
</tr>

<tr>
<td colspan="2">--token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.</p></td>
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

### The join workflow {#join-workflow}

`kubeadm join` bootstraps a Kubernetes worker node or a control-plane node and adds it to the cluster.
This action consists of the following steps for worker nodes:

1. kubeadm downloads necessary cluster information from the API server.
   By default, it uses the bootstrap token and the CA key hash to verify the
   authenticity of that data. The root CA can also be discovered directly via a
   file or URL.

1. Once the cluster information is known, kubelet can start the TLS bootstrapping
   process.

   The TLS bootstrap uses the shared token to temporarily authenticate
   with the Kubernetes API server to submit a certificate signing request (CSR); by
   default the control plane signs this CSR request automatically.

1. Finally, kubeadm configures the local kubelet to connect to the API
   server with the definitive identity assigned to the node.

For control-plane nodes additional steps are performed:

1. Downloading certificates shared among control-plane nodes from the cluster
  (if explicitly requested by the user).

1. Generating control-plane component manifests, certificates and kubeconfig.

1. Adding new local etcd member.

### Using join phases with kubeadm {#join-phases}

Kubeadm allows you join a node to the cluster in phases using `kubeadm join phase`.

To view the ordered list of phases and sub-phases you can call `kubeadm join --help`. The list will be located
at the top of the help screen and each phase will have a description next to it.
Note that by calling `kubeadm join` all of the phases and sub-phases will be executed in this exact order.

Some phases have unique flags, so if you want to have a look at the list of available options add `--help`, for example:

```shell
kubeadm join phase kubelet-start --help
```

Similar to the [kubeadm init phase](/reference/setup-tools/kubeadm/kubeadm-init.md#init-phases)
command, `kubeadm join phase` allows you to skip a list of phases using the `--skip-phases` flag.

For example:

```shell
sudo kubeadm join --skip-phases=preflight --config=config.yaml
```

**Feature state:** Kubernetes v1.22 [beta]

Alternatively, you can use the `skipPhases` field in `JoinConfiguration`.

### Discovering what cluster CA to trust

The kubeadm discovery has several options, each with security tradeoffs.
The right method for your environment depends on how you provision nodes and the
security expectations you have about your network and node lifecycles.

#### Token-based discovery with CA pinning

This is the default mode in kubeadm. In this mode, kubeadm downloads
the cluster configuration (including root CA) and validates it using the token
as well as validating that the root CA public key matches the provided hash and
that the API server certificate is valid under the root CA.

The CA key hash has the format `sha256:<hex_encoded_hash>`.
By default, the hash value is printed at the end of the `kubeadm init` command or
in the output from the `kubeadm token create --print-join-command` command.
It is in a standard format (see [RFC7469](https://tools.ietf.org/html/rfc7469#section-2.4))
and can also be calculated by 3rd party tools or provisioning systems.
For example, using the OpenSSL CLI:

```shell
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
```

**Example `kubeadm join` commands:**

For worker nodes:

```shell
kubeadm join --discovery-token abcdef.1234567890abcdef --discovery-token-ca-cert-hash sha256:1234..cdef 1.2.3.4:6443
```

For control-plane nodes:

```shell
kubeadm join --discovery-token abcdef.1234567890abcdef --discovery-token-ca-cert-hash sha256:1234..cdef --control-plane 1.2.3.4:6443
```

You can also call `join` for a control-plane node with `--certificate-key` to copy certificates to this node,
if the `kubeadm init` command was called with `--upload-certs`.

**Advantages:**

- Allows bootstrapping nodes to securely discover a root of trust for the
  control-plane node even if other worker nodes or the network are compromised.

- Convenient to execute manually since all of the information required fits
  into a single `kubeadm join` command.

**Disadvantages:**

- The CA hash is not normally known until the control-plane node has been provisioned,
  which can make it more difficult to build automated provisioning tools that
  use kubeadm. By generating your CA in beforehand, you may workaround this
  limitation.

#### Token-based discovery without CA pinning

This mode relies only on the symmetric token to sign
(HMAC-SHA256) the discovery information that establishes the root of trust for
the control-plane. To use the mode the joining nodes must skip the hash validation of the
CA public key, using `--discovery-token-unsafe-skip-ca-verification`. You should consider
using one of the other modes if possible.

**Example `kubeadm join` command:**

```shell
kubeadm join --token abcdef.1234567890abcdef --discovery-token-unsafe-skip-ca-verification 1.2.3.4:6443
```

**Advantages:**

- Still protects against many network-level attacks.

- The token can be generated ahead of time and shared with the control-plane node and
  worker nodes, which can then bootstrap in parallel without coordination. This
  allows it to be used in many provisioning scenarios.

**Disadvantages:**

- If an attacker is able to steal a bootstrap token via some vulnerability,
  they can use that token (along with network-level access) to impersonate the
  control-plane node to other bootstrapping nodes. This may or may not be an appropriate
  tradeoff in your environment.

#### File or HTTPS-based discovery

This provides an out-of-band way to establish a root of trust between the control-plane node
and bootstrapping nodes. Consider using this mode if you are building automated provisioning
using kubeadm. The format of the discovery file is a regular Kubernetes
[kubeconfig](/tasks/access-application-cluster/configure-access-multiple-clusters.md) file.

In case the discovery file does not contain credentials, the TLS discovery token will be used.

**Example `kubeadm join` commands:**

- `kubeadm join --discovery-file path/to/file.conf` (local file)

- `kubeadm join --discovery-file https://url/file.conf` (remote HTTPS URL)

**Advantages:**

- Allows bootstrapping nodes to securely discover a root of trust for the
  control-plane node even if the network or other worker nodes are compromised.

**Disadvantages:**

- Requires that you have some way to carry the discovery information from
  the control-plane node to the bootstrapping nodes. If the discovery file contains credentials
  you must keep it secret and transfer it over a secure channel. This might be possible with your
  cloud provider or provisioning tool.

#### Use of custom kubelet credentials with `kubeadm join`

To allow `kubeadm join` to use predefined kubelet credentials and skip client TLS bootstrap
and CSR approval for a new node:

1. From a working control plane node in the cluster that has `/etc/kubernetes/pki/ca.key`
   execute `kubeadm kubeconfig user --org system:nodes --client-name system:node:$NODE > kubelet.conf`.
   `$NODE` must be set to the name of the new node.
2. Modify the resulted `kubelet.conf` manually to adjust the cluster name and the server endpoint,
   or run `kubeadm kubeconfig user --config` (it accepts `InitConfiguration`).

If your cluster does not have the `ca.key` file, you must sign the embedded certificates in
the `kubelet.conf` externally. For additional information, see
[PKI certificates and requirements](https://kubernetes.io/docs/setup/best-practices/certificates/) and
[Certificate Management with kubeadm](/tasks/administer-cluster/kubeadm/kubeadm-certs.md#external-ca-mode).

1. Copy the resulting `kubelet.conf` to `/etc/kubernetes/kubelet.conf` on the new node.
2. Execute `kubeadm join` with the flag
   `--ignore-preflight-errors=FileAvailable--etc-kubernetes-kubelet.conf` on the new node.

### Securing your installation even more {#securing-more}

The defaults for kubeadm may not work for everyone. This section documents how to tighten up a kubeadm installation
at the cost of some usability.

#### Turning off auto-approval of node client certificates

By default, there is a CSR auto-approver enabled that basically approves any client certificate request
for a kubelet when a Bootstrap Token was used when authenticating. If you don't want the cluster to
automatically approve kubelet client certs, you can turn it off by executing this command:

```shell
kubectl delete clusterrolebinding kubeadm:node-autoapprove-bootstrap
```

After that, `kubeadm join` will block until the admin has manually approved the CSR in flight:

1. Using `kubectl get csr`, you can see that the original CSR is in the Pending state.

   ```shell
   kubectl get csr
   ```

   The output is similar to this:

   ```
   NAME                                                   AGE       REQUESTOR                 CONDITION
   node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ   18s       system:bootstrap:878f07   Pending
   ```

2. `kubectl certificate approve` allows the admin to approve CSR.This action tells a certificate signing
   controller to issue a certificate to the requestor with the attributes requested in the CSR.

   ```shell
   kubectl certificate approve node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ
   ```

   The output is similar to this:

   ```
   certificatesigningrequest "node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ" approved
   ```

3. This would change the CSR resource to Active state.

   ```shell
   kubectl get csr
   ```

   The output is similar to this:

   ```
   NAME                                                   AGE       REQUESTOR                 CONDITION
   node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ   1m        system:bootstrap:878f07   Approved,Issued
   ```

This forces the workflow that `kubeadm join` will only succeed if `kubectl certificate approve` has been run.

#### Turning off public access to the `cluster-info` ConfigMap

In order to achieve the joining flow using the token as the only piece of validation information, a
 ConfigMap with some data needed for validation of the control-plane node's identity is exposed publicly by
default. While there is no private data in this ConfigMap, some users might wish to turn
it off regardless. Doing so will disable the ability to use the `--discovery-token` flag of the
`kubeadm join` flow. Here are the steps to do so:

* Fetch the `cluster-info` file from the API Server:

```shell
kubectl -n kube-public get cm cluster-info -o jsonpath='{.data.kubeconfig}' | tee cluster-info.yaml
```

The output is similar to this:

```yaml
apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority-data: <ca-cert>
    server: https://<ip>:<port>
  name: ""
contexts: []
current-context: ""
preferences: {}
users: []
```

* Use the `cluster-info.yaml` file as an argument to `kubeadm join --discovery-file`.

* Turn off public access to the `cluster-info` ConfigMap:

```shell
kubectl -n kube-public delete rolebinding kubeadm:bootstrap-signer-clusterinfo
```

These commands should be run after `kubeadm init` but before `kubeadm join`.

### Using kubeadm join with a configuration file {#config-file}

**Caution:**
The config file is still considered beta and may change in future versions.

It's possible to configure `kubeadm join` with a configuration file instead of command
line flags, and some more advanced features may only be available as
configuration file options. This file is passed using the `--config` flag and it must
contain a `JoinConfiguration` structure. Mixing `--config` with others flags may not be
allowed in some cases.

The default configuration can be printed out using the
[kubeadm config print](/reference/setup-tools/kubeadm/kubeadm-config.md#cmd-config-print) command.

If your configuration is not using the latest version it is **recommended** that you migrate using
the [kubeadm config migrate](/reference/setup-tools/kubeadm/kubeadm-config.md#cmd-config-migrate) command.

For more information on the fields and usage of the configuration you can navigate to our
[API reference](/reference/config-api/kubeadm-config.v1beta4.md).

## What's next

* [kubeadm init](/reference/setup-tools/kubeadm/kubeadm-init.md) to bootstrap a Kubernetes control-plane node.
* [kubeadm token](/reference/setup-tools/kubeadm/kubeadm-token.md) to manage tokens for `kubeadm join`.
* [kubeadm reset](/reference/setup-tools/kubeadm/kubeadm-reset.md) to revert any changes made to this host by `kubeadm init` or `kubeadm join`.
