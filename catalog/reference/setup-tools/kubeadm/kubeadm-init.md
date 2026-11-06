---
type: concept
title: kubeadm init
description: This command initializes a Kubernetes control plane node.
resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/
tags: [kubernetes, reference, setup-tools]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-05T00:00:00Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/2fb3b0d636e8b49d100b77d554443b41571dbcd4/content/en/docs/reference/setup-tools/kubeadm/kubeadm-init.md
    title: 'Kubernetes website source: reference/setup-tools/kubeadm/kubeadm-init.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/
    title: kubeadm init | Kubernetes documentation
kubernetes_version: v1.37
source_commit: 2fb3b0d636e8b49d100b77d554443b41571dbcd4
source_path: content/en/docs/reference/setup-tools/kubeadm/kubeadm-init.md
---

# kubeadm init

This command initializes a Kubernetes control plane node.

### Synopsis

Run this command in order to set up the Kubernetes control plane

The "init" command executes the following phases:
```
preflight                     Run pre-flight checks
certs                         Certificate generation
  /ca                           Generate the self-signed Kubernetes CA to provision identities for other Kubernetes components
  /apiserver                    Generate the certificate for serving the Kubernetes API
  /apiserver-kubelet-client     Generate the certificate for the API server to connect to kubelet
  /front-proxy-ca               Generate the self-signed CA to provision identities for front proxy
  /front-proxy-client           Generate the certificate for the front proxy client
  /etcd-ca                      Generate the self-signed CA to provision identities for etcd
  /etcd-server                  Generate the certificate for serving etcd
  /etcd-peer                    Generate the certificate for etcd nodes to communicate with each other
  /etcd-healthcheck-client      Generate the certificate for liveness probes to healthcheck etcd
  /apiserver-etcd-client        Generate the certificate the apiserver uses to access etcd
  /sa                           Generate a private key for signing service account tokens along with its public key
kubeconfig                    Generate all kubeconfig files necessary to establish the control plane and the admin kubeconfig file
  /admin                        Generate a kubeconfig file for the admin to use and for kubeadm itself
  /super-admin                  Generate a kubeconfig file for the super-admin
  /kubelet                      Generate a kubeconfig file for the kubelet to use *only* for cluster bootstrapping purposes
  /controller-manager           Generate a kubeconfig file for the controller manager to use
  /scheduler                    Generate a kubeconfig file for the scheduler to use
etcd                          Generate static Pod manifest file for local etcd
  /local                        Generate the static Pod manifest file for a local, single-node local etcd instance
control-plane                 Generate all static Pod manifest files necessary to establish the control plane
  /apiserver                    Generates the kube-apiserver static Pod manifest
  /controller-manager           Generates the kube-controller-manager static Pod manifest
  /scheduler                    Generates the kube-scheduler static Pod manifest
kubelet-start                 Write kubelet settings and (re)start the kubelet
wait-control-plane            Wait for the control plane to start
upload-config                 Upload the kubeadm and kubelet configuration to a ConfigMap
  /kubeadm                      Upload the kubeadm ClusterConfiguration to a ConfigMap
  /kubelet                      Upload the kubelet component config to a ConfigMap
upload-certs                  Upload certificates to kubeadm-certs
mark-control-plane            Mark a node as a control-plane
bootstrap-token               Generates bootstrap tokens used to join a node to a cluster
kubelet-finalize              Updates settings relevant to the kubelet after TLS bootstrap
  /enable-client-cert-rotation  Enable kubelet client certificate rotation
addon                         Install required addons for passing conformance tests
  /coredns                      Install the CoreDNS addon to a Kubernetes cluster
  /kube-proxy                   Install the kube-proxy addon to a Kubernetes cluster
show-join-command             Show the join command for control-plane and worker node
```

