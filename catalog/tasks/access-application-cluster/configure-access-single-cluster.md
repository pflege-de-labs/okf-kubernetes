---
type: task
title: Configure Access to a Single Cluster
description: Set up a kubeconfig file to access a single existing Kubernetes cluster with kubectl.
resource: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-single-cluster/
tags: [kubernetes, tasks, access-application-cluster]
generated: {by: 'process:okf-kubernetes-okfgen', at: "2026-11-23T15:56:50Z"}
sources:
  - id: k8s-docs-source
    resource: https://github.com/kubernetes/website/blob/a81f5383d57959d479d946ae3ad2c620ceb11e6a/content/en/docs/tasks/access-application-cluster/configure-access-single-cluster.md
    title: 'Kubernetes website source: tasks/access-application-cluster/configure-access-single-cluster.md'
  - id: k8s-docs
    resource: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-single-cluster/
    title: Configure Access to a Single Cluster | Kubernetes documentation
kubernetes_version: v1.37
source_commit: a81f5383d57959d479d946ae3ad2c620ceb11e6a
source_path: content/en/docs/tasks/access-application-cluster/configure-access-single-cluster.md
---

# Configure Access to a Single Cluster

This page shows how to configure kubectl to access one existing cluster using
a supplied kubeconfig file or your cluster's connection details. For background,
see [Organizing Cluster Access Using kubeconfig Files](/concepts/configuration/organize-cluster-access-kubeconfig.md).

## Before you begin