```
kubeadm init [flags]
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
<td colspan="2">--apiserver-cert-extra-sans strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Optional extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names.</p></td>
</tr>

<tr>
<td colspan="2">--cert-dir string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "/etc/kubernetes/pki"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The path where to save and store the certificates.</p></td>
</tr>

<tr>
<td colspan="2">--certificate-key string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Key used to encrypt the control-plane certificates in the kubeadm-certs Secret. The certificate key is a hex encoded string that is an AES key of size 32 bytes.</p></td>
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
<td colspan="2">--cri-socket string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.</p></td>
</tr>

<tr>
<td colspan="2">--dry-run</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Don't apply any changes; just output what would be done.</p></td>
</tr>

<tr>
<td colspan="2">--feature-gates string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>A set of key=value pairs that describe feature gates for various features. Options are:<br/>RootlessControlPlane=true|false (ALPHA - default=false)</p></td>
</tr>

<tr>
<td colspan="2">-h, --help</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>help for init</p></td>
</tr>

<tr>
<td colspan="2">--ignore-preflight-errors strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.</p></td>
</tr>

<tr>
<td colspan="2">--image-repository string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "registry.k8s.io"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Choose a container registry to pull control plane images from</p></td>
</tr>

<tr>
<td colspan="2">--kubernetes-version string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "stable-1"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Choose a specific Kubernetes version for the control plane.</p></td>
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
<td colspan="2">--pod-network-cidr string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.</p></td>
</tr>

<tr>
<td colspan="2">--service-cidr string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "10.96.0.0/12"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Use alternative range of IP address for service VIPs.</p></td>
</tr>

<tr>
<td colspan="2">--service-dns-domain string&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: "cluster.local"</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Use alternative domain for services, e.g. &quot;myorg.internal&quot;.</p></td>
</tr>

<tr>
<td colspan="2">--skip-certificate-key-print</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Don't print the key used to encrypt the control-plane certificates.</p></td>
</tr>

<tr>
<td colspan="2">--skip-phases strings</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>List of phases to be skipped</p></td>
</tr>

<tr>
<td colspan="2">--skip-token-print</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Skip printing of the default bootstrap token generated by 'kubeadm init'.</p></td>
</tr>

<tr>
<td colspan="2">--token string</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The token to use for establishing bidirectional trust between nodes and control-plane nodes. The format is [a-z0-9]{6}.[a-z0-9]{16} - e.g. abcdef.0123456789abcdef</p></td>
</tr>

<tr>
<td colspan="2">--token-ttl duration&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Default: 24h0m0s</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire</p></td>
</tr>

<tr>
<td colspan="2">--upload-certs</td>
</tr>
<tr>
<td></td><td style="line-height: 130%; word-wrap: break-word;"><p>Upload control-plane certificates to the kubeadm-certs Secret.</p></td>
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

### Init workflow {#init-workflow}

`kubeadm init` bootstraps a Kubernetes control plane node by executing the
following steps:

1. Runs a series of pre-flight checks to validate the system state
   before making changes. Some checks only trigger warnings, others are
   considered errors and will exit kubeadm until the problem is corrected or the
   user specifies `--ignore-preflight-errors=<list-of-errors>`.

1. Generates a self-signed CA to set up identities for each component in the cluster. The user can provide their
   own CA cert and/or key by dropping it in the cert directory configured via `--cert-dir`
   (`/etc/kubernetes/pki` by default).
   The API server certs will have additional SAN entries for any `--apiserver-cert-extra-sans`
   arguments, lowercased if necessary.

1. Writes kubeconfig files in `/etc/kubernetes/` for the kubelet, the controller-manager, and the
   scheduler to connect to the API server, each with its own identity. Also
   additional kubeconfig files are written, for kubeadm as administrative entity (`admin.conf`)
   and for a super admin user that can bypass RBAC (`super-admin.conf`).

1. Generates static Pod manifests for the API server,
   controller-manager and scheduler. In case an external etcd is not provided,
   an additional static Pod manifest is generated for etcd.

   Static Pod manifests are written to `/etc/kubernetes/manifests`; the kubelet
   watches this directory for Pods to create on startup.

   Once control plane Pods are up and running, the `kubeadm init` sequence can continue.

1. Apply labels and taints to the control plane node so that no additional workloads will
   run there.

1. Generates the token that additional nodes can use to register
   themselves with a control plane in the future. Optionally, the user can provide a
   token via `--token`, as described in the
   [kubeadm token](/reference/setup-tools/kubeadm/kubeadm-token.md) documents.

1. Makes all the necessary configurations for allowing node joining with the
   [Bootstrap Tokens](/reference/access-authn-authz/bootstrap-tokens.md) and
   [TLS Bootstrap](/reference/access-authn-authz/kubelet-tls-bootstrapping.md)
   mechanism:

   - Write a ConfigMap for making available all the information required
     for joining, and set up related RBAC access rules.

   - Let Bootstrap Tokens access the CSR signing API.

   - Configure auto-approval for new CSR requests.

   See [kubeadm join](/reference/setup-tools/kubeadm/kubeadm-join.md) for additional information.

1. Installs a DNS server (CoreDNS) and the kube-proxy addon components via the API server.
   Please note that although the DNS server is deployed, it will not be scheduled until CNI is installed.

### Using init phases with kubeadm {#init-phases}

kubeadm allows you to create a control plane node in phases using the `kubeadm init phase` command.

To view the ordered list of phases and sub-phases you can call `kubeadm init --help`. The list
will be located at the top of the help screen and each phase will have a description next to it.
Note that by calling `kubeadm init` all of the phases and sub-phases will be executed in this exact order.

Some phases have unique flags, so if you want to have a look at the list of available options add
`--help`, for example:

```shell
sudo kubeadm init phase control-plane controller-manager --help
```

You can also use `--help` to see the list of sub-phases for a certain parent phase:

```shell
sudo kubeadm init phase control-plane --help
```

`kubeadm init` also exposes a flag called `--skip-phases` that can be used to skip certain phases.
The flag accepts a list of phase names and the names can be taken from the above ordered list.

An example:

```shell
sudo kubeadm init phase control-plane all --config=configfile.yaml
sudo kubeadm init phase etcd local --config=configfile.yaml
# you can now modify the control plane and etcd manifest files
sudo kubeadm init --skip-phases=control-plane,etcd --config=configfile.yaml
```

What this example would do is write the manifest files for the control plane and etcd in
`/etc/kubernetes/manifests` based on the configuration in `configfile.yaml`. This allows you to
modify the files and then skip these phases using `--skip-phases`. By calling the last command you
will create a control plane node with the custom manifest files.

**Feature state:** Kubernetes v1.22 [beta]

Alternatively, you can use the `skipPhases` field under `InitConfiguration`.

### Using kubeadm init with a configuration file {#config-file}

**Caution:**
The configuration file is still considered beta and may change in future versions.

It's possible to configure `kubeadm init` with a configuration file instead of command
line flags, and some more advanced features may only be available as
configuration file options. This file is passed using the `--config` flag and it must
contain a `ClusterConfiguration` structure and optionally more structures separated by `---\n`.
Mixing `--config` with others flags may not be allowed in some cases.

The default configuration can be printed out using the
[kubeadm config print](/reference/setup-tools/kubeadm/kubeadm-config.md) command.

If your configuration is not using the latest version it is **recommended** that you migrate using
the [kubeadm config migrate](/reference/setup-tools/kubeadm/kubeadm-config.md) command.

For more information on the fields and usage of the configuration you can navigate to our
[API reference page](/reference/config-api/kubeadm-config.v1beta4.md).

### Using kubeadm init with feature gates {#feature-gates}

kubeadm supports a set of feature gates that are unique to kubeadm and can only be applied
during cluster creation with `kubeadm init`. These features can control the behavior
of the cluster. Feature gates are removed after a feature graduates to GA.

To pass a feature gate you can either use the `--feature-gates` flag for
`kubeadm init`, or you can add items into the `featureGates` field when you pass
a [configuration file](/reference/config-api/kubeadm-config.v1beta4.md#kubeadm-k8s-io-v1beta4-ClusterConfiguration)
using `--config`.

Passing [feature gates for core Kubernetes components](/reference/command-line-tools-reference/feature-gates.md)
directly to kubeadm is not supported. Instead, it is possible to pass them by
[Customizing components with the kubeadm API](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/control-plane-flags/).

List of feature gates:

*kubeadm feature gates*

Feature | Default | Alpha | Beta | GA | Deprecated
:-------|:--------|:------|:-----|:---|:----------
`RootlessControlPlane` | `false` | 1.22 | - | - | 1.31

Feature gate descriptions:

`RootlessControlPlane`
: Setting this flag configures the kubeadm deployed control plane component static Pod containers
  for `kube-apiserver`, `kube-controller-manager`, `kube-scheduler` and `etcd` to run as non-root users.
  If the flag is not set, those components run as root. You can change the value of this feature gate before
  you upgrade to a newer version of Kubernetes.

List of removed feature gates:

*kubeadm removed feature gates*

Feature | Alpha | Beta | GA | Removed
:-------|:------|:-----|:---|:-------
`ControlPlaneKubeletLocalMode` | 1.31 | 1.33 | 1.35 | 1.36
`EtcdLearnerMode` | 1.27 | 1.29 | 1.32 | 1.33
`IPv6DualStack` | 1.16 | 1.21 | 1.23 | 1.24
`NodeLocalCRISocket` | 1.32 | 1.34 | 1.36 | 1.37
`PublicKeysECDSA` | 1.19 | - | - | 1.37
`UnversionedKubeletConfigMap` | 1.22 | 1.23 | 1.25 | 1.26
`UpgradeAddonsBeforeControlPlane` | 1.28 | - | - | 1.31
`WaitForAllControlPlaneComponents` | 1.30 | 1.33 | 1.34 | 1.35

Feature gate descriptions:

`ControlPlaneKubeletLocalMode`
: With this feature gate enabled, when joining a new control plane node, kubeadm will configure the kubelet
  to connect to the local kube-apiserver. This ensures that there will not be a violation of the version skew
  policy during rolling upgrades.

`EtcdLearnerMode`
: When joining a new control plane node, a new etcd member will be created
as a learner and promoted to a voting member only after the etcd data are fully aligned.

`IPv6DualStack`
: This flag helps to configure components dual stack when the feature is in progress. For more details on Kubernetes
  dual-stack support see [Dual-stack support with kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/dual-stack-support/).

`NodeLocalCRISocket`
: With this feature gate enabled, kubeadm will read/write the CRI socket for each node from/to the file
  `/var/lib/kubelet/instance-config.yaml` instead of reading/writing it from/to the annotation
  `kubeadm.alpha.kubernetes.io/cri-socket` on the Node object. The new file is applied as an instance
  configuration patch, before any other user managed patches are applied when the `--patches` flag
  is used. It contains a single field `containerRuntimeEndpoint` from the
  [KubeletConfiguration file format](/reference/config-api/kubelet-config.v1beta1.md). If the feature gate
  is enabled during upgrade, but the file `/var/lib/kubelet/instance-config.yaml` does not exist yet,
  kubeadm will attempt to read the CRI socket value from the file `/var/lib/kubelet/kubeadm-flags.env`.

`PublicKeysECDSA`
: Can be used to create a cluster that uses ECDSA certificates instead of the default RSA algorithm.
  Renewal of existing ECDSA certificates is also supported using `kubeadm certs renew`, but you cannot
  switch between the RSA and ECDSA algorithms on the fly or during upgrades. Kubernetes versions before v1.31
  had a bug where keys in generated kubeconfig files were set use RSA, even when you had enabled the
  `PublicKeysECDSA` feature gate. This feature gate is deprecated in favor of the `encryptionAlgorithm`
  functionality available in kubeadm v1beta4.

`UnversionedKubeletConfigMap`
: This flag controls the name of the [ConfigMap](/reference/glossary/configmap.md) where kubeadm stores
  kubelet configuration data. With this flag not specified or set to `true`, the ConfigMap is named `kubelet-config`.
  If you set this flag to `false`, the name of the ConfigMap includes the major and minor version for Kubernetes
  (for example: `kubelet-config-1.37`). Kubeadm ensures that RBAC rules for reading and writing
  that ConfigMap are appropriate for the value you set. When kubeadm writes this ConfigMap (during `kubeadm init`
  or `kubeadm upgrade apply`), kubeadm respects the value of `UnversionedKubeletConfigMap`. When reading that ConfigMap
  (during `kubeadm join`, `kubeadm reset`, `kubeadm upgrade`...), kubeadm attempts to use unversioned ConfigMap name first.
  If that does not succeed, kubeadm falls back to using the legacy (versioned) name for that ConfigMap.

`UpgradeAddonsBeforeControlPlane`
: This feature gate has been removed. It was introduced in v1.28 as a deprecated feature and then removed in v1.31.
  For documentation on older versions, please switch to the corresponding website version.

`WaitForAllControlPlaneComponents`
: With this feature gate enabled, kubeadm will wait for all control plane components (kube-apiserver,
  kube-controller-manager, kube-scheduler) on a control plane node to report status 200 on their `/livez`
  or `/healthz` endpoints. These checks are performed on `https://ADDRESS:PORT/ENDPOINT`.

  - `PORT` is taken from `--secure-port` of a component.
  - `ADDRESS` is `--advertise-address` for kube-apiserver and `--bind-address` for the
     kube-controller-manager and kube-scheduler.
  - `ENDPOINT` is only `/healthz` for kube-controller-manager until it supports `/livez` as well.

  If you specify custom `ADDRESS` or `PORT` in the kubeadm configuration they will be respected.
  Without the feature gate enabled, kubeadm will only wait for the kube-apiserver
  on a control plane node to become ready. The wait process starts right after the kubelet on the host
  is started by kubeadm. You are advised to enable this feature gate in case you wish to observe a ready
  state from all control plane components during the `kubeadm init` or `kubeadm join` command execution.