* Have a running Kubernetes cluster. If you do not have one, create a local
  cluster with [Minikube](https://minikube.sigs.k8s.io/docs/start/) or choose
  another [learning environment](https://kubernetes.io/docs/setup/learning-environment/).
* [Install kubectl](/tasks/tools.md). Its version should be
  [within one minor version](/releases/version-skew-policy.md#kubectl) of your
  cluster's API server. Check your client version with `kubectl version --client`.
* Ensure your computer can reach the cluster's API server, using a VPN if required.
* Have a text editor and a private working directory.

Shell examples use Bash or Zsh on Linux and macOS, and PowerShell on Windows.

## Obtain your cluster access information

Ask your cluster administrator or provider for a personal kubeconfig file or
the information below. The example values match the commands later in this task;
replace them with the values for your cluster.

*Examples and sources of cluster access information*

| Information | Example | Where to obtain it |
| :--- | :--- | :--- |
| API server address | `https://api.example.com:6443` | Your administrator or the cluster provider's connection instructions. |
| Trusted certificate authority (CA) certificate | A PEM file named `ca.crt` | Ask your administrator for the file. A provider-generated kubeconfig already records the required CA file or data. Omit a separate CA file if the server certificate is trusted by your operating system. |
| Authentication method | A credential plugin such as `example-credential-plugin`, a signed client certificate, or a bearer token | Your provider's login instructions or your administrator's credential enrollment process. |
| Namespace and permissions | Namespace `default`, with permission to list Pods | Your administrator or the provider's access policy. |

For clusters managed on your behalf by a provider, follow the provider's kubeconfig instructions.
These instructions also explain any required login tools.

For kubeadm clusters, administrators can refer to
[Generating kubeconfig files for additional users](/tasks/administer-cluster/kubeadm/kubeadm-certs.md#kubeconfig-additional-users).

**Warning:**
Only use trusted kubeconfig files. They can execute credential plugins or read
local files. Inspect unfamiliar files before use.

### Try the task with Minikube

For a local example, install Minikube and a supported container or virtual-machine
manager using the [Minikube getting started guide](https://minikube.sigs.k8s.io/docs/start/).

1. Start the default Minikube cluster:

    ```shell
    minikube start --profile=minikube
    ```

    Minikube creates a `minikube` context with the API server address and local
    certificate settings, so you do not need to collect these values manually.

1. In your private working directory, choose an unused filename such as
    `cluster-config` and export that context:

   
   **Linux and macOS:**

   ```shell
   kubectl config view --context=minikube --minify --flatten --raw > cluster-config
   ```

   
   **Windows:**

   ```powershell
   kubectl config view --context=minikube --minify --flatten --raw | Set-Content -Path cluster-config -Encoding utf8
   ```

   
   

    `--flatten` embeds the certificate files, and `--raw` preserves credentials
    instead of redacting them. This file contains a private key; keep it private.

Continue with [Verify access to the cluster](#verify-access-to-the-cluster).
Your file is `cluster-config` and your context is `minikube`; skip the manual
configuration sections below.

## Use a supplied kubeconfig file

Follow your provider's setup instructions, including installing and signing in
to any credential plugin. For certificate authentication, use a locally generated
private key; see the [client certificate workflow](#configure-credentials).

1. Save the kubeconfig as `cluster-config` in your working directory. If it
    already exists, choose another name and use it below. Use absolute paths
    for any referenced files or credential plugin executables.

1. List the contexts in the supplied file:

    ```shell
    kubectl --kubeconfig=cluster-config config get-contexts
    ```

1. Select your cluster's context, replacing `CONTEXT_NAME` with a name from the
    output:

    ```shell
    kubectl --kubeconfig=cluster-config config use-context CONTEXT_NAME
    ```

Continue with [Verify access to the cluster](#verify-access-to-the-cluster).

## Create a kubeconfig from connection information

Start in a directory without a `cluster-config` file. These commands create it
without changing your existing kubeconfig. Replace the server URL, certificate
paths, and namespace with your cluster's values. You can keep the local names
`my-cluster`, `my-user`, and `my-context` if they are not already used in your
existing kubeconfig files.

### Configure the cluster

Save your cluster's trusted CA certificate as `ca.crt`, then record it and the
API server address using
[`kubectl config set-cluster`](/reference/kubectl/generated/kubectl_config/kubectl_config_set-cluster.md):

```shell
kubectl --kubeconfig=cluster-config config set-cluster my-cluster --server=https://api.example.com:6443 --certificate-authority=ca.crt
```

This records a reference to `ca.crt`, relative to the kubeconfig file's
directory. If the server's certificate is trusted by your operating system,
omit `--certificate-authority`.

### Configure credentials

Use
[`kubectl config set-credentials`](/reference/kubectl/generated/kubectl_config/kubectl_config_set-credentials.md)
to create an entry named `my-user` under `users` in `cluster-config`. This entry
tells kubectl which credential plugin, certificate and key, or token to use. Choose one of the following authentication methods.

**Credential plugin:**

A credential plugin retrieves credentials when kubectl connects to the API server.

1. Install the plugin and sign in using your provider's instructions. Its
    executable must be on `PATH`, or specified by its full path.

1. Configure the plugin, replacing `example-credential-plugin` with its executable:

    ```shell
    kubectl --kubeconfig=cluster-config config set-credentials my-user --exec-command=example-credential-plugin --exec-api-version=client.authentication.k8s.io/v1 --exec-interactive-mode=IfAvailable
    ```

    Use the API version and interactive mode required by your provider. Add
    `--exec-arg=VALUE` and `--exec-env=NAME=VALUE` for required arguments and
    environment variables. The plugin's response must match `--exec-api-version`.
    See [client-go credential plugins](/reference/access-authn-authz/authentication.md#client-go-credential-plugins)
    for details.

**Client certificate:**

Use this method if your cluster supports X.509 client certificates. It requires
the `openssl` command-line tool.

1. If you need a new certificate, generate a private key and certificate signing
    request (CSR). Use unused filenames, and replace `YOUR_USERNAME` with the
    username agreed with your administrator, such as `jane`:

    ```shell
    openssl genrsa -out my-user.key 3072
    openssl req -new -key my-user.key -out my-user.csr -subj "/CN=YOUR_USERNAME"
    ```

1. If you generated a CSR, send only `my-user.csr` to your administrator; keep
    `my-user.key` locally. Save the signed certificate as `my-user.crt`.

    Your administrator can follow the
    [CertificateSigningRequest workflow](/tasks/tls/certificate-issue-client-csr.md)
    where supported, using existing cluster access and the required permissions.

1. Record your certificate and key:

    ```shell
    kubectl --kubeconfig=cluster-config config set-credentials my-user --client-certificate=my-user.crt --client-key=my-user.key --embed-certs=true
    ```

    This embeds the certificate and private key in `cluster-config`.
    Protect that file as you would the private key.

**Long-lived bearer token:**

Use this method if your administrator or provider supplies a long-lived bearer
token for your identity that the cluster accepts.

Replace `YOUR_TOKEN` with that token:

```shell
kubectl --kubeconfig=cluster-config config set-credentials my-user --token=YOUR_TOKEN
```

The token is stored in `cluster-config`. Anyone who obtains it can authenticate
as that identity until the token expires or is revoked. Keep the file private
and follow your administrator's instructions for replacing the token.

### Configure and select a context

A context saves the cluster, credential entry, and namespace to use together.

1. Create a context that connects to `my-cluster` using the credentials saved as
    `my-user`, with `default` as its namespace. Use
    [`kubectl config set-context`](/reference/kubectl/generated/kubectl_config/kubectl_config_set-context.md):

    ```shell
    kubectl --kubeconfig=cluster-config config set-context my-context --cluster=my-cluster --user=my-user --namespace=default
    ```

    Here, `--user=my-user` selects the entry from the previous section. If that
    entry uses a certificate issued to `jane`, the cluster authenticates you as
    `jane`; `my-user` simply names the stored credentials in your file.

1. Select it:

    ```shell
    kubectl --kubeconfig=cluster-config config use-context my-context
    ```

1. Inspect the configuration:

    ```shell
    kubectl --kubeconfig=cluster-config config view
    ```

    For the credential-plugin example, the output is similar to:

    ```yaml
    apiVersion: v1
    clusters:
    - cluster:
        certificate-authority: ca.crt
        server: https://api.example.com:6443
      name: my-cluster
    contexts:
    - context:
        cluster: my-cluster
        namespace: default
        user: my-user
      name: my-context
    current-context: my-context
    kind: Config
    users:
    - name: my-user
      user:
        exec:
          apiVersion: client.authentication.k8s.io/v1
          args: null
          command: example-credential-plugin
          env: null
          interactiveMode: IfAvailable
          provideClusterInfo: false
    ```

    Certificate authentication shows certificate and key fields instead of
    `exec`; bearer token authentication uses a `token` field.

**Note:**
Add `--minify` to `kubectl config view` to show only the configuration associated
with the current context.

## Verify access to the cluster

Use your assigned namespace in place of `default`.

1. Check the identity recognized by the cluster:

    ```shell
    kubectl --kubeconfig=cluster-config auth whoami
    ```

    Confirm that the username and groups are as expected.

1. Check whether you can list Pods in your namespace:

    ```shell
    kubectl --kubeconfig=cluster-config auth can-i list pods --namespace=default
    ```

    If it returns `no`, ask your administrator for the required permissions.
    Authentication alone does not grant access to resources.

1. If the response is `yes`, list the Pods:

    ```shell
    kubectl --kubeconfig=cluster-config get pods --namespace=default
    ```

    `No resources found` is normal if the namespace has no Pods.

## Save the kubeconfig in the default location

After following any of the setup options above, save `cluster-config` as your
default kubeconfig: `$HOME/.kube/config` on Linux and macOS, or
`$env:USERPROFILE\.kube\config` on Windows. Kubectl reads this file automatically
when `KUBECONFIG` env variable is unset or empty and you do not specify `--kubeconfig`.
You do not need to set `KUBECONFIG` to use the default file.

If your kubeconfig uses `certificate-authority: ca.crt`, change that value to
the CA file's absolute path before copying the kubeconfig.

If the default file does not exist, copy `cluster-config` from your working
directory:

**Linux and macOS:**

```shell
mkdir -p "$HOME/.kube"
cp cluster-config "$HOME/.kube/config"
chmod 600 "$HOME/.kube/config"
```

**Windows:**

```powershell
New-Item -ItemType Directory -Path "$env:USERPROFILE\.kube" -Force | Out-Null
Copy-Item cluster-config "$env:USERPROFILE\.kube\config"
```

## Use additional kubeconfig files (optional)

If you prefer to keep configurations in separate files, use `KUBECONFIG` to
select the files kubectl reads. This is optional; your saved default kubeconfig
works without it. Paths in `KUBECONFIG` are separated by `:` on Linux and macOS,
or `;` on Windows. An explicit `--kubeconfig` flag selects only that file instead.

The following examples save the current setting, append the default file and
`cluster-config`, and make a request. Run them
from the directory containing `cluster-config`. Replace `CONTEXT_NAME` with
`my-context` for the manual example, `minikube` for the Minikube example, or the
context selected in your supplied file. Replace `default` with your assigned
namespace:

**Linux and macOS:**

```shell
export KUBECONFIG_SAVED="$KUBECONFIG"
export KUBECONFIG="${KUBECONFIG}:$HOME/.kube/config:$PWD/cluster-config"
kubectl --context=CONTEXT_NAME get pods --namespace=default
```

**Windows:**

```powershell
$env:KUBECONFIG_SAVED = $env:KUBECONFIG
$env:KUBECONFIG = "$env:KUBECONFIG;$env:USERPROFILE\.kube\config;$PWD\cluster-config"
kubectl --context=CONTEXT_NAME get pods --namespace=default
```

Earlier files take precedence for duplicate entries. For more information, see
[Configure Access to Multiple Clusters](/tasks/access-application-cluster/configure-access-multiple-clusters.md).

If you followed the optional section above, run the following in the same
terminal to restore your previous `KUBECONFIG` setting:

**Linux and macOS:**

```shell
export KUBECONFIG="$KUBECONFIG_SAVED"
```

**Windows:**

```powershell
$env:KUBECONFIG = $env:KUBECONFIG_SAVED
```

## Troubleshoot access

If a command fails, see
[Troubleshooting kubectl](/tasks/debug.md)
for help with kubeconfig settings, connectivity, authentication, and TLS errors.

## Other uses of kubeconfig

Components such as the
[kubelet](/reference/command-line-tools-reference/kubelet.md) and
[kube-proxy](/reference/command-line-tools-reference/kube-proxy.md) also use
kubeconfig files. Follow their documentation for the required identity and
file settings.

## What's next

### Customize kubectl preferences

Your kubeconfig provides cluster connection and authentication settings.
A separate, optional `kuberc` file configures kubectl defaults and aliases;
you still use kubeconfig for cluster access.

If authentication uses an exec credential plugin, kuberc's
[credential plugin policy](/reference/kubectl/kuberc.md#credential-plugin-policy)
can control whether that plugin may run. This helps protect you against
malicious kubeconfig files.

See
[Kubectl user preferences (kuberc)](/reference/kubectl/kuberc.md)
for instructions on writing a separate `kuberc` file.

### Additional reading

* [Configure Access to Multiple Clusters](/tasks/access-application-cluster/configure-access-multiple-clusters.md)
* [Organizing Cluster Access Using kubeconfig Files](/concepts/configuration/organize-cluster-access-kubeconfig.md)
* [kubeconfig (v1) reference](/reference/config-api/kubeconfig.v1.md)
* [`kubectl config` reference](/reference/kubectl/generated/kubectl_config.md)