### Adding kube-proxy parameters {#kube-proxy}

For information about kube-proxy parameters in the kubeadm configuration see:

- [kube-proxy reference](/reference/config-api/kube-proxy-config.v1alpha1.md)

For information about enabling IPVS mode with kubeadm see:

- [IPVS](https://github.com/kubernetes/kubernetes/blob/master/pkg/proxy/ipvs/README.md)

### Passing custom flags to control plane components {#control-plane-flags}

For information about passing flags to control plane components see:

- [control-plane-flags](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/control-plane-flags/)

### Running kubeadm without an Internet connection {#without-internet-connection}

For running kubeadm without an Internet connection you have to pre-pull the required control plane images.

You can list and pull the images using the `kubeadm config images` sub-command:

```shell
kubeadm config images list
kubeadm config images pull
```

You can pass `--config` to the above commands with a [kubeadm configuration file](#config-file)
to control the `kubernetesVersion` and `imageRepository` fields.

All default `registry.k8s.io` images that kubeadm requires support multiple architectures.

### Using custom images {#custom-images}

By default, kubeadm pulls images from `registry.k8s.io`. If the
requested Kubernetes version is a CI label (such as `ci/latest`)
`gcr.io/k8s-staging-ci-images` is used.

You can override this behavior by using [kubeadm with a configuration file](#config-file).
Allowed customization are:

* To provide `kubernetesVersion` which affects the version of the images.
* To provide an alternative `imageRepository` to be used instead of
  `registry.k8s.io`.
* To provide a specific `imageRepository` and `imageTag` for etcd or CoreDNS.

Image paths between the default `registry.k8s.io` and a custom repository specified using
`imageRepository` may differ for backwards compatibility reasons. For example,
one image might have a subpath at `registry.k8s.io/subpath/image`, but be defaulted
to `my.customrepository.io/image` when using a custom repository.

To ensure you push the images to your custom repository in paths that kubeadm
can consume, you must:

* Pull images from the defaults paths at `registry.k8s.io` using `kubeadm config images {list|pull}`.
* Push images to the paths from `kubeadm config images list --config=config.yaml`,
  where `config.yaml` contains the custom `imageRepository`, and/or `imageTag` for etcd and CoreDNS.
* Pass the same `config.yaml` to `kubeadm init`.

#### Custom sandbox (pause) images {#custom-pause-image}

To set a custom image for these you need to configure this in your
[container runtime](/reference/glossary/container-runtime.md) to use the image.
Consult the documentation for your container runtime to find out how to change this setting;
for selected container runtimes, you can also find advice within the
[Container Runtimes](https://kubernetes.io/docs/setup/production-environment/container-runtimes/) topic.

### Uploading control plane certificates to the cluster

By adding the flag `--upload-certs` to `kubeadm init` you can temporary upload
the control plane certificates to a Secret in the cluster. Please note that this Secret
will expire automatically after 2 hours. The certificates are encrypted using
a 32byte key that can be specified using `--certificate-key`. The same key can be used
to download the certificates when additional control plane nodes are joining, by passing
`--control-plane` and `--certificate-key` to `kubeadm join`.

The following phase command can be used to re-upload the certificates after expiration:

```shell
kubeadm init phase upload-certs --upload-certs --config=SOME_YAML_FILE
```

**Note:**
A predefined `certificateKey` can be provided in `InitConfiguration` when passing the
[configuration file](/reference/config-api/kubeadm-config.v1beta4.md) with `--config`.

If a predefined certificate key is not passed to `kubeadm init` and
`kubeadm init phase upload-certs` a new key will be generated automatically.

The following command can be used to generate a new key on demand:

```shell
kubeadm certs certificate-key
```

### Certificate management with kubeadm

For detailed information on certificate management with kubeadm see
[Certificate Management with kubeadm](/tasks/administer-cluster/kubeadm/kubeadm-certs.md).
The document includes information about using external CA, custom certificates
and certificate renewal.

### Managing the kubeadm drop-in file for the kubelet {#kubelet-drop-in}

The `kubeadm` package ships with a configuration file for running the `kubelet` by `systemd`.
Note that the kubeadm CLI never touches this drop-in file. This drop-in file is part of the kubeadm
DEB/RPM package.

For further information, see
[Managing the kubeadm drop-in file for systemd](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/kubelet-integration/#the-kubelet-drop-in-file-for-systemd).

### Use kubeadm with CRI runtimes

By default, kubeadm attempts to detect your container runtime. For more details on this detection,
see the [kubeadm CRI installation guide](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#installing-runtime).

### Setting the node name

By default, kubeadm assigns a node name based on a machine's host address.
You can override this setting with the `--node-name` flag.
The flag passes the appropriate [`--hostname-override`](/reference/command-line-tools-reference/kubelet.md#options)
value to the kubelet.

Be aware that overriding the hostname can
[interfere with cloud providers](https://github.com/kubernetes/website/pull/8873).

### Automating kubeadm

Rather than copying the token you obtained from `kubeadm init` to each node, as
in the [basic kubeadm tutorial](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/),
you can parallelize the token distribution for easier automation. To implement this automation,
you must know the IP address that the control plane node will have after it is started, or use a
DNS name or an address of a load balancer.

1. Generate a token. This token must have the form `<6 character string>.<16 character string>`.
   More formally, it must match the regex: `[a-z0-9]{6}\.[a-z0-9]{16}`.

   kubeadm can generate a token for you:

   ```shell
   kubeadm token generate
   ```

1. Start both the control plane node and the worker nodes concurrently with this token.
   As they come up they should find each other and form the cluster. The same
   `--token` argument can be used on both `kubeadm init` and `kubeadm join`.

1. Similar can be done for `--certificate-key` when joining additional control plane
   nodes. The key can be generated using:

   ```shell
   kubeadm certs certificate-key
   ```

Once the cluster is up, you can use the `/etc/kubernetes/admin.conf` file from
a control plane node to talk to the cluster with administrator credentials or
[Generating kubeconfig files for additional users](/tasks/administer-cluster/kubeadm/kubeadm-certs.md#kubeconfig-additional-users).

Note that this style of bootstrap has some relaxed security guarantees because
it does not allow the root CA hash to be validated with
`--discovery-token-ca-cert-hash` (since it's not generated when the nodes are provisioned).
For details, see the [kubeadm join](/reference/setup-tools/kubeadm/kubeadm-join.md).

## What's next

* [kubeadm init phase](/reference/setup-tools/kubeadm/kubeadm-init-phase.md) to understand more about
  `kubeadm init` phases
* [kubeadm join](/reference/setup-tools/kubeadm/kubeadm-join.md) to bootstrap a Kubernetes
  worker node and join it to the cluster
* [kubeadm upgrade](/reference/setup-tools/kubeadm/kubeadm-upgrade.md) to upgrade a Kubernetes
  cluster to a newer version
* [kubeadm reset](/reference/setup-tools/kubeadm/kubeadm-reset.md) to revert any changes made
  to this host by `kubeadm init` or `kubeadm join`
